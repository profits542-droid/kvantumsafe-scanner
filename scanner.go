package main

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
		.navbar h1 { margin: 0; font-size: 30px; font-weight: 600; }
		.navbar p { margin: 8px 0 15px 0; color: #cbd5e1; font-size: 16px; }
		.lang-switcher { display: flex; justify-content: center; gap: 12px; margin-top: 15px; }
		.lang-switcher a { color: #cbd5e1; text-decoration: none; font-weight: bold; font-size: 14px; background: rgba(255,255,255,0.15); padding: 6px 14px; border-radius: 6px; }
		.lang-switcher a:hover { color: white; background: #004d40; }
		.content { max-width: 850px; background: white; margin: 40px auto; padding: 40px; border-radius: 12px; box-shadow: 0 4px 25px rgba(0,0,0,0.05); }
		.scan-container { background: #e0f2f1; padding: 35px; border-radius: 10px; border: 2px dashed #004d40; text-align: center; margin: 20px 0; }
		.scan-input { width: 60%; padding: 14px; font-size: 16px; border: 1px solid #cbd5e1; border-radius: 6px; margin-right: 10px; outline: none; }
		.scan-btn { background: #004d40; color: white; padding: 14px 30px; font-size: 16px; border: none; border-radius: 6px; font-weight: bold; cursor: pointer; }
		.feature-card { background: #f8fafc; padding: 25px; border-left: 4px solid #d4af37; margin-bottom: 20px; border-radius: 0 8px 8px 0; }
		html[dir="rtl"] .feature-card { border-left: none; border-right: 4px solid #d4af37; border-radius: 8px 0 0 8px; }
		.feature-title { font-weight: bold; color: #0a2540; font-size: 18px; margin-bottom: 8px; }
		.tariff-table { width: 100%; border-collapse: collapse; margin-top: 25px; text-align: center; }
		.tariff-table th, .tariff-table td { padding: 20px; border: 1px solid #cbd5e1; font-size: 15px; }
		.tariff-free { background-color: #f0fdf4; color: #16a34a; }
		.tariff-premium { background-color: #faf5ff; color: #7c3aed; }
		.footer { text-align: center; margin-top: 40px; color: #64748b; font-size: 14px; border-top: 1px solid #e2e8f0; padding-top: 20px; }
	</style>
</head>
<body>
	<div class='navbar'>
		<h1>` + getTranslation("nav_title", lang) + `</h1>
		<p>📊 ` + getTranslation("nav_sub", lang) + `</p>
		<div class='lang-switcher'>
			<a href='/?lang=ru'>RU</a>
			<a href='/?lang=en'>EN</a>
			<a href='/?lang=kz'>KZ</a>
			<a href='/?lang=ar'>AR</a>
		</div>
	</div>
	<div class='content'>
		<h2>` + getTranslation("about_title", lang) + `</h2>
		<p>` + getTranslation("about_desc", lang) + `</p>
		
		<div class='scan-container'>
			<h3>` + getTranslation("scan_title", lang) + `</h3>
			<p style='color: #004d40; font-size: 15px; margin-bottom: 20px;'>` + getTranslation("scan_desc", lang) + `</p>
			<form action='/report/mbank?lang=` + lang + `' method='GET'>
				<input type='text' name='domain' class='scan-input' placeholder='bank.com...'>
				<button type='submit' class='scan-btn'>` + getTranslation("scan_btn", lang) + `</button>
			</form>
		</div>

		<h2>` + getTranslation("prod_title", lang) + `</h2>
		<div class='feature-card'><div class='feature-title'>` + getTranslation("p1_title", lang) + `</div><div>` + getTranslation("p1_desc", lang) + `</div></div>
		<div class='feature-card' style='border-left-color: #004d40; border-right-color: #004d40;'><div class='feature-title'>` + getTranslation("p2_title", lang) + `</div><div>` + getTranslation("p2_desc", lang) + `</div></div>
		<div class='feature-card' style='border-left-color: #2563eb; border-right-color: #2563eb;'><div class='feature-title'>` + getTranslation("p3_title", lang) + `</div><div>` + getTranslation("p3_desc", lang) + `</div></div>

		<h2>` + getTranslation("legal_title", lang) + `</h2><p>` + getTranslation("legal_desc", lang) + `</p>

		<h2>` + getTranslation("saas_title", lang) + `</h2>
		<table class='tariff-table'>
			<tr>
				<th class='tariff-free'><h3>Global Scanner</h3><div style='font-size:20px; font-weight:bold;'>FREE / $0</div></th>
				<th><h3>Compliance Pro</h3><div style='font-size:20px; font-weight:bold;'>$15,000 / год</div></th>
				<th class='tariff-premium'><h3>Quantum Web3</h3><div style='font-size:20px; font-weight:bold;'>$35,000 / год</div></th>
			</tr>
		</table>
		<div class='footer'><p>© 2026 KvantumSafe Pro. Международная ИТ-платформа комплаенса органов контроля.</p></div>
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
