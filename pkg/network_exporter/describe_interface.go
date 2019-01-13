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
	// interface metrics
	ifaceName = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "iface", "name"),
		"The name of an interface. The value is always set to 1.",
		[]string{
			"node",
			"iface",
			"name",
		}, nil,
	)
	ifaceLocalIndex = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "iface", "local_index"),
		"The local index of an interface.",
		[]string{
			"node",
			"iface",
		}, nil,
	)
	ifaceDescription = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "iface", "descr"),
		"The description attached to an interface. The value is the checksum of the description",
		[]string{
			"node",
			"iface",
			"description",
		}, nil,
	)
	// routing metrics
	ifaceMetricBandwidth = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "iface", "bandwidth"),
		"The bandwith routing metric of an interface.",
		[]string{
			"node",
			"iface",
		}, nil,
	)
	ifaceMetricDelay = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "iface", "delay"),
		"The delay routing metric of an interface.",
		[]string{
			"node",
			"iface",
		}, nil,
	)
	ifaceMetricReliability = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "iface", "reliability"),
		"The reliability routing metric of an interface.",
		[]string{
			"node",
			"iface",
		}, nil,
	)
	ifaceMetricRxload = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "iface", "rx_load"),
		"The rx_load routing metric of an interface.",
		[]string{
			"node",
			"iface",
		}, nil,
	)
	ifaceMetricTxload = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "iface", "tx_load"),
		"The tx_load routing metric of an interface.",
		[]string{
			"node",
			"iface",
		}, nil,
	)
	// interface counters
	ifaceCounterBabbles = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "iface", "babbles"),
		"The babbles counter of an interface.",
		[]string{
			"node",
			"iface",
		}, nil,
	)
	ifaceCounterBadEtherTypeDrops = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "iface", "bad_ethtype_drops"),
		"The bad_ethtype_drops counter of an interface.",
		[]string{
			"node",
			"iface",
		}, nil,
	)
	ifaceCounterBadProtocolDrops = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "iface", "bad_proto_drops"),
		"The bad_proto_drops counter of an interface.",
		[]string{
			"node",
			"iface",
		}, nil,
	)
	ifaceCounterNoCarrier = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "iface", "no_carrier"),
		"The no_carrier counter of an interface.",
		[]string{
			"node",
			"iface",
		}, nil,
	)
	ifaceCounterDribble = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "iface", "dribble"),
		"The dribble counter of an interface.",
		[]string{
			"node",
			"iface",
		}, nil,
	)
	ifaceCounterInputFrameErrors = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "iface", "input_frame_errors"),
		"The input_frame_errors counter of an interface.",
		[]string{
			"node",
			"iface",
		}, nil,
	)
	ifaceCounterInputDiscards = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "iface", "input_discards"),
		"The input_frame_errors counter of an interface.",
		[]string{
			"node",
			"iface",
		}, nil,
	)
	ifaceCounterInputErrors = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "iface", "input_errors"),
		"The input_errors counter of an interface.",
		[]string{
			"node",
			"iface",
		}, nil,
	)
	ifaceCounterInputPause = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "iface", "input_pause"),
		"The input_pause counter of an interface.",
		[]string{
			"node",
			"iface",
		}, nil,
	)
	ifaceCounterInputOverruns = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "iface", "input_overruns"),
		"The input_overruns counter of an interface.",
		[]string{
			"node",
			"iface",
		}, nil,
	)
	ifaceCounterInputIfaceDownDrops = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "iface", "input_iface_down_drops"),
		"The input if-down drops counter of an interface.",
		[]string{
			"node",
			"iface",
		}, nil,
	)
	ifaceCounterInputBytes = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "iface", "input_bytes"),
		"The input_bytes counter of an interface.",
		[]string{
			"node",
			"iface",
		}, nil,
	)
	ifaceCounterInputUnicastBytes = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "iface", "input_ucast_bytes"),
		"The input_ucast_bytes counter of an interface.",
		[]string{
			"node",
			"iface",
		}, nil,
	)
	ifaceCounterInputPackets = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "iface", "input_packets"),
		"The input_packets counter of an interface.",
		[]string{
			"node",
			"iface",
		}, nil,
	)
	ifaceCounterInputUnicastPackets = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "iface", "input_ucast_packets"),
		"The input_ucast_packets counter of an interface.",
		[]string{
			"node",
			"iface",
		}, nil,
	)
	ifaceCounterInputBroadcastPackets = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "iface", "input_bcast_packets"),
		"The input_bcast_packets counter of an interface.",
		[]string{
			"node",
			"iface",
		}, nil,
	)
	ifaceCounterInputMulticastPackets = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "iface", "input_mcast_packets"),
		"The input_mcast_packets counter of an interface.",
		[]string{
			"node",
			"iface",
		}, nil,
	)
	ifaceCounterInputJumboPackets = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "iface", "input_jumbo_packets"),
		"The input_jumbo_packets counter of an interface.",
		[]string{
			"node",
			"iface",
		}, nil,
	)
	ifaceCounterInputCompressed = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "iface", "input_compressed"),
		"The input_compressed counter of an interface.",
		[]string{
			"node",
			"iface",
		}, nil,
	)
	ifaceCounterInputFifo = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "iface", "input_fifo"),
		"The input_fifo counter of an interface.",
		[]string{
			"node",
			"iface",
		}, nil,
	)
	ifaceCounterLateCollisions = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "iface", "late_collisions"),
		"The late_collisions counter of an interface.",
		[]string{
			"node",
			"iface",
		}, nil,
	)
	ifaceCounterLostCarrier = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "iface", "lost_carrier"),
		"The lost_carrier counter of an interface.",
		[]string{
			"node",
			"iface",
		}, nil,
	)
	ifaceCounterOutputDiscards = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "iface", "output_discards"),
		"The output_discards counter of an interface.",
		[]string{
			"node",
			"iface",
		}, nil,
	)
	ifaceCounterOutputErrors = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "iface", "output_errors"),
		"The output_errors counter of an interface.",
		[]string{
			"node",
			"iface",
		}, nil,
	)
	ifaceCounterOutputPause = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "iface", "output_pause"),
		"The output_pause counter of an interface.",
		[]string{
			"node",
			"iface",
		}, nil,
	)
	ifaceCounterOutputUnderruns = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "iface", "output_underruns"),
		"The output_underruns counter of an interface.",
		[]string{
			"node",
			"iface",
		}, nil,
	)
	ifaceCounterOutputBytes = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "iface", "output_bytes"),
		"The output_bytes counter of an interface.",
		[]string{
			"node",
			"iface",
		}, nil,
	)
	ifaceCounterOutputUnicastBytes = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "iface", "output_ucast_bytes"),
		"The output_ucast_bytes counter of an interface.",
		[]string{
			"node",
			"iface",
		}, nil,
	)
	ifaceCounterOutputPackets = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "iface", "output_packets"),
		"The output_packets counter of an interface.",
		[]string{
			"node",
			"iface",
		}, nil,
	)
	ifaceCounterOutputUnicastPackets = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "iface", "output_ucast_packets"),
		"The output_ucast_packets counter of an interface.",
		[]string{
			"node",
			"iface",
		}, nil,
	)
	ifaceCounterOutputBroadcastPackets = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "iface", "output_bcast_packets"),
		"The output_bcast_packets counter of an interface.",
		[]string{
			"node",
			"iface",
		}, nil,
	)
	ifaceCounterOutputMulticastPackets = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "iface", "output_mcast_packets"),
		"The output_mcast_packets counter of an interface.",
		[]string{
			"node",
			"iface",
		}, nil,
	)
	ifaceCounterOutputJumboPackets = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "iface", "output_jumbo_packets"),
		"The output_jumbo_packets counter of an interface.",
		[]string{
			"node",
			"iface",
		}, nil,
	)
	ifaceCounterOutputCarrierErrors = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "iface", "output_carrier_errors"),
		"The output_carrier_errors counter of an interface.",
		[]string{
			"node",
			"iface",
		}, nil,
	)
	ifaceCounterCollisions = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "iface", "collisions"),
		"The collisions counter of an interface.",
		[]string{
			"node",
			"iface",
		}, nil,
	)
	ifaceCounterOutputFifo = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "iface", "output_fifo"),
		"The output_fifo counter of an interface.",
		[]string{
			"node",
			"iface",
		}, nil,
	)
	ifaceCounterWatchdog = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "iface", "watchdog"),
		"The watchdog counter of an interface.",
		[]string{
			"node",
			"iface",
		}, nil,
	)
	ifaceCounterStormSuppression = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "iface", "storm_suppression"),
		"The storm_suppression counter of an interface.",
		[]string{
			"node",
			"iface",
		}, nil,
	)
	ifaceCounterIgnored = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "iface", "ignored"),
		"The ignored counter of an interface.",
		[]string{
			"node",
			"iface",
		}, nil,
	)
	ifaceCounterRunts = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "iface", "runts"),
		"The runts counter of an interface.",
		[]string{
			"node",
			"iface",
		}, nil,
	)
	ifaceCounterCrcErrors = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "iface", "crc_errors"),
		"The crc_errors counter of an interface.",
		[]string{
			"node",
			"iface",
		}, nil,
	)
	ifaceCounterDeferred = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "iface", "deferred"),
		"The deferred counter of an interface.",
		[]string{
			"node",
			"iface",
		}, nil,
	)
	ifaceCounterNoBufferReceivedErrors = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "iface", "no_buffer"),
		"The no_buffer counter of an interface.",
		[]string{
			"node",
			"iface",
		}, nil,
	)
	ifaceCounterResets = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "iface", "resets"),
		"The resets counter of an interface.",
		[]string{
			"node",
			"iface",
		}, nil,
	)

	ifacePropBeaconEnabled = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "iface", "beacon_enabled"),
		"Whether beacon is enabled (1) or disabled (0) on an interface.",
		[]string{
			"node",
			"iface",
		}, nil,
	)
	ifacePropAutoNegotiationEnabled = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "iface", "auto_negotiation_enabled"),
		"Whether auto negotiation is enabled (1) or disabled (0) on an interface.",
		[]string{
			"node",
			"iface",
		}, nil,
	)
	ifacePropMdixEnabled = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "iface", "mdix_enabled"),
		"Whether auto MDIX is enabled (1) or disabled (0) on an interface.",
		[]string{
			"node",
			"iface",
		}, nil,
	)
	ifacePropMTU = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "iface", "mtu"),
		"The MTU of an interface.",
		[]string{
			"node",
			"iface",
		}, nil,
	)
	ifacePropSpeed = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "iface", "speed"),
		"The speed (in Mb/s) of an interface. If the value is 0, then it is auto.",
		[]string{
			"node",
			"iface",
		}, nil,
	)
	ifacePropDuplex = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "iface", "duplex"),
		"The duplex of an interface. Values are auto (3), full (2), half (1), other (0)",
		[]string{
			"node",
			"iface",
		}, nil,
	)
	ifacePropEncapsulatedVlan = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "iface", "encapsulated_vlan"),
		"The encapsulated VLAN associated with an interface.",
		[]string{
			"node",
			"iface",
		}, nil,
	)
	ifacePropState = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "iface", "state"),
		"The state of an interface. Values are up (1), any other value (0).",
		[]string{
			"node",
			"iface",
		}, nil,
	)
	ifacePropAdminState = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "iface", "admin_state"),
		"The state of an interface. Values are up (1), any other value (0).",
		[]string{
			"node",
			"iface",
		}, nil,
	)
	ifaceIsSubinterface = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "iface", "subinterface"),
		"Indicates whether an interface is a sub-interface. Values are yes (1), no (0).",
		[]string{
			"node",
			"iface",
		}, nil,
	)
	ifaceIsRoutedMode = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "iface", "routed_mode"),
		"Indicates whether an interface is in routed (L3-configured) mode. Values are yes (1), no (0).",
		[]string{
			"node",
			"iface",
		}, nil,
	)
	ifaceIsAccessMode = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "iface", "access_mode"),
		"Indicates whether an interface is in access (L2-configured) mode. Values are yes (1), no (0).",
		[]string{
			"node",
			"iface",
		}, nil,
	)
	ifacePropIPAddress = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "iface", "ip_address"),
		"The IP address associated with an interface. The value is always 1.",
		[]string{
			"node",
			"iface",
			"ip_address",
		}, nil,
	)
	ifacePropHWAddress = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "iface", "hw_address"),
		"The MAC address associated with an interface. The value is always 1.",
		[]string{
			"node",
			"iface",
			"hw_address",
		}, nil,
	)
)
