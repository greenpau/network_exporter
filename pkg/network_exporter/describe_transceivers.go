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
	transceiverUp = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "interface", "transceiver"),
		"The serial number and vendor of a transceiver attached to an interface are the labels of this metric. The value of the metric is always set to 1.",
		[]string{
			"node",
			"iface_name",
			"serial",
			"vendor",
		}, nil,
	)
	transceiverLaneTemperature = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "interface", "transceiver_lane_temperature"),
		"The temperature of a transceiver lane.",
		[]string{
			"node",
			"iface_name",
			"lane_id",
		}, nil,
	)
	transceiverLaneVoltage = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "interface", "transceiver_lane_voltage"),
		"The voltage of a transceiver lane.",
		[]string{
			"node",
			"iface_name",
			"lane_id",
		}, nil,
	)
	transceiverLaneCurrent = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "interface", "transceiver_lane_current"),
		"The current of a transceiver lane.",
		[]string{
			"node",
			"iface_name",
			"lane_id",
		}, nil,
	)
	transceiverLaneTxPower = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "interface", "transceiver_lane_tx_power"),
		"The transmit power of a transceiver lane.",
		[]string{
			"node",
			"iface_name",
			"lane_id",
		}, nil,
	)
	transceiverLaneRxPower = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "interface", "transceiver_lane_rx_power"),
		"The receive power of a transceiver lane.",
		[]string{
			"node",
			"iface_name",
			"lane_id",
		}, nil,
	)
	transceiverLaneErrors = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "interface", "transceiver_lane_errors"),
		"The number of errors with a transceiver lane.",
		[]string{
			"node",
			"iface_name",
			"lane_id",
		}, nil,
	)
)
