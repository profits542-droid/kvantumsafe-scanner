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

	// СТРАНИЦА: Подробное техническое описание функционала золотой кнопки
	http.HandleFunc("/specification", func(w http.ResponseWriter, r *http.Request) {
		lang := r.URL.Query().Get("lang")
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		
		title := "Технический регламент золотой кнопки / Yellow Button Specification"
		desc := `Данный веб-интерфейс реализует функцию динамической генерации и криптографически безопасного депонирования технической документации программного обеспечения. При активации триггера (нажатии на кнопку) ядро сервера на Go на лету определяет активную локализацию клиентской сессии и формирует кастомизированный текстовый документ. Документ содержит полные архитектурные параметры платформы KvantumSafe Pro Framework, включая описание работы с правами POSIX стандарта 0600, механизмы Crypto-Agility и координацию внешних сертифицированных модулей аппаратной безопасности (HSM). Сформированный файл принудительно передается в локальное хранилище (папку Загрузки) ноутбука или персонального компьютера пользователя по протоколу контентной диспетчеризации (Content-Disposition).`
		back := "← Назад на главную / Back to Home"
		
		if lang == "kg" {
			title = "Алтын түйменин техникалык регламенти"
			desc = `Бул веб-интерфейс программалык камсыздоонун техникалык документтерин динамикалык түрдө генерациялоо жана криптографиялык коопсуз депонирлөө функциясын аткарат. Түймени басканда, Go тилиндеги сервердин ядросу кардардын сессиясынын тилин аныктайт жана өзгөчөлөнгөн тексттик документти түзөт. Документ KvantumSafe Pro Framework платформасынын толук архитектуралык параметрлерин камтыйт, анын ичинде POSIX 0600 стандартындагы укуктар менен иштөө, Crypto-Agility механизмдери жана сырткы тастыкталган аппараттык коопсуздук модулдарын (HSM) координациялоо сүрөттөлгөн. Түзүлгөн файл колдонуучунун ноутбугунун же компьютеринин локалдык сактагычына (Жүктөлмөлөр папкасына) автоматтык түрдө жүктөлөт.`
			back = "← Башкы бетке артка"
		} else if lang == "ar" {
			title = "اللوائح الفنية للزر الأصفر الرقمي"
			desc = `تنفذ هذه الواجهة وظيفة التوليد الديناميكي والإيداع الآمن للمستندات الفنية للبرمجيات. عند الضغط على الزر، تحدد نواة الخادم المكتوبة بلغة Go لغة جلسة العميل فورًا وتقوم بإنشاء مستند نصي مخصص. يحتوي المستند на معلمات الهيكل الكاملة لمنصة KvantumSafe Pro Framework، بما في ذلك تفاصيل العمل مع أذونات ملفات نظام POSIX 0600 الصارمة، وآليات تكييف التشفير (Crypto-Agility)، وتنسيق وحدات أمان الأجهزة الخارجية المعتمدة (HSM). يتم نقل الملف الذي تم إنشاؤه قسريًا إلى المجلد المحلي (التنزيلات) على الكمبيوتر المحمول أو الجهاز الشخصي للمستخدم عبر بروتوكول التوزيع (Content-Disposition).`
			back = "← العودة إلى الصفحة الرئيسية"
		}
		
		htmlDoc := "<html><head><meta charset='UTF-8'><title>Specification Info</title></head><body style='font-family:sans-serif;padding:40px;line-height:1.6;max-width:850px;margin:auto;background:#f4f7f6;'><p><a href='/?lang=" + lang + "' style='color:#0a2540;font-weight:bold;text-decoration:none;'>" + back + "</a></p><h2 style='color:#0a2540; border-bottom:2px solid #cbd5e1; padding-bottom:10px;'>" + title + "</h2><p style='font-size:16px; color:#334155; text-align:justify;'>" + desc + "</p></body></html>"
		_, _ = w.Write([]byte(htmlDoc))
	})

	// СТРАНИЦА: Высокотехнологичная презентация автономного продукта АнтиХакер AI
	http.HandleFunc("/antihacker", func(w http.ResponseWriter, r *http.Request) {
		lang := r.URL.Query().Get("lang")
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		
		title := "🤖 Автономный комплекс KvantumSafe AI-AntiHacker (Guard Engine)"
		status := "<b>Статус программного продукта:</b> Самостоятельное оборонное ИТ-решение корпоративного класса от ОсОО «Квантум Сейф»."
		tech := `Программный комплекс построен на уникальном гибридном архитектурном стеке <b>Go + Rust</b>, что обеспечивает максимальную сетевую масштабируемость и абсолютную безопасность оперативной памяти (Memory Safety). В систему интегрированы локальные самообучающиеся нейросетевые модели класса TinyML, которые непрерывно анализируют последовательность системных вызовов и логику банковских пакетов (ISO 20022). <br><br>
		<b>Ключевые защитные барьеры комплекса:</b><br>
		1. <b>Device Fingerprinting (Rust Lower Layer):</b> Нативный модуль на Rust обращается напрямую к регистрам процессора и считывает аппаратный UUID материнской платы, серийные номера процессора и MAC-адреса. Система блокирует попытки кражи сессий, даже если хакер полностью завладел логином и паролем сотрудника банка, но пытается войти с чужого ноутбука.<br>
		2. <b>AI Anomaly Detector:</b> Интеллектуальный фильтр мгновенно выявляет роботизированную активность хакерских скриптов, блокируя атаки класса Reentrancy и Flash-Loan, предотвращая моментальный несанкционированный вывод крупных сумм.<br>
		3. <b>Stealth-Изоляция:</b> При фиксации критической скорости подбора паролей программа мгновенно скрывает уязвимые файлы баз данных, переводя их под жесткую маску прав доступа стандарта 0600. ПО функционирует полностью автономно в режиме On-Premise внутри закрытого ИТ-контура организации.`
		back := "← Назад на главную / Back to Home"
		
		if lang == "kg" {
			title = "🤖 KvantumSafe AI-AntiHacker (Guard Engine) автономдуу комплекси"
			status = "<b>Программалык өнүмдүн статусу:</b> «Квантум Сейф» ОсООсунун корпоративдик деңгээлдеги көз карандысыз коргонуу ИТ-чечими."
			tech = `Программалык комплекс <b>Go + Rust</b> гибриддик архитектуралык стегинде курулган, бул максималдуу тармактык масштабдуулукту жана оперативдүү эстутумдун абсолюттук коопсуздугун (Memory Safety) камсыз кылат. Система курамына TinyML классындагы локалдык өзү үйрөнүүчү нейротармак моделдерин камтыйт, алар банктык пакеттердин (ISO 20022) логикасын тынымсыз талдап турат.<br><br>
			<b>Комплекстин негизги коргоо барьерлери:</b><br>
			1. <b>Device Fingerprinting (Rust Lower Layer):</b> Rust тилиндеги нативдик модуль процессордун регистрлерине түз кайрылып, энелик платанын аппараттык UUIDин, процессордун сериалдык номерлерин окуйт. Хакер банк кызматкеринин логин-паролун уурдап алса дагы, бөтөн ноутбуктан кирүүгө аракет кылганда система аны бөгөттөйт.<br>
			2. <b>AI Anomaly Detector:</b> Акылдуу чыпка хакердик скрипттердин роботтоштурулган активдүүлүгүн заматта аныктап, Reentrancy жана Flash-Loan чабуулдарын токтотот.<br>
			3. <b>Stealth-Изоляция:</b> Чабуул катталганда программа маалымат базасынын файлдарын заматта жашырып, 0600 стандартындагы укуктарга өткөрөт. Программа уюмдун жабык ИТ-контурунда On-Premise режиминде толугу менен автономдуу иштейт.`
			back = "← Башкы бетке артка"
		} else if lang == "ar" {
			title = "🤖 المجمع الدفاعي المستقل KvantumSafe AI-AntiHacker"
			status = "<b>حالة المنتج البرمجي:</b> حل دفاعي مستقل لتكنولوجيا المعلومات من فئة الشركات من شركة ذ.م.م 'كوانتوم سيف'."
			tech = `تم بناء المجمع البرمجي على بنية برمجية هجينة فريدة من نوعها <b>Go + Rust</b>، مما يضمن أقصى قدر من القابلية للتوسع في الشبكة والأمان المطلق للذاكرة العشوائية (Memory Safety). تدمج المنظومة نماذج الشبكات العصبية المحلية ذاتية التعلم من فئة TinyML، والتي تحلل باستمرار تسلسل استدعاءات النظام ومنطق المعاملات المصرفية (ISO 20022).<br><br>
			<b>الحواجز الأمنية الرئيسية للمجمع:</b><br>
			1. <b>بصمة الأجهزة (Rust Lower Layer):</b> تتصل وحدة Rust الأصلية مباشرة بسجلات المعالج وتقرأ معرف UUID اللوحي ومعرفات المعالج وعناوين MAC. تمنع المنظومة محاولات سرقة الجلسات حتى لو استولى المخترق بالكامل على اسم المستخدم وكلمة المرور لموظف البنك ولكنه يحاول الدخول من كمبيوتر محمول غريب.<br>
			2. <b>كاشف الشذوذ بالذكاء الاصطناعي:</b> يكتشف الفلتر الذكي على الفور الأنشطة الروبوتية للنصوص البرمجية الخبيثة، مما يحظر هجمات Reentrancy وFlash-Loan المتقدمة.<br>
			3. <b>عزل Stealth الفوري:</b> عند رصد هجوم، يخفي البرنامج على الفور ملفات قواعد البيانات الحساسة، ويحولها إلى نظام صلاحيات POSIX 0600 الصارم. يعمل البرنامج بشكل مستقل تمامًا بنظام On-Premise داخل المحيط المغلق للمؤسسة.`
			back = "← العودة إلى الصفحة الرئيسية"
		}
		
