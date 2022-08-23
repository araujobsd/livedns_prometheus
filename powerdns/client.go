package powerdns

import (
	"io"
	"net"

	"gitlab.corp.gandi.net/devops/livedns-prometheus/powerdns/parser"
)

type powerdnsClient struct {
	socket string
}

func (p *powerdnsClient) ask(command string) ([]byte, error) {
	conn, err := net.Dial("unix", p.socket)
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	command = command + string('\n')

	_, err = conn.Write([]byte(command))
	if err != nil {
		return nil, err
	}

	var out []byte

	for {
		buf := make([]byte, 1500)
		n, err := conn.Read(buf)
		if err != nil && err != io.EOF {
			return nil, err
		}
		if err == io.EOF {
			out = append(out, buf...)
			break
		}
		if n == 0 {
			break
		}

		out = append(out, buf...)
	}

	return out, nil
}

func (p *powerdnsClient) Stats() (*parser.PowerdnsStats, error) {
	// Trailing space is not a typo here
	buf, err := p.ask("SHOW * ")
	if err != nil {
		return nil, err
	}
	stats := parser.ParseStats(buf)

	return &stats, nil
}

func (p *powerdnsClient) Qtypes() (map[string]int64, error) {
	buf, err := p.ask("qtypes")
	if err != nil {
		return nil, err
	}
	qtypes := parser.ParseQtypes(buf)

	return qtypes, nil
}

func (p *powerdnsClient) Respsizes() (map[string]int64, error) {
	buf, err := p.ask("respsizes")
	if err != nil {
		return nil, err
	}
	respsizes := parser.ParseRespsizes(buf)

	return respsizes, nil
}

