package objectstorage

import (
	"fmt"

	"k8s.io/api"
	"k8s.io/apimachinery"
	"k8s.io/apiserver"
	"k8s.io/client-go"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func newKubeClient() (*kubernetes.Clientset, error) {
	cfg, err := rest.InClusterConfig()
	if err != nil {
		return nil, fmt.Errorf("in cluster config: %w", err)
	}
	return kubernetes.NewForConfig(cfg)
}
