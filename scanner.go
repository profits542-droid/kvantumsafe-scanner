package main

func GetDashboardHTML(clientName, serverID, lang string) string {
	direction := "ltr"
	textAlign := "left"
	if lang == "ar" { 
		direction = "rtl" 
		textAlign = "right"
	}

	return `<!DOCTYPE html>
<html lang='` + lang + `' dir='` + direction + `'>
<head>
	<meta charset='UTF-8'>
	<title>KvantumSafe Pro</title>
	<style>
		body { font-family: 'Segoe UI', sans-serif; background-color: #f4f7f6; color: #333; margin: 0; padding: 0; line-height: 1.6; }
		.navbar { background-color: #0a2540; color: white; padding: 35px 40px; text-align: center; position: relative; box-shadow: 0 4px 10px rgba(0,0,0,0.1); }
		.navbar h1 { margin: 0; font-size: 26px; font-weight: 600; }
		.navbar p { margin: 8px 0 15px 0; color: #cbd5e1; font-size: 15px; }
		.lang-switcher { display: flex; justify-content: center; gap: 12px; margin-top: 15px; }
		.lang-switcher a { color: #cbd5e1; text-decoration: none; font-weight: bold; font-size: 14px; background: rgba(255,255,255,0.15); padding: 6px 14px; border-radius: 6px; }
		.lang-switcher a:hover, .lang-switcher a.active { color: white; background: #004d40; }
		.content { max-width: 850px; background: white; margin: 40px auto; padding: 40px; border-radius: 12px; box-shadow: 0 4px 25px rgba(0,0,0,0.05); }
		.scan-container { background: #e0f2f1; padding: 35px; border-radius: 10px; border: 2px dashed #004d40; text-align: center; margin: 20px 0; }
		.scan-input { width: 60%; padding: 14px; font-size: 16px; border: 1px solid #cbd5e1; border-radius: 6px; margin-right: 10px; outline: none; }
		.scan-btn { background: #004d40; color: white; padding: 14px 30px; font-size: 16px; border: none; border-radius: 6px; font-weight: bold; cursor: pointer; }
		.footer { text-align: center; margin-top: 40px; color: #64748b; font-size: 14px; border-top: 1px solid #e2e8f0; padding-top: 20px; }
		.tariff-table { width: 100%; border-collapse: collapse; margin-top: 25px; text-align: center; direction: ` + direction + `; }
		.tariff-table th, .tariff-table td { padding: 15px; border: 1px solid #cbd5e1; font-size: 14px; }
		.tariff-free { background-color: #f0fdf4; color: #16a34a; }
		.tariff-premium { background-color: #faf5ff; color: #7c3aed; }
		.btn-doc { display: block; width: 80%; margin: 30px auto 10px auto; background-color: #d4af37; color: #0a2540; padding: 15px; text-align: center; border-radius: 8px; font-weight: bold; text-decoration: none; font-size: 15px; box-shadow: 0 4px 12px rgba(0,0,0,0.05); }
		.btn-doc:hover { background-color: #f3cd44; }
		.build-box { background-color: #f8fafc; padding: 25px; border-radius: 8px; border: 1px solid #e2e8f0; margin-top: 30px; text-align: ` + textAlign + `; }
		.status-badge { background-color: #16a34a; color: white; padding: 3px 8px; border-radius: 4px; font-weight: bold; font-size: 13px; }
		.links-container { display: flex; justify-content: center; gap: 20px; margin-top: 15px; }
		.links-container a { color: #2563eb; font-weight: bold; text-decoration: none; font-size: 14px; }
		.links-container a:hover { text-decoration: underline; }
	</style>
</head>
<body>
	<div class='navbar'>
		<h1>` + getTranslation("nav_title", lang) + `</h1>
		<p>` + getTranslation("nav_sub", lang) + `</p>
		<div class='lang-switcher'>
			<a href='/?lang=kg' class='` + activeClass(lang, "kg") + `'>KG</a>
			<a href='/?lang=ru' class='` + activeClass(lang, "ru") + `'>RU</a>
			<a href='/?lang=en' class='` + activeClass(lang, "en") + `'>EN</a>
			<a href='/?lang=kz' class='` + activeClass(lang, "kz") + `'>KZ</a>
			<a href='/?lang=ar' class='` + activeClass(lang, "ar") + `'>AR</a>
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
		
		<h2 style='text-align:` + textAlign + `; color:#0a2540; border-bottom:2px solid #e2e8f0; padding-bottom:10px;'>` + getTranslation("saas_title", lang) + `</h2>
		<table class='tariff-table'>
			<tr style='background:#f8fafc;'>
				<th class='tariff-free'><b>Global Scanner</b><br><span>` + getTranslation("t1_free", lang) + `</span></th>
				<th><b>Compliance Pro</b><br><span>` + getTranslation("t2_pro", lang) + `</span></th>
				<th class='tariff-premium'><b>Quantum Web3</b><br><span>` + getTranslation("t3_web3", lang) + `</span></th>
			</tr>
		</table>
		
		<a href='/download?lang=` + lang + `' class='btn-doc'>` + getTranslation("doc_btn", lang) + `</a>
		
		<div class='links-container'>
			<a href='/specification?lang=` + lang + `'>` + getTranslation("link_spec", lang) + `</a>
			<a href='/antihacker?lang=` + lang + `' style='color:#7c3aed;'>` + getTranslation("link_anti", lang) + `</a>
		</div>

		<div class='build-box'>
			<h4 style='margin-top:0; color:#0a2540; font-size:16px;'>` + getTranslation("build_title", lang) + ` <span class='status-badge'>✓ Live</span></h4>
			<p style='font-size:14px; color:#475569; margin-bottom:0;'>` + getTranslation("build_desc", lang) + `</p>
		</div>
		<div class='footer'><p>` + getTranslation("footer_text", lang) + `</p></div>
	</div>
</body>
</html>`
}

