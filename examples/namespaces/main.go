// Copyright (c) 2021 SIGHUP s.r.l All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/sighupio/fip-commons/pkg/kube"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func main() {
	// Get the arguments of the program. Check there is only one
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("error. Ensure passing just one argument, the kubeconfig path")
		os.Exit(1)
	}

	// Get the kubeconfig from the arguments
	kubeConfigPath := args[0]
	if kubeConfigPath == "" {
		fmt.Println("error. You must pass the kubeconfig path as argument")
		os.Exit(1)
	}

	// Create the client using the library
	k := kube.KubernetesClient{KubeConfig: kubeConfigPath}
	err := k.Init()

	if err != nil {
		fmt.Println("error. Something happened while trying to get connection to the API Server")
		os.Exit(1)
	}

	ctx := context.TODO()

	// Checking the health of the Kubernetes API Server with our client
	err = k.Healthz(&ctx)
	if err != nil {
		fmt.Println("error. Cluster seems to be not healthy")
		os.Exit(1)
	}

	fmt.Println("KubernetesClient is ready!")

	// Querying the namespaces in the cluster
	nsList, err := k.Client.CoreV1().Namespaces().List(ctx, v1.ListOptions{})
	if err != nil {
		fmt.Println("error. Can not get the list of namespaces")
		os.Exit(1)
	}

	fmt.Println("Namespaces in the cluster: ")

	for _, ns := range nsList.Items {
		fmt.Printf(" - %v\n", ns.Name)
	}
}
