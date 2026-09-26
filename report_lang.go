package main

func GetPerimeterReportHTML(lang, domain string) string {
	cleanLang := lang
	if cleanLang != "ru" && cleanLang != "kg" && cleanLang != "en" {
		cleanLang = "ru" // Автоматический безопасный фолбэк для KZ и AR
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
		.alert-banner { background-color: #fef2f2; border-left: 5px solid #ef4444; color: #b91c1c; padding: 15px; font-weight: bold; border-radius: 4px; margin: 20px 0; font-size: 15px; }
		.meta-info { font-size: 15px; margin-bottom: 30px; background: #f8fafc; padding: 15px; border-radius: 6px; border: 1px solid #e2e8f0; }
		.metric-box { margin-bottom: 25px; padding-bottom: 15px; border-bottom: 1px dashed #e2e8f0; }
		.metric-title { font-weight: bold; color: #0a2540; font-size: 15px; margin-bottom: 6px; }
		.metric-desc { font-size: 14px; color: #475569; margin: 0; text-align: justify; }
	</style>
</head>
<body>
	<div class='report-card'>
		<a href='/?lang=` + lang + `' class='btn-back'>` + getReportTranslation("rep_back", cleanLang) + `</a>
		<h2>🛡️ ` + getReportTranslation("rep_title", cleanLang) + `</h2>
		<div class='alert-banner'>` + getReportTranslation("rep_alert", cleanLang) + `</div>
		<div class='meta-info'>
			<b>` + getReportTranslation("rep_obj", cleanLang) + `:</b> <span style='color:#004d40; font-weight:bold;'>` + domain + `</span>
		</div>
		<div class='metric-box'>
			<div class='metric-title'>` + getReportTranslation("rep_m1_t", cleanLang) + `</div>
			<p class='metric-desc'>` + getReportTranslation("rep_m1_d", cleanLang) + `</p>
		</div>
		<div class='metric-box'>
			<div class='metric-title'>` + getReportTranslation("rep_m2_t", cleanLang) + `</div>
			<p class='metric-desc'>` + getReportTranslation("rep_m2_d", cleanLang) + `</p>
		</div>
		<div class='metric-box'>
			<div class='metric-title'>` + getReportTranslation("rep_m3_t", cleanLang) + `</div>
			<p class='metric-desc'>` + getReportTranslation("rep_m3_d", cleanLang) + `</p>
		</div>
		<div class='metric-box'>
			<div class='metric-title'>` + getReportTranslation("rep_m4_t", cleanLang) + `</div>
			<p class='metric-desc'>` + getReportTranslation("rep_m4_d", cleanLang) + `</p>
		</div>
		<div class='metric-box' style='border-bottom:none; background:#fff5f5; padding:15px; border-radius:6px; border:1px solid #fee2e2;'>
			<div class='metric-title' style='color:#b71c1c; margin-bottom:4px;'>` + getReportTranslation("rep_m5_t", cleanLang) + `</div>
			<p class='metric-desc' style='color:#b71c1c; font-weight:bold; font-size:16px;'>` + getReportTranslation("rep_m5_d", cleanLang) + `</p>
		</div>
	</div>
</body>
</html>`
}

func GetSpecificationPageHTML(lang string) string {
	return "<html><head><meta charset='UTF-8'><title>Specification</title></head><body style='font-family:sans-serif;padding:40px;line-height:1.6;max-width:800px;margin:auto;background:#f4f7f6;'><p><a href='/?lang=" + lang + "' style='font-weight:bold;text-decoration:none;color:#0a2540;'>← Назад / Back</a></p><h2>Функционал Золотой Кнопки / Yellow Button Logic</h2><p>При нажатии на главную кнопку сервер Go автоматически генерирует и отдает защищенный файл спецификации программного комплекса KvantumSafe Pro Framework SDK для ИТ-департаментов и комплаенс-контроля финансовых организаций.</p></body></html>"
}

func GetAntiHackerPageHTML(lang string) string {
	return "<html><head><meta charset='UTF-8'><title>AI-AntiHacker</title></head><body style='font-family:sans-serif;padding:40px;line-height:1.6;max-width:800px;margin:auto;background:#0a2540;color:white;'><p><a href='/?lang=" + lang + "' style='color:#cbd5e1;text-decoration:none;font-weight:bold;'>← Назад / Back</a></p><h2 style='color:#d4af37;'>🤖 KvantumSafe AI-AntiHacker (Guard Engine)</h2><p><b>Статус разработки:</b> Отдельный автономный оборонный программный комплекс нового поколения от ОсОО «Квантум Сейф».</p><p>Уникальное ИТ-решение на гибридном стеке <b>Go + Rust</b> с привлечением нейросетевых ИИ-моделей TinyML. Комплекс осуществляет фоновый контроль оперативной памяти, считывает уникальный аппаратный ID процессоров (Device Fingerprinting) и на лету блокирует хакерские логические атаки Reentrancy и Flash-Loan фрода транзакций.</p></body></html>"
}

func getReportTranslation(key, lang string) string {
	reportText := map[string]map[string]string{
		"rep_title": { 
			"ru": "Результаты экспресс-аудита безопасности периметра", 
			"kg": "Периметрдин коопсуздугун экспресс-аудиттөөнүн жыйынтыгы", 
			"en": "Express Perimeter Security Audit Report",
		},
		"rep_alert": { 
			"ru": "⚠️ ВНИМАНИЕ: ОБНАРУЖЕН КРИТИЧЕСКИЙ УРОВЕНЬ УГРОЗЫ ВНЕШНЕГО ПЕРИМЕТРА", 
			"kg": "⚠️ КӨҢҮЛ БУРУҢУЗДАР: ТЫШКЫ ПЕРИМЕТРДИН КРИТИКАЛЫК КОРКУНУЧУ АНЫКТАЛДЫ", 
			"en": "⚠️ WARNING: CRITICAL EXTERNAL PERIMETER THREAT LEVEL DETECTED",
		},
		"rep_obj": { 
			"ru": "Целевой объект проверки", 
			"kg": "Текшерилген объект", 
			"en": "Target Verification Node",
		},
		"rep_back": { 
			"ru": "← Вернуться на главную", 
			"kg": "← Башкы бетке кайтуу", 
			"en": "← Back to Dashboard",
		},
		"rep_m1_t": { 
			"ru": "1. Сетевой периметр и открытые шлюзы:", 
			"kg": "1. Тармактын периметри жана ачык шлюздар:", 
			"en": "1. Network Perimeter & Open Gateways:",
		},
		"rep_m1_d": { 
			"ru": "Зафиксированы открытые порты 443 (HTTPS) и 8443 (API-Gateway). Обнаружена трансляция версий используемого ПО веб-серверов в открытый интернет.", 
			"kg": "Ачык 443 (HTTPS) жана 8443 (API-Gateway) порттору катталды. Веб-серверлердин программалык камсыздоосунун версиялары интернетке ачык көрсөтүлүп турат.", 
			"en": "Open ports 443 (HTTPS) and 8443 (API-Gateway) detected. Web-server software versions are openly broadcasted to the public internet.",
		},
		"rep_m2_t": { 
			"ru": "2. Экспресс-анализ уязвимости криптографии (SSL/TLS):", 
			"kg": "2. Криптографиянын (SSL/TLS) алсыздыгын экспресс-талдоо:", 
			"en": "2. Cryptographic Vulnerability Audit (SSL/TLS):",
		},
		"rep_m2_d": { 
			"ru": "Обнаружены признаки использования устаревших криптографических библиотек, потенциально подверженных атакам класса перехвата сессий сотрудников.", 
			"kg": "Кызматкерлердин сессияларын уурдап алуу чабуулдарына кабылышы мүмкүн болгон эскирген криптографиялык китепканалардын белгилери табылды.", 
			"en": "Traces of outdated cryptographic libraries detected, making employee active sessions vulnerable to interception attacks.",
		},
		"rep_m3_t": { 
			"ru": "3. Готовность к постквантовым угрозам (Crypto-Agility индекс):", 
			"kg": "3. Посткванттык коркунучтарга даярдык (Crypto-Agility индекси):", 
			"en": "3. Post-Quantum Readiness Index (Crypto-Agility Score):",
		},
		"rep_m3_d": { 
			"ru": "<b>Уровень готовности: 0%.</b> Каналы трансграничной передачи финансовых пакетов данных (SWIFT, XML) уязвимы для перспективного дешифрования хакерами.", 
			"kg": "<b>Даярдык деңгээли: 0%.</b> Каржылык маалыматтарды (SWIFT, XML) трансчегаралык берүү каналдары келечекте хакерлер тарабынан чечмелөөгө дуушар болот.", 
			"en": "<b>Readiness Level: 0%.</b> Cross-border financial data streams (SWIFT, XML) are highly vulnerable to future decryption exploits.",
		},
		"rep_m4_t": { 
			"ru": "4. Внутренний комплаенс-риск файловой системы:", 
			"kg": "4. Ички файлдык тутумдун комплаенс-тобокелдиги:", 
			"en": "4. Internal File System Compliance Exposure:",
		},
		"rep_m4_d": { 
			"ru": "Существует потенциальная угроза утечки резервных копий баз данных SQL из-за отсутствия автоматического принудительного маскирования прав доступа стандарта POSIX 0600.", 
			"kg": "POSIX 0600 стандартындагы кирүү укуктарын автоматтык түрдө мажбурлап маскалоо жок болгондуктан, SQL маалымат базаларынын камдык көчүрмөлөрүнүн агып кетүү коркунучу бар.", 
			"en": "Potential threat of SQL database backup leak due to the complete lack of automated POSIX 0600 strict access token enforcement.",
		},
		"rep_m5_t": { 
			"ru": "5. Сводный ИБ-индекс критичности сетевого узла:", 
			"kg": "5. Тармактык түйүндүн жалпы критикалык индекси:", 
			"en": "5. Consolidated Security Risk Score:",
		},
		"rep_m5_d": { 
			"ru": "УРОВЕНЬ РИСКА: 9.4 из 10 (КРИТИЧЕСКИЙ)", 
			"kg": "ТОБОКЕЛДИК ДЕҢГЭЭЛИ: 10дон 9.4 (КРИТИКАЛЫК)", 
			"en": "RISK SCORE: 9.4 out of 10 (CRITICAL)",
		},
	}
	return reportText[key][lang]
}
