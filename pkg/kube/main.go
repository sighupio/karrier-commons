// Copyright (c) 2021 SIGHUP s.r.l All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package kube

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// KubernetesClient represents the Kubernetes configuration of the project.
type KubernetesClient struct {
	ctx *context.Context

	KubeConfig string
	Client     kubernetes.Interface
}

// Init initializes the Kubernetes client-go.
func (kc *KubernetesClient) Init(ctx *context.Context) error {
	var (
		config *rest.Config
		err    error
	)

	kc.ctx = ctx

	if kc.KubeConfig != "" {
		config, err = kc.getConfigFromFile(kc.KubeConfig)
	} else {
		// if no kubeconfigfile is provided creates the in-cluster config
		config, err = kc.inClusterConfig()

		if err != nil {
			// If inCluster config does not work, try with the default kube config path
			config, err = kc.extClusterConfig()
		}
	}

	if err != nil {
		return err
	}
	// return k8s client and err
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		return err
	}

	kc.Client = client

	return nil
}

func (kc *KubernetesClient) inClusterConfig() (*rest.Config, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}

	return config, nil
}

func (kc *KubernetesClient) extClusterConfig() (*rest.Config, error) {
	if home := os.Getenv("HOME"); home != "" {
		kubeConfigPath := filepath.Join(home, ".kube", "config")
		config, err := kc.getConfigFromFile(kubeConfigPath)

		if err != nil {
			return nil, err
		}

		return config, nil
	}

	return nil, fmt.Errorf("can not configure external cluster configuration from the default $HOME/.kube/config path")
}

func (kc *KubernetesClient) getConfigFromFile(kubeConfigPath string) (*rest.Config, error) {
	kubeConfigContent, err := ioutil.ReadFile(kubeConfigPath)
	if err != nil {
		return nil, err
	}

	return clientcmd.RESTConfigFromKubeConfig(kubeConfigContent)
}
