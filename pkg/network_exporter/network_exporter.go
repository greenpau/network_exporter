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
	//"github.com/davecgh/go-spew/spew"
	ansible "github.com/greenpau/go-ansible-db/pkg/db"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/log"
	"github.com/prometheus/common/version"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	namespace = "net"
)

var (
	appName    = "network-exporter"
	appVersion = "[untracked]"
	gitBranch  string
	gitCommit  string
	buildUser  string // whoami
	buildDate  string // date -u
)

// Exporter holds the inventory and credentials for accessing network nodes.
// It also contains exporters for each of the nodes.
type Exporter struct {
	sync.RWMutex
	timeout       int
	pollInterval  int64
	InventoryFile string
	VaultFile     string
	VaultKeyFile  string
	Inventory     *ansible.Inventory
	Vault         *ansible.Vault
	Modules       map[string]bool
	Subsystems    map[string]bool
	Nodes         map[string]*NetworkNode
	Tokens        map[string]bool
}

// Options are the options for the initialization of an instance of the
// Exporter.
type Options struct {
	Timeout       int
	InventoryFile string
	VaultFile     string
	VaultKeyFile  string
}

// NewExporter returns an initialized Exporter.
func NewExporter(opts Options) (*Exporter, error) {
	version.Version = appVersion
	version.Revision = gitCommit
	version.Branch = gitBranch
	version.BuildUser = buildUser
	version.BuildDate = buildDate
	e := Exporter{
		timeout:       opts.Timeout,
		InventoryFile: opts.InventoryFile,
		VaultFile:     opts.VaultFile,
		VaultKeyFile:  opts.VaultKeyFile,
		Modules:       make(map[string]bool),
		Subsystems:    make(map[string]bool),
		Nodes:         make(map[string]*NetworkNode),
		Tokens:        make(map[string]bool),
		Inventory:     ansible.NewInventory(),
		Vault:         ansible.NewVault(),
	}
	e.Modules["cisco_nxos"] = true
	e.Subsystems["interfaces"] = true   // interfaces
	e.Subsystems["transceivers"] = true // fiber optics
	e.Subsystems["vlans"] = true        // VLANs
	e.Subsystems["bgp"] = true          // BGP
	e.Subsystems["resources"] = true    // CPU and Memory
	if err := e.updateInventory(); err != nil {
		return nil, err
	}
	for _, n := range e.Nodes {
		if n.timeout == 0 {
			n.timeout = e.timeout
		}
	}
	log.Debugf("NewExporter() initialized successfully")
	return &e, nil
}

func (e *Exporter) updateInventory() error {
	if err := e.Inventory.LoadFromFile(e.InventoryFile); err != nil {
		return fmt.Errorf("error reading inventory: %s", err)
	}
	if err := e.Vault.LoadPasswordFromFile(e.VaultKeyFile); err != nil {
		return fmt.Errorf("error reading vault key file: %s", err)
	}
	if err := e.Vault.LoadFromFile(e.VaultFile); err != nil {
		return fmt.Errorf("error reading vault: %s", err)
	}
	hosts, err := e.Inventory.GetHosts()
	if err != nil {
		return fmt.Errorf("error getting hosts from the inventory: %s", err)
	}
	if len(hosts) < 1 {
		return fmt.Errorf("the inventory has no hosts")
	}
	for _, h := range hosts {
		if nos, exists := h.Variables["os"]; !exists {
			log.Debugf("The host '%s' was not added to exporter because it lacks 'os' atribute", h.Name)
			continue
		} else {
			if _, supported := e.Modules[nos]; !supported {
				log.Debugf("The host '%s' was not added to exporter because 'os' atribute value '%s' is unsupported", h.Name, nos)
				continue
			}
		}
		if _, exists := e.Nodes[h.Name]; !exists {
			hash := sha1.New()
			hash.Write([]byte(h.Name))
			n := &NetworkNode{
				Name:                 h.Name,
				UUID:                 fmt.Sprintf("%x", hash.Sum(nil)),
				result:               "unknown",
				module:               "unknown",
				timestamp:            "unknown",
				nextCollectionTicker: 0,
				errors:               0,
				Variables:            make(map[string]string),
				credentials:          []*credential{},
				Interfaces:           make(map[string]string),
				Vlans:                make(map[string]string),
			}
			for k, v := range h.Variables {
				n.Variables[k] = v
			}
			if nos, exists := n.Variables["os"]; exists {
				n.module = nos
			}
			if target, exists := n.Variables["host_overwrite"]; exists {
				n.target = target
			} else {
				n.target = h.Name
			}
			if apiPort, exists := n.Variables["api_port"]; exists {
				if i, err := strconv.Atoi(apiPort); err == nil {
					n.port = i
				}
			}
			if apiProto, exists := n.Variables["api_proto"]; exists {
				if apiProto == "http" || apiProto == "https" {
					n.proto = apiProto
				} else {
					log.Debugf("The host '%s' was not added to exporter because 'api_proto' atribute value '%s' is unsupported", h.Name, apiProto)
					continue
				}
			}
			e.Nodes[h.Name] = n
		}
	}

	for _, n := range e.Nodes {
		creds, err := e.Vault.GetCredentials(n.Name)
		if err != nil {
			return fmt.Errorf("error getting credentials for host %s: %s", n.Name, err)
		}
		n.credentials = []*credential{}
		for _, c := range creds {
			nc := &credential{
				Username: c.Username,
				Password: c.Password,
				Failed:   false,
			}
			n.credentials = append(n.credentials, nc)
		}
	}
	return nil
}

