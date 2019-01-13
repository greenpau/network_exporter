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
	fanUp = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "node", "fan_up"),
		"The status of a fan. 1 (up, Ok), 0 (down)",
		[]string{
			"node",
			"fan",
		}, nil,
	)
	powerSupplyUp = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "node", "ps_up"),
		"The status of a power supply. 1 (up, Ok), 0 (down)",
		[]string{
			"node",
			"power_supply",
		}, nil,
	)
	powerSupplyPowerInput = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "node", "ps_pwr_input"),
		"The power input of a power supply.",
		[]string{
			"node",
			"power_supply",
		}, nil,
	)
	powerSupplyPowerOutput = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "node", "ps_pwr_output"),
		"The power output of a power supply.",
		[]string{
			"node",
			"power_supply",
		}, nil,
	)
	powerSupplyPowerCapacity = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "node", "ps_pwr_capacity"),
		"The power capacity of a power supply.",
		[]string{
			"node",
			"power_supply",
		}, nil,
	)

	sensorUp = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "node", "sensor_up"),
		"The status of a sensor. 1 (up, Ok), 0 (down)",
		[]string{
			"node",
			"sensor",
		}, nil,
	)
	sensorTemperature = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "node", "sensor_temperature"),
		"The temperature of a sensor.",
		[]string{
			"node",
			"sensor",
		}, nil,
	)
	sensorTemperatureThresholdHigh = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "node", "sensor_temperature_threshold_high"),
		"The alarm upper threshold for the temperature of a sensor.",
		[]string{
			"node",
			"sensor",
		}, nil,
	)
	sensorTemperatureThresholdLow = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "node", "sensor_temperature_threshold_low"),
		"The alarm lower threshold for the temperature of a sensor.",
		[]string{
			"node",
			"sensor",
		}, nil,
	)
)
