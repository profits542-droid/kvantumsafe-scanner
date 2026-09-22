package main

import (
	"encoding/base64"
	"fmt"
	"html"
	"net/http"
	"strings"
	"time"
)

type LicenseInfo struct {
	ClientName string
	ServerID   string
	IsValid    bool
	DaysLeft   int
}

func VerifyLicenseKey(encodedKey string) LicenseInfo {
	info := LicenseInfo{IsValid: false, DaysLeft: 0}
	decodedBytes, err := base64.StdEncoding.DecodeString(encodedKey)
	if err != nil {
		return info
	}
	parts := strings.Split(string(decodedBytes), "|")
	if len(parts) == 3 {
		info.ClientName = parts[0]
		info.ServerID = parts[2]
		if expTime, err := time.Parse("2006-01-02", parts[1]); err == nil {
			if time.Now().Before(expTime) {
				info.IsValid = true
				info.DaysLeft = int(expTime.Sub(time.Now()).Hours() / 24)
			}
		}
	}
	return info
}

func main() {
	premiumLicenseKey := "T0FPIEtvbW1lcmNoZXNreWkgYmFuayBLWVJHWVpTVEFOIChNQkFOSyl8MjAyNy0wOS0yMHxDb3JlLU5vZGUtMDE="
	license := VerifyLicenseKey(premiumLicenseKey)

	// Главная страница с поддержкой мультиязычности
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" { http.NotFound(w, r); return }
		lang := r.URL.Query().Get("lang")
		if lang == "" { lang = "ru" }
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		_, _ = w.Write([]byte(GetDashboardHTML(html.EscapeString(license.ClientName), html.EscapeString(license.ServerID), lang)))
	})

	// Страница отчета сканирования периметра (считывает введенный домен)
http.HandleFunc("/report/mbank", func(w http.ResponseWriter, r *http.Request) {
    lang := r.URL.Query().Get("lang")
    if lang == "" { lang = "ru" }
    domainName := r.URL.Query().Get("domain")
    if domainName == "" { domainName = "unknown-node.com" }
    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    _, _ = w.Write([]byte(GetPerimeterReportHTML(lang, html.EscapeString(domainName))))
})

	// Страница биллинга
	http.HandleFunc("/billing", func(w http.ResponseWriter, r *http.Request) {
		lang := r.URL.Query().Get("lang")
		if lang == "" { lang = "ru" }
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		_, _ = w.Write([]byte(GetBillingPageHTML(html.EscapeString(license.ClientName), lang)))
	})

	// Обработчик логотипа
	http.HandleFunc("/static/logo.jpg", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./logo.jpg")
	})

	fmt.Println(">>> Глобальная мультиязычная платформа KvantumSafe SDK успешно запущена <<<")
	server := &http.Server{Addr: ":8080", ReadHeaderTimeout: 3 * time.Second}
	_ = server.ListenAndServe()
}
