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
			"kz": "Желілік қауіпсіздікті автоматты түрде тексеру и деректерді қорғауды ақылды басқару",
			"ar": "التدقيق الآلي لأمن الشبكات والإدارة الذكية لحماية البيانات",
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
		.feature-card { background: #f8fafc; padding: 20px; border-left: 4px solid #d4af37; margin-bottom: 15px; border-radius: 4px; text-align: left; }
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
			<h3>🔍 Проверить безопасность сайта / Check Website Security</h3>
			<form action='/report/mbank?lang=` + lang + `' method='GET'>
				<input type='text' name='domain' class='scan-input' placeholder='domain.com...'>
				<button type='submit' class='scan-btn'>Сканировать</button>
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
		
		<h2 style='color:#0a2540; border-bottom:2px solid #e2e8f0; padding-bottom:10px; margin-top:30px;'>Уникальность и ключевые продукты платформы:</h2>
		<div class='feature-card'>
			<div style='font-weight:bold; color:#0a2540; margin-bottom:5px;'>🌐 Продукт 1. Автоматический аудит периметра (Защита данных в пути)</div>
			<div style='font-size:14px; color:#475569;'>Программа выполняет роль бдительного цифрового ревизора. Она сканирует внешние порты системы, проверяет сетевые шлюзы приложений и мгновенно выявляет уязвимости, предотвращая перехват данных снаружи.</div>
		</div>
		<div class='feature-card' style='border-left-color: #004d40;'>
			<div style='font-weight:bold; color:#004d40; margin-bottom:5px;'>🗄️ Продукт 2. Внутренний комплаенс-контроль серверов (Защита сохраненных данных)</div>
			<div style='font-size:14px; color:#475569;'>Модуль проводит тотальную ревизию файловой системы серверов. При обнаружении критических ошибок сотрудников, KvantumSafe автоматически изолирует угрозу, присваивая файлам права доступа банковского стандарта 0600.</div>
		</div>
		<div class='feature-card' style='border-left-color: #2563eb;'>
			<div style='font-weight:bold; color:#2563eb; margin-bottom:5px;'>🔀 Продукт 3. Постквантовый координатор маршрутов (Crypto-Agility)</div>
			<div style='font-size:14px; color:#475569;'>Флагманское решение для трансграничных переводов. Наша система выступает как умный транзитный диспетчер — она упаковывает потоки данных в защищенные контейнеры нового поколения по международным стандартам решеток (NIST ML-KEM).</div>
		</div>
		
		<h2 style='color:#0a2540; border-bottom:2px solid #e2e8f0; padding-bottom:10px; margin-top:30px;'>⚖️ Юридическая чистота</h2>
		<p style='font-size:14px; color:#475569; text-align:left;'>Платформа не осуществляет самостоятельную разработку криптографических алгоритмов и не является средством шифрования. Программа лишь оркестрирует и управляет теми защитными модулями, которые уже сертифицированы и встроены в серверное оборудование вашей организации. Это полностью снимает любые вопросы контролирующих и проверяющих органов касательно лицензирования.</p>
		
		<div class='footer'>
			<p>© 2026 KvantumSafe Pro. Международная ИТ-платформа комплаенса и контроля защитных шлюзов.</p>
		</div>
	</div>
</body>
</html>`
}

func GetPerimeterReportHTML(lang string) string {
	return "<html><head><meta charset='UTF-8'></head><body style='font-family:sans-serif;padding:40px;background:#f4f7f6;'><p><a href='/?lang=" + lang + "' style='color:#0a2540;font-weight:bold;text-decoration:none;'>← Назад / Back</a></p><h2>Результаты экспресс-аудита / Audit Report</h2><p style='color:#b71c1c;font-weight:bold;'>⚠️ КРИТИЧЕСКИЙ УРОВЕНЬ УГРОЗЫ / CRITICAL THREAT DETECTED</p><p>Обнаружены устаревшие версии криптографических библиотек OpenSSL на порту 443.</p></body></html>"
}

func GetBillingPageHTML(clientName, lang string) string {
	return "<html><head><meta charset='UTF-8'></head><body style='font-family:sans-serif;padding:40px;'><p><a href='/?lang=" + lang + "'>← Назад / Back</a></p><h2>Billing Panel</h2><p>Organization: <b>" + clientName + "</b></p></html>"
}
