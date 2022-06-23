// Copyright (c) 2021 SIGHUP s.r.l All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package v1

// ClusterMetadata represents the basic information of a cluster.
type ClusterMetadata struct {
	Version    string    `json:"version"`    // The version of the Control Plane.
	Nodes      []Node    `json:"nodes"`      // List of Nodes of the cluster.
	Namespaces []string  `json:"namespaces"` // Available namespaces in the cluster.
	Ingresses  []Ingress `json:"ingresses"`  // Hostnames managed by the cluster.
}

type Ingress struct {
	Name     string `json:"name"`
	Hostname string `json:"hostname"`
}

// Node represents the basic information of a cluster node.
type Node struct {
	Version string `json:"version"` // The Kubelet version of the node.
	Name    string `json:"name"`    // Name of the node.
	OS      string `json:"os"`      // OS os the node.
	CRI     string `json:"cri"`     // The container runtime version.
	Kernel  string `json:"kernel"`  // Kernel version.
	IP      string `json:"ip"`      // The internal IP of the node.
	Pods    []Pod  `json:"pods"`    // The list of pods the node is running.
	CPU     string `json:"cpu"`     // Amount of allocable CPU.
	Memory  string `json:"memory"`  // The amount of allocable Memory.
	Ready   bool   `json:"ready"`   // Mark if the node is in Ready state.
	Master  bool   `json:"master"`  // Mark if it belong to the control plane.
	Role    string `json:"role"`    // Role from label.
}

// Pod represents the basic information of a pod.
type Pod struct {
	Namespace string `json:"namespace"` // The namespace.
	Name      string `json:"name"`      // The pod name.
}
