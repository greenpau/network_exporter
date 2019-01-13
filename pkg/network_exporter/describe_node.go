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
	nodeUp = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "node", "up"),
		"Is node up and responding to queries (1) or is it down (0).",
		[]string{
			"node",
		}, nil,
	)
	nodeHostname = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "node", "name"),
		"The Ansible inventory name for the device. The value is always set to 1.",
		[]string{
			"node",
			"name",
		}, nil,
	)
	nodeErrors = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "node", "failed_req_count"),
		"The number of failed requests for a network node.",
		[]string{"node"}, nil,
	)
	nodeNextScrape = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "node", "next_poll"),
		"The timestamp of the next potential scrape of the node.",
		[]string{"node"}, nil,
	)
	nodeScrapeTime = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "node", "scrape_time"),
		"The amount of time it took to scrape the node.",
		[]string{"node"}, nil,
	)
)
