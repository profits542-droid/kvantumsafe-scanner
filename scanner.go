package main

func getTranslation(key, lang string) string {
	translations := map[string]map[string]string{
		"nav_title": {
			"ru": "🛡️ Международная ИТ-Платформа KvantumSafe Pro",
			"en": "🛡️ KvantumSafe Pro International IT Platform",
			"kz": "🛡️ KvantumSafe Pro Халықаралық ИТ-Платформасы",
			"ar": "🛡️ منصة كيفانتوم سيف برو الدولية لتكنولوجيا المعلومات",
		},
		"nav_sub": {
			"ru": "Автоматическая проверка сетевой безопасности и умное управление защитой данных",
			"en": "Automatic network security audit and smart data protection management",
			"kz": "Желілік қауіпсіздікті автоматты түрде тексеру және деректерді қорғауды ақылды басқару",
			"ar": "التدقيق الآلي لأمن الشبكات والإدارة الذكية لحماية البيانات",
		},
		"scan_title": {
			"ru": "🔍 Проверить безопасность сайта / Check Website Security",
			"en": "🔍 Check Website Security / Verify Perimeter",
			"kz": "🔍 Желілік қауіпсіздікті тексеру / Периметрді аудиттеу",
			"ar": "🔍 تحقق من أمن الموقع / تدقيق محيط الشبكة",
		},
		"scan_btn": {
			"ru": "Сканировать",
			"en": "Scan Now",
			"kz": "Сканерлеу",
			"ar": "فحص الآن",
		},
		"prod_title": {
			"ru": "Уникальность и ключевые продукты платформы:",
			"en": "Uniqueness and Key Platform Products:",
			"kz": "Платформаның бірегейлігі мен негізгі өнімдері:",
			"ar": "التفرد والمنتجات الرئيسية للمنصة:",
		},
		"p1_title": {
			"ru": "🌐 Продукт 1. Автоматический аудит периметра (Защита данных в пути)",
			"en": "🌐 Product 1. Automatic Perimeter Audit (Data in Transit Protection)",
			"kz": "🌐 Өнім 1. Периметрді автоматты аудиттеу (Транзиттегі деректерді қорғау)",
			"ar": "🌐 المنتج 1. التدقيق الآلي لمحيط الشبكة (حماية البيانات أثناء الانتقال)",
		},
		"p1_desc": {
			"ru": "Программа выполняет роль бдительного цифрового ревизора. Она сканирует внешние порты системы, проверяет сетевые шлюзы приложений и мгновенно выявляет уязвимости, предотвращая перехват данных снаружи.",
			"en": "The software acts as a vigilant digital auditor. It scans external system ports, verifies application gateways, and instantly identifies vulnerabilities, preventing data interception from the outside.",
			"kz": "Бағдарлама қырағы цифрлық ревизор рөлін атқарады. Ол жүйенің сыртқы порттарын сканерлейді, қосымшалардың желілік шлюздерін тексереді және осалдықтарды дереу анықтап, деректерді сырттан ұстап алудың алдын алады.",
			"ar": "يعمل البرنامج كمدقق رقمي يقظ. يقوم بفحص منافذ النظام الخارجية، والتحقق من بوابات تطبيقات الشبكة، والتعرف الفوري على الثغرات الأمنية، مما يمنع اعتراض البيانات من الخارج.",
		},
		"p2_title": {
			"ru": "🗄️ Продукт 2. Внутренний комплаенс-контроль серверов (Защита сохраненных данных)",
			"en": "🗄️ Product 2. Internal Server Compliance Control (Data at Rest Protection)",
			"kz": "🗄️ Өнім 2. Серверлерді ішкі комплаенс-бақылау (Сақталған деректерді қорғау)",
			"ar": "🗄️ المنتج 2. الرقابة الداخلية لامتثال الخوادم (حماية البيانات المخزنة)",
		},
		"p2_desc": {
			"ru": "Модуль проводит тотальную ревизию файловой системы серверов. При обнаружении критических ошибок сотрудников, KvantumSafe автоматически изолирует угрозу, присваивая файлам права доступа банковского стандарта 0600.",
			"en": "The module performs a total audit of the server file system. Upon detecting critical employee errors, KvantumSafe automatically isolates the threat, assigning strict banking-standard 0600 access rights.",
			"kz": "Модуль серверлердің файлдық жүйесіне толық ревизия жүргізеді. Қызметкерлердің қателері анықталған кезде, KvantumSafe файлдарға банктік стандарттағы 0600 қатаң қолжетімділік құқықтарын бере отырып, қауіпті автоматты түрде оқшаулайды.",
			"ar": "يقوم الموديل بإجراء تدقيق شامل لنظام ملفات الخادم. عند اكتشاف أخطاء حرجة من الموظفين، يقوم البرنامج تلقائيًا وعزل التهديد بصلاحيات 0600 الصارمة.",
		},
		"p3_title": {
			"ru": "🔀 Продукт 3. Постквантовый координатор маршрутов (Crypto-Agility)",
			"en": "🔀 Product 3. Post-Quantum Route Coordinator (Crypto-Agility)",
			"kz": "🔀 Өнім 3. Посткванттық бағыт үйлестірушісі (Crypto-Agility)",
			"ar": "🔀 المنتج 3. منسق المسارات بعد العصر الكمي (Crypto-Agility)",
		},
		"p3_desc": {
			"ru": "Флагманское решение для трансграничных переводов. Наша система выступает как умный транзитный диспетчер — она упаковывает потоки данных в защищенные контейнеры нового поколения по международным стандартам решеток (NIST ML-KEM).",
			"en": "A flagship solution for cross-border transfers. Our system acts as a smart transit dispatcher — it packs data streams into next-generation secure containers using international lattice standards (NIST ML-KEM).",
			"kz": "Трансшекаралық аударымдарға арналған флагмандық шешім. Біздің жүйе ақылды транзиттік диспетчер ретінде әрекет етеді — ол торлардың халықаралық стандарттары (NIST ML-KEM) бойынша деректер ағынын жаңа буынның қорғалған контейнерлеріне жинақтайды.",
			"ar": "الحل الرائد للتحويلات عبر الحدود. يعمل نظامنا كمنسق نقل ذكي، حيث يدمج تدفقات البيانات في حاويات آمنة من الجيل الجديد وفقًا لمعايير الشبكة الدولية (NIST ML-KEM).",
		},
		"legal_title": {
			"ru": "⚖️ Юридическая чистота",
			"en": "⚖️ Legal Purity",
			"kz": "⚖️ Заңды тазалық",
			"ar": "⚖️ النزاهة القانونية",
		},
		"legal_desc": {
			"ru": "Платформа не осуществляет самостоятельную разработку криптографических алгоритмов и не является средством шифрования. Программа лишь оркестрирует и управляет теми защитными модулями, которые уже сертифицированы и встроены в серверное оборудование вашей организации. Это полностью снимает любые вопросы контролирующих и проверяющих органов касательно лицензирования.",
			"en": "The platform does not independently develop cryptographic algorithms and is not an encryption tool. The program only orchestrates and manages protective modules that are already certified and embedded in your organization's server hardware. This completely eliminates any questions from regulatory and checking authorities regarding licensing.",
			"kz": "Платформа криптографиялық алгоритмдерді дербес әзірлеуді жүзеге асырмайды және шифрлау құралы болып табылмайды. Бағдарлама тек ұйымыңыздың серверлік жабдығына ендірілген және сертификатталған қорғаныс модульдерін басқарады. Бұл бақылаушы және тексеруші органдардың арнайы лицензиялауға қатысты кез келген сұрақтарын толығымен алып тастайды.",
			"ar": "المنصة لا تطور خوارزميات التشفير بشكل مستقل وليست أداة تشفير. يقوم البرنامج فقط بتنسيق وإدارة وحدات الحماية المعتمدة بالفعل والمدمجة في أجهزة خادم مؤسستك. هذا يزيل تمامًا أي أسئلة من السلطات الرقابية والتدقيقية فيما يتعلق بالترخيص.",
		},
	}
	return translations[key][lang]
}

