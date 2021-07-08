// Copyright (c) 2021 SIGHUP s.r.l All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package kube

import (
    "context"
	"testing"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
)

func TestGetService(t *testing.T) {
	svcObj := &corev1.Service{ObjectMeta: metav1.ObjectMeta{Name: "mySvc", Namespace: "myNs"}}
	fakeKC := &KubernetesClient{Client: fake.NewSimpleClientset(svcObj)}

	// Get service of name `mySvc` in ns `myNs`. Should not be err
	var ctx context.Context
	svc, err01 := fakeKC.GetService(&ctx, "mySvc", "myNs")
	if svc.Name != "mySvc" {
		t.Fatal(err01.Error())
	}
	if err01 != nil {
		t.Fatal(err01.Error())
	}

	// Test 02 Get service of name `invalidSvc` in ns `myNs`. Should be err
	_, err02 := fakeKC.GetService(&ctx, "invalidSvc", "myNs")
	if err02 == nil {
		t.Fatal("Test 02 failed. invalidSvc should not be found in myNs")
	}

	// Test 03 Get service of name `mySvc` in ns `invalidNs`. Should be err
	_, err03 := fakeKC.GetService(&ctx, "mySvc", "invalidNs")
	if err03 == nil {
		t.Fatal("Test 03 failed. mySvc should not be found in invalidNs")
	}
}

func TestGetEndpoints(t *testing.T) {
	var ctx context.Context
	labels := make(map[string]string)
	labels["app"] = "sample"
	labels02 := make(map[string]string)
	labels02["app2"] = "invalid"

	podObj := &corev1.Pod{ObjectMeta: metav1.ObjectMeta{Name: "my-pod",
		Namespace: "myNs", Labels: labels}}

	svcObj := &corev1.Service{ObjectMeta: metav1.ObjectMeta{Name: "mySvc", Namespace: "myNs"},
		Spec: corev1.ServiceSpec{Selector: labels}}
	badSvcObj := &corev1.Service{ObjectMeta: metav1.ObjectMeta{Name: "mySvc2", Namespace: "myNs"},
		Spec: corev1.ServiceSpec{Selector: labels02}}

	epObj := &corev1.Endpoints{ObjectMeta: metav1.ObjectMeta{Name: "mySvc",
		Namespace: "myNs", Labels: labels}, Subsets: []corev1.EndpointSubset{}}
	fakeKC := &KubernetesClient{Client: fake.NewSimpleClientset(podObj,
		svcObj, epObj, badSvcObj)}

	// Test 06 Get the endpoints corresponsponding to svc in namespace
	// `myNs`. should not be err
	_, err06 := fakeKC.GetEndpoints(&ctx, svcObj, "myNs")
	if err06 != nil {
		t.Fatal(err06)
	}
}
