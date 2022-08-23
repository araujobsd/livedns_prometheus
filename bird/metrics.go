package bird

import (
	"github.com/czerwonk/bird_exporter/client"
	"github.com/czerwonk/bird_exporter/metrics"
	"github.com/czerwonk/bird_exporter/protocol"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"
)

type MetricCollector struct {
	exporters        map[int][]metrics.MetricExporter
	client           *client.BirdClient
	enabledProtocols int
}

func NewMetricCollector(enabledProtocols int, birdSocket, bird6Socket string, birdEnabled, bird6Enabled, birdV2 bool, hostname, datacenter, platform string) *MetricCollector {
	c := getClient(birdSocket, bird6Socket, bird6Enabled, birdEnabled, birdV2)
	var e map[int][]metrics.MetricExporter

	e = exportersForDefault(c, hostname, datacenter, platform)

	return &MetricCollector{exporters: e, client: c, enabledProtocols: enabledProtocols}
}

func getClient(birdSocket string, bird6Socket string, bird6Enabled bool, birdEnabled bool, birdV2 bool) *client.BirdClient {
	o := &client.BirdClientOptions{
		BirdSocket:   birdSocket,
		Bird6Socket:  bird6Socket,
		Bird6Enabled: bird6Enabled,
		BirdEnabled:  birdEnabled,
		BirdV2:       birdV2,
	}

	return &client.BirdClient{Options: o}
}

func exportersForDefault(c *client.BirdClient, hostname, datacenter, platform string) map[int][]metrics.MetricExporter {
	l := NewLivednsLabelStrategy(hostname, datacenter, platform)
	e := metrics.NewGenericProtocolMetricExporter("bird_protocol", true, &l)

	return map[int][]metrics.MetricExporter{
		protocol.BGP:    {e},
		protocol.Direct: {e},
		protocol.Kernel: {e},
		protocol.OSPF:   {e, metrics.NewOspfExporter("bird_", c)},
		protocol.Static: {e},
	}
}

func (m *MetricCollector) Describe(ch chan<- *prometheus.Desc) {
	for _, v := range m.exporters {
		for _, e := range v {
			e.Describe(ch)
		}
	}
}

func (m *MetricCollector) Collect(ch chan<- prometheus.Metric) {
	protocols, err := m.client.GetProtocols()
	if err != nil {
		log.Errorln(err)
		return
	}

	for _, p := range protocols {
		if p.Proto == protocol.PROTO_UNKNOWN || (m.enabledProtocols&p.Proto != p.Proto) {
			continue
		}

		for _, e := range m.exporters[p.Proto] {
			e.Export(p, ch)
		}
	}
}
