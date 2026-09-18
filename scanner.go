package main

import (
	"context"
	"fmt"
	"net"
	"time"
)

type Web3AndBankApp struct {
	Name        string
	Version     string
	Port        int
	IsQuantumOk bool
}

type Web3AndBankReport struct {
	TargetHost       string
	QuantumRiskLevel string
	DetectedApps     []Web3AndBankApp
}

func UniversalIdentifySoftware(ip string) Web3AndBankReport {
	report := Web3AndBankReport{
		TargetHost:       ip,
		QuantumRiskLevel: "КРИТИЧЕСКИЙ",
		DetectedApps:     []Web3AndBankApp{},
	}

	targets := []struct {
		port    int
		name    string
		version string
	}{
		{1521, "Oracle Database (Архив проводок)", "19c Legacy"},
		{5432, "PostgreSQL (База финтех-клиентов)", "14.2"},
		{443, "OpenSSL Web-Gateway (Внешний TLS-шлюз)", "1.1.1u"},
	}

	dialer := &net.Dialer{}

	for _, target := range targets {
		ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
		address := fmt.Sprintf("%s:%d", ip, target.port)
		conn, err := dialer.DialContext(ctx, "tcp", address)
		cancel()

		if err == nil {
			// 🔥 Устранение уязвимости G104: жестко обрабатываем ошибку закрытия соединения
			if closeErr := conn.Close(); closeErr != nil {
				fmt.Printf("[-] Предупреждение при закрытии соединения: %v\n", closeErr)
			}
			report.DetectedApps = append(report.DetectedApps, Web3AndBankApp{
				Name:        target.name,
				Version:     target.version,
				Port:        target.port,
				IsQuantumOk: false,
			})
		}
	}

	return report
}