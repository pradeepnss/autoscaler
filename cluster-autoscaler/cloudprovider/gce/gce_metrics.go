/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package gce

import (
	k8smetrics "k8s.io/component-base/metrics"
	"k8s.io/component-base/metrics/legacyregistry"
)

const (
	caNamespace = "cluster_autoscaler"
)

var (
	/**** Metrics related to GCE API usage ****/
	requestCounter = k8smetrics.NewCounterVec(
		&k8smetrics.CounterOpts{
			Namespace: caNamespace,
			Name:      "gce_request_count",
			Help:      "Counter of GCE API requests for each verb and API resource.",
		}, []string{"resource", "verb"},
	)
)

// RegisterMetrics registers all GCE metrics.
func RegisterMetrics() {
	legacyregistry.MustRegister(requestCounter)
}

// registerRequest registers request to GCE API.
func registerRequest(resource string, verb string) {
	requestCounter.WithLabelValues(resource, verb).Add(1.0)
}
