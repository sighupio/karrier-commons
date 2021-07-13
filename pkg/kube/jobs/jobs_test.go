package jobs

import (
	"context"
	"reflect"
	"testing"

	"github.com/sighupio/fip-commons/pkg/kube"
	batchv1 "k8s.io/api/batch/v1"
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
			want:    []corev1.Pod{},
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
			want:    []corev1.Pod{},
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
