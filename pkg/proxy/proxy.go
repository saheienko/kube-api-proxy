package proxy

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/saheienko/kube-api-proxy/pkg/kube"
)

type Proxy struct {
	addr string
	svc  *kube.Service
}

func New(host, port string) *Proxy {
	if port == "" {
		port = "8080"
	}

	return &Proxy{
		addr: fmt.Sprintf("%s:%s", host, port),
		svc:  kube.NewService(),
	}
}

func (p *Proxy) Run() error {
	r := mux.NewRouter()

	r.HandleFunc("/kubes", p.svc.CreateKube).Methods(http.MethodPost)
	r.HandleFunc("/kubes", p.svc.ListKubes).Methods(http.MethodGet)
	r.HandleFunc("/kubes/{id}", p.svc.GetKube).Methods(http.MethodGet)
	r.HandleFunc("/kubes/{id}", p.svc.DeleteKube).Methods(http.MethodDelete)
	r.HandleFunc("/kubes/{id}/list", p.svc.ListResources).Methods(http.MethodGet)
	r.HandleFunc("/kubes/{id}/resources/{resource}", p.svc.GetResource).Methods(http.MethodGet)

	s := http.Server{
		Addr:         p.addr,
		Handler:      r,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}

	log.Printf("proxy listens on: %s\n", s.Addr)
	return s.ListenAndServe()
}
