package main

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"net"
	"net/http"
	"time"
	)

type Tls struct {
	config tls.Config
}

func SetupTLS(ca, key, crt string) (*Tls, error) {
	fail := func(err error) (*Tls, error) { return nil, err }

	cert, err := tls.LoadX509KeyPair(crt, key)
	if (err != nil) {
		return fail(err)
	}

	clientCert, err := loadClientCertificate(ca)
	if (err != nil) {
		return fail(err)
	}

	var certs []tls.Certificate
	certs = append(certs, cert)

	config := tls.Config {
		Certificates: certs,

		// Client authentication
		ClientAuth: tls.RequireAndVerifyClientCert,
		ClientCAs: clientCert,
	}

	srv := Tls {
		config: config,
	}
	return &srv, nil
}

func (t *Tls) ListenAndServeTLS(addr string, handler http.Handler) error {
	if addr == "" {
		addr = ":https"
	}

	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	defer ln.Close()

	tlsListener := tls.NewListener(tcpKeepAliveListener{ln.(*net.TCPListener)}, &t.config)
	return http.Serve(tlsListener, handler)
}

type tcpKeepAliveListener struct {
	*net.TCPListener
}

func (ln tcpKeepAliveListener) Accept() (net.Conn, error) {
	tc, err := ln.AcceptTCP()
	if err != nil {
		return nil, err
	}
	tc.SetKeepAlive(true)
	tc.SetKeepAlivePeriod(3 * time.Minute)
	return tc, nil
}

func loadClientCertificate(ca string) (*x509.CertPool, error) {
	fail := func(err error) (*x509.CertPool, error) { return nil, err }

	certPEMBlock, err := ioutil.ReadFile(ca)
	if err != nil {
		return fail(err)
	}

	var cert [][]byte
	var skippedBlockTypes []string
	for {
		var certDERBlock *pem.Block
		certDERBlock, certPEMBlock = pem.Decode(certPEMBlock)
		if certDERBlock == nil {
			break
		}
		if certDERBlock.Type == "CERTIFICATE" {
			cert = append(cert, certDERBlock.Bytes)
		} else {
			skippedBlockTypes = append(skippedBlockTypes, certDERBlock.Type)
		}
	}

	x509Cert, err := x509.ParseCertificate(cert[0])
	if err != nil {
		return fail(err)
	}

	var pool = x509.NewCertPool()
	pool.AddCert(x509Cert)

	return pool, nil
}
