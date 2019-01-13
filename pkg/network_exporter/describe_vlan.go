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
	// vlan metrics
	vlanID = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "vlan", "id"),
		"The Vlan ID of a VLAN. The value is always set to 1.",
		[]string{
			"node",
			"vlan",
		}, nil,
	)
	vlanName = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "vlan", "name"),
		"The name of a VLAN. The value is always set to 1.",
		[]string{
			"node",
			"vlan",
			"name",
		}, nil,
	)
	vlanState = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "vlan", "state"),
		"The state of a VLAN. Values are active (1), any other value (0).",
		[]string{
			"node",
			"vlan",
		}, nil,
	)
	vlanShutdownState = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "vlan", "shutdown_state"),
		"The shutdown state of a VLAN. Values are noshutdown (1), any other value (0).",
		[]string{
			"node",
			"vlan",
		}, nil,
	)
)
