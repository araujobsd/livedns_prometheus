package bird

import (
	"github.com/czerwonk/bird_exporter/protocol"
)

type LivednsLabelStrategy struct {
	hostname   string
	datacenter string
	platform   string
}

func NewLivednsLabelStrategy(hostname, datacenter, platform string) LivednsLabelStrategy {
	return LivednsLabelStrategy{
		hostname:   hostname,
		datacenter: datacenter,
		platform:   platform,
	}
}

func (l *LivednsLabelStrategy) LabelNames() []string {
	return []string{"name", "proto", "ip_version", "hostname", "datacenter", "platform"}
}

func (l *LivednsLabelStrategy) LabelValues(p *protocol.Protocol) []string {
	return []string{p.Name, protoString(p), p.IpVersion, l.hostname, l.datacenter, l.platform}
}

func protoString(p *protocol.Protocol) string {
	switch p.Proto {
	case protocol.BGP:
		return "BGP"
	case protocol.OSPF:
		if p.IpVersion == "4" {
			return "OSPF"
		}
		return "OSPFv3"
	case protocol.Static:
		return "Static"
	case protocol.Kernel:
		return "Kernel"
	case protocol.Direct:
		return "Direct"
	}

	return ""
}
