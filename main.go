package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/czerwonk/bird_exporter/protocol"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/log"

	"github.com/araujobsd/livedns_prometheus/bird"
	"github.com/araujobsd/livedns_prometheus/powerdns"
)

const version string = "0.0.3"
const metricsPath string = "/metrics"

var (
	showVersion   = flag.Bool("version", false, "Print version information.")
	listenAddress = flag.String("web.listen-address", ":9324", "Address on which to expose metrics and web interface.")
	birdSocket    = flag.String("bird.socket", "/run/bird/bird.ctl", "Socket to communicate with bird routing daemon")
	birdV2        = flag.Bool("bird.v2", false, "Bird major version >= 2.0 (multi channel protocols)")
	// pre bird 2.0
	bird6Socket  = flag.String("bird.socket6", "/run/bird/bird6.ctl", "Socket to communicate with bird6 routing daemon (not compatible with -bird.v2)")
	birdEnabled  = flag.Bool("bird.ipv4", true, "Get protocols from bird (not compatible with -bird.v2)")
	bird6Enabled = flag.Bool("bird.ipv6", true, "Get protocols from bird6 (not compatible with -bird.v2)")

	// TLS
	ca  = flag.String("tls.ca", "/etc/gandi/monitoring.pem", "CA to authenticate clients")
	key = flag.String("tls.key", "/etc/gandi/hostname.key", "Private key for serving requests")
	crt = flag.String("tls.crt", "/etc/gandi/hostname.crt", "Certificate for serving requests")

	// Livedns
	hostname   = flag.String("hostname", "", "Hostname to report in metrics")
	datacenter = flag.String("datacenter", "", "Datacenter to report in metrics")
	platform   = flag.String("platform", "livedns", "livedns platform to report in metrics (livedns, abc, premiumdns, ...)")

	//Powerdns
	powerdnsSocket = flag.String("powerdns.socket", "/var/run/pdns.controlsocket", "Socket to communicate with powerdns daemon")
)

func init() {
	flag.Usage = func() {
		fmt.Println("Usage: bird_exporter [ ... ]\n\nParameters:")
		fmt.Println()
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()

	if *showVersion {
		printVersion()
		os.Exit(0)
	}

	tlsConfig, err := SetupTLS(*ca, *key, *crt)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	startServer(tlsConfig)
}

func printVersion() {
	fmt.Println("livedns_exporter")
	fmt.Printf("Version: %s\n", version)
	fmt.Println("Author(s): Arthur Gautier, Daniel Czerwonk")
	fmt.Println("Metric exporter for livedns & bird routing daemon")
}

func startServer(tlsConfig *Tls) {
	log.Infof("Starting livedns exporter (Version: %s)\n", version)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
			<head><title>Livedns Exporter (Version ` + version + `)</title></head>
			<body>
			<h1>Livedns Exporter</h1>
			<p><a href="` + metricsPath + `">Metrics</a></p>
			</body>
			</html>`))
	})
	http.HandleFunc(metricsPath, handleMetricsRequest)

	log.Infof("Listening for %s on %s\n", metricsPath, *listenAddress)
	log.Fatal(tlsConfig.ListenAndServeTLS(*listenAddress, nil))
}

func handleMetricsRequest(w http.ResponseWriter, r *http.Request) {
	reg := prometheus.NewRegistry()
	{
		p := enabledProtocols()
		c := bird.NewMetricCollector(p, *birdSocket, *bird6Socket, *birdEnabled, *bird6Enabled, *birdV2, *hostname, *datacenter, *platform)
		reg.MustRegister(c)
	}
	{
		c := powerdns.NewMetricCollector(*powerdnsSocket, *hostname, *datacenter, *platform)
		reg.MustRegister(&c)
	}

	promhttp.HandlerFor(reg, promhttp.HandlerOpts{
		ErrorLog:      log.NewErrorLogger(),
		ErrorHandling: promhttp.ContinueOnError}).ServeHTTP(w, r)
}
func enabledProtocols() int {
	res := 0

	res |= protocol.BGP
	res |= protocol.OSPF
	res |= protocol.Kernel
	res |= protocol.Static
	res |= protocol.Direct

	return res
}
