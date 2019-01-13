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
	//"github.com/davecgh/go-spew/spew"
	api "github.com/greenpau/go-cisco-nx-api/pkg/client"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"
	"sync"
	"time"
)

// GatherMetrics collect data from a network node and stores them
// as Prometheus metrics.
func (n *NetworkNode) GatherMetrics() {
	n.Lock()
	defer n.Unlock()
	log.Debugf("%s: GatherMetrics() locked for %s", n.UUID, n.Name)
	if time.Now().Unix() < n.nextCollectionTicker {
		return
	}
	start := time.Now()
	if len(n.metrics) > 0 {
		n.metrics = n.metrics[:0]
		log.Debugf("%s: GatherMetrics() cleared metrics", n.UUID)
	}
	upValue := 1

	cli := api.NewClient()
	cli.SetHost(n.target)
	if n.port != 0 {
		cli.SetPort(n.port)
	}
	if n.proto != "" {
		cli.SetProtocol(n.proto)
	}

	var info *api.SysInfo
	var workingCredential *credential
	// test all available credentials
	tryFailed := false
	failedCredentials := make(map[int]bool)
	for {
		for i, c := range n.credentials {
			if c.Failed && !tryFailed {
				continue
			}
			if _, exists := failedCredentials[i]; exists {
				continue
			}
			cli.SetUsername(c.Username)
			cli.SetPassword(c.Password)
			data, err := cli.GetSystemInfo()
			if err != nil {
				failedCredentials[i] = true
				log.Debugf("%s: GetSystemInfo() failed (host: %s, target: %s, username: %s): %s", n.UUID, n.Name, n.target, c.Username, err)
				c.Failed = true
				continue
			}
			c.Failed = false
			info = data
			workingCredential = c
			break
		}
		if workingCredential == nil {
			if tryFailed {
				break
			}
			tryFailed = true
			continue
		}
		break
	}

	if workingCredential == nil || info == nil {
		n.IncrementErrorCounter()
		upValue = 0
	} else {
		log.Debugf("%s: hostname: %s, chassis id: %s", n.UUID, info.Hostname, info.ChassisID)
		// General Metrics
		n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
			nodeSystemHostname,
			prometheus.GaugeValue,
			1,
			n.UUID,
			info.Hostname,
		))
		n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
			nodeSystemIdentifier,
			prometheus.GaugeValue,
			1,
			n.UUID,
			info.ProcessorBoardID,
		))

		var wg sync.WaitGroup
		wg.Add(6)

		go func() {
			defer wg.Done()
			n.GetInterfaces(cli)
		}()

		go func() {
			defer wg.Done()
			n.GetVlans(cli)
		}()

		go func() {
			defer wg.Done()
			n.GetSystemEnvironment(cli)
		}()

		go func() {
			defer wg.Done()
			n.GetSystemResources(cli)
		}()

		go func() {
			defer wg.Done()
			n.GetTransceivers(cli)
		}()

		go func() {
			defer wg.Done()
			n.GetRoutingBgp(cli)
		}()

		wg.Wait()
	}

	// Generic Metrics
	n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
		nodeUp,
		prometheus.GaugeValue,
		float64(upValue),
		n.UUID,
	))
	n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
		nodeHostname,
		prometheus.GaugeValue,
		1,
		n.UUID,
		n.Name,
	))
	n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
		nodeErrors,
		prometheus.CounterValue,
		float64(n.errors),
		n.UUID,
	))
	n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
		nodeNextScrape,
		prometheus.CounterValue,
		float64(n.nextCollectionTicker),
		n.UUID,
	))
	n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
		nodeScrapeTime,
		prometheus.GaugeValue,
		time.Since(start).Seconds(),
		n.UUID,
	))

	n.nextCollectionTicker = time.Now().Add(time.Duration(n.pollInterval) * time.Second).Unix()

	if upValue > 0 {
		n.result = "success"
	} else {
		n.result = "failure"
	}
	n.timestamp = time.Now().Format(time.RFC3339)

	log.Debugf("%s: GatherMetrics() returns", n.UUID)
	return
}
