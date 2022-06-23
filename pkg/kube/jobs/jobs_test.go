// Copyright (c) 2021 SIGHUP s.r.l All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build unit

package jobs

import (
	"context"
	"reflect"
	"testing"

	"github.com/sighupio/fip-commons/pkg/kube"
	batchv1 "k8s.io/api/batch/v1"
	batchv1beta1 "k8s.io/api/batch/v1beta1"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
)

func TestGetPodsfromJob(t *testing.T) {
	backoff := int32(4)
	podObj := &corev1.Pod{ObjectMeta: v1.ObjectMeta{Name: "hello-12345", Namespace: "default", Labels: map[string]string{"job-name": "hello"}}}
	type args struct {
		ctx context.Context
		kc  *kube.KubernetesClient
		job batchv1.Job
	}
	tests := []struct {
		name    string
		args    args
		want    []corev1.Pod
		wantErr bool
	}{
		{
			name: "Empty List",
			args: args{
				ctx: context.TODO(),
				kc: &kube.KubernetesClient{
					Client: fake.NewSimpleClientset(),
				},
				job: batchv1.Job{
					TypeMeta: v1.TypeMeta{},
					ObjectMeta: v1.ObjectMeta{
						Name: "demo",
					},
					Spec: batchv1.JobSpec{
						BackoffLimit: &backoff,
					},
					Status: batchv1.JobStatus{},
				},
			},
			want:    nil,
			wantErr: false,
		}, {
			name: "One Pod",
			args: args{
				ctx: context.TODO(),
				kc: &kube.KubernetesClient{
					Client: fake.NewSimpleClientset(podObj),
				},
				job: batchv1.Job{
					TypeMeta: v1.TypeMeta{},
					ObjectMeta: v1.ObjectMeta{
						Name: "hello",
					},
					Spec: batchv1.JobSpec{
						BackoffLimit: &backoff,
					},
					Status: batchv1.JobStatus{},
				},
			},
			want:    []corev1.Pod{*podObj},
			wantErr: false,
		}, {
			name: "No Pods",
			args: args{
				ctx: context.TODO(),
				kc: &kube.KubernetesClient{
					Client: fake.NewSimpleClientset(podObj),
				},
				job: batchv1.Job{
					TypeMeta: v1.TypeMeta{},
					ObjectMeta: v1.ObjectMeta{
						Name: "bye",
					},
					Spec: batchv1.JobSpec{
						BackoffLimit: &backoff,
					},
					Status: batchv1.JobStatus{},
				},
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetPodsfromJob(tt.args.ctx, tt.args.kc, tt.args.job)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPodsfromJob() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPodsfromJob() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetJob(t *testing.T) {
	backoff := int32(4)
	job1 := batchv1.Job{
		TypeMeta: v1.TypeMeta{},
		ObjectMeta: v1.ObjectMeta{
			Name:      "hello",
			Namespace: "default",
		},
		Spec: batchv1.JobSpec{
			BackoffLimit: &backoff,
		},
		Status: batchv1.JobStatus{},
	}
	type args struct {
		ctx       context.Context
		kc        *kube.KubernetesClient
		namespace string
		name      string
	}
	tests := []struct {
		name    string
		args    args
		want    *batchv1.Job
		wantErr bool
	}{
		{
			name: "Not Found error",
			args: args{
				ctx: context.TODO(),
				kc: &kube.KubernetesClient{
					Client: fake.NewSimpleClientset(),
				},
				namespace: "default",
				name:      "hello",
			},
			want:    nil,
			wantErr: true,
		}, {
			name: "Found",
			args: args{
				ctx: context.TODO(),
				kc: &kube.KubernetesClient{
					Client: fake.NewSimpleClientset(&job1),
				},
				namespace: "default",
				name:      "hello",
			},
			want:    &job1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetJob(tt.args.ctx, tt.args.kc, tt.args.namespace, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetJob() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetJob() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetCronJobFromJob(t *testing.T) {
	backoff := int32(4)
	job1 := batchv1.Job{
		TypeMeta: v1.TypeMeta{},
		ObjectMeta: v1.ObjectMeta{
			Name:      "hello-12345",
			Namespace: "default",
			OwnerReferences: []v1.OwnerReference{
				{
					Kind: "CronJob",
					Name: "hello",
				},
			},
		},
		Spec: batchv1.JobSpec{
			BackoffLimit: &backoff,
		},
		Status: batchv1.JobStatus{},
	}
	cronjob1 := batchv1beta1.CronJob{
		TypeMeta: v1.TypeMeta{},
		ObjectMeta: v1.ObjectMeta{
			Name:      "hello",
			Namespace: "default",
		},
		Spec: batchv1beta1.CronJobSpec{
			Schedule: "* * * * *",
		},
		Status: batchv1beta1.CronJobStatus{},
	}
	type args struct {
		ctx context.Context
		kc  *kube.KubernetesClient
		job batchv1.Job
	}
	tests := []struct {
		name    string
		args    args
		want    *batchv1beta1.CronJob
		wantErr bool
	}{
		{
			name: "Simple",
			args: args{
				ctx: context.TODO(),
				kc: &kube.KubernetesClient{
					Client: fake.NewSimpleClientset(&job1, &cronjob1),
				},
				job: job1,
			},
			want: &cronjob1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetCronJobFromJob(tt.args.ctx, tt.args.kc, tt.args.job)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCronJobFromJob() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCronJobFromJob() = %v, want %v", got, tt.want)
			}
		})
	}
}
