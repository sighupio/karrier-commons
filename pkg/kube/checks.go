// Copyright (c) 2021 SIGHUP s.r.l All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package kube

import (
	"context"
	"errors"
	"fmt"
)

var ErrHealtzResponse = errors.New("expected healthz response's content to be 'ok'")

func (kc *KubernetesClient) Healthz(ctx *context.Context) error {
	content, err := kc.Client.Discovery().RESTClient().Get().AbsPath("/healthz").DoRaw(*ctx)
	if err != nil {
		return err
	}

	if string(content) != "ok" {
		return fmt.Errorf("%w, got '%s'", ErrHealtzResponse, string(content))
	}

	return nil
}
