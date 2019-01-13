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

// GetTransceivers collects interface fiber transceiver related metrics.
func (n *NetworkNode) GetTransceivers(cli *api.Client) {
	trs, err := cli.GetTransceivers()
	if err != nil {
		log.Debugf("%s: GetTransceivers() failed (host: %s, target: %s): %s", n.UUID, n.Name, n.target, err)
		n.IncrementErrorCounter()
		return
	}
	for _, t := range trs {
		n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
			transceiverUp,
			prometheus.GaugeValue,
			1,
			n.UUID,
			t.Interface,
			t.SerialNumber,
			t.Name,
		))
		for _, lane := range t.Lanes {
			laneID := fmt.Sprintf("%d", lane.ID)
			n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
				transceiverLaneTemperature,
				prometheus.GaugeValue,
				lane.Temperature,
				n.UUID,
				t.Interface,
				laneID,
			))
			n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
				transceiverLaneVoltage,
				prometheus.GaugeValue,
				lane.Voltage,
				n.UUID,
				t.Interface,
				laneID,
			))
			n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
				transceiverLaneCurrent,
				prometheus.GaugeValue,
				lane.Current,
				n.UUID,
				t.Interface,
				laneID,
			))
			n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
				transceiverLaneTxPower,
				prometheus.GaugeValue,
				lane.TxPower,
				n.UUID,
				t.Interface,
				laneID,
			))
			n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
				transceiverLaneRxPower,
				prometheus.GaugeValue,
				lane.RxPower,
				n.UUID,
				t.Interface,
				laneID,
			))
			n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
				transceiverLaneErrors,
				prometheus.CounterValue,
				lane.Errors,
				n.UUID,
				t.Interface,
				laneID,
			))
		}
	}
}
