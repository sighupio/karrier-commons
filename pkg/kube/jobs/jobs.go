// Copyright (c) 2021 SIGHUP s.r.l All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package jobs

import (
	"context"
	"errors"

	"github.com/sighupio/fip-commons/pkg/kube"
	batchv1 "k8s.io/api/batch/v1"
	batchv1beta1 "k8s.io/api/batch/v1beta1"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
)

// GetJob returns the job by namespace and name.
func GetJob(ctx context.Context, kc *kube.KubernetesClient, namespace string, name string) (*batchv1.Job, error) {
	job, err := kc.Client.BatchV1().Jobs(namespace).Get(ctx, name, v1.GetOptions{})
	if err != nil {
		return nil, err
	}

	return job, nil
}

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

// GetCronJobFromJob returns the parent cronjob of the job.
func GetCronJobFromJob(ctx context.Context, kc *kube.KubernetesClient, job batchv1.Job) (*batchv1beta1.CronJob, error) {
	if len(job.OwnerReferences) == 0 {
		return nil, errors.New("job does not belongs to a cronjob")
	}

	var cronJobName string

	for _, or := range job.OwnerReferences {
		if or.Kind == "CronJob" {
			cronJobName = or.Name
		}
	}

	cj, err := kc.Client.BatchV1beta1().CronJobs(job.Namespace).Get(ctx, cronJobName, v1.GetOptions{})
	if err != nil {
		return nil, err
	}

	return cj, nil
}
