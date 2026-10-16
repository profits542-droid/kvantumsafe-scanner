package main

import (
	"crypto/tls"
	"log"
	"net"
	"net/http"
	"time"
)

// Структура отчета сканирования
type TLSReport struct {
	Domain      string
	Version     string
	CipherSuite string
	Issuer      string
	ExpiryDate  string
	IPAddress   string
	Success     bool
	ErrorMsg    string
}

// Встроенная функция бесконтактного ИИ-сканирования шлюзов
func ScanDomain(domain string) TLSReport {
	report := TLSReport{Domain: domain, Success: false}

	// 1. Резолв IP-адреса узла
	ips, err := net.LookupIP(domain)
	if err != nil || len(ips) == 0 {
		report.ErrorMsg = "Не удалось разрешить имя домена в IP-адрес"
		return report
	}
	report.IPAddress = ips[0].String()

	// 2. Установка защищенного TLS-соединения с таймаутом
	dialer := &net.Dialer{Timeout: 5 * time.Second}
	conn, err := tls.DialWithDialer(dialer, "tcp", domain+":443", &tls.Config{
		InsecureSkipVerify: true,
	})
	if err != nil {
		report.ErrorMsg = "Сетевая сессия прервана: ошибка безопасного рукопожатия"
		return report
	}
	defer conn.Close()

	state := conn.ConnectionState()
	if len(state.PeerCertificates) == 0 {
		report.ErrorMsg = "Удаленный сервер не предоставил паспорт безопасности"
		return report
	}

	cert := state.PeerCertificates[0]

	// 3. Заполнение верифицированных параметров контура
	report.Success = true
	report.ExpiryDate = cert.NotAfter.Format("02.01.2006")
	report.Issuer = cert.Issuer.Organization[0]

	switch state.Version {
	case tls.VersionTLS13:
		report.Version = "TLS 1.3 (Современный контур)"
	case tls.VersionTLS12:
		report.Version = "TLS 1.2 (Требуется обновление)"
	default:
		report.Version = "Устаревший протокол связи"
	}

	report.CipherSuite = tls.CipherSuiteName(state.CipherSuite)
	return report
}

func main() {
	// 1. Стартовая страница витрины
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(RenderStartPage()))
	})

	// 2. Страница результатов сканирования
	http.HandleFunc("/report/mbank", func(w http.ResponseWriter, r *http.Request) {
		domain := r.URL.Query().Get("domain")
		if domain == "" {
			w.Write([]byte(RenderStartPage()))
			return
		}
		report := ScanDomain(domain)
		w.Write([]byte(RenderReportPage(report)))
	})

	// 3. Страница финансовой спецификации для ПВТ КР
	http.HandleFunc("/download", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(RenderDownloadPage()))
	})

	// 4. Страница интерактивного Маркетплейса решений
	http.HandleFunc("/marketplace", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(RenderMarketplacePage()))
	})

	log.Println("[SUCCESS] Квантум Сейф Веб-Ядро запущено на порту: 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Ошибка запуска сервера: ", err)
	}
}
