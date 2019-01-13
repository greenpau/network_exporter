# Prometheus Exporter for Networking 

Export network device data accessible via network APIs, e.g. Cisco Nexus API,
to Prometheus.

## Table of Contents

* [Introduction](#introduction)
* [Authentication](#authentication)
* [Getting Started](#getting-started)
* [Building From Source](#building-from-source)
* [Exported Metrics](#exported-metrics)
* [Exporter Flags](#exporter-flags)
* [Prometheus Configuration](#prometheus-configuration)

## Introduction

This exporter exports metrics from the following APIs:
* Cisco NX-OS API
  - Interface
  - VLAN
  - Environment: fans, power supplies, and sensors
  - Resources: CPU, memory, and processes
  - Fiber Transceivers

The following is a screenshot for the exporter's `/metrics` page.

[![Metrics Page](https://raw.githubusercontent.com/greenpau/network_exporter/master/assets/images/exporter_metrics_page.png "Metrics Page")](https://raw.githubusercontent.com/greenpau/network_exporter/master/assets/images/exporter_metrics_page.png)

[:arrow_up: Back to Top](#table-of-contents)

## Authentication

The exporter is designed such that all of the network devices managed by the
exporter must be stored in [Ansible Inventory](https://docs.ansible.com/ansible/latest/user_guide/intro_inventory.html).
Additionally, the credentials necessary to access the devices via API are located in
[Ansible Vault](https://docs.ansible.com/ansible/latest/user_guide/vault.html)
file.

This repository contains a number of expamples of Ansible Inventory and Vault.
See `assets/demo` directory.

The `assets/ansible/hosts` file contains two nodes: `ny-sw01` and `ny-sw02`:

```
$ cat assets/ansible/hosts
#
# Cisco Nexus API Managed devices
#

[cisco:children]
cisco-api-switches

[cisco-api-switches]
ny-sw01 os=cisco_nxos host_overwrite=127.0.0.1 api_port=55443
ny-sw02 os=cisco_nxos host_overwrite=127.0.0.1 api_port=80 api_proto=http

[all:vars]
contact_person=Paul Greenberg @greenpau
```

The `ny-sw01` is listening to API requests on IP address `127.0.0.1` and TCP port `55443`.
The protocol is HTTPS (default).

The `ny-sw02` is listening to API requests on IP address `127.0.0.1` and TCP port `80`.
However, in this case, `api_proto` is being used to overwrite HTTPS with HTTP.

The credentials to access `ny-sw01` are located in Ansible Vault file
`assets/demo/default/ansible/vault.yml`. The key to the vault file is in
`assets/demo/default/ansible/vault.key`. The following command is being used to view
the credentials:

```
$ ansible-vault view assets/demo/default/ansible/vault.yml --vault-password-file assets/demo/default/ansible/vault.key
---
credentials:
- regex: ny-sw0[1-9]
  username: admin
  password: 'cisco'
  password_enable: 'cisco'
  priority: 10
  description: 'NY switch password #1'
- regex: ny-sw0[1-9]
  username: admin
  password: 'cisco123'
  password_enable: 'cisco123'
  priority: 5
  description: 'NY switch password #2'
- default: yes
  username: root
  password: 'QWERTY'
  password_enable: 'QWERTY'
  priority: 1
  description: my default password
- default: yes
  username: root
  password: 'root123'
  password_enable: 'root123'
  priority: 5
  description: my default password
```

The first credential in the `credentials` list will match both `ny-sw01` and
`ny-sw02` name and the exporter will use `admin/admin` for accessing the
device. The last credential in the list is a "catch-all" one.

[:arrow_up: Back to Top](#table-of-contents)

## Getting Started

First, create a configuration file for the exporter:

```
cat << EOF > /etc/sysconfig/network-exporter
OPTIONS="-log.level info -api.poll-interval 15 \
-api.inventory /etc/network-exporter/hosts \
-api.vault /etc/network-exporter/vault.yml \
-api.vault.key /etc/network-exporter/vault.key"
EOF
```

Run the following commands to install it:

```bash
wget https://github.com/greenpau/network_exporter/releases/download/v1.0.0/network-exporter-1.0.0.linux-amd64.tar.gz
tar xvzf network-exporter-1.0.0.linux-amd64.tar.gz
cd network-exporter*
./install.sh
cd ..
rm -rf network-exporter-1.0.0.linux-amd64*
systemctl status network-exporter -l
curl "http://localhost:9516/metrics?x-token=anonymous"
```

[:arrow_up: Back to Top](#table-of-contents)

## Building From Source

Run the following commands to build the exporter from source and test it:

```bash
cd $GOPATH/src
mkdir -p github.com/greenpau
cd github.com/greenpau
git clone https://github.com/greenpau/network_exporter.git
cd network_exporter
make
make qtest
```

[:arrow_up: Back to Top](#table-of-contents)

## Exported Metrics

| **Metric** | **Description** | **Labels** |
| ------ | ------- | ------ |
`net_node_up` | Is node up and responding to queries (1) or is it down (0). | `node` |
`net_node_name` | The Ansible inventory name for the device. The value is always set to 1. | `name`, `node` |
`net_node_failed_req_count` | The number of failed requests for a network node. | `node` |
`net_node_next_poll` | The timestamp of the next potential scrape of the node. | `node` |
`net_node_scrape_time` | The amount of time it took to scrape the node. | `node` |
`net_node_hostname` | The configured hostname on the device itself. The value is always set to 1. | `hostname`, `node` |
`net_node_id` | The unique identifier for the physical device, e.g. a serial number. The value is always set to 1. | `id`, `node` |
`net_iface_name` | The name of an interface. The value is always set to 1. | `iface`, `name`, `node` |
`net_iface_local_index` | The local index of an interface. | `iface`, `node` |
`net_iface_descr` | The description attached to an interface. The value is the checksum of the description | `description`, `iface`, `node` |
`net_iface_bandwidth` | The bandwith routing metric of an interface. | `iface`, `node` |
`net_iface_delay` | The delay routing metric of an interface. | `iface`, `node` |
`net_iface_reliability` | The reliability routing metric of an interface. | `iface`, `node` |
`net_iface_rx_load` | The rx_load routing metric of an interface. | `iface`, `node` |
`net_iface_tx_load` | The tx_load routing metric of an interface. | `iface`, `node` |
`net_iface_babbles` | The babbles counter of an interface. | `iface`, `node` |
`net_iface_bad_ethtype_drops` | The bad_ethtype_drops counter of an interface. | `iface`, `node` |
`net_iface_bad_proto_drops` | The bad_proto_drops counter of an interface. | `iface`, `node` |
`net_iface_no_carrier` | The no_carrier counter of an interface. | `iface`, `node` |
`net_iface_dribble` | The dribble counter of an interface. | `iface`, `node` |
`net_iface_input_frame_errors` | The input_frame_errors counter of an interface. | `iface`, `node` |
`net_iface_input_discards` | The input_frame_errors counter of an interface. | `iface`, `node` |
`net_iface_input_errors` | The input_errors counter of an interface. | `iface`, `node` |
`net_iface_input_pause` | The input_pause counter of an interface. | `iface`, `node` |
`net_iface_input_overruns` | The input_overruns counter of an interface. | `iface`, `node` |
`net_iface_input_iface_down_drops` | The input if-down drops counter of an interface. | `iface`, `node` |
`net_iface_input_bytes` | The input_bytes counter of an interface. | `iface`, `node` |
`net_iface_input_ucast_bytes` | The input_ucast_bytes counter of an interface. | `iface`, `node` |
`net_iface_input_packets` | The input_packets counter of an interface. | `iface`, `node` |
`net_iface_input_ucast_packets` | The input_ucast_packets counter of an interface. | `iface`, `node` |
`net_iface_input_bcast_packets` | The input_bcast_packets counter of an interface. | `iface`, `node` |
`net_iface_input_mcast_packets` | The input_mcast_packets counter of an interface. | `iface`, `node` |
`net_iface_input_jumbo_packets` | The input_jumbo_packets counter of an interface. | `iface`, `node` |
`net_iface_input_compressed` | The input_compressed counter of an interface. | `iface`, `node` |
`net_iface_input_fifo` | The input_fifo counter of an interface. | `iface`, `node` |
`net_iface_late_collisions` | The late_collisions counter of an interface. | `iface`, `node` |
`net_iface_lost_carrier` | The lost_carrier counter of an interface. | `iface`, `node` |
`net_iface_output_discards` | The output_discards counter of an interface. | `iface`, `node` |
`net_iface_output_errors` | The output_errors counter of an interface. | `iface`, `node` |
`net_iface_output_pause` | The output_pause counter of an interface. | `iface`, `node` |
`net_iface_output_underruns` | The output_underruns counter of an interface. | `iface`, `node` |
`net_iface_output_bytes` | The output_bytes counter of an interface. | `iface`, `node` |
`net_iface_output_ucast_bytes` | The output_ucast_bytes counter of an interface. | `iface`, `node` |
`net_iface_output_packets` | The output_packets counter of an interface. | `iface`, `node` |
`net_iface_output_ucast_packets` | The output_ucast_packets counter of an interface. | `iface`, `node` |
`net_iface_output_bcast_packets` | The output_bcast_packets counter of an interface. | `iface`, `node` |
`net_iface_output_mcast_packets` | The output_mcast_packets counter of an interface. | `iface`, `node` |
`net_iface_output_jumbo_packets` | The output_jumbo_packets counter of an interface. | `iface`, `node` |
`net_iface_output_carrier_errors` | The output_carrier_errors counter of an interface. | `iface`, `node` |
`net_iface_collisions` | The collisions counter of an interface. | `iface`, `node` |
`net_iface_output_fifo` | The output_fifo counter of an interface. | `iface`, `node` |
`net_iface_watchdog` | The watchdog counter of an interface. | `iface`, `node` |
`net_iface_storm_suppression` | The storm_suppression counter of an interface. | `iface`, `node` |
`net_iface_ignored` | The ignored counter of an interface. | `iface`, `node` |
`net_iface_runts` | The runts counter of an interface. | `iface`, `node` |
`net_iface_crc_errors` | The crc_errors counter of an interface. | `iface`, `node` |
`net_iface_deferred` | The deferred counter of an interface. | `iface`, `node` |
`net_iface_no_buffer` | The no_buffer counter of an interface. | `iface`, `node` |
`net_iface_resets` | The resets counter of an interface. | `iface`, `node` |
`net_iface_beacon_enabled` | Whether beacon is enabled (1) or disabled (0) on an interface. | `iface`, `node` |
`net_iface_auto_negotiation_enabled` | Whether auto negotiation is enabled (1) or disabled (0) on an interface. | `iface`, `node` |
`net_iface_mdix_enabled` | Whether auto MDIX is enabled (1) or disabled (0) on an interface. | `iface`, `node` |
`net_iface_mtu` | The MTU of an interface. | `iface`, `node` |
`net_iface_speed` | The speed (in Mb/s) of an interface. If the value is 0, then it is auto. | `iface`, `node` |
`net_iface_duplex` | The duplex of an interface. Values are auto (3), full (2), half (1), other (0) | `iface`, `node` |
`net_iface_encapsulated_vlan` | The encapsulated VLAN associated with an interface. | `iface`, `node` |
`net_iface_state` | The state of an interface. Values are up (1), any other value (0). | `iface`, `node` |
`net_iface_admin_state` | The state of an interface. Values are up (1), any other value (0). | `iface`, `node` |
`net_iface_subinterface` | Indicates whether an interface is a sub-interface. Values are yes (1), no (0). | `iface`, `node` |
`net_iface_routed_mode` | Indicates whether an interface is in routed (L3-configured) mode. Values are yes (1), no (0). | `iface`, `node` |
`net_iface_access_mode` | Indicates whether an interface is in access (L2-configured) mode. Values are yes (1), no (0). | `iface`, `node` |
`net_iface_ip_address` | The IP address associated with an interface. The value is always 1. | `iface`, `ip_address`, `node` |
`net_iface_hw_address` | The MAC address associated with an interface. The value is always 1. | `hw_address`, `iface`, `node` |
`net_vlan_id` | The Vlan ID of a VLAN. The value is always set to 1. | `node`, `vlan` |
`net_vlan_name` | The name of a VLAN. The value is always set to 1. | `name`, `node`, `vlan` |
`net_vlan_state` | The state of a VLAN. Values are active (1), any other value (0). | `node`, `vlan` |
`net_vlan_shutdown_state` | The shutdown state of a VLAN. Values are noshutdown (1), any other value (0). | `node`, `vlan` |
`net_node_fan_up` | The status of a fan. 1 (up, Ok), 0 (down) | `fan`, `node` |
`net_node_ps_up` | The status of a power supply. 1 (up, Ok), 0 (down) | `node`, `power_supply` |
`net_node_ps_pwr_input` | The power input of a power supply. | `node`, `power_supply` |
`net_node_ps_pwr_output` | The power output of a power supply. | `node`, `power_supply` |
`net_node_ps_pwr_capacity` | The power capacity of a power supply. | `node`, `power_supply` |
`net_node_sensor_up` | The status of a sensor. 1 (up, Ok), 0 (down) | `node`, `sensor` |
`net_node_sensor_temperature` | The temperature of a sensor. | `node`, `sensor` |
`net_node_sensor_temperature_threshold_high` | The alarm upper threshold for the temperature of a sensor. | `node`, `sensor` |
`net_node_sensor_temperature_threshold_low` | The alarm lower threshold for the temperature of a sensor. | `node`, `sensor` |
`net_node_running_process_count` | The number of running processes. | `node` |
`net_node_total_process_count` | The number of total processes. | `node` |
`net_node_memory_total` | The amount of total memory available. | `node` |
`net_node_memory_free` | The amount of free memory available. | `node` |
`net_node_memory_used` | The amount of memory used. | `node` |
`net_node_total_cpu_idle` | The amount of CPU time in idle state. | `node` |
`net_node_total_cpu_kernel` | The amount of CPU time in kernel state. | `node` |
`net_node_total_cpu_user` | The amount of CPU time in user state. | `node` |
`net_node_cpu_idle` | The amount of CPU time in idle state on per CPU basis. | `cpu_id`, `node` |
`net_node_cpu_kernel` | The amount of CPU time in kernel state on per CPU basis. | `cpu_id`, `node` |
`net_node_cpu_user` | The amount of CPU time in user state on per CPU basis. | `cpu_id`, `node` |
`net_interface_transceiver` | The serial number and vendor of a transceiver attached to an interface are the labels of this metric. The value of the metric is always set to 1. | `iface_name`, `node`, `serial`, `vendor` |
`net_interface_transceiver_lane_temperature` | The temperature of a transceiver lane. | `iface_name`, `lane_id`, `node` |
`net_interface_transceiver_lane_voltage` | The voltage of a transceiver lane. | `iface_name`, `lane_id`, `node` |
`net_interface_transceiver_lane_current` | The current of a transceiver lane. | `iface_name`, `lane_id`, `node` |
`net_interface_transceiver_lane_tx_power` | The transmit power of a transceiver lane. | `iface_name`, `lane_id`, `node` |
`net_interface_transceiver_lane_rx_power` | The receive power of a transceiver lane. | `iface_name`, `lane_id`, `node` |
`net_interface_transceiver_lane_errors` | The number of errors with a transceiver lane. | `iface_name`, `lane_id`, `node` |

For example:

```bash
$ curl "http://localhost:9516/metrics?node=ny-sw01&module=cisco_nxos&x-token=anonymous"
```

[:arrow_up: Back to Top](#table-of-contents)

## Exporter Flags

```bash
$ ./bin/network-exporter --help
```

[:arrow_up: Back to Top](#table-of-contents)

## Prometheus Configuration

The following is a sample Prometheus configuration for the exporter.
Here, the exporter resides on a Prometheus server itself.
Hence, the `127.0.0.1`.
The scrape collects data from 4 network switches: `ny-sw01`, `ny-sw01`,
`ny-sw03`, and `ny-sw04`.

```
---
# Prometheus configuration file
global:
  scrape_interval: 1m
  scrape_timeout: 30s

rule_files:
  - "/etc/prometheus/rules/*.yml"
  - "/etc/prometheus/conf/rules/*.yml"
  - "/etc/prometheus/config/rules/*.yml"

alerting:
  alertmanagers:
  - scheme: http
    static_configs:
    - targets:
      - "127.0.0.1:9093"

scrape_configs:
  - job_name: devnet_nxos_nodes
    params:
      module:
        - cisco_nxos
      x_token:
        - anonymous
    scrape_interval: 1m
    scrape_timeout: 1m
    metrics_path: /metrics
    scheme: http
    static_configs:
    - targets:
      - ny-sw01
      - ny-sw02
      - ny-sw03
      - ny-sw04
      labels:
        region: us
    relabel_configs:
    - source_labels: [__address__]
      target_label: __param_target
    - source_labels: [__param_target]
      target_label: instance
    - target_label: __address__
      # the exporter's hostname:port
      replacement: 127.0.0.1:9516
```

[:arrow_up: Back to Top](#table-of-contents)