func activeClass(current, target string) string {
	if current == target { return "active" }
	return ""
}

func GetPerimeterReportHTML(lang, domain string) string {
	return "<html><head><meta charset='UTF-8'></head><body style='font-family:sans-serif;padding:40px;background:#f4f7f6;'><p><a href='/?lang=" + lang + "' style='color:#0a2540;font-weight:bold;text-decoration:none;'>← Назад / Back</a></p><h2>Жыйынтык / Audit Report</h2><p style='color:#b71c1c;font-weight:bold;font-size:18px;'>⚠️ КОРКУНУЧ ТУУДУРГАН АБАЛ / CRITICAL THREAT DETECTED</p><p><b>Объект:</b> " + domain + "</p></body></html>"
}

func GetSpecificationPageHTML(lang string) string {
	return "<html><head><meta charset='UTF-8'><title>Specification</title></head><body style='font-family:sans-serif;padding:40px;line-height:1.6;max-width:800px;margin:auto;background:#f4f7f6;'><p><a href='/?lang=" + lang + "' style='font-weight:bold;text-decoration:none;color:#0a2540;'>← Назад / Back</a></p><h2>Функционал Золотой Кнопки / Yellow Button Logic</h2><p>При нажатии на главную кнопку сервер Go автоматически генерирует и отдает защищенный файл спецификации программного комплекса KvantumSafe Pro Framework SDK для ИТ-департаментов и комплаенс-контроля финансовых организаций.</p></body></html>"
}

func GetAntiHackerPageHTML(lang string) string {
	return "<html><head><meta charset='UTF-8'><title>AI-AntiHacker</title></head><body style='font-family:sans-serif;padding:40px;line-height:1.6;max-width:800px;margin:auto;background:#0a2540;color:white;'><p><a href='/?lang=" + lang + "' style='color:#cbd5e1;text-decoration:none;font-weight:bold;'>← Назад / Back</a></p><h2 style='color:#d4af37;'>🤖 KvantumSafe AI-AntiHacker (Guard Engine)</h2><p><b>Статус разработки:</b> Отдельный автономный оборонный программный комплекс нового поколения от ОсОО «Квантум Сейф».</p><p>Уникальное ИТ-решение на гибридном стеке <b>Go + Rust</b> с привлечением нейросетевых ИИ-моделей TinyML. Комплекс осуществляет фоновый контроль оперативной памяти, считывает уникальный аппаратный ID процессоров (Device Fingerprinting) и на лету блокирует хакерские логические атаки Reentrancy и Flash-Loan фрода транзакций.</p></body></html>"
}
