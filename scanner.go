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
		.navbar p { margin: 8px 0 15px 0; color: #94a3b8; font-size: 14px; max-width: 800px; margin-left: auto; margin-right: auto; line-height: 1.5; background: rgba(255,255,255,0.08); padding: 12px; border-radius: 6px; border: 1px dashed rgba(255,255,255,0.2); }
		.lang-switcher { display: flex; justify-content: center; gap: 12px; margin-top: 15px; }
		.lang-switcher a { color: #cbd5e1; text-decoration: none; font-weight: bold; font-size: 14px; background: rgba(255,255,255,0.15); padding: 6px 14px; border-radius: 6px; }
		.lang-switcher a:hover, .lang-switcher a.active { color: white; background: #004d40; }
		.content { max-width: 850px; background: white; margin: 40px auto; padding: 40px; border-radius: 12px; box-shadow: 0 4px 25px rgba(0,0,0,0.05); }
		.scan-container { background: #e0f2f1; padding: 35px; border-radius: 10px; border: 2px dashed #004d40; text-align: center; margin: 20px 0; }
		.scan-input { width: 60%; padding: 14px; font-size: 16px; border: 1px solid #cbd5e1; border-radius: 6px; margin-right: 10px; outline: none; }
		.scan-btn { background: #004d40; color: white; padding: 14px 30px; font-size: 16px; border: none; border-radius: 6px; font-weight: bold; cursor: pointer; }
		.tariff-table { width: 100%; border-collapse: collapse; margin-top: 25px; text-align: center; direction: ` + direction + `; }
		.tariff-table th, .tariff-table td { padding: 15px; border: 1px solid #cbd5e1; font-size: 14px; }
		.tariff-free { background-color: #f8fafc; color: #64748b; }
		.tariff-pro { background-color: #f0fdf4; color: #16a34a; font-weight: bold; }
		.tariff-premium { background-color: #faf5ff; color: #7c3aed; }
		.btn-doc { display: block; width: 80%; margin: 30px auto 10px auto; background-color: #d4af37; color: #0a2540; padding: 15px; text-align: center; border-radius: 8px; font-weight: bold; text-decoration: none; font-size: 15px; box-shadow: 0 4px 12px rgba(0,0,0,0.05); }
		.btn-doc:hover { background-color: #f3cd44; }
		.build-box { background-color: #f8fafc; padding: 25px; border-radius: 8px; border: 1px solid #e2e8f0; margin-top: 30px; text-align: ` + textAlign + `; }
		.status-badge { background-color: #16a34a; color: white; padding: 3px 8px; border-radius: 4px; font-weight: bold; font-size: 13px; }
		.links-container { display: flex; justify-content: center; gap: 20px; margin-top: 15px; margin-bottom: 25px; }
		.links-container a { color: #2563eb; font-weight: bold; text-decoration: none; font-size: 14px; }
		.links-container a:hover { text-decoration: underline; }

		/* СТИЛЬНЫЙ СТРУКТУРИРОВАННЫЙ B2B ПОДВАЛ */
		.footer-corporate { background-color: #0f172a; color: #94a3b8; padding: 40px; border-radius: 10px; margin-top: 40px; border-top: 3px solid #d4af37; text-align: left; }
		html[dir="rtl"] .footer-corporate { text-align: right; }
		.footer-grid { display: flex; flex-wrap: wrap; gap: 30px; justify-content: space-between; }
		.footer-section { flex: 1; min-width: 220px; }
		.footer-section h4 { color: #f8fafc; font-size: 15px; margin-top: 0; margin-bottom: 15px; text-transform: uppercase; letter-spacing: 0.5px; border-left: 3px solid #d4af37; padding-left: 8px; }
		html[dir="rtl"] .footer-section h4 { border-left: none; border-right: 3px solid #d4af37; padding-left: 0; padding-right: 8px; }
		.footer-section p { margin: 6px 0; font-size: 13.5px; }
		.footer-link { color: #38bdf8; text-decoration: none; font-weight: 500; }
		.footer-link:hover { text-decoration: underline; }
		.btn-messenger { display: inline-flex; align-items: center; padding: 8px 16px; border-radius: 6px; color: white; text-decoration: none; font-size: 13px; font-weight: bold; margin-top: 8px; margin-right: 8px; box-shadow: 0 2px 8px rgba(0,0,0,0.15); transition: 0.2s; }
		.btn-messenger:hover { transform: translateY(-1px); opacity: 0.9; }
		.btn-tg { background-color: #0284c7; }
		.btn-wa { background-color: #16a34a; }
		.footer-bottom { border-top: 1px solid #334155; margin-top: 30px; padding-top: 20px; text-align: center; font-size: 12px; color: #64748b; }
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
				<th class='tariff-free' style='width:33%;'><b>Global Scanner</b><br><span>` + getTranslation("t1_free", lang) + `</span></th>
				<th class='tariff-pro' style='width:33%; border: 2px solid #16a34a;'><b>ПО АнтиХакер AI</b><br><span>` + getTranslation("t2_pro", lang) + `</span></th>
				<th class='tariff-premium' style='width:33%;'><b>Quantum Web3 (SDK v2.6)</b><br><span>` + getTranslation("t3_web3", lang) + `</span></th>
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

		<!-- НОВЫЙ СОЛИДНЫЙ КОРПОРАТИВНЫЙ ПОДВАЛ -->
		<div class='footer-corporate'>
			<div class='footer-grid'>
				<div class='footer-section'>
					<h4>Правообладатель</h4>
					<p style='font-weight:bold; color:#f1f5f9;'>ОсОО «Квантум Сейф»</p>
					<p>Государственная регистрация финансово-оборонного софта нового поколения.</p>
					<p>г. Ош, Кыргызская Республика</p>
				</div>
				<div class='footer-section'>
					<h4>Официальная связь</h4>
					<p>B2B Департамент: <a href='mailto:info@kvantumsafe.tech' class='footer-link'>info@kvantumsafe.tech</a></p>
					<p>Центральный узел: +996 (777) 57-99-70</p>
				</div>
				<div class='footer-section'>
					<h4>Каналы прямого отклика</h4>
					<p style='font-size:12px; margin-bottom:8px;'>Быстрая фиксация времени пилотных тестов:</p>
					<a href='https://t.me' target='_blank' class='btn-messenger btn-tg'>Telegram</a>
					<a href='https://wa.me' target='_blank' class='btn-messenger btn-wa'>WhatsApp</a>
				</div>
			</div>
			<div class='footer-bottom'>
				<p>© 2026 ОсОО «Квантум Сейф». Все права защищены. Разработано в соответствии с международными стандартами безопасной архитектуры распределенных систем ядра.</p>
			</div>
		</div>

	</div>
</body>
</html>`
}

func activeClass(current, target string) string {
	if current == target { return "active" }
	return ""
}
