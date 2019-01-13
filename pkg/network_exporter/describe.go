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

// Describe describes all the metrics ever exported by the exporter. It
// implements prometheus.Collector.
func (n *NetworkNode) Describe(ch chan<- *prometheus.Desc) {
	ch <- nodeUp
	ch <- nodeHostname
	ch <- nodeErrors
	ch <- nodeNextScrape
	ch <- nodeScrapeTime
	ch <- nodeSystemHostname
	ch <- nodeSystemIdentifier
	ch <- ifaceName
	ch <- ifaceLocalIndex
	ch <- ifaceDescription
	ch <- ifaceMetricBandwidth
	ch <- ifaceMetricDelay
	ch <- ifaceMetricReliability
	ch <- ifaceMetricRxload
	ch <- ifaceMetricTxload
	ch <- ifaceCounterBabbles
	ch <- ifaceCounterBadEtherTypeDrops
	ch <- ifaceCounterBadProtocolDrops
	ch <- ifaceCounterNoCarrier
	ch <- ifaceCounterDribble
	ch <- ifaceCounterInputFrameErrors
	ch <- ifaceCounterInputDiscards
	ch <- ifaceCounterInputErrors
	ch <- ifaceCounterInputPause
	ch <- ifaceCounterInputOverruns
	ch <- ifaceCounterInputIfaceDownDrops
	ch <- ifaceCounterInputBytes
	ch <- ifaceCounterInputUnicastBytes
	ch <- ifaceCounterInputPackets
	ch <- ifaceCounterInputUnicastPackets
	ch <- ifaceCounterInputBroadcastPackets
	ch <- ifaceCounterInputMulticastPackets
	ch <- ifaceCounterInputJumboPackets
	ch <- ifaceCounterInputCompressed
	ch <- ifaceCounterInputFifo
	ch <- ifaceCounterLateCollisions
	ch <- ifaceCounterLostCarrier
	ch <- ifaceCounterOutputDiscards
	ch <- ifaceCounterOutputErrors
	ch <- ifaceCounterOutputPause
	ch <- ifaceCounterOutputUnderruns
	ch <- ifaceCounterOutputBytes
	ch <- ifaceCounterOutputUnicastBytes
	ch <- ifaceCounterOutputPackets
	ch <- ifaceCounterOutputUnicastPackets
	ch <- ifaceCounterOutputBroadcastPackets
	ch <- ifaceCounterOutputMulticastPackets
	ch <- ifaceCounterOutputJumboPackets
	ch <- ifaceCounterOutputCarrierErrors
	ch <- ifaceCounterCollisions
	ch <- ifaceCounterOutputFifo
	ch <- ifaceCounterWatchdog
	ch <- ifaceCounterStormSuppression
	ch <- ifaceCounterIgnored
	ch <- ifaceCounterRunts
	ch <- ifaceCounterCrcErrors
	ch <- ifaceCounterDeferred
	ch <- ifaceCounterNoBufferReceivedErrors
	ch <- ifaceCounterResets
	ch <- ifacePropBeaconEnabled
	ch <- ifacePropAutoNegotiationEnabled
	ch <- ifacePropMdixEnabled
	ch <- ifacePropMTU
	ch <- ifacePropSpeed
	ch <- ifacePropDuplex
	ch <- ifacePropEncapsulatedVlan
	ch <- ifacePropState
	ch <- ifacePropAdminState
	ch <- ifaceIsSubinterface
	ch <- ifaceIsRoutedMode
	ch <- ifaceIsAccessMode
	ch <- ifacePropIPAddress
	ch <- ifacePropHWAddress
	ch <- vlanID
	ch <- vlanName
	ch <- vlanState
	ch <- vlanShutdownState

	ch <- fanUp
	ch <- powerSupplyUp
	ch <- powerSupplyPowerInput
	ch <- powerSupplyPowerOutput
	ch <- powerSupplyPowerCapacity
	ch <- sensorUp
	ch <- sensorTemperature
	ch <- sensorTemperatureThresholdHigh
	ch <- sensorTemperatureThresholdLow

	ch <- processUsageRunning
	ch <- processUsageTotal
	ch <- memoryUsageTotal
	ch <- memoryUsageFree
	ch <- memoryUsageUsed
	ch <- cpuUsageTotalIdle
	ch <- cpuUsageTotalKernel
	ch <- cpuUsageTotalUser
	ch <- cpuUsagePerCPUIdle
	ch <- cpuUsagePerCPUKernel
	ch <- cpuUsagePerCPUUser

	ch <- transceiverUp
	ch <- transceiverLaneTemperature
	ch <- transceiverLaneVoltage
	ch <- transceiverLaneCurrent
	ch <- transceiverLaneTxPower
	ch <- transceiverLaneRxPower
	ch <- transceiverLaneErrors
}
