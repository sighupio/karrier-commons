# Namespaces

This example project has been developed to showcase how to use the `KubernetesClient` struct exposed at
the `github.com/sighupio/fip-commons/pkg/kube` package.

## TL;DR

The code is just representation of the required sentences to develop a go client with this library.

```go
package main

import (
    "context"

    "github.com/sighupio/fip-commons/pkg/kube"
    v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func main() {
    // Create the `KubernetesClient`
    k := kube.KubernetesClient{KubeConfig: "/tmp/my-kubeconfig-file"}
    // Init the client
    err := k.Init()
    // Get a context
    ctx := context.TODO()
    // Healthz checks the connection to the Kubernetes API server
    err = k.Healthz(&ctx)
    // Then you can use the Clien attribute to use the Kubernetes client-go library
    nsList, err := k.Client.CoreV1().Namespaces().List(ctx, v1.ListOptions{})
}
```

Change it to fit your own requirements. [Take a look to the code](main.go) implemented in this example to see a proper
working example.

## Run it

```bash
$ cd examples/namespaces
$ go run main.go /tmp/kubeconfig
KubernetesClient is ready!
Namespaces in the cluster:
 - default
 - kube-node-lease
 - kube-public
 - kube-system
 - local-path-storage
```
