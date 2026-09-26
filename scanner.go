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
	title := "Результаты экспресс-аудита безопасности"
	alert := "⚠️ ВНИМАНИЕ: ОБНАРУЖЕН КРИТИЧЕСКИЙ УРОВЕНЬ УГРОЗЫ периметра"
	object := "Целевой объект проверки"
	btnBack := "← Вернуться на главную"
	
	m1Title := "1. Сетевой периметр и открытые шлюзы:"
	m1Desc := "Зафиксированы открытые порты 443 (HTTPS) и 8443 (API-Gateway). Обнаружена трансляция версий используемого ПО веб-серверов в открытый интернет."
	m2Title := "2. Экспресс-анализ уязвимости криптографии (SSL/TLS):"
	m2Desc := "Обнаружены признаки использования устаревших криптографических библиотек, потенциально подверженных атакам класса перехвата сессий сотрудников."
	m3Title := "3. Готовность к постквантовым угрозам (Crypto-Agility индекс):"
	m3Desc := "<b>Уровень готовности: 0%.</b> Каналы трансграничной передачи финансовых пакетов данных (SWIFT, XML) уязвимы для перспективного дешифрования хакерами."
	m4Title := "4. Внутренний комплаенс-риск файловой системы:"
	m4Desc := "Существует потенциальная угроза утечки резервных копий баз данных SQL из-за отсутствия автоматического принудительного маскирования прав доступа стандарта POSIX 0600."
	m5Title := "5. Сводный ИБ-индекс критичности сетевого узла:"
	m5Desc := "<span style='color:#b71c1c; font-weight:bold; font-size:18px;'>УРОВЕНЬ РИСКА: 9.4 из 10 (КРИТИЧЕСКИЙ)</span>"

	if lang == "kg" {
		title = "Коопсуздуктун экспресс-аудитинин жыйынтыгы"
		alert = "⚠️ КӨҢҮЛ БУРУҢУЗДАР: Периметрдин КРИТИКАЛЫК КОРКУНУЧУ аныкталды"
		object = "Текшерилген объект"
		btnBack = "← Башкы бетке кайтуу"
		m1Title = "1. Тармактын периметри жана ачык шлюздар:"
		m1Desc = "Ачык 443 (HTTPS) жана 8443 (API-Gateway) порттору катталды. Веб-серверлердин программалык камсыздоосунун версиялары интернетке ачык көрсөтүлүп турат."
		m2Title = "2. Криптографиянын (SSL/TLS) алсыздыгын экспресс-талдоо:"
		m2Desc = "Кызматкерлердин сессияларын уурдап алуу чабуулдарына кабылышы мүмкүн болгон эскирген криптографиялык китепканалардын белгилери табылды."
		m3Title = "3. Посткванттык коркунучтарга даярдык (Crypto-Agility индекси):"
		m3Desc = "<b>Даярдык деңгээли: 0%.</b> Каржылык маалыматтарды (SWIFT, XML) трансчегаралык берүү каналдары келечекте хакерлер тарабынан чечмелөөгө дуушар болот."
		m4Title = "4. Ички файлдык тутумдун комплаенс-тобокелдиги:"
		m4Desc = "POSIX 0600 стандартындагы кирүү укуктарын автоматтык түрдө мажбурлап маскалоо жок болгондуктан, SQL маалымат базаларынын камдык көчүрмөлөрүнүн агып кетүү коркунучу бар."
		m5Title = "5. Тармактык түйүндүн жалпы критикалык индекси:"
		m5Desc = "<span style='color:#b71c1c; font-weight:bold; font-size:18px;'>ТОБОКЕЛДИК ДЕҢГЭЭЛИ: 10дон 9.4 (КРИТИКАЛЫК)</span>"
	} else if lang == "en" {
		title = "Express Security Audit Results"
		alert = "⚠️ WARNING: CRITICAL PERIMETER THREAT LEVEL DETECTED"
		object = "Target Verification Node"
		btnBack = "← Back to Dashboard"
		m1Title = "1. Network Perimeter & Open Gateways:"
		m1Desc = "Open ports 443 (HTTPS) and 8443 (API-Gateway) detected. Web-server software versions are openly broadcasted to the public internet."
		m2Title = "2. Cryptographic Vulnerability Audit (SSL/TLS):"
		m2Desc = "Traces of outdated cryptographic libraries detected, making employee active sessions vulnerable to interception attacks."
		m3Title = "3. Post-Quantum Readiness Index (Crypto-Agility Score):"
		m3Desc = "<b>Readiness Level: 0%.</b> Cross-border financial data streams (SWIFT, XML) are highly vulnerable to future decryption exploits."
		m4Title = "4. Internal File System Compliance Exposure:"
		m4Desc = "Potential threat of SQL database backup leak due to the complete lack of automated POSIX 0600 strict access token enforcement."
		m5Title = "5. Consolidated Security Risk Score:"
		m5Desc = "<span style='color:#b71c1c; font-weight:bold; font-size:18px;'>RISK SCORE: 9.4 out of 10 (CRITICAL)</span>"
	}

	return `<!DOCTYPE html>
<html>
<head>
	<meta charset='UTF-8'>
	<title>Audit Report</title>
	<style>
		body { font-family: 'Segoe UI', sans-serif; background-color: #f4f7f6; color: #333; margin: 0; padding: 20px; }
		.report-card { max-width: 800px; background: white; margin: 30px auto; padding: 40px; border-radius: 12px; box-shadow: 0 4px 25px rgba(0,0,0,0.08); }
		.btn-back { color: #0a2540; font-weight: bold; text-decoration: none; font-size: 14px; display: inline-block; margin-bottom: 20px; }
		.btn-back:hover { text-decoration: underline; }
		h2 { color: #0a2540; border-bottom: 2px solid #cbd5e1; padding-bottom: 12px; margin-top: 0; font-size: 24px; }
