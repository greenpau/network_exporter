// Copyright 2018 Paul Greenberg (greenpau@outlook.com)
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package exporter

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	// system info metrics
	nodeSystemHostname = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "node", "hostname"),
		"The configured hostname on the device itself. The value is always set to 1.",
		[]string{
			"node",
			"hostname",
		}, nil,
	)
	nodeSystemIdentifier = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "node", "id"),
		"The unique identifier for the physical device, e.g. a serial number. The value is always set to 1.",
		[]string{
			"node",
			"id",
		}, nil,
	)
)
