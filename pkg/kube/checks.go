// Copyright (c) 2021 SIGHUP s.r.l All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package kube

func (kc *KubernetesClient) Healthz() error {
	path := "/healthz"
	content, err := kc.Client.Discovery().RESTClient().Get().AbsPath(path).DoRaw(kc.ctx)

	if err != nil {
		kc.Log.Println("can not query the api server")
		kc.Log.Println(err)

		return err
	}

	contentStr := string(content)
	if contentStr != "ok" {
		kc.Log.Printf("api server response not ok %v\n", contentStr)
		kc.Log.Println(err)

		return err
	}

	kc.Log.Println("kube client healthy")

	return nil
}
