package magicalroleapi

import (
	"errors"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

// ClusterConn is the type of connection to the K8S cluster
type ClusterConn int

const (
	// InCluster references the cluster the API is executed within.
	InCluster ClusterConn = iota
	// OutOfCluster references external cluster.
	OutOfCluster
)

type k8sClientSource struct {
	mode ClusterConn
}

func (cs *k8sClientSource) client() (*kubernetes.Clientset, error) {
	if cs.mode != InCluster {
		return nil, errors.New("only in-cluster client supported at this time")
	}
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	return clientset, nil
}
