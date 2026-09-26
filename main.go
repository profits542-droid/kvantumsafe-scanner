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
		info.ClientName = parts[0] // Индекс 0 - имя клиента
		info.ServerID = parts[2]   // Индекс 2 - ID сервера
		
		if expTime, err := time.Parse("2006-01-02", parts[1]); err == nil { // Индекс 1 - дата
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
		
		title := "Функционал Золотой Кнопки / Yellow Button Logic"
		desc := "Настоящий веб-интерфейс реализует функцию динамической генерации и криптографически безопасного депонирования технической документации. При активации триггера (нажатии на кнопку) ядро сервера на Go на лету определяет активную языковую локализацию клиентской сессии и формирует официальный защищенный файл спецификации (TXT) для предоставления ИТ-департаментам и комплаенс-контролю финансовых организаций.<br><br>Сгенерированный документ содержит полные архитектурные параметры платформы KvantumSafe Pro Framework, подробную разбивку коммерческой тарифной сетки (Global Scanner, Compliance Pro, Quantum Web3) и легальное обоснование юридической чистоты софта перед государственными регуляторами. Файл принудительно сохраняется в локальное хранилище (папку Загрузки) ноутбука или персонального компьютера пользователя по протоколу контентной диспетчеризации (Content-Disposition)."
		back := "← Назад / Back"
		
		if lang == "kg" {
			title = "Алтын Түйнөктүн Функционалы"
			desc = "Бул веб-интерфейс программалык камсыздоонун техникалык документтерин динамикалык түрдө генерациялоо жана криптографиялык коопсуз депонирлөө функциясын аткарат. Түймени басканда, Go тилиндеги сервердин ядросу кардардын сессиясынын тилин аныктайт жана өзгөчөлөнгөн тексттик документти (TXT) түзөт.<br><br>Документ KvantumSafe Pro Framework платформасынын толук архитектуралык параметрлерин, коммерциялык тарифтик торду (Global Scanner, Compliance Pro, Quantum Web3) жана мамлекеттик жөнгө салуучулардын алдында программалык камсыздоонун юридикалык тазалыгынын укуктук негиздемесин камтыйт. Файл автоматтык түрдө колдонуучунун ноутбугунун же компьютериний локалдык сактагычына (Жүктөлмөлөр папкасына) жүктөлөт."
			back = "← Артка"
		} else if lang == "ar" {
			title = "منطق الزر الأصفر العملي"
			desc = "تنفذ هذه الواجهة وظيفة التوليد الديناميكي والإيداع الآمن للمستندات الفنية. عند الضغط على الزر، تحدد نواة الخادم المكتوبة بلغة Go لغة جلسة العميل فورًا وتقوم بإنشاء مستند نصي مخصص (TXT) لتقديمه إلى أقسام تكنولوجيا المعلومات والامتثال في البنوك.<br><br>يحتوي المستند على معلمات الهيكل الكاملة لمنصة KvantumSafe Pro Framework، وتفاصيل مفصلة لشبكة الأسعار التجارية (Global Scanner, Compliance Pro, Quantum Web3) والمبررات القانونية لسلامة البرمجيات أمام الجهات التنظيمية الحكومية. يتم نقل الملف تلقائيًا إلى مجلد التنزيلات على جهاز الكمبيوتر الخاص للمستخدم عبر بروتوكول (Content-Disposition)."
			back = "← عودة"
		}
		
		htmlDoc := "<html><head><meta charset='UTF-8'><title>Specification</title></head><body style='font-family:sans-serif;padding:40px;line-height:1.6;max-width:850px;margin:auto;background:#f4f7f6;'><p><a href='/?lang=" + lang + "' style='font-weight:bold;text-decoration:none;color:#0a2540;'>" + back + "</a></p><h2 style='color:#0a2540; border-bottom:2px solid #cbd5e1; padding-bottom:10px;'>" + title + "</h2><p style='font-size:16px; color:#334155; text-align:justify;'>" + desc + "</p></body></html>"
		_, _ = w.Write([]byte(htmlDoc))
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
			doc = "ОФИЦИАЛЬНАЯ ТЕХНИЧЕСКАЯ СПЕЦИФИКАЦИЯ И ТАРИФНАЯ СЕТКА B2B-ОФФЕРА\nПравообладатель: Общество с ограниченной ответственностью «Квантум Сейф» (ОсОО «Квантум Сейф», г. Ош, КР)\n\n" +
				"ДОСТУПНЫЕ КОРПОРАТИВНЫЕ ЛИЦЕНЗИИ И СТОИМОСТЬ ПОДПИСКИ:\n\n" +
				"1. ТАРИФ «GLOBAL SCANNER» — БЕСПЛАТНО / $0\n" +
				"Базовый инструмент для экспресс-анализа внешних сетевых шлюзов ИТ-инфраструктуры организации.\n\n" +
				"2. ТАРИФ «COMPLIANCE PRO» — 15 000 долларов США / год\n" +
				"Включает в себя Продукт 1 (Внешний цифровой ревизор портов) и Продукт 2 (Внутренний сейф комплаенс-контроля серверов с принудительной изоляцией уязвимых бэкапов под POSIX-права 0600).\n\n" +
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
