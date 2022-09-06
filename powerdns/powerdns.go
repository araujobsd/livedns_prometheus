package powerdns

import (
	"github.com/prometheus/client_golang/prometheus"
	// BEGIN LiveDNS platform a, b and c
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strings"
	"strconv"
	// END LiveDNS platform a, b and c
)

type PowerdnsCollector struct {
	client     powerdnsClient
	hostname   string
	datacenter string
	platform   string

	qtypes   *prometheus.Desc
	respsize *prometheus.Desc

	// Define metric for LiveDNS abc platform
	livednsNetwork *prometheus.Desc

	corruptPackets             *prometheus.Desc
	deferredCacheInserts       *prometheus.Desc
	deferredCacheLookup        *prometheus.Desc
	deferredPacketcacheInserts *prometheus.Desc
	deferredPacketcacheLookup  *prometheus.Desc
	dnsupdateAnswers           *prometheus.Desc
	dnsupdateChanges           *prometheus.Desc
	dnsupdateQueries           *prometheus.Desc
	dnsupdateRefused           *prometheus.Desc
	incomingNotifications      *prometheus.Desc
	overloadDrops              *prometheus.Desc
	packetcacheHit             *prometheus.Desc
	packetcacheMiss            *prometheus.Desc
	packetcacheSize            *prometheus.Desc
	queryCacheHit              *prometheus.Desc
	queryCacheMiss             *prometheus.Desc
	queryCacheSize             *prometheus.Desc
	rdQueries                  *prometheus.Desc
	recursingAnswers           *prometheus.Desc
	recursingQuestions         *prometheus.Desc
	recursionUnanswered        *prometheus.Desc
	securityStatus             *prometheus.Desc
	servfailPackets            *prometheus.Desc
	signatures                 *prometheus.Desc
	answers                    *prometheus.Desc
	answersBytes               *prometheus.Desc
	queries                    *prometheus.Desc
	timedoutPackets            *prometheus.Desc
	udpDoQueries               *prometheus.Desc
	fdUsage                    *prometheus.Desc
	keyCacheSize               *prometheus.Desc
	latency                    *prometheus.Desc
	metaCacheSize              *prometheus.Desc
	qsizeQ                     *prometheus.Desc
	realMemoryUsage            *prometheus.Desc
	signatureCacheSize         *prometheus.Desc
	sysMsec                    *prometheus.Desc
	udpInErrors                *prometheus.Desc
	udpNoportErrors            *prometheus.Desc
	udpRecvbufErrors           *prometheus.Desc
	udpSndbufErrors            *prometheus.Desc
	uptime                     *prometheus.Desc
	userMsec                   *prometheus.Desc
}

