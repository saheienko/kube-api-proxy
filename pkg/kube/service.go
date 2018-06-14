package kube

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/rest"
)

type Kube struct {
	ID      string `json:"id"`
	APIHost string `json:"apiHost"`
	APIPort string `json:"apiPort"`
	Auth    Auth   `json:"auth"`
}

type Auth struct {
	Username string `json:"username"`
	Token    string `json:"token"`
	CA       string `json:"ca"`
	Cert     string `json:"cert"`
	Key      string `json:"key"`
}

type Service struct {
	discoveryClientFn func(k *Kube) (*discovery.DiscoveryClient, error)
	clientForGroupFn  func(k *Kube, gv schema.GroupVersion) (rest.Interface, error)

	mu    sync.RWMutex
	kubes map[string]Kube
}

func NewService() *Service {
	return &Service{
		clientForGroupFn:  RESTClientForGroupVersion,
		discoveryClientFn: DiscoveryClient,
		mu:                sync.RWMutex{},
		kubes:             map[string]Kube{},
	}
}

func (s *Service) CreateKube(w http.ResponseWriter, r *http.Request) {
	k := &Kube{}
	err := json.NewDecoder(r.Body).Decode(k)
	if err != nil {
		log.Printf("createKube: %v", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	k.ID = getHash(k.APIHost + k.APIPort + k.Auth.Username)
	if k.ID == "" {
		log.Printf("createKube: kube id: %q", k.ID)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	if !s.storeKube(k) {
		http.Error(w, "has already exists", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("has been added: " + k.ID))
}

func (s *Service) GetKube(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if vars == nil {
		http.Error(w, "id not found", http.StatusBadRequest)
	}

	id := vars["id"]

	k, ok := s.getKube(id)
	if !ok {
		http.Error(w, "not found", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(k)
}

func (s *Service) DeleteKube(w http.ResponseWriter, r *http.Request) {
	notImplemented(w, r)
}

func (s *Service) ListKubes(w http.ResponseWriter, r *http.Request) {
	notImplemented(w, r)
}

func (s *Service) ListResources(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if vars == nil {
		http.Error(w, "id not found", http.StatusBadRequest)
		return
	}

	id := vars["id"]

	kube, ok := s.getKube(id)
	if !ok {
		http.Error(w, "not found", http.StatusInternalServerError)
		return
	}

	client, err := s.discoveryClientFn(kube)
	if err != nil {
		log.Printf("get client for %s kube: %v", id, err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	apiResourceLists, err := client.ServerResources()
	if err != nil {
		log.Printf("kube %s: get server resouces: %v", id, err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	discoveredExpansions := map[string][]schema.GroupVersion{}
	for _, apiResourceList := range apiResourceLists {
		gv, err := schema.ParseGroupVersion(apiResourceList.GroupVersion)
		if err != nil {
			continue
		}
		// Collect GroupVersions by categories
		for _, apiResource := range apiResourceList.APIResources {
			if _, ok := discoveredExpansions[apiResource.Kind]; !ok {
				discoveredExpansions[apiResource.Name] = append(discoveredExpansions[apiResource.Kind], gv)
			}
		}
	}

	raw, err := json.Marshal(discoveredExpansions)
	if err != nil {
		log.Printf("kube %s: marshal discovered expansions: %v", id, err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	w.Write(raw)
}

func (s *Service) GetResource(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if vars == nil {
		http.Error(w, "id not found", http.StatusBadRequest)
	}

	id := vars["id"]
	rs := vars["resource"]
	ns := r.URL.Query().Get("namespace")
	log.Printf("getResource: kube=%s ns=%s rs=%s", id, ns, rs)

	kube, ok := s.getKube(id)
	if !ok {
		http.Error(w, "not found", http.StatusInternalServerError)
		return
	}

	gv, err := s.groupForResource(kube, rs)
	if err != nil {
		log.Printf("get group version for %s kube: %v", id, err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	client, err := s.clientForGroupFn(kube, gv)
	if err != nil {
		log.Printf("get client for %s kube: %v", id, err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	raw, err := client.Get().Resource(rs).Namespace(ns).DoRaw()
	if err != nil {
		log.Printf("kube %s: namespace %s: resource %s: %v", id, ns, rs, err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	w.Write(raw)
}

func (s *Service) getKube(id string) (*Kube, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	k, ok := s.kubes[id]
	return &k, ok
}

func (s *Service) storeKube(k *Kube) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.kubes[k.ID]; ok {
		return false
	}

	s.kubes[k.ID] = *k
	return true
}

func (s *Service) groupForResource(kube *Kube, resourceName string) (schema.GroupVersion, error) {
	client, err := s.discoveryClientFn(kube)
	if err != nil {
		return schema.GroupVersion{}, err
	}

	apiResourceLists, err := client.ServerResources()
	if err != nil {
		return schema.GroupVersion{}, err
	}

	for _, apiResourceList := range apiResourceLists {
		gv, err := schema.ParseGroupVersion(apiResourceList.GroupVersion)
		if err != nil {
			continue
		}
		// Collect GroupVersions by categories
		for _, apiResource := range apiResourceList.APIResources {
			if apiResource.Name == resourceName {
				return gv, nil
			}
		}
	}
	return schema.GroupVersion{}, errors.New("not found")
}

func notImplemented(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "not implemented!\n")
}

func getHash(text string) string {
	hasher := sha256.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
