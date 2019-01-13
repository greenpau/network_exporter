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
	"crypto/sha1"
	"fmt"
	api "github.com/greenpau/go-cisco-nx-api/pkg/client"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"
	"strconv"
)

// GetVlans collects VLAN related metrics.
func (n *NetworkNode) GetVlans(cli *api.Client) {
	vlans, err := cli.GetVlans()
	if err != nil {
		log.Debugf("%s: GetVlans() failed (host: %s, target: %s): %s", n.UUID, n.Name, n.target, err)
		n.IncrementErrorCounter()
		return
	}
	for _, vlan := range vlans {
		var _uuid string
		// VLAN UUID
		if v, exists := n.Vlans[vlan.ID]; !exists {
			hash := sha1.New()
			hash.Write([]byte(n.UUID))
			hash.Write([]byte(vlan.ID))
			_uuid = fmt.Sprintf("%x", hash.Sum(nil))
			n.Vlans[vlan.ID] = _uuid
		} else {
			_uuid = v
		}
		// Metrics
		if v, err := strconv.Atoi(vlan.ID); err == nil {
			n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
				vlanID,
				prometheus.GaugeValue,
				float64(v),
				n.UUID,
				_uuid,
			))
		}
		n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
			vlanName,
			prometheus.GaugeValue,
			1,
			n.UUID,
			_uuid,
			vlan.Name,
		))
		// vlan state: active (1), any other value (0)
		var _vlanState float64
		if vlan.State == "active" {
			_vlanState = 1
		}
		n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
			vlanState,
			prometheus.GaugeValue,
			_vlanState,
			n.UUID,
			_uuid,
		))
		// vlan shutdown state: noshutdown (1), any other value (0)
		var _vlanShutdownState float64
		if vlan.ShutdownState == "noshutdown" {
			_vlanShutdownState = 1
		}
		n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
			vlanShutdownState,
			prometheus.GaugeValue,
			_vlanShutdownState,
			n.UUID,
			_uuid,
		))
	}
}
