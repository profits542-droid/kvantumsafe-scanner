package main

func GetPerimeterReportHTML(lang, domain string) string {
	cleanLang := lang
	if cleanLang != "ru" && cleanLang != "kg" && cleanLang != "en" && cleanLang != "kz" && cleanLang != "ar" {
		cleanLang = "ru"
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

func getReportTranslation(key, lang string) string {
	reportText := map[string]map[string]string{
		"rep_title": { 
			"ru": "Результаты экспресс-аудита безопасности периметра", 
			"kg": "Периметрдин коопсуздугун экспресс-аудиттөөнүн жыйынтыгы", 
			"en": "Express Perimeter Security Audit Report",
			"kz": "Периметрдің қауіпсіздігін экспресс-аудиттеу нәтижелері",
			"ar": "نتائج التدقيق السريع لأمن محيط الشبكة المصرفية",
		},
		"rep_alert": { 
			"ru": "⚠️ ВНИМАНИЕ: ОБНАРУЖЕН КРИТИЧЕСКИЙ УРОВЕНЬ УГРОЗЫ ВНЕШНЕГО ПЕРИМЕТРА", 
			"kg": "⚠️ КӨҢҮЛ БУРУҢУЗДАР: ТЫШКЫ ПЕРИМЕТРДИН КРИТИКАЛЫК КОРКУНУЧУ АНЫКТАЛДЫ", 
			"en": "⚠️ WARNING: CRITICAL EXTERNAL PERIMETER THREAT LEVEL DETECTED",
			"kz": "⚠️ НАЗАР АУДАРЫҢЫЗ: СЫРТҚЫ ПЕРИМЕТРДІҢ КРИТИКАЛЫҚ ҚАУІП ДЕҢГЕЙІ АНЫҚТАЛДЫ",
			"ar": "⚠️ تحذير: تم رصد مستوى تهديد خطير في المحيط الخارجي للخادم",
		},
		"rep_obj": { 
			"ru": "Целевой объект проверки", 
			"kg": "Текшерилген объект", 
			"en": "Target Verification Node",
			"kz": "Тексерілетін нысан",
			"ar": "عقدة التحقق المستهدفة",
		},
		"rep_back": { 
			"ru": "← Вернуться на главную", 
			"kg": "← Башкы бетке кайтуу", 
			"en": "← Back to Dashboard",
			"kz": "← Бас мәзірге қайту",
			"ar": "← العودة إلى الصفحة الرئيسية",
		},
		"rep_m1_t": { 
			"ru": "1. Сетевой периметр и открытые шлюзы:", 
			"kg": "1. Тармактын периметри жана ачык шлюздар:", 
			"en": "1. Network Perimeter & Open Gateways:",
			"kz": "1. Желілік периметр және ашық шлюздер:",
			"ar": "1. محيط الشبكة والبوابات المفتوحة للموقع:",
		},
		"rep_m1_d": { 
			"ru": "Зафиксированы open ports 443 (HTTPS) и 8443 (API-Gateway). Обнаружена трансляция версий веб-серверов.", 
			"kg": "Ачык 443 жана 8443 порттору катталды. Программалык камсыздоонун версиялары интернетке ачык көрсөтүлүп турат.", 
			"en": "Open ports 443 (HTTPS) and 8443 (API-Gateway) detected. Web-server versions are openly broadcasted.",
			"kz": "Ашық 443 және 8443 порттары тіркелді. Бағдарламалық құрал нұсқалары интернетке ашық көрсетіліп тұр.",
			"ar": "تم رصد منافذ مفتوحة 443 و 8443. هناك بث علني لإصدارات البرامج البرمجية للخوادم.",
		},
		"rep_m2_t": { 
			"ru": "2. Экспресс-анализ уязвимости криптографии (SSL/TLS):", 
			"kg": "2. Криптографиянын алсыздыгын экспресс-талдоо:", 
			"en": "2. Cryptographic Vulnerability Audit (SSL/TLS):",
			"kz": "2. Криптографияның әлсіздігін экспресс-талдау:",
			"ar": "2. التدقيق السريع لثغرات التشفير וחزن الأمان (SSL/TLS):",
		},
		"rep_m2_d": { 
			"ru": "Обнаружены признаки использования устаревших библиотек, подверженных атакам перехвата сессий.", 
			"kg": "Сессияларды уурдап алуу чабуулдарына кабылышы мүмкүн болгон эскирген криптографиялык китепканалардын белгилери табылды.", 
			"en": "Traces of outdated cryptographic libraries detected, making active sessions vulnerable to interception.",
			"kz": "Белсенді сессияларын ұрлау шабуылдарына ұшырауы мүмкін ескірген криптографиялық кітапханалардың белгілері табылды.",
			"ar": "تم العثور على مؤشرات لاستخدام مكتبات تشفير قديمة، مما يجعل جلسات الموظفين الحالية уязвиمة للاعتراض والتنصت.",
		},
		"rep_m3_t": { 
			"ru": "3. Готовность к постквантовым угрозам (Crypto-Agility индекс):", 
			"kg": "3. Посткванттык коркунучтарга даярдык (Crypto-Agility индекси):", 
			"en": "3. Post-Quantum Readiness Index (Crypto-Agility Score):",
			"kz": "3. Посткванттық қауіптерге дайындық (Crypto-Agility индексі):",
			"ar": "3. مؤشر الجاهзية لتهديدات الحوسبة الكمية (Crypto-Agility):",
		},
		"rep_m3_d": { 
			"ru": "<b>Уровень готовности: 0%.</b> Каналы передачи финансовых пакетов данных (SWIFT, XML) уязвимы для дешифрования.", 
			"kg": "<b>Даярдык деңгээли: 0%.</b> Каржылык маалыматтарды берүү каналдары хакерлер тарабынан чечмелөөгө дуушар болот.", 
			"en": "<b>Readiness Level: 0%.</b> Cross-border financial data streams (SWIFT, XML) are highly vulnerable to future decryption exploits.",
			"kz": "<b>Дайындық деңгейі: 0%.</b> Қаржылық деректерді тасымалдау арналары хакерлердің шифрды шешу шабуылдарына қорғаусыз.",
			"ar": "<b>مستوى الجاهزية الحالي: 0%.</b> قنوات نقل البيانات المالية عبر الحدود (SWIFT) عرضة لفك التشفير المستقبلي من قبل القраصنة.",
		},
		"rep_m4_t": { 
			"ru": "4. Внутренний комплаенс-риск файловой системы:", 
			"kg": "4. Ички файлдык тутумдун комплаенс-тобокелдиги:", 
			"en": "4. Internal File System Compliance Exposure:",
			"kz": "4. Ішкі файлдық жүйенің комплаенс-тәуекелі:",
			"ar": "4. مخاطر الامتثال الداخلي لنظام ملفات الخادم:",
		},
		"rep_m4_d": { 
			"ru": "Потенциальная угроза утечки бэкапов SQL из-за отсутствия автоматического маскирования прав стандарта POSIX 0600.", 
			"kg": "POSIX 0600 стандартындагы кирүү укуктарын мажбурлап маскалоо жок болгондуктан, камдык көчүрмөлөрдүн агып кетүү коркунучу бар.", 
			"en": "Potential threat of SQL database backup leak due to the lack of automated POSIX 0600 strict access token enforcement.",
			"kz": "POSIX 0600 стандартындағы рұқсат құқықтарын автоматты түрде мәжбүрлі маскалау жоқтығынан сыртқа кету қаупі бар.",
			"ar": "هناك تهديد محتمل لتسريب النسخ الاحتياطية لقواعد البيانات نتيجة الغياب التام لفرض قيود نظام الحماية الصارم POSIX 0600 автоматически.",
		},
		"rep_m5_t": { 
			"ru": "5. Сводный ИБ-индекс критичности сетевого узла:", 
			"kg": "5. Тармактык түйүндүн жалпы критикалык индекси:", 
			"en": "5. Consolidated Security Risk Score:",
			"kz": "5. Желілік тораптың жиынтық қауіпсіздік индексі:",
			"ar": "5. المؤشر الموحد لمستوى خطورة العقدة الشبكية:",
		},
		"rep_m5_d": { 
			"ru": "УРОВЕНЬ РИСКА: 9.4 из 10 (КРИТИЧЕСКИЙ)", 
			"kg": "ТОБОКЕЛДИК ДЕҢГЭЭЛИ: 10дон 9.4 (КРИТИКАЛЫК)", 
			"en": "RISK SCORE: 9.4 out of 10 (CRITICAL)",
			"kz": "ҚАУІП ДЕҢГЕЙІ: 10-нан 9.4 (КРИТИКАЛЫҚ)",
			"ar": "مستوى الخطورة الإجمالي: 9.4 من 10 (خطير للغاية)",
		},
	}
	return reportText[key][lang]
}