func GetDashboardHTML(clientName, serverID, lang string) string {
	direction := "ltr"
	if lang == "ar" { direction = "rtl" }

	return `<!DOCTYPE html>
<html lang='` + lang + `' dir='` + direction + `'>
<head>
	<meta charset='UTF-8'>
	<title>KvantumSafe Pro</title>
	<style>
		body { font-family: 'Segoe UI', sans-serif; background-color: #f4f7f6; color: #333; margin: 0; padding: 0; line-height: 1.6; }
		.navbar { background-color: #0a2540; color: white; padding: 35px 40px; text-align: center; box-shadow: 0 4px 10px rgba(0,0,0,0.1); }
		.navbar h1 { margin: 0; font-size: 28px; font-weight: 600; }
		.navbar p { margin: 8px 0 15px 0; color: #cbd5e1; font-size: 15px; }
		.lang-switcher { display: flex; justify-content: center; gap: 12px; margin-top: 15px; }
		.lang-switcher a { color: #cbd5e1; text-decoration: none; font-weight: bold; font-size: 14px; background: rgba(255,255,255,0.15); padding: 6px 14px; border-radius: 6px; }
		.lang-switcher a:hover { color: white; background: #004d40; }
		.content { max-width: 850px; background: white; margin: 40px auto; padding: 40px; border-radius: 12px; box-shadow: 0 4px 25px rgba(0,0,0,0.05); }
		.scan-container { background: #e0f2f1; padding: 35px; border-radius: 10px; border: 2px dashed #004d40; text-align: center; margin: 20px 0; }
		.scan-input { width: 60%; padding: 14px; font-size: 16px; border: 1px solid #cbd5e1; border-radius: 6px; margin-right: 10px; outline: none; }
		.scan-btn { background: #004d40; color: white; padding: 14px 30px; font-size: 16px; border: none; border-radius: 6px; font-weight: bold; cursor: pointer; }
		.footer { text-align: center; margin-top: 40px; color: #64748b; font-size: 14px; border-top: 1px solid #e2e8f0; padding-top: 20px; }
		.tariff-table { width: 100%; border-collapse: collapse; margin-top: 25px; text-align: center; }
		.tariff-table th, .tariff-table td { padding: 15px; border: 1px solid #cbd5e1; font-size: 14px; }
		.feature-card { background: #f8fafc; padding: 20px; border-left: 4px solid #d4af37; margin-bottom: 15px; border-radius: 4px; text-align: ` + leftOrRight(direction) + `; }
		html[dir="rtl"] .feature-card { border-left: none; border-right: 4px solid #d4af37; }
		.legal-block { text-align: ` + leftOrRight(direction) + `; }
	</style>
</head>
<body>
	<div class='navbar'>
		<h1>` + getTranslation("nav_title", lang) + `</h1>
		<p>` + getTranslation("nav_sub", lang) + `</p>
		<div class='lang-switcher'>
			<a href='/?lang=ru'>RU</a>
			<a href='/?lang=en'>EN</a>
			<a href='/?lang=kz'>KZ</a>
			<a href='/?lang=ar'>AR</a>
		</div>
	</div>
	<div class='content'>
		<div class='scan-container'>
			<h3>` + getTranslation("scan_title", lang) + `</h3>
			<form action='/report/mbank?lang=` + lang + `' method='GET'>
				<input type='text' name='domain' class='scan-input' placeholder='domain.com...'>
				<button type='submit' class='scan-btn'>` + getTranslation("scan_btn", lang) + `</button>
			</form>
		</div>
		
		<h2 style='color:#0a2540; border-bottom:2px solid #e2e8f0; padding-bottom:10px;'>💰 SaaS Лицензии / Licenses</h2>
		<table class='tariff-table'>
			<tr style='background:#f8fafc;'>
				<th><b>Global Scanner</b><br><span style='color:#16a34a;'>FREE / $0</span></th>
				<th><b>Compliance Pro</b><br><span>$15,000 / год</span></th>
				<th style='background:#faf5ff;'><b>Quantum Web3</b><br><span style='color:#7c3aed;'>$35,000 / год</span></th>
			</tr>
		</table>
		