func NewMetricCollector(socket, hostname, datacenter, platform string) PowerdnsCollector {
	p := PowerdnsCollector{
		client: powerdnsClient{
			socket: socket,
		},
		hostname:   hostname,
		datacenter: datacenter,
		platform:   platform,
	}

	// Collect LiveDNS abc platform metrics
	labels := []string{"datacenter", "platform", "hostname"}
	p.livednsNetwork = prometheus.NewDesc("livedns_platform_count", "dns request per platform", labels, nil)

	labels = []string{"datacenter", "platform", "hostname", "qtype"}
	p.qtypes = prometheus.NewDesc("livedns_qtypes_count", "qtypes in requests", labels, nil)

	labels = []string{"datacenter", "platform", "hostname", "bucket_size"}
	p.respsize = prometheus.NewDesc("livedns_respsize_count", "response by size buckets", labels, nil)

	labels = []string{"datacenter", "platform", "hostname"}
	p.corruptPackets = prometheus.NewDesc("livedns_pdns_stats_corrupt_packets_count", "CorruptPackets", labels, nil)
	p.deferredCacheInserts = prometheus.NewDesc("livedns_pdns_stats_deferred_cache_inserts_count", "DeferredCacheInserts", labels, nil)
	p.deferredCacheLookup = prometheus.NewDesc("livedns_pdns_stats_deferred_cache_lookup_count", "DeferredCacheLookup", labels, nil)
	p.deferredPacketcacheInserts = prometheus.NewDesc("livedns_pdns_stats_deferred_packetcache_inserts_count", "DeferredPacketcacheInserts", labels, nil)
	p.deferredPacketcacheLookup = prometheus.NewDesc("livedns_pdns_stats_deferred_packetcache_lookup_count", "DeferredPacketcacheLookup", labels, nil)
	p.dnsupdateAnswers = prometheus.NewDesc("livedns_pdns_stats_dnsupdate_answers_count", "DnsupdateAnswers", labels, nil)
	p.dnsupdateChanges = prometheus.NewDesc("livedns_pdns_stats_dnsupdate_changes_count", "DnsupdateChanges", labels, nil)
	p.dnsupdateQueries = prometheus.NewDesc("livedns_pdns_stats_dnsupdate_queries_count", "DnsupdateQueries", labels, nil)
	p.dnsupdateRefused = prometheus.NewDesc("livedns_pdns_stats_dnsupdate_refused_count", "DnsupdateRefused", labels, nil)
	p.incomingNotifications = prometheus.NewDesc("livedns_pdns_stats_incoming_notifications_count", "IncomingNotifications", labels, nil)
	p.overloadDrops = prometheus.NewDesc("livedns_pdns_stats_overload_drops_count", "OverloadDrops", labels, nil)
	p.packetcacheHit = prometheus.NewDesc("livedns_pdns_stats_packetcache_hit_count", "PacketcacheHit", labels, nil)
	p.packetcacheMiss = prometheus.NewDesc("livedns_pdns_stats_packetcache_miss_count", "PacketcacheMiss", labels, nil)
	p.packetcacheSize = prometheus.NewDesc("livedns_pdns_stats_packetcache_size_count", "PacketcacheSize", labels, nil)
	p.queryCacheHit = prometheus.NewDesc("livedns_pdns_stats_query_cache_hit_count", "QueryCacheHit", labels, nil)
	p.queryCacheMiss = prometheus.NewDesc("livedns_pdns_stats_query_cache_miss_count", "QueryCacheMiss", labels, nil)
	p.queryCacheSize = prometheus.NewDesc("livedns_pdns_stats_query_cache_size_count", "QueryCacheSize", labels, nil)
	p.rdQueries = prometheus.NewDesc("livedns_pdns_stats_rd_queries_count", "RdQueries", labels, nil)
	p.recursingAnswers = prometheus.NewDesc("livedns_pdns_stats_recursing_answers_count", "RecursingAnswers", labels, nil)
	p.recursingQuestions = prometheus.NewDesc("livedns_pdns_stats_recursing_questions_count", "RecursingQuestions", labels, nil)
	p.recursionUnanswered = prometheus.NewDesc("livedns_pdns_stats_recursion_unanswered_count", "RecursionUnanswered", labels, nil)
	p.securityStatus = prometheus.NewDesc("livedns_pdns_stats_security_status_count", "SecurityStatus", labels, nil)
	p.servfailPackets = prometheus.NewDesc("livedns_pdns_stats_servfail_packets_count", "ServfailPackets", labels, nil)
	p.signatures = prometheus.NewDesc("livedns_pdns_stats_signatures_count", "Signatures", labels, nil)

	labels = []string{"datacenter", "platform", "hostname", "ip_version", "protocol"}
	p.answers = prometheus.NewDesc("livedns_pdns_stats_answers_count", "Answers", labels, nil)
	p.answersBytes = prometheus.NewDesc("livedns_pdns_stats_answers_bytes", "AnswersBytes", labels, nil)
	p.queries = prometheus.NewDesc("livedns_pdns_stats_queries_count", "Queries", labels, nil)

	labels = []string{"datacenter", "platform", "hostname"}
	p.timedoutPackets = prometheus.NewDesc("livedns_pdns_stats_timedout_packets_count", "TimedoutPackets", labels, nil)
	p.udpDoQueries = prometheus.NewDesc("livedns_pdns_stats_udp_do_queries_count", "UdpDoQueries", labels, nil)
	p.fdUsage = prometheus.NewDesc("livedns_pdns_stats_fd_usage_count", "FdUsage", labels, nil)
	p.keyCacheSize = prometheus.NewDesc("livedns_pdns_stats_key_cache_size_count", "KeyCacheSize", labels, nil)
	p.latency = prometheus.NewDesc("livedns_pdns_stats_latency_seconds", "Latency", labels, nil)
	p.metaCacheSize = prometheus.NewDesc("livedns_pdns_stats_meta_cache_size_count", "MetaCacheSize", labels, nil)
	p.qsizeQ = prometheus.NewDesc("livedns_pdns_stats_qsize_q_count", "QsizeQ", labels, nil)
	p.realMemoryUsage = prometheus.NewDesc("livedns_pdns_stats_real_memory_usage_count", "RealMemoryUsage", labels, nil)
	p.signatureCacheSize = prometheus.NewDesc("livedns_pdns_stats_signature_cache_size_count", "SignatureCacheSize", labels, nil)
	p.sysMsec = prometheus.NewDesc("livedns_pdns_stats_icpu_sys_seconds", "SysMsec", labels, nil)
	p.udpInErrors = prometheus.NewDesc("livedns_pdns_stats_udp_in_errors_count", "UdpInErrors", labels, nil)
	p.udpNoportErrors = prometheus.NewDesc("livedns_pdns_stats_udp_noport_errors_count", "UdpNoportErrors", labels, nil)
	p.udpRecvbufErrors = prometheus.NewDesc("livedns_pdns_stats_udp_recvbuf_errors_count", "UdpRecvbufErrors", labels, nil)
	p.udpSndbufErrors = prometheus.NewDesc("livedns_pdns_stats_udp_sndbuf_errors_count", "UdpSndbufErrors", labels, nil)
	p.uptime = prometheus.NewDesc("livedns_pdns_stats_uptime_seconds", "Uptime", labels, nil)
	p.userMsec = prometheus.NewDesc("livedns_pdns_stats_cpu_user_seconds", "UserMsec", labels, nil)

	return p
}

