package objectstorage

import (
	"context"
	"fmt"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

// Orchestrate handles the orchestration of Worker k8s clusters.
// These sit outside the main cluster containing frontend, backend etc...

//TODO: refactor to ensure newKubeClient returns a ref to &rest.Config, host, tls, authToken, CAData

// Creates new Kubernetes client.
// return kubernetes.Clientset or error.
func newKubeClient() (*kubernetes.Clientset, error) {
	cfg, err := rest.InClusterConfig()
	if err != nil {
		return nil, fmt.Errorf("in cluster config: %w", err)
	}
	return kubernetes.NewForConfig(cfg)
}

func CreateBucketPod(ctx context.Context, bucketID string, size uint64) error {
	panic("Not implemented yet.")
}
