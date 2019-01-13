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

// GetSystemEnvironment collects system environment related metrics,
// e.g. fans, power supplies, sensors, etc.
func (n *NetworkNode) GetSystemEnvironment(cli *api.Client) {
	envt, err := cli.GetSystemEnvironment()
	if err != nil {
		log.Debugf("%s: GetSystemEnvironment() failed (host: %s, target: %s): %s", n.UUID, n.Name, n.target, err)
		n.IncrementErrorCounter()
		return
	}
	for _, fan := range envt.Fans {
		var fanStatus float64
		if fan.Status == "Ok" || fan.Status == "OK" {
			fanStatus = 1
		}
		n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
			fanUp,
			prometheus.GaugeValue,
			fanStatus,
			n.UUID,
			fan.Name,
		))
	}
	for _, ps := range envt.PowerSupplies {
		psName := fmt.Sprintf("%s %d", ps.Model, ps.ID)
		var psStatus float64
		if ps.Status == "Ok" || ps.Status == "OK" {
			psStatus = 1
		}
		n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
			fanUp,
			prometheus.GaugeValue,
			psStatus,
			n.UUID,
			psName,
		))
		n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
			powerSupplyPowerInput,
			prometheus.GaugeValue,
			ps.PowerInput,
			n.UUID,
			psName,
		))
		n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
			powerSupplyPowerOutput,
			prometheus.GaugeValue,
			ps.PowerOutput,
			n.UUID,
			psName,
		))
		n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
			powerSupplyPowerCapacity,
			prometheus.GaugeValue,
			ps.PowerCapacity,
			n.UUID,
			psName,
		))
	}
	for _, sensor := range envt.Sensors {
		sensorName := fmt.Sprintf("%s %d", sensor.Name, sensor.Module)
		var sensorStatus float64
		if sensor.Status == "Ok" || sensor.Status == "OK" {
			sensorStatus = 1
		}
		n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
			sensorUp,
			prometheus.GaugeValue,
			sensorStatus,
			n.UUID,
			sensorName,
		))
		n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
			sensorTemperature,
			prometheus.GaugeValue,
			sensor.Temperature,
			n.UUID,
			sensorName,
		))
		n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
			sensorTemperatureThresholdHigh,
			prometheus.GaugeValue,
			sensor.ThresholdHigh,
			n.UUID,
			sensorName,
		))
		n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
			sensorTemperatureThresholdLow,
			prometheus.GaugeValue,
			sensor.ThresholdLow,
			n.UUID,
			sensorName,
		))
	}
}
