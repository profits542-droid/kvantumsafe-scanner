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
		"doc_btn": {
			"ru": "📄 Скачать спецификацию ПО и B2B-оффер (PDF)",
			"en": "📄 Download Software Specification & B2B Offer (PDF)",
			"kz": "📄 Бағдарламалық құралдың сипаттамасын және B2B ұсынысын жүктеу (PDF)",
			"ar": "📄 تنزيل مواصفات البرامج وعرض B2B (PDF)",
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
		.navbar h1 { margin: 0; font-size: 26px; font-weight: 600; }
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
		.btn-doc { display: block; width: 80%; margin: 30px auto 10px auto; background-color: #d4af37; color: #0a2540; padding: 15px; text-align: center; border-radius: 8px; font-weight: bold; text-decoration: none; font-size: 16px; box-shadow: 0 4px 12px rgba(0,0,0,0.05); transition: 0.2s; }
		.btn-doc:hover { background-color: #f3cd44; }
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
			<form action='/report/mbank' method='GET'>
				<input type='hidden' name='lang' value='` + lang + `'>
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
		
		<a href='/static/logo.jpg' download class='btn-doc'>` + getTranslation("doc_btn", lang) + `</a>
		
		<div class='footer'>
			<p>© 2026 KvantumSafe Pro. Международная ИТ-платформа комплаенса и контроля защитных шлюзов.</p>
		</div>
	</div>
</body>
</html>`
}

func GetPerimeterReportHTML(lang, domain string) string {
	return "<html><head><meta charset='UTF-8'></head><body style='font-family:sans-serif;padding:40px;background:#f4f7f6;'><p><a href='/?lang=" + lang + "' style='color:#0a2540;font-weight:bold;text-decoration:none;'>← Назад / Back</a></p><h2>Результаты экспресс-аудита / Audit Report</h2><p style='color:#b71c1c;font-weight:bold;font-size:18px;'>⚠️ СТАТУС: ОБНАРУЖЕН КРИТИЧЕСКИЙ УРОВЕНЬ УГРОЗЫ / CRITICAL THREAT DETECTED</p><p><b>Объект сканирования / Target:</b> " + domain + "</p><p>Обнаружены устаревшие версии криптографических библиотек OpenSSL на порту 443 Вашего сетевого узла. Рекомендуется интеграция стандартов комплаенса KvantumSafe Pro.</p></body></html>"
}

func GetBillingPageHTML(clientName, lang string) string {
	return "<html><head><meta charset='UTF-8'></head><body style='font-family:sans-serif;padding:40px;'><h2>Billing Panel (" + lang + ")</h2><p>Organization: " + clientName + "</p></html>"
}
