package jobs

import (
	"context"

	"github.com/sighupio/fip-commons/pkg/kube"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
)

// GetPodsfromJob returns a list of Pods from a specific job.
func GetPodsfromJob(ctx context.Context, kc *kube.KubernetesClient, job batchv1.Job) ([]corev1.Pod, error) {
	labelSelector := v1.LabelSelector{MatchLabels: map[string]string{"job-name": job.Name}}
	podList, err := kc.Client.CoreV1().Pods(job.Namespace).List(ctx, v1.ListOptions{
		LabelSelector: labels.Set(labelSelector.MatchLabels).String(),
		Limit:         int64(*job.Spec.BackoffLimit),
	})

	if err != nil {
		return nil, err
	}

	if len(podList.Items) > 0 {
		return podList.Items, nil
	}

	return make([]corev1.Pod, 0), nil
}