func (p *PowerdnsCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- p.qtypes
	ch <- p.respsize

	// Channel for LiveDNS abc platform metrics.
	ch <- p.livednsNetwork

	ch <- p.corruptPackets
	ch <- p.deferredCacheInserts
	ch <- p.deferredCacheLookup
	ch <- p.deferredPacketcacheInserts
	ch <- p.deferredPacketcacheLookup
	ch <- p.dnsupdateAnswers
	ch <- p.dnsupdateChanges
	ch <- p.dnsupdateQueries
	ch <- p.dnsupdateRefused
	ch <- p.incomingNotifications
	ch <- p.overloadDrops
	ch <- p.packetcacheHit
	ch <- p.packetcacheMiss
	ch <- p.packetcacheSize
	ch <- p.queryCacheHit
	ch <- p.queryCacheMiss
	ch <- p.queryCacheSize
	ch <- p.rdQueries
	ch <- p.recursingAnswers
	ch <- p.recursingQuestions
	ch <- p.recursionUnanswered
	ch <- p.securityStatus
	ch <- p.servfailPackets
	ch <- p.signatures
	ch <- p.answers
	ch <- p.answersBytes
	ch <- p.queries
	ch <- p.timedoutPackets
	ch <- p.udpDoQueries
	ch <- p.fdUsage
	ch <- p.keyCacheSize
	ch <- p.latency
	ch <- p.metaCacheSize
	ch <- p.qsizeQ
	ch <- p.realMemoryUsage
	ch <- p.signatureCacheSize
	ch <- p.sysMsec
	ch <- p.udpInErrors
	ch <- p.udpNoportErrors
	ch <- p.udpRecvbufErrors
	ch <- p.udpSndbufErrors
	ch <- p.uptime
	ch <- p.userMsec
}

