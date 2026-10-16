package main

import (
	"crypto/tls"
	"fmt"
	"net"
	"time"
)

type TLSReport struct {
	Domain      string
	Success     bool
	ErrorMsg    string
	Version     string
	CipherSuite string
	Issuer      string
	ExpiryDate  string
	IPAddress   string
}

func CheckDomainSecurity(domain string) TLSReport {
	report := TLSReport{Domain: domain, Success: false}

	ips, err := net.LookupIP(domain)
	if err != nil || len(ips) == 0 {
		report.ErrorMsg = "Ошибка DNS: Не удалось разрешить сетевой адрес домена."
		return report
	}
	report.IPAddress = ips[0].String()

	dialer := &net.Dialer{Timeout: 3 * time.Second}
	conn, err := tls.DialWithDialer(dialer, "tcp", domain+":443", &tls.Config{InsecureSkipVerify: true})
	if err != nil {
		report.ErrorMsg = fmt.Sprintf("Ошибка сетевого шлюза: Порт 443 не отвечает (%v)", err)
		return report
	}
	defer conn.Close()

	state := conn.ConnectionState()
	report.Success = true
	
	switch state.Version {
	case tls.VersionTLS13:
		report.Version = "TLS 1.3 (Высший протокол квантовой гибкости)"
	case tls.VersionTLS12:
		report.Version = "TLS 1.2 (Стандарт финтеха / Рекомендована верификация)"
	default:
		report.Version = "Устаревший протокол передачи данных"
	}

	report.CipherSuite = tls.CipherSuiteName(state.CipherSuite)
	if len(state.PeerCertificates) > 0 {
		cert := state.PeerCertificates[0]
		if len(cert.Issuer.Organization) > 0 {
			report.Issuer = cert.Issuer.Organization[0]
		} else {
			report.Issuer = cert.Issuer.CommonName
		}
		report.ExpiryDate = cert.NotAfter.Format("2006-01-02")
	}

	return report
}
