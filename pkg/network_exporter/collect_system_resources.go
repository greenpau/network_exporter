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
	"fmt"
	api "github.com/greenpau/go-cisco-nx-api/pkg/client"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"
)

// GetSystemResources collects system resource usage metrics.
// That includes data about CPU, memory, and processes.
func (n *NetworkNode) GetSystemResources(cli *api.Client) {
	rsc, err := cli.GetSystemResources()
	if err != nil {
		log.Debugf("%s: GetSystemResources() failed (host: %s, target: %s): %s", n.UUID, n.Name, n.target, err)
		n.IncrementErrorCounter()
		return
	}
	n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
		processUsageRunning,
		prometheus.GaugeValue,
		float64(rsc.Processes.Running),
		n.UUID,
	))
	n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
		processUsageTotal,
		prometheus.GaugeValue,
		float64(rsc.Processes.Total),
		n.UUID,
	))
	n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
		memoryUsageTotal,
		prometheus.GaugeValue,
		float64(rsc.Memory.Total),
		n.UUID,
	))
	n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
		memoryUsageFree,
		prometheus.GaugeValue,
		float64(rsc.Memory.Free),
		n.UUID,
	))
	n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
		memoryUsageUsed,
		prometheus.GaugeValue,
		float64(rsc.Memory.Used),
		n.UUID,
	))
	n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
		cpuUsageTotalIdle,
		prometheus.GaugeValue,
		rsc.CPU.Idle,
		n.UUID,
	))
	n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
		cpuUsageTotalKernel,
		prometheus.GaugeValue,
		rsc.CPU.Kernel,
		n.UUID,
	))
	n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
		cpuUsageTotalUser,
		prometheus.GaugeValue,
		rsc.CPU.User,
		n.UUID,
	))
	for _, c := range rsc.CPUs {
		n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
			cpuUsagePerCPUIdle,
			prometheus.GaugeValue,
			c.Usage.Idle,
			n.UUID,
			fmt.Sprintf("%d", c.ID),
		))
		n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
			cpuUsagePerCPUKernel,
			prometheus.GaugeValue,
			c.Usage.Kernel,
			n.UUID,
			fmt.Sprintf("%d", c.ID),
		))
		n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
			cpuUsagePerCPUUser,
			prometheus.GaugeValue,
			c.Usage.User,
			n.UUID,
			fmt.Sprintf("%d", c.ID),
		))
	}
}
