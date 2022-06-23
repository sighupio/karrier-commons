// Copyright (c) 2021 SIGHUP s.r.l All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package kube

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (kc *KubernetesClient) GetService(ctx *context.Context, svcName string,
	namespace string,
) (*corev1.Service, error) {
	service, err := kc.Client.CoreV1().Services(namespace).Get(*ctx,
		svcName, metav1.GetOptions{})

	return service, err
}

func (kc *KubernetesClient) GetEndpoints(ctx *context.Context, service *corev1.Service,
	namespace string,
) (*corev1.Endpoints, error) {
	// Retrieve all the endpoints corresponding to the service
	// Name of the endpoint will always match that of the svc
	endpoint, err := kc.Client.CoreV1().Endpoints(namespace).Get(*ctx,
		service.Name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	return endpoint, err
}
