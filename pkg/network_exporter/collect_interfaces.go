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
	"strings"
)

// GetInterfaces collects interface related metrics.
func (n *NetworkNode) GetInterfaces(cli *api.Client) {
	ifaces, err := cli.GetInterfaces()
	if err != nil {
		log.Debugf("%s: GetInterfaces() failed (host: %s, target: %s): %s", n.UUID, n.Name, n.target, err)
		n.IncrementErrorCounter()
		return
	}
	// Interface metrics
	for _, iface := range ifaces {
		var _uuid string
		// Interface UUID
		if v, exists := n.Interfaces[iface.Name]; !exists {
			hash := sha1.New()
			hash.Write([]byte(n.UUID))
			hash.Write([]byte(iface.Name))
			_uuid = fmt.Sprintf("%x", hash.Sum(nil))
			n.Interfaces[iface.Name] = _uuid
		} else {
			_uuid = v
		}
		/*
		   log.Debugf(
		       "%s: host: %s, interface: %s, Mode: %s, Speed: %s, Duplex: %s",
		       n.UUID, n.Name, iface.Name, iface.Props.Mode, iface.Props.Speed, iface.Props.Duplex,
		   )
		*/

		// Metrics
		n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
			ifaceName,
			prometheus.GaugeValue,
			1,
			n.UUID,
			_uuid,
			iface.Name,
		))
		n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
			ifaceLocalIndex,
			prometheus.GaugeValue,
			float64(iface.LocalIndex),
			n.UUID,
			_uuid,
		))
		var _ifaceDescription string
		if iface.Description == "" {
			_ifaceDescription = "empty"
		} else {
			_ifaceDescription = iface.Description
		}
		n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
			ifaceDescription,
			prometheus.GaugeValue,
			1,
			n.UUID,
			_uuid,
			_ifaceDescription,
		))
		// Various Metrics
		n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
			ifaceMetricBandwidth,
			prometheus.GaugeValue,
			float64(iface.Metrics.Bandwidth),
			n.UUID,
			_uuid,
		))
		n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
			ifaceMetricDelay,
			prometheus.GaugeValue,
			float64(iface.Metrics.Delay),
			n.UUID,
			_uuid,
		))
		n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
			ifaceMetricReliability,
			prometheus.GaugeValue,
			float64(iface.Metrics.Reliability),
			n.UUID,
			_uuid,
		))
		n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
			ifaceMetricRxload,
			prometheus.GaugeValue,
			float64(iface.Metrics.Rxload),
			n.UUID,
			_uuid,
		))
		n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
			ifaceMetricTxload,
			prometheus.GaugeValue,
			float64(iface.Metrics.Txload),
			n.UUID,
			_uuid,
		))
		// Various Counters
		n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
			ifaceCounterBabbles,
			prometheus.CounterValue,
			float64(iface.Counters.Babbles),
			n.UUID,
			_uuid,
		))
		n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
			ifaceCounterBadEtherTypeDrops,
			prometheus.CounterValue,
			float64(iface.Counters.BadEtherTypeDrops),
			n.UUID,
			_uuid,
		))
		n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
			ifaceCounterBadProtocolDrops,
			prometheus.CounterValue,
			float64(iface.Counters.BadProtocolDrops),
			n.UUID,
			_uuid,
		))
		n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
			ifaceCounterNoCarrier,
			prometheus.CounterValue,
			float64(iface.Counters.NoCarrier),
			n.UUID,
			_uuid,
		))
		n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
			ifaceCounterDribble,
			prometheus.CounterValue,
			float64(iface.Counters.Dribble),
			n.UUID,
			_uuid,
		))
		n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
			ifaceCounterInputFrameErrors,
			prometheus.CounterValue,
			float64(iface.Counters.InputFrameErrors),
			n.UUID,
			_uuid,
		))
		n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
			ifaceCounterInputDiscards,
			prometheus.CounterValue,
			float64(iface.Counters.InputDiscards),
			n.UUID,
			_uuid,
		))
		n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
			ifaceCounterInputErrors,
			prometheus.CounterValue,
			float64(iface.Counters.InputErrors),
			n.UUID,
			_uuid,
		))
		n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
			ifaceCounterInputPause,
			prometheus.CounterValue,
			float64(iface.Counters.InputPause),
			n.UUID,
			_uuid,
		))
		n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
			ifaceCounterInputOverruns,
			prometheus.CounterValue,
			float64(iface.Counters.InputOverruns),
			n.UUID,
			_uuid,
		))
		n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
			ifaceCounterInputIfaceDownDrops,
			prometheus.CounterValue,
			float64(iface.Counters.InputIfaceDownDrops),
			n.UUID,
			_uuid,
		))
		n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
			ifaceCounterInputBytes,
			prometheus.CounterValue,
			float64(iface.Counters.InputBytes),
			n.UUID,
			_uuid,
		))
		n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
			ifaceCounterInputUnicastBytes,
			prometheus.CounterValue,
			float64(iface.Counters.InputUnicastBytes),
			n.UUID,
			_uuid,
		))
		n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
			ifaceCounterInputPackets,
			prometheus.CounterValue,
			float64(iface.Counters.InputPackets),
			n.UUID,
			_uuid,
		))
		n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
			ifaceCounterInputUnicastPackets,
			prometheus.CounterValue,
			float64(iface.Counters.InputUnicastPackets),
			n.UUID,
			_uuid,
		))
		n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
			ifaceCounterInputBroadcastPackets,
			prometheus.CounterValue,
			float64(iface.Counters.InputBroadcastPackets),
			n.UUID,
			_uuid,
		))
		n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
			ifaceCounterInputMulticastPackets,
			prometheus.CounterValue,
			float64(iface.Counters.InputMulticastPackets),
			n.UUID,
			_uuid,
		))
		n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
			ifaceCounterInputJumboPackets,
			prometheus.CounterValue,
			float64(iface.Counters.InputJumboPackets),
			n.UUID,
			_uuid,
		))
		n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
			ifaceCounterInputCompressed,
			prometheus.CounterValue,
			float64(iface.Counters.InputCompressed),
			n.UUID,
			_uuid,
		))
		n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
			ifaceCounterInputFifo,
			prometheus.CounterValue,
			float64(iface.Counters.InputFifo),
			n.UUID,
			_uuid,
		))
		n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
			ifaceCounterLateCollisions,
			prometheus.CounterValue,
			float64(iface.Counters.LateCollisions),
			n.UUID,
			_uuid,
		))
		n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
			ifaceCounterLostCarrier,
			prometheus.CounterValue,
			float64(iface.Counters.LostCarrier),
			n.UUID,
			_uuid,
		))
		n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
			ifaceCounterOutputDiscards,
			prometheus.CounterValue,
			float64(iface.Counters.OutputDiscards),
			n.UUID,
			_uuid,
		))
		n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
			ifaceCounterOutputErrors,
			prometheus.CounterValue,
			float64(iface.Counters.OutputErrors),
			n.UUID,
			_uuid,
		))
		n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
			ifaceCounterOutputPause,
			prometheus.CounterValue,
			float64(iface.Counters.OutputPause),
			n.UUID,
			_uuid,
		))
		n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
			ifaceCounterOutputUnderruns,
			prometheus.CounterValue,
			float64(iface.Counters.OutputUnderruns),
			n.UUID,
			_uuid,
		))
		n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
			ifaceCounterOutputBytes,
			prometheus.CounterValue,
			float64(iface.Counters.OutputBytes),
			n.UUID,
			_uuid,
		))
		n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
			ifaceCounterOutputUnicastBytes,
			prometheus.CounterValue,
			float64(iface.Counters.OutputUnicastBytes),
			n.UUID,
			_uuid,
		))
		n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
			ifaceCounterOutputPackets,
			prometheus.CounterValue,
			float64(iface.Counters.OutputPackets),
			n.UUID,
			_uuid,
		))
		n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
			ifaceCounterOutputUnicastPackets,
			prometheus.CounterValue,
			float64(iface.Counters.OutputUnicastPackets),
			n.UUID,
			_uuid,
		))
		n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
			ifaceCounterOutputBroadcastPackets,
			prometheus.CounterValue,
			float64(iface.Counters.OutputBroadcastPackets),
			n.UUID,
			_uuid,
		))
		n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
			ifaceCounterOutputMulticastPackets,
			prometheus.CounterValue,
			float64(iface.Counters.OutputMulticastPackets),
			n.UUID,
			_uuid,
		))
		n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
			ifaceCounterOutputJumboPackets,
			prometheus.CounterValue,
			float64(iface.Counters.OutputJumboPackets),
			n.UUID,
			_uuid,
		))
		n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
			ifaceCounterOutputCarrierErrors,
			prometheus.CounterValue,
			float64(iface.Counters.OutputCarrierErrors),
			n.UUID,
			_uuid,
		))
		n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
			ifaceCounterCollisions,
			prometheus.CounterValue,
			float64(iface.Counters.Collisions),
			n.UUID,
			_uuid,
		))
		n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
			ifaceCounterOutputFifo,
			prometheus.CounterValue,
			float64(iface.Counters.OutputFifo),
			n.UUID,
			_uuid,
		))
		n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
			ifaceCounterWatchdog,
			prometheus.CounterValue,
			float64(iface.Counters.Watchdog),
			n.UUID,
			_uuid,
		))
		n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
			ifaceCounterStormSuppression,
			prometheus.CounterValue,
			float64(iface.Counters.StormSuppression),
			n.UUID,
			_uuid,
		))
		n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
			ifaceCounterIgnored,
			prometheus.CounterValue,
			float64(iface.Counters.Ignored),
			n.UUID,
			_uuid,
		))
		n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
			ifaceCounterRunts,
			prometheus.CounterValue,
			float64(iface.Counters.Runts),
			n.UUID,
			_uuid,
		))
		n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
			ifaceCounterCrcErrors,
			prometheus.CounterValue,
			float64(iface.Counters.CrcErrors),
			n.UUID,
			_uuid,
		))
		n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
			ifaceCounterDeferred,
			prometheus.CounterValue,
			float64(iface.Counters.Deferred),
			n.UUID,
			_uuid,
		))
		n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
			ifaceCounterNoBufferReceivedErrors,
			prometheus.CounterValue,
			float64(iface.Counters.NoBufferReceivedErrors),
			n.UUID,
			_uuid,
		))
		n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
			ifaceCounterResets,
			prometheus.CounterValue,
			float64(iface.Counters.Resets),
			n.UUID,
			_uuid,
		))
		// Various boolean props
		var _ifacePropsBeaconEnabled float64
		if iface.Props.BeaconEnabled {
			_ifacePropsBeaconEnabled = 1
		}
		n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
			ifacePropBeaconEnabled,
			prometheus.GaugeValue,
			_ifacePropsBeaconEnabled,
			n.UUID,
			_uuid,
		))
		var _ifacePropsAutoNegotiationEnabled float64
		if iface.Props.AutoNegotiationEnabled {
			_ifacePropsAutoNegotiationEnabled = 1
		}
		n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
			ifacePropAutoNegotiationEnabled,
			prometheus.GaugeValue,
			_ifacePropsAutoNegotiationEnabled,
			n.UUID,
			_uuid,
		))
		var _ifacePropsMdixEnabled float64
		if iface.Props.MdixEnabled {
			_ifacePropsMdixEnabled = 1
		}
		n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
			ifacePropMdixEnabled,
			prometheus.GaugeValue,
			_ifacePropsMdixEnabled,
			n.UUID,
			_uuid,
		))
		n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
			ifacePropMTU,
			prometheus.GaugeValue,
			float64(iface.Props.MTU),
			n.UUID,
			_uuid,
		))
		var _ifacePropsDuplex float64
		switch iface.Props.Duplex {
		case "auto":
			_ifacePropsDuplex = 3
		case "full":
			_ifacePropsDuplex = 2
		case "half":
			_ifacePropsDuplex = 1
		default:
			_ifacePropsDuplex = 0
		}
		n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
			ifacePropDuplex,
			prometheus.GaugeValue,
			_ifacePropsDuplex,
			n.UUID,
			_uuid,
		))
		var _ifacePropsSpeed float64
		switch iface.Props.Speed {
		case "auto-speed":
			_ifacePropsSpeed = 0
		default:
			// Multipliers: Gb/s
			var speedVal float64
			var speedMulti float64
			speedArr := strings.Split(iface.Props.Speed, " ")
			if len(speedArr) > 1 {
				if v, err := strconv.ParseFloat(speedArr[0], 64); err == nil {
					speedVal = v
				}
				switch speedArr[1] {
				case "Tb/s":
					speedMulti = 1000000
				case "Gb/s":
					speedMulti = 1000
				case "Mb/s":
					speedMulti = 1
				default:
					speedMulti = 1
				}
			}
			if speedVal > 0 && speedMulti > 0 {
				_ifacePropsSpeed = speedVal * speedMulti
			}
		}
		n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
			ifacePropSpeed,
			prometheus.GaugeValue,
			_ifacePropsSpeed,
			n.UUID,
			_uuid,
		))
		n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
			ifacePropEncapsulatedVlan,
			prometheus.GaugeValue,
			float64(iface.Props.EncapsulatedVlan),
			n.UUID,
			_uuid,
		))
		// interface state
		var _ifacePropsState float64
		if iface.Props.State == "up" {
			_ifacePropsState = 1
		}
		n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
			ifacePropState,
			prometheus.GaugeValue,
			_ifacePropsState,
			n.UUID,
			_uuid,
		))
		var _ifacePropsAdminState float64
		if iface.Props.AdminState == "up" {
			_ifacePropsAdminState = 1
		}
		n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
			ifacePropAdminState,
			prometheus.GaugeValue,
			_ifacePropsAdminState,
			n.UUID,
			_uuid,
		))
		var _ifaceIsSubinterface float64
		if iface.Props.ParentInterface != "" {
			_ifaceIsSubinterface = 1
		}
		n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
			ifaceIsSubinterface,
			prometheus.GaugeValue,
			_ifaceIsSubinterface,
			n.UUID,
			_uuid,
		))
		var _ifaceIsRoutedMode float64
		if iface.Props.IPAddress != "" {
			_ifaceIsRoutedMode = 1
		}
		n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
			ifaceIsRoutedMode,
			prometheus.GaugeValue,
			_ifaceIsRoutedMode,
			n.UUID,
			_uuid,
		))
		var _ifaceIsAccessMode float64
		if iface.Props.Mode == "access" {
			_ifaceIsAccessMode = 1
		}
		n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
			ifaceIsAccessMode,
			prometheus.GaugeValue,
			_ifaceIsAccessMode,
			n.UUID,
			_uuid,
		))
		if iface.Props.IPAddress != "" {
			n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
				ifacePropIPAddress,
				prometheus.GaugeValue,
				1,
				n.UUID,
				_uuid,
				fmt.Sprintf("%s/%d", iface.Props.IPAddress, iface.Props.IPMask),
			))
		}
		if iface.Props.HwAddr != "" {
			n.metrics = append(n.metrics, prometheus.MustNewConstMetric(
				ifacePropHWAddress,
				prometheus.GaugeValue,
				1,
				n.UUID,
				_uuid,
				iface.Props.HwAddr,
			))
		}
	}
}
