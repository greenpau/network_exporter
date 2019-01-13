package main

import (
	"flag"
	"fmt"
	exporter "github.com/greenpau/network_exporter/pkg/network_exporter"
	"github.com/prometheus/common/log"
	"net/http"
	"os"
)

func main() {
	var listenAddress string
	var metricsPath string
	var pollTimeout int
	var pollInterval int
	var isShowMetrics bool
	var isShowVersion bool
	var logLevel string
	var apiInventory string
	var apiVault string
	var apiVaultKey string
	var authToken string

	flag.StringVar(&listenAddress, "web.listen-address", ":9533", "Address to listen on for web interface and telemetry.")
	flag.StringVar(&metricsPath, "web.telemetry-path", "/metrics", "Path under which to expose metrics.")
	flag.IntVar(&pollTimeout, "api.timeout", 5, "Timeout on requests to network devices.")
	flag.IntVar(&pollInterval, "api.poll-interval", 15, "The minimum interval (in seconds) between collections from a network device.")
	flag.StringVar(&apiInventory, "api.inventory", "/etc/network-exporter/hosts", "Node inventory file")
	flag.StringVar(&apiVault, "api.vault", "/etc/network-exporter/vault.yml", "Node credentials vault")
	flag.StringVar(&apiVaultKey, "api.vault.key", "/etc/network-exporter/vault.key", "The key to the vault")
	flag.StringVar(&authToken, "auth.token", "anonymous", "The X-Token for accessing the exporter itself")
	flag.BoolVar(&isShowMetrics, "metrics", false, "Display available metrics")
	flag.BoolVar(&isShowVersion, "version", false, "version information")
	flag.StringVar(&logLevel, "log.level", "info", "logging severity level")

	var usageHelp = func() {
		fmt.Fprintf(os.Stderr, "\n%s - Prometheus Exporter for Networking\n\n", exporter.GetExporterName())
		fmt.Fprintf(os.Stderr, "Usage: %s [arguments]\n\n", exporter.GetExporterName())
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nDocumentation: https://github.com/greenpau/network_exporter/\n\n")
	}
	flag.Usage = usageHelp
	flag.Parse()

	opts := exporter.Options{
		Timeout:       pollTimeout,
		InventoryFile: apiInventory,
		VaultFile:     apiVault,
		VaultKeyFile:  apiVaultKey,
	}

	if err := log.Base().SetLevel(logLevel); err != nil {
		log.Errorf(err.Error())
		os.Exit(1)
	}

	if isShowVersion {
		fmt.Fprintf(os.Stdout, "%s %s", exporter.GetExporterName(), exporter.GetVersion())
		if exporter.GetRevision() != "" {
			fmt.Fprintf(os.Stdout, ", commit: %s\n", exporter.GetRevision())
		} else {
			fmt.Fprint(os.Stdout, "\n")
		}
		os.Exit(0)
	}

	if isShowMetrics {
		e := &exporter.NetworkNode{}
		fmt.Fprintf(os.Stdout, "%s\n", e.GetMetricsTable())
		os.Exit(0)
	}

	log.Infof("Starting %s %s", exporter.GetExporterName(), exporter.GetVersionInfo())
	log.Infof("Build context %s", exporter.GetVersionBuildContext())

	e, err := exporter.NewExporter(opts)
	if err != nil {
		log.Errorf("%s failed to init properly: %s", exporter.GetExporterName(), err)
		os.Exit(1)
	}
	e.SetPollInterval(int64(pollInterval))
	if err := e.AddAuthenticationToken(authToken); err != nil {
		log.Errorf("%s failed to add authentication token: %s", exporter.GetExporterName(), err)
		os.Exit(1)
	}

	log.Infof("Inventory file: %s", e.InventoryFile)
	log.Infof("Vault file: %s", e.VaultFile)
	log.Infof("Vault key file: %s", e.VaultKeyFile)
	log.Infof("Minimal scrape interval: %d seconds", e.GetPollInterval())

	http.HandleFunc(metricsPath, func(w http.ResponseWriter, r *http.Request) {
		e.Scrape(w, r)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		e.Summary(metricsPath, w, r)
	})

	log.Infoln("Listening on", listenAddress)
	log.Fatal(http.ListenAndServe(listenAddress, nil))
}
