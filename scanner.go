package main

import "fmt"

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
		.navbar { background-color: #0a2540; color: white; padding: 35px 40px; text-align: center; position: relative; box-shadow: 0 4px 10px rgba(0,0,0,0.1); }
		.navbar h1 { margin: 0; font-size: 30px; font-weight: 600; }
		.navbar p { margin: 8px 0 0 0; color: #cbd5e1; font-size: 16px; }
		.lang-switcher { position: absolute; top: 15px; right: 40px; display: flex; gap: 10px; }
		.lang-switcher a { color: #cbd5e1; text-decoration: none; font-weight: bold; font-size: 14px; background: rgba(255,255,255,0.1); padding: 5px 10px; border-radius: 4px; }
		.lang-switcher a:hover, .lang-switcher a.active { color: white; background: #004d40; }
		.content { max-width: 850px; background: white; margin: 40px auto; padding: 40px; border-radius: 12px; box-shadow: 0 4px 25px rgba(0,0,0,0.05); }
		.scan-container { background: #e0f2f1; padding: 35px; border-radius: 10px; border: 2px dashed #004d40; text-align: center; margin: 20px 0; }
		.scan-input { width: 60%; padding: 14px; font-size: 16px; border: 1px solid #cbd5e1; border-radius: 6px; margin-right: 10px; outline: none; }
		.scan-btn { background: #004d40; color: white; padding: 14px 30px; font-size: 16px; border: none; border-radius: 6px; font-weight: bold; cursor: pointer; }
		.footer { text-align: center; margin-top: 40px; color: #64748b; font-size: 14px; }
	</style>
</head>
<body>
	<div class='navbar'>
		<div class='lang-switcher'>
			<a href='/?lang=ru'>RU</a>
			<a href='/?lang=en'>EN</a>
			<a href='/?lang=kz'>KZ</a>
			<a href='/?lang=ar'>AR</a>
		</div>
		<h1>` + getTranslation("nav_title", lang) + `</h1>
		<p>` + getTranslation("nav_sub", lang) + `</p>
	</div>
	<div class='content'>
		<div class='scan-container'>
			<h3 style='margin-top:0; color:#004d40;'>🔍 Проверить безопасность сайта / Check Website Security</h3>
			<form action='/report/mbank?lang=` + lang + `' method='GET'>
				<input type='text' name='domain' class='scan-input' placeholder='Введите домен / Enter domain...'>
				<button type='submit' class='scan-btn'>Сканировать</button>
			</form>
		</div>
		<div class='footer'>
			<p>© 2026 KvantumSafe Pro. Международная ИТ-платформа комплаенса органов контроля.</p>
		</div>
	</div>
</body>
</html>`
}

func GetPerimeterReportHTML(lang string) string {
	return "<html><body><h2>Audit Report (" + lang + ")</h2><p><a href='/?lang=" + lang + "'>← Back</a></p></body></html>"
}

func GetBillingPageHTML(clientName, lang string) string {
	return "<html><body><h2>Billing Panel (" + lang + ")</h2><p><a href='/?lang=" + lang + "'>← Back</a></p></body></html>"
}
