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
	"github.com/prometheus/common/log"
	"sync"
	"sync/atomic"
	"time"
)

type credential struct {
	Username string
	Password string
	Failed   bool
}

// NetworkNode is an instance of a managed network node, e.g. a router or switch.
type NetworkNode struct {
	sync.RWMutex
	Name                 string
	UUID                 string
	Variables            map[string]string
	Interfaces           map[string]string
	Vlans                map[string]string
	target               string
	port                 int
	proto                string
	credentials          []*credential
	result               string
	module               string
	timestamp            string
	pollInterval         int64
	timeout              int
	errors               int64
	errorsLocker         sync.RWMutex
	nextCollectionTicker int64
	metrics              []prometheus.Metric
	subsystems           []string
}

// IncrementErrorCounter increases the counter of failed queries
// to a network node.
func (n *NetworkNode) IncrementErrorCounter() {
	n.errorsLocker.Lock()
	defer n.errorsLocker.Unlock()
	atomic.AddInt64(&n.errors, 1)
}

// Collect implements prometheus.Collector.
func (n *NetworkNode) Collect(ch chan<- prometheus.Metric) {
	start := time.Now()
	log.Debugf("%s: subsystems: %s", n.UUID, n.subsystems)
	n.GatherMetrics()
	log.Debugf("%s: Collect() calls RLock()", n.UUID)
	n.RLock()
	defer n.RUnlock()
	log.Debugf("%s: Collect() successful RLock()", n.UUID)
	if len(n.metrics) == 0 {
		log.Debugf("%s: Collect() no metrics found", n.UUID)
		ch <- prometheus.MustNewConstMetric(
			nodeUp,
			prometheus.GaugeValue,
			0,
			n.UUID,
		)
		ch <- prometheus.MustNewConstMetric(
			nodeHostname,
			prometheus.GaugeValue,
			1,
			n.UUID,
			n.Name,
		)
		ch <- prometheus.MustNewConstMetric(
			nodeErrors,
			prometheus.CounterValue,
			float64(n.errors),
			n.UUID,
		)
		ch <- prometheus.MustNewConstMetric(
			nodeNextScrape,
			prometheus.CounterValue,
			float64(n.nextCollectionTicker),
			n.UUID,
		)

		ch <- prometheus.MustNewConstMetric(
			nodeScrapeTime,
			prometheus.GaugeValue,
			time.Since(start).Seconds(),
			n.UUID,
		)
		return
	}
	log.Debugf("%s: Collect() sends %d metrics to a shared channel", n.UUID, len(n.metrics))
	for _, m := range n.metrics {
		ch <- m
	}
}
