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
		
		htmlPage := GetDashboardHTML(html.EscapeString(license.ClientName), html.EscapeString(license.ServerID), lang)
		
		direction := "right: 25px;"
		if lang == "ar" { direction = "left: 25px;" }
		
		// БАЗОВЫЕ НАСТРОЙКИ ПЕРЕМЕННЫХ НА РУССКОМ ЯЗЫКЕ (DEFAULT)
		welcomeMessage := "Здравствуйте! Я ИИ-консультант ОсОО «Квантум Сейф». Если у вас есть вопросы по нашему ПО, стандартам NIST или азиатским алгоритмам GmSSL, я готов помочь!"
		placeholderText := "Задать вопрос ИИ..."
		ansNist := "Ядро KvantumSafe Pro оркестрирует постквантовые алгоритмы решеток стандарта <b>NIST ML-KEM</b> и суверенные азиатские криптопротоколы <b>GmSSL (SM4-GCM)</b>. System автоматически выбирает оптимальный маршрут данных, исключая риски дешифрования транзакций хакерами."
		ansMeet := "Отличное решение! Наш Генеральный директор готов провести личную техническую презентацию контура безопасности. Оставьте ваши контакты, и мы зафиксируем удобное время встречи."
		ansDefault := "Платформа, разработанная ОсОО «Квантум Сейф» на языке Go, обеспечивает автоматический комплаенс-контроль, сканирование внешнего периметра (порт 443) и фоновую изоляцию файлов под права POSIX 0600. Программа юридически чиста и не требует лицензий СКЗИ."

		// ДИНАМИЧЕСКАЯ ПОДСТАНОВКА ЛОКАЛИЗАЦИЙ ДЛЯ РОБОТА
		if lang == "kg" {
			welcomeMessage = "Саламатсызбы! Мен ОсОО «Квантум Сейф» ИИ-консультантымын. Программалык камсыздоо, NIST посткванттык стандарттары же азиялык GmSSL алгоритмдери боюнча суроолоруңуз болсо, берсеңиз болот."
			placeholderText = "Текст жазыңыз..."
			ansNist = "KvantumSafe Pro ядросу <b>NIST ML-KEM</b> посткванттык алгоритмдерин жана азиялык <b>GmSSL (SM4-GCM)</b> криптопротоколдорун оркестрациялайт. Система транзакцияларды хакерлерден ишенимдүү коргойт."
			ansMeet = "Абдан жакшы чечим! Биздин Башкы директорубуз коопсуздук контуру боюнча сизге жеке бетме-бет презентация өткөрүүгө даяр. Байланыш маалыматыңызды калтырыңыз, биз жолугушуу убактысын белгилейбиз."
			ansDefault = "«Квантум Сейф» ОсОО тарабынан Go тилинде иштелип чыккан платформа автоматтык комплаенс-контролду, сырткы периметрди сканерлөөнү (443-порт) жана файлдарды 0600 стандартына жашыруун которууну камсыз кылат."
		} else if lang == "en" {
			welcomeMessage = "Hello! I am the AI Assistant of Quantum Safe LLC. Feel free to ask any questions about our software, NIST post-quantum standards, or Asian GmSSL protocols."
			placeholderText = "Ask AI Assistant..."
			ansNist = "The KvantumSafe Pro core orchestrates next-generation post-quantum lattice algorithms (<b>NIST ML-KEM</b>) and Asian sovereign protocols (<b>GmSSL SM4-GCM</b>), ensuring absolute transaction safety."
			ansMeet = "Excellent choice! Our General Director is ready to conduct a personal technical presentation of our security framework. Please leave your contact details to schedule a B2B meeting."
			ansDefault = "The platform developed by Quantum Safe LLC in Go provides automated compliance control, external perimeter audits (port 443), and stealth file isolation under strict POSIX 0600 tokens."
		} else if lang == "ar" {
			welcomeMessage = "مرحباً! أنا المستشار الذكي لشركة 'كوانتوم سيف'. لا تتردد في طرح أي أسئلة حول برامجنا أو معايير NIST بعد العصر الكمي أو بروتوكول GmSSL الآسيوي."
			placeholderText = "اسأل المستشار الذكي..."
			ansNist = "تنسق نواة KvantumSafe Pro خوارزميات الشبكة لما بعد العصر الكمي المعتمدة من قبل <b>NIST ML-KEM</b> وبروتوكولات <b>GmSSL (SM4-GCM)</b> الآسيوية السيادية لحماية البيانات."
			ansMeet = "قرار ممتاز! مديرنا العام مستعد لإجراء عرض فني شخصي لمنظومة الأمان. يرجى ترك معلومات الاتصال الخاصة بك لتحديد موعد الاجتماع."
			ansDefault = "توفر المنصة التي طورتها شركة 'كوانتوم سيف' بلغة Go رقابة تلقائية على الامتثال، وفحص المحيط الخارجي (المنفذ 443) وعزل الملفات تحت صلاحيات POSIX 0600 الصارمة."
		} else if lang == "kz" {
			welcomeMessage = "Сәлеметсіз бе! Мен «Квантум Сейф» ЖШС ИИ-консультантымын. Бағдарламалық құрал, NIST немесе GmSSL алгоритмдері туралы сұрақтарыңыз болса, қоя аласыз."
			placeholderText = "Сұрақ қою..."
			ansNist = "KvantumSafe Pro ядросы <b>NIST ML-KEM</b> посткванттық тор алгоритмдерін және <b>GmSSL (SM4-GCM)</b> азиялық егеменді криптопротоколдарын үйлестіреді. Жүйе транзакцияларды хакерлерден сенімді қорғайды."
			ansMeet = "Өте жақсы шешім! Біздің Бас директорымыз қауіпсіздік контуры бойынша сізге жеке техникалық таныстырылым өткізуге дайын. Байланыс мәліметтеріңізді қалдырыңыз, біз кездесу уақытын белгілейміз."
			ansDefault = "«Квантум Сейф» ЖШС Go тілінде әзірлеген платформа автоматты комплаенс-бақылауды, сырткы периметрді сканерлеуді (443-порт) және файлдарды 0600 стандартына жасырын ауыстыруды қамтамасыз етеді."
		}

		botWidget := `
	<div style='position: fixed; bottom: 25px; ` + direction + ` width: 65px; height: 65px; background: #7c3aed; color: white; border-radius: 50%; display: flex; align-items: center; justify-content: center; font-size: 28px; cursor: pointer; box-shadow: 0 4px 16px rgba(124,58,237,0.4); z-index: 1000; transition: 0.3s;' onclick='toggleChat()'>🤖</div>
	<div style='position: fixed; bottom: 100px; ` + direction + ` width: 370px; height: 480px; background: white; border-radius: 12px; box-shadow: 0 8px 32px rgba(0,0,0,0.15); display: none; flex-direction: column; z-index: 1000; overflow: hidden; border: 1px solid #e2e8f0; font-size: 14px;' id='chatWindow'>
		<div style='background: #0a2540; color: white; padding: 15px; font-weight: bold; display: flex; justify-content: space-between; align-items: center;'>
			<div>KvantumSafe AI Assistant <span style='font-size: 12px; background: #16a34a; padding: 2px 6px; border-radius: 4px; margin-left: 10px;'>Online</span></div>
			<div style='cursor:pointer;' onclick='toggleChat()'>✕</div>
		</div>
		<div style='flex: 1; padding: 15px; overflow-y: auto; background: #f8fafc; display: flex; flex-direction: column; gap: 10px;' id='chatBody'>
			<div style='max-width: 80%; padding: 10px 14px; border-radius: 8px; line-height: 1.4; background: #e2e8f0; color: #1e293b; align-self: flex-start;'>` + welcomeMessage + `</div>
		</div>
		<div style='padding: 10px; border-top: 1px solid #e2e8f0; display: flex; background: white;'>
			<input type='text' style='flex: 1; border: none; padding: 10px; outline: none; font-size: 14px;' id='chatInput' placeholder='` + placeholderText + `' onkeypress='handleKey(event)'>
			<button style='background: #0a2540; color: white; border: none; padding: 0 20px; font-weight: bold; cursor: pointer;' onclick='sendMessage()'>&gt;</button>
		</div>
	</div>
	<script>
		function toggleChat() {
			var win = document.getElementById("chatWindow");
			win.style.display = (win.style.display === "flex") ? "none" : "flex";
		}
		function handleKey(e) {
			if (e.key === "Enter") sendMessage();
		}
		function sendMessage() {
			var input = document.getElementById("chatInput");
			var text = input.value.trim();
			if (!text) return;
			var body = document.getElementById("chatBody");
			var userMsg = document.createElement("div");
			userMsg.style = "max-width: 80%; padding: 10px 14px; border-radius: 8px; line-height: 1.4; background: #7c3aed; color: white; align-self: flex-end;";
			userMsg.innerText = text;
			body.appendChild(userMsg);
			input.value = "";
			body.scrollTop = body.scrollHeight;
			setTimeout(function() {
				var botMsg = document.createElement("div");
				botMsg.style = "max-width: 80%; padding: 10px 14px; border-radius: 8px; line-height: 1.4; background: #e2e8f0; color: #1e293b; align-self: flex-start;";
				var lowText = text.toLowerCase();
				if (lowText.includes("nist") || lowText.includes("алгоритм") || lowText.includes("квант") || lowText.includes("algorithm")) {
					botMsg.innerHTML = "` + ansNist + `";
				} else if (lowText.includes("встреч") || lowText.includes("купить") || lowText.includes("цена") || lowText.includes("meet") || lowText.includes("buy")) {
					botMsg.innerHTML = "` + ansMeet + `";
				} else {
					botMsg.innerHTML = "` + ansDefault + `";
				}
				body.appendChild(botMsg);
				body.scrollTop = body.scrollHeight;
			}, 800);
		}
	</script>
</body>
</html>`
		
		htmlPage = strings.Replace(htmlPage, "</body>\n</html>", botWidget, 1)
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		_, _ = w.Write([]byte(htmlPage))
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

