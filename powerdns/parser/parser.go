package parser

import (
	"bufio"
	"bytes"
	"regexp"
	"strconv"
)

var (
	qtypeRegex *regexp.Regexp
	statRegex  *regexp.Regexp
)

func init() {
	qtypeRegex = regexp.MustCompile(`^([^\t]+)\t([0-9]+)$`)
	statRegex = regexp.MustCompile(`^([^=]+)=([0-9]+)$`)
}

func ParseQtypes(data []byte) map[string]int64 {
	reader := bytes.NewReader(data)
	scanner := bufio.NewScanner(reader)

	out := make(map[string]int64)

	for scanner.Scan() {
		line := scanner.Text()
		match := qtypeRegex.FindStringSubmatch(line)

		if match == nil {
			continue
		}

		key := match[1]
		value := match[2]

		intValue, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			continue
		}

		out[key] = intValue
	}

	return out
}

func ParseRespsizes(data []byte) map[string]int64 {
	return ParseQtypes(data)
}

type PowerdnsStats struct {
	CorruptPackets int64
	DeferredCacheInserts int64
	DeferredCacheLookup int64
	DeferredPacketcacheInserts int64
	DeferredPacketcacheLookup int64
	DnsupdateAnswers int64
	DnsupdateChanges int64
	DnsupdateQueries int64
	DnsupdateRefused int64
	IncomingNotifications int64
	OverloadDrops int64
	PacketcacheHit int64
	PacketcacheMiss int64
	PacketcacheSize int64
	QueryCacheHit int64
	QueryCacheMiss int64
	QueryCacheSize int64
	RdQueries int64
	RecursingAnswers int64
	RecursingQuestions int64
	RecursionUnanswered int64
	SecurityStatus int64
	ServfailPackets int64
	Signatures int64
	TcpAnswers int64
	TcpAnswersBytes int64
	TcpQueries int64
	Tcp4Answers int64
	Tcp4AnswersBytes int64
	Tcp4Queries int64
	Tcp6Answers int64
	Tcp6AnswersBytes int64
	Tcp6Queries int64
	TimedoutPackets int64
	UdpAnswers int64
	UdpAnswersBytes int64
	UdpDoQueries int64
	UdpQueries int64
	Udp4Answers int64
	Udp4AnswersBytes int64
	Udp4Queries int64
	Udp6Answers int64
	Udp6AnswersBytes int64
	Udp6Queries int64
	FdUsage int64
	KeyCacheSize int64
	Latency int64
	MetaCacheSize int64
	QsizeQ int64
	RealMemoryUsage int64
	SignatureCacheSize int64
	SysMsec int64
	UdpInErrors int64
	UdpNoportErrors int64
	UdpRecvbufErrors int64
	UdpSndbufErrors int64
	Uptime int64
	UserMsec int64
}

