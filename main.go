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
		info.ServerID = parts[1]
		if expTime, err := time.Parse("2006-01-02", parts[2]); err == nil {
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
		
		doc := "OFFICIAL B2B COMMERCIAL OFFER & PRICING\nCompany: Quantum Safe LLC (Osh, Kyrgyz Republic)\n\n" +
			"AVAILABLE SAAS LICENSES AND PRICING TIERS:\n\n" +
			"1. GLOBAL SCANNER — FREE / $0\n" +
			"Basic external vulnerability scanning of the corporate network infrastructure.\n\n" +
			"2. COMPLIANCE PRO — $15,000 / year\n" +
			"Includes Product 1 (External perimeter audit) + Product 2 (Internal server vault file protection under POSIX 0600 strict policy).\n\n" +
			"3. QUANTUM WEB3 (ENTERPRISE CORE) — $35,000 / year\n" +
			"Full software protection suite. Includes Product 1 (External Audit), Product 2 (Internal Vault), and Product 3 (Post-quantum Crypto-Agility route orchestrator for secure international transfers)."
			
		if lang == "ru" {
			doc = "ОФИЦИАЛЬНАЯ ТЕХНИЧЕСКАЯ СПЕЦИФИКАЦИЯ И ТАРИФНАЯ СЕТКА B2B-ОФФЕРА\nПравообладатель: ОсОО «Квантум Сейф» (г. Ош, Кыргызская Республика)\n\n" +
				"ДОСТУПНЫЕ КОРПОРАТИВНЫЕ ЛИЦЕНЗИИ И СТОИМОСТЬ ПОДПИСКИ:\n\n" +
				"1. ТАРИФ «GLOBAL SCANNER» — БЕСПЛАТНО / $0\n" +
				"Базовый инструмент для экспресс-анализа внешних сетевых шлюзов ИТ-инфраструктуры организации.\n\n" +
				"2. ТАРИФ «COMPLIANCE PRO» — 15 000 долларов США / год\n" +
				"Включает в себя Продукт 1 (Внешний цифровой ревизор портов) и Продукт 2 (Внутренний невидимый сейф комплаенс-контроля серверов с принудительной изоляцией уязвимых бэкапов под POSIX-права 0600).\n\n" +
				"3. ТАРИФ «QUANTUM WEB3» (МАКСИМАЛЬНАЯ БЕЗОПАСНОСТЬ) — 35 000 долларов США / год\n" +
				"Полный оборонный комплекс программного ядра Framework SDK. Включает в себя все три уровня защиты: Продукт 1 (Внешний аудит периметра), Продукт 2 (Внутренний сейф защиты файлов памяти) и Продукт 3 (Интеллектуальный транзитный диспетчер и постквантовый оркестратор трансграничных платежей NIST ML-KEM)."
		} else if lang == "kg" {
			doc = "ТЕХНИКАЛЫК СПЕЦИФИКАЦИЯСЫ ЖАНА ТАРИФТИК B2B СУНУШУ\nУкук ээси: «Квантум Сейф» ОсООсу (Ош ш., Кыргыз Республикасы)\n\n" +
				"ЖЕТКИЛИКТҮҮ КОРПОРАТИВДИК ЛИЦЕНЗИЯЛАР ЖАНА БААЛАРЫ:\n\n" +
				"1. «GLOBAL SCANNER» ТАРИФИ — БЕСПЛАТНО / $0\n" +
				"Уюмдун ИТ-инфраструктурасынын тышкы тармактык шлюздарын экспресс-анализдөө үчүн базалык курал.\n\n" +
				"2. «COMPLIANCE PRO» ТАРИФИ — $15,000 / жыл\n" +
				"Продукт 1 (Тышкы санариптик ревизор) жана Продукт 2 (Ички файлдык тутумду комплаенс-контролдоо жана файлдарды 0600 коопсуздук стандартына которуу) кызматтарын камтыйт.\n\n" +
				"3. «QUANTUM WEB3» ТАРИФИ (МАКСИМАЛДУУ КОРГОО) — $35,000 / жыл\n" +
				"Программалык камсыздоонун толук коргонуу комплекси. Курамына Продукт 1 (Тышкы аудит), Продукт 2 (Ички сейф) жана Продукт 3 (Трансчегаралык төлөмдөрдү жаңы посткванттык коопсуздук контейнерлерине NIST ML-KEM салып коопсуз багыттоочу акылдуу диспетчер) кызматтары толугу менен кирет."
		}
		
		_, _ = w.Write([]byte(doc))
	})

	server := &http.Server{Addr: ":8080", ReadHeaderTimeout: 3 * time.Second}
	_ = server.ListenAndServe()
}
