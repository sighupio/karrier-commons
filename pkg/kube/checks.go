// Copyright (c) 2021 SIGHUP s.r.l All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package kube

func (kc *KubernetesClient) Healthz() error {
	path := "/healthz"
	content, err := kc.Client.Discovery().RESTClient().Get().AbsPath(path).DoRaw(*kc.ctx)

	if err != nil {
		return err
	}

	contentStr := string(content)
	if contentStr != "ok" {
		return err
	}

	return nil
}