func (p *PowerdnsCollector) Collect(ch chan<- prometheus.Metric) {
	qtypes, err := p.client.Qtypes()
	if err == nil {
		for k := range qtypes {
			labels := []string{p.datacenter, p.platform, p.hostname, k}
			ch <- prometheus.MustNewConstMetric(p.qtypes, prometheus.GaugeValue, float64(qtypes[k]), labels...)
		}
	}

	respsizes, err := p.client.Respsizes()
	if err == nil {
		for k := range respsizes {
			labels := []string{p.datacenter, p.platform, p.hostname, k}
			ch <- prometheus.MustNewConstMetric(p.respsize, prometheus.GaugeValue, float64(respsizes[k]), labels...)
		}
	}

	// BEGIN LiveDNS platform.
	// LiveDNS platform variables to aggregate the data.
	var row []string
	var rows [][]string

	// Perform request to PDNS webserver metrics
	resp, err := http.Get("http://127.0.0.1:8081")
	if err != nil {
		print(err)
		return
	}
	// Close resp when we are out of scope.
	defer resp.Body.Close()

	// Parse the response data.
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	// Collect the div.panel from webserver metrics for LiveDNS platform.
	doc.Find("div.panel").Each(func(i int, s *goquery.Selection) {
		if strings.Contains(s.Text(), "Queries received on platform") {
			s.Find("tr td").Each(func(i int, td *goquery.Selection) {
				row = append(row, strings.TrimSpace(td.Text()))
				if len(row) == 3 {
					rows = append(rows, row)
					row = nil
				}

			})
		}
		row = nil
	})

	// Export the metrics per LiveDNS platform.
	for _, data := range rows {
		if strings.Contains(data[0], "queries-recv-address-a") {
			labels := []string{p.datacenter, "livedns_a", p.hostname}
			count,_ := strconv.ParseFloat(data[1], 64)
			ch <- prometheus.MustNewConstMetric(p.livednsNetwork, prometheus.GaugeValue, float64(count), labels...)
		} else if strings.Contains(data[0], "queries-recv-address-b") {
			labels := []string{p.datacenter, "livedns_b", p.hostname}
			count,_ := strconv.ParseFloat(data[1], 64)
			ch <- prometheus.MustNewConstMetric(p.livednsNetwork, prometheus.GaugeValue, float64(count), labels...)
		} else if strings.Contains(data[0], "queries-recv-address-c") {
			labels := []string{p.datacenter, "livedns_c", p.hostname}
			count,_ := strconv.ParseFloat(data[1], 64)
			ch <- prometheus.MustNewConstMetric(p.livednsNetwork, prometheus.GaugeValue, float64(count), labels...)
		}
	}
	// END LiveDNS platform.

	stats, err := p.client.Stats()
	if err == nil {
		labels := []string{p.datacenter, p.platform, p.hostname}

		ch <- prometheus.MustNewConstMetric(p.corruptPackets, prometheus.GaugeValue, float64(stats.CorruptPackets), labels...)
		ch <- prometheus.MustNewConstMetric(p.deferredCacheInserts, prometheus.GaugeValue, float64(stats.DeferredCacheInserts), labels...)
		ch <- prometheus.MustNewConstMetric(p.deferredCacheLookup, prometheus.GaugeValue, float64(stats.DeferredCacheLookup), labels...)
		ch <- prometheus.MustNewConstMetric(p.deferredPacketcacheInserts, prometheus.GaugeValue, float64(stats.DeferredPacketcacheInserts), labels...)
		ch <- prometheus.MustNewConstMetric(p.deferredPacketcacheLookup, prometheus.GaugeValue, float64(stats.DeferredPacketcacheLookup), labels...)
		ch <- prometheus.MustNewConstMetric(p.dnsupdateAnswers, prometheus.GaugeValue, float64(stats.DnsupdateAnswers), labels...)
		ch <- prometheus.MustNewConstMetric(p.dnsupdateChanges, prometheus.GaugeValue, float64(stats.DnsupdateChanges), labels...)
		ch <- prometheus.MustNewConstMetric(p.dnsupdateQueries, prometheus.GaugeValue, float64(stats.DnsupdateQueries), labels...)
		ch <- prometheus.MustNewConstMetric(p.dnsupdateRefused, prometheus.GaugeValue, float64(stats.DnsupdateRefused), labels...)
		ch <- prometheus.MustNewConstMetric(p.incomingNotifications, prometheus.GaugeValue, float64(stats.IncomingNotifications), labels...)
		ch <- prometheus.MustNewConstMetric(p.overloadDrops, prometheus.GaugeValue, float64(stats.OverloadDrops), labels...)
		ch <- prometheus.MustNewConstMetric(p.packetcacheHit, prometheus.GaugeValue, float64(stats.PacketcacheHit), labels...)
		ch <- prometheus.MustNewConstMetric(p.packetcacheMiss, prometheus.GaugeValue, float64(stats.PacketcacheMiss), labels...)
		ch <- prometheus.MustNewConstMetric(p.packetcacheSize, prometheus.GaugeValue, float64(stats.PacketcacheSize), labels...)
		ch <- prometheus.MustNewConstMetric(p.queryCacheHit, prometheus.GaugeValue, float64(stats.QueryCacheHit), labels...)
		ch <- prometheus.MustNewConstMetric(p.queryCacheMiss, prometheus.GaugeValue, float64(stats.QueryCacheMiss), labels...)
		ch <- prometheus.MustNewConstMetric(p.queryCacheSize, prometheus.GaugeValue, float64(stats.QueryCacheSize), labels...)
		ch <- prometheus.MustNewConstMetric(p.rdQueries, prometheus.GaugeValue, float64(stats.RdQueries), labels...)
		ch <- prometheus.MustNewConstMetric(p.recursingAnswers, prometheus.GaugeValue, float64(stats.RecursingAnswers), labels...)
		ch <- prometheus.MustNewConstMetric(p.recursingQuestions, prometheus.GaugeValue, float64(stats.RecursingQuestions), labels...)
		ch <- prometheus.MustNewConstMetric(p.recursionUnanswered, prometheus.GaugeValue, float64(stats.RecursionUnanswered), labels...)
		ch <- prometheus.MustNewConstMetric(p.securityStatus, prometheus.GaugeValue, float64(stats.SecurityStatus), labels...)
		ch <- prometheus.MustNewConstMetric(p.servfailPackets, prometheus.GaugeValue, float64(stats.ServfailPackets), labels...)
		ch <- prometheus.MustNewConstMetric(p.signatures, prometheus.GaugeValue, float64(stats.Signatures), labels...)
		ch <- prometheus.MustNewConstMetric(p.timedoutPackets, prometheus.GaugeValue, float64(stats.TimedoutPackets), labels...)
		ch <- prometheus.MustNewConstMetric(p.udpDoQueries, prometheus.GaugeValue, float64(stats.UdpDoQueries), labels...)
		ch <- prometheus.MustNewConstMetric(p.fdUsage, prometheus.GaugeValue, float64(stats.FdUsage), labels...)
		ch <- prometheus.MustNewConstMetric(p.keyCacheSize, prometheus.GaugeValue, float64(stats.KeyCacheSize), labels...)
		ch <- prometheus.MustNewConstMetric(p.latency, prometheus.GaugeValue, float64(stats.Latency/100000.), labels...)
		ch <- prometheus.MustNewConstMetric(p.metaCacheSize, prometheus.GaugeValue, float64(stats.MetaCacheSize), labels...)
		ch <- prometheus.MustNewConstMetric(p.qsizeQ, prometheus.GaugeValue, float64(stats.QsizeQ), labels...)
		ch <- prometheus.MustNewConstMetric(p.realMemoryUsage, prometheus.GaugeValue, float64(stats.RealMemoryUsage), labels...)
		ch <- prometheus.MustNewConstMetric(p.signatureCacheSize, prometheus.GaugeValue, float64(stats.SignatureCacheSize), labels...)
		ch <- prometheus.MustNewConstMetric(p.sysMsec, prometheus.GaugeValue, float64(stats.SysMsec/1000.), labels...)
		ch <- prometheus.MustNewConstMetric(p.udpInErrors, prometheus.GaugeValue, float64(stats.UdpInErrors), labels...)
		ch <- prometheus.MustNewConstMetric(p.udpNoportErrors, prometheus.GaugeValue, float64(stats.UdpNoportErrors), labels...)
		ch <- prometheus.MustNewConstMetric(p.udpRecvbufErrors, prometheus.GaugeValue, float64(stats.UdpRecvbufErrors), labels...)
		ch <- prometheus.MustNewConstMetric(p.udpSndbufErrors, prometheus.GaugeValue, float64(stats.UdpSndbufErrors), labels...)
		ch <- prometheus.MustNewConstMetric(p.uptime, prometheus.GaugeValue, float64(stats.Uptime), labels...)
		ch <- prometheus.MustNewConstMetric(p.userMsec, prometheus.GaugeValue, float64(stats.UserMsec/1000.), labels...)
		labels = []string{p.datacenter, p.platform, p.hostname, "v4", "tcp"}
		ch <- prometheus.MustNewConstMetric(p.answers, prometheus.GaugeValue, float64(stats.Tcp4Answers), labels...)
		ch <- prometheus.MustNewConstMetric(p.answersBytes, prometheus.GaugeValue, float64(stats.Tcp4AnswersBytes), labels...)
		ch <- prometheus.MustNewConstMetric(p.queries, prometheus.GaugeValue, float64(stats.Tcp4Queries), labels...)
		labels = []string{p.datacenter, p.platform, p.hostname, "v6", "tcp"}
		ch <- prometheus.MustNewConstMetric(p.answers, prometheus.GaugeValue, float64(stats.Tcp6Answers), labels...)
		ch <- prometheus.MustNewConstMetric(p.answersBytes, prometheus.GaugeValue, float64(stats.Tcp6AnswersBytes), labels...)
		ch <- prometheus.MustNewConstMetric(p.queries, prometheus.GaugeValue, float64(stats.Tcp6Queries), labels...)
		labels = []string{p.datacenter, p.platform, p.hostname, "v4", "udp"}
		ch <- prometheus.MustNewConstMetric(p.answers, prometheus.GaugeValue, float64(stats.Udp4Answers), labels...)
		ch <- prometheus.MustNewConstMetric(p.answersBytes, prometheus.GaugeValue, float64(stats.Udp4AnswersBytes), labels...)
		ch <- prometheus.MustNewConstMetric(p.queries, prometheus.GaugeValue, float64(stats.Udp4Queries), labels...)
		labels = []string{p.datacenter, p.platform, p.hostname, "v6", "udp"}
		ch <- prometheus.MustNewConstMetric(p.answers, prometheus.GaugeValue, float64(stats.Udp6Answers), labels...)
		ch <- prometheus.MustNewConstMetric(p.answersBytes, prometheus.GaugeValue, float64(stats.Udp6AnswersBytes), labels...)
		ch <- prometheus.MustNewConstMetric(p.queries, prometheus.GaugeValue, float64(stats.Udp6Queries), labels...)
	}
}
