package kube

import (
	"fmt"

	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	clientcmddapi "k8s.io/client-go/tools/clientcmd/api"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/discovery"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func RESTClientForGroupVersion(k *Kube, gv schema.GroupVersion) (rest.Interface, error) {
	cfg, err := buildConfig(getAddr(k.APIHost, k.APIPort), k.Auth)
	if err != nil {
		return nil, err
	}

	err = setGroupDefaults(cfg, gv)
	if err != nil {
		return nil, err
	}

	return rest.RESTClientFor(cfg)
}

func DiscoveryClient(k *Kube) (*discovery.DiscoveryClient, error) {
	cfg, err := buildConfig(getAddr(k.APIHost, k.APIPort), k.Auth)
	if err != nil {
		return nil, err
	}

	return discovery.NewDiscoveryClientForConfig(cfg)
}

func buildConfig(addr string, auth Auth) (*rest.Config, error) {
	cfg := clientcmddapi.Config{
		AuthInfos: map[string]*clientcmddapi.AuthInfo{
			auth.Username: {
				Token: auth.Token,
				ClientCertificateData: []byte(auth.Cert),
				ClientKeyData:         []byte(auth.Key),
			},
		},
		Clusters: map[string]*clientcmddapi.Cluster{
			"sg-" + auth.Username: {
				Server: addr,
				CertificateAuthorityData: []byte(auth.CA),
			},
		},
		Contexts: map[string]*clientcmddapi.Context{
			"sg-" + auth.Username: {
				AuthInfo: auth.Username,
				Cluster:  "sg-" + auth.Username,
			},
		},
		CurrentContext: "sg-" + auth.Username,
	}
	return clientcmd.NewNonInteractiveClientConfig(cfg, "sg-" + auth.Username, &clientcmd.ConfigOverrides{}, nil).ClientConfig()
}

func getAddr(host, port string) string {
	return fmt.Sprintf("https://%s:%s", host, port)
}

//TODO: review options
func setGroupDefaults(config *rest.Config, gv schema.GroupVersion) error {
	config.GroupVersion = &gv
	if len(gv.Group) == 0 {
		config.APIPath = "/api"
	} else {
		config.APIPath = "/apis"
	}
	config.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: scheme.Codecs}
	if len(config.UserAgent) == 0 {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}
	return nil
}
