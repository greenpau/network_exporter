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
	processUsageRunning = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "node", "running_process_count"),
		"The number of running processes.",
		[]string{
			"node",
		}, nil,
	)
	processUsageTotal = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "node", "total_process_count"),
		"The number of total processes.",
		[]string{
			"node",
		}, nil,
	)
	memoryUsageTotal = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "node", "memory_total"),
		"The amount of total memory available.",
		[]string{
			"node",
		}, nil,
	)
	memoryUsageFree = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "node", "memory_free"),
		"The amount of free memory available.",
		[]string{
			"node",
		}, nil,
	)
	memoryUsageUsed = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "node", "memory_used"),
		"The amount of memory used.",
		[]string{
			"node",
		}, nil,
	)
	cpuUsageTotalIdle = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "node", "total_cpu_idle"),
		"The amount of CPU time in idle state.",
		[]string{
			"node",
		}, nil,
	)
	cpuUsageTotalKernel = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "node", "total_cpu_kernel"),
		"The amount of CPU time in kernel state.",
		[]string{
			"node",
		}, nil,
	)
	cpuUsageTotalUser = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "node", "total_cpu_user"),
		"The amount of CPU time in user state.",
		[]string{
			"node",
		}, nil,
	)
	cpuUsagePerCPUIdle = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "node", "cpu_idle"),
		"The amount of CPU time in idle state on per CPU basis.",
		[]string{
			"node",
			"cpu_id",
		}, nil,
	)
	cpuUsagePerCPUKernel = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "node", "cpu_kernel"),
		"The amount of CPU time in kernel state on per CPU basis.",
		[]string{
			"node",
			"cpu_id",
		}, nil,
	)
	cpuUsagePerCPUUser = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "node", "cpu_user"),
		"The amount of CPU time in user state on per CPU basis.",
		[]string{
			"node",
			"cpu_id",
		}, nil,
	)
)
