package main

import (
	"encoding/base64"
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

	http.HandleFunc("/specification", func(w http.ResponseWriter, r *http.Request) {
		lang := r.URL.Query().Get("lang")
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		_, _ = w.Write([]byte(GetSpecificationPageHTML(lang)))
	})

	http.HandleFunc("/antihacker", func(w http.ResponseWriter, r *http.Request) {
		lang := r.URL.Query().Get("lang")
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		_, _ = w.Write([]byte(GetAntiHackerPageHTML(lang)))
	})

	http.HandleFunc("/download", func(w http.ResponseWriter, r *http.Request) {
		lang := r.URL.Query().Get("lang")
		w.Header().Set("Content-Disposition", "attachment; filename=KvantumSafe_Specification.txt")
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		
		doc := "OFFICIAL B2B COMMERCIAL OFFER\nCompany: Quantum Safe LLC (Osh, Kyrgyz Republic)\n\n" +
			"1. External Auditor Layer: Scanning perimeter on port 443.\n" +
			"2. Internal Vault Layer: Hardening files under POSIX 0600 strict policy.\n" +
			"3. Route Orchestrator Layer: Post-quantum Crypto-Agility dispatcher (NIST ML-KEM).\n\n" +
			"Annual Corporate SaaS License Cost: $35,000 USD."
			
		if lang == "ru" {
			doc = "ОФИЦИАЛЬНАЯ ТЕХНИЧЕСКАЯ СПЕЦИФИКАЦИЯ И КОРПОРАТИВНЫЙ B2B-ОФФЕР\nПравообладатель: ОсОО «Квантум Сейф» (г. Ош, КР)\n\n" +
				"1. Продукт 1. Внешний цифровой ревизор: Сканирует сетевые шлюзы банковских приложений и выявляет уязвимости OpenSSL.\n" +
				"2. Продукт 2. Внутренний невидимый сейф: Изолирует критические бэкапы баз данных под POSIX-права стандарта 0600.\n" +
				"3. Продукт 3. Транзитный диспетчер: Оркестрирует трансграничные переводы по постквантовым алгоритмам NIST ML-KEM.\n\n" +
				"Стоимость годовой корпоративной лицензии контура «Quantum Web3»: 35 000 долларов США."
		} else if lang == "kg" {
			doc = "ТЕХНИКАЛЫК СПЕЦИФИКАЦИЯСЫ ЖАНА КОРПОРАТИВДИК B2B СУНУШУ\nУкук ээси: «Квантум Сейф» ОсООсу (Кыргыз Республикасы)\n\n" +
				"1. Тышкы ревизор: Банктын тиркемелерин сырттан чабуулдардан коргойт.\n" +
				"2. Ички санариптик сейф: Сервердеги паролдорду заматта жашырып, катуу 0600 стандартына өткөрөт.\n" +
				"3. Акылдуу транзиттик диспетчер: Эл аралык которууларды жаңы посткванттык коопсуздук контейнерлерине (NIST ML-KEM) салып багыттайт.\n\n" +
				"Жылдык корпоративдик лицензиянын баасы: $35,000 АКШ доллары."
		}
		
		_, _ = w.Write([]byte(doc))
	})

	server := &http.Server{Addr: ":8080", ReadHeaderTimeout: 3 * time.Second}
	_ = server.ListenAndServe()
}