// GetVersionInfo returns exporter info.
func GetVersionInfo() string {
	return version.Info()
}

// GetVersionBuildContext returns exporter build context.
func GetVersionBuildContext() string {
	return version.BuildContext()
}

// GetVersion returns exporter version.
func GetVersion() string {
	return version.Version
}

// GetRevision returns exporter revision.
func GetRevision() string {
	return version.Revision
}

// GetExporterName returns exporter name.
func GetExporterName() string {
	return appName
}

// SetPollInterval sets exporter's minimal polling/scraping interval.
func (e *Exporter) SetPollInterval(i int64) {
	e.pollInterval = i
	for _, n := range e.Nodes {
		if n.pollInterval == 0 {
			n.pollInterval = i
		}
	}
}

// GetPollInterval returns exporters minimal polling/scraping interval.
func (e *Exporter) GetPollInterval() int64 {
	return e.pollInterval
}

// Scrape scrapes individual nodes.
func (e *Exporter) Scrape(w http.ResponseWriter, r *http.Request) {
	if _, authorized := e.authorize(r); !authorized {
		http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
		return
	}
	var nodeName string
	if r.URL.Query().Get("target") != "" {
		nodeName = r.URL.Query().Get("target")
	} else {
		if r.URL.Query().Get("node") != "" {
			nodeName = r.URL.Query().Get("node")
		}
	}
	if nodeName == "" {
		http.Error(w, "node parameter is required", http.StatusBadRequest)
		return
	}
	node, exists := e.Nodes[nodeName]
	if !exists {
		http.Error(w, fmt.Sprintf("unknown node %q", nodeName), http.StatusBadRequest)
		return
	}
	moduleName := r.URL.Query().Get("module")
	if moduleName == "" {
		node.result = "failure"
		node.timestamp = time.Now().Format(time.RFC3339)
		http.Error(w, "module parameter is required", http.StatusBadRequest)
		return
	}
	if _, supported := e.Modules[moduleName]; !supported {
		node.result = "failure"
		node.timestamp = time.Now().Format(time.RFC3339)
		http.Error(w, fmt.Sprintf("unsupported module %q", moduleName), http.StatusBadRequest)
		return
	}
	subsystemName := r.URL.Query().Get("subsystem")
	if subsystemName == "" {
		subsystemName = "interfaces"
	} else if subsystemName == "all" {
		subsystemName = "interfaces,resources"
	} else {
		// do nothing
	}
	for _, s := range strings.Split(subsystemName, ",") {
		if _, supported := e.Subsystems[s]; !supported {
			node.result = "failure"
			node.timestamp = time.Now().Format(time.RFC3339)
			http.Error(w, fmt.Sprintf("unsupported subsystem %q", s), http.StatusBadRequest)
			return
		}
	}
	subsystems := strings.Split(subsystemName, ",")

	log.Debugf("%s: calls Scrape() for node '%s' and module '%s'", node.UUID, node.Name, moduleName)
	start := time.Now()
	registry := prometheus.NewRegistry()
	node.module = moduleName
	node.subsystems = subsystems
	registry.MustRegister(node)
	h := promhttp.HandlerFor(registry, promhttp.HandlerOpts{})
	h.ServeHTTP(w, r)
	duration := time.Since(start).Seconds()
	log.Debugf(
		"%s: Scrape() for node '%s', module '%s', subsystems '%s' took %f seconds",
		node.UUID, node.Name, moduleName, subsystems, duration,
	)
}
