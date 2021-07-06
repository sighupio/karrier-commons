# Examples

This directory contains a sample golang project that uses the library to list namespaces.

## Namespaces

In the [namespaces](./namespaces) directory you'll find an example golang project that creates a Kubernetes Client
using the library exposed in this project.

Feel free to run it by running the following command from the project root:

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