func ParseStats(data []byte) PowerdnsStats {
	reader := bytes.NewReader(data)
	scanner := bufio.NewScanner(reader)

	onComma := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		for i := 0; i < len(data); i++ {
			if data[i] == ',' {
				return i + 1, data[:i], nil
			}
		}
		// There is one final token to be delivered, which may be the empty string.
		// Returning bufio.ErrFinalToken here tells Scan there are no more tokens after this
		// but does not trigger an error to be returned from Scan itself.
		return 0, data, bufio.ErrFinalToken
	}

	scanner.Split(onComma)

	out := PowerdnsStats{}

	for scanner.Scan() {
		line := scanner.Text()
		match := statRegex.FindStringSubmatch(line)

		if match == nil {
			continue
		}

		key := match[1]
		value := match[2]

		intValue, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			continue
		}

		switch key {
		case "corrupt-packets":
			out.CorruptPackets = intValue
		case "deferred-cache-inserts":
			out.DeferredCacheInserts = intValue
		case "deferred-cache-lookup":
			out.DeferredCacheLookup = intValue
		case "deferred-packetcache-inserts":
			out.DeferredPacketcacheInserts = intValue
		case "deferred-packetcache-lookup":
			out.DeferredPacketcacheLookup = intValue
		case "dnsupdate-answers":
			out.DnsupdateAnswers = intValue
		case "dnsupdate-changes":
			out.DnsupdateChanges = intValue
		case "dnsupdate-queries":
			out.DnsupdateQueries = intValue
		case "dnsupdate-refused":
			out.DnsupdateRefused = intValue
		case "incoming-notifications":
			out.IncomingNotifications = intValue
		case "overload-drops":
			out.OverloadDrops = intValue
		case "packetcache-hit":
			out.PacketcacheHit = intValue
		case "packetcache-miss":
			out.PacketcacheMiss = intValue
		case "packetcache-size":
			out.PacketcacheSize = intValue
		case "query-cache-hit":
			out.QueryCacheHit = intValue
		case "query-cache-miss":
			out.QueryCacheMiss = intValue
		case "query-cache-size":
			out.QueryCacheSize = intValue
		case "rd-queries":
			out.RdQueries = intValue
		case "recursing-answers":
			out.RecursingAnswers = intValue
		case "recursing-questions":
			out.RecursingQuestions = intValue
		case "recursion-unanswered":
			out.RecursionUnanswered = intValue
		case "security-status":
			out.SecurityStatus = intValue
		case "servfail-packets":
			out.ServfailPackets = intValue
		case "signatures":
			out.Signatures = intValue
		case "tcp-answers":
			out.TcpAnswers = intValue
		case "tcp-answers-bytes":
			out.TcpAnswersBytes = intValue
		case "tcp-queries":
			out.TcpQueries = intValue
		case "tcp4-answers":
			out.Tcp4Answers = intValue
		case "tcp4-answers-bytes":
			out.Tcp4AnswersBytes = intValue
		case "tcp4-queries":
			out.Tcp4Queries = intValue
		case "tcp6-answers":
			out.Tcp6Answers = intValue
		case "tcp6-answers-bytes":
			out.Tcp6AnswersBytes = intValue
		case "tcp6-queries":
			out.Tcp6Queries = intValue
		case "timedout-packets":
			out.TimedoutPackets = intValue
		case "udp-answers":
			out.UdpAnswers = intValue
		case "udp-answers-bytes":
			out.UdpAnswersBytes = intValue
		case "udp-do-queries":
			out.UdpDoQueries = intValue
		case "udp-queries":
			out.UdpQueries = intValue
		case "udp4-answers":
			out.Udp4Answers = intValue
		case "udp4-answers-bytes":
			out.Udp4AnswersBytes = intValue
		case "udp4-queries":
			out.Udp4Queries = intValue
		case "udp6-answers":
			out.Udp6Answers = intValue
		case "udp6-answers-bytes":
			out.Udp6AnswersBytes = intValue
		case "udp6-queries":
			out.Udp6Queries = intValue
		case "fd-usage":
			out.FdUsage = intValue
		case "key-cache-size":
			out.KeyCacheSize = intValue
		case "latency":
			out.Latency = intValue
		case "meta-cache-size":
			out.MetaCacheSize = intValue
		case "qsize-q":
			out.QsizeQ = intValue
		case "real-memory-usage":
			out.RealMemoryUsage = intValue
		case "signature-cache-size":
			out.SignatureCacheSize = intValue
		case "sys-msec":
			out.SysMsec = intValue
		case "udp-in-errors":
			out.UdpInErrors = intValue
		case "udp-noport-errors":
			out.UdpNoportErrors = intValue
		case "udp-recvbuf-errors":
			out.UdpRecvbufErrors = intValue
		case "udp-sndbuf-errors":
			out.UdpSndbufErrors = intValue
		case "uptime":
			out.Uptime = intValue
		case "user-msec":
			out.UserMsec = intValue
		}
	}

	return out
}
