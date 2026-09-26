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
	if err != nil { return info }
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

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" { http.NotFound(w, r); return }
		lang := r.URL.Query().Get("lang")
		if lang == "" { lang = "kg" }
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		_, _ = w.Write([]byte(GetDashboardHTML(html.EscapeString(license.ClientName), html.EscapeString(license.ServerID), lang)))
	})

	http.HandleFunc("/report/mbank", func(w http.ResponseWriter, r *http.Request) {
		lang := r.URL.Query().Get("lang")
		if lang == "" { lang = "kg" }
		domainName := r.URL.Query().Get("domain")
		if domainName == "" { domainName = "unknown-node.com" }
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		_, _ = w.Write([]byte(GetPerimeterReportHTML(lang, html.EscapeString(domainName))))
	})

	// Страница описания золотой кнопки (Вшита прямо в ядро)
	http.HandleFunc("/specification", func(w http.ResponseWriter, r *http.Request) {
		lang := r.URL.Query().Get("lang")
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		htmlDoc := "<html><head><meta charset='UTF-8'><title>Specification Info</title></head><body style='font-family:sans-serif;padding:40px;line-height:1.6;max-width:800px;margin:auto;'><p><a href='/?lang=" + lang + "'>← Назад / Back</a></p><h2>Функционал Золотой Кнопки / Yellow Button Logic</h2><p>При нажатии на главную кнопку система автоматически генерирует официальный защищенный файл спецификации софта (TXT) для предоставления ИТ-департаментам банков.</p></body></html>"
		_, _ = w.Write([]byte(htmlDoc))
	})

	// Страница презентации АнтиХакер AI (Вшита прямо в ядро)
	http.HandleFunc("/antihacker", func(w http.ResponseWriter, r *http.Request) {
		lang := r.URL.Query().Get("lang")
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		htmlDoc := "<html><head><meta charset='UTF-8'><title>KvantumSafe AI-AntiHacker</title></head><body style='font-family:sans-serif;padding:40px;line-height:1.6;max-width:800px;margin:auto;background:#0a2540;color:white;'><p><a href='/?lang=" + lang + "' style='color:#cbd5e1;'>← Назад / Back</a></p><h2 style='color:#d4af37;'>🤖 KvantumSafe AI-AntiHacker (Guard Engine)</h2><p><b>Статус разработки:</b> Автономный оборонный комплекс нового поколения от ОсОО «Квантум Сейф».</p><p>ПО построено на гибридном стеке <b>Go + Rust</b> с привлечением локальных самообучающихся ИИ-моделей TinyML. Комплекс осуществляет низкоуровневый контроль оперативной памяти, выполняет Device Fingerprinting и блокирует любые попытки логического фрода транзакций в реальном времени.</p></body></html>"
		_, _ = w.Write([]byte(htmlDoc))
	})

	http.HandleFunc("/download", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Disposition", "attachment; filename=KvantumSafe_Specification.txt")
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		englishDoc := "TECHNICAL SPECIFICATION & B2B COMMERCIAL OFFER\n" +
			"Company: Kvantum Safe LLC\n\n" +
			"KvantumSafe Pro functions as an intelligent network coordinator. " +
			"The software does not independently develop cryptographic algorithms and is not an encryption tool."
		_, _ = w.Write([]byte(englishDoc))
	})

	server := &http.Server{Addr: ":8080", ReadHeaderTimeout: 3 * time.Second}
	_ = server.ListenAndServe()
}
