package main

func getTableTranslation(key, lang string) string {
	tableTexts := map[string]map[string]string{
		"tbl_main_title": {
			"ru": "📊 Сравнение технологий: Обычные антивирусы и Программное обеспечение «АнтиХакер AI»",
			"kg": "📊 Технологияларды салыштыруу: Жөнөкөй антивирустар жана «АнтиХакер AI» Программалык камсыздоосу",
			"en": "📊 Technology Comparison: Traditional Antiviruses vs 'AI-AntiHacker' Software Suite",
			"kz": "📊 Technology Салыштыруу",
			"ar": "📊 مقارنة التقنيات",
		},
		"tbl_h1": { "ru": "Критерий защиты инфраструктуры", "kg": "Инфраструктураны коргоо критерийи", "en": "Infrastructure Protection Criterion" },
		"tbl_h2": { "ru": "Обычный антивирус / Сканер", "kg": "Жөнөкөй антивирус / Сканер", "en": "Traditional Antivirus / Scanner" },
		"tbl_h3": { "ru": "Программное обеспечение «АнтиХакер AI»", "kg": "«АнтиХакер AI» Программалык камсыздоосу", "en": "'AI-AntiHacker' Software Suite" },
		"tbl_r1_c1": { "ru": "Принцип обнаружения угроз", "kg": "Коркунучтарды аныктоо принциби", "en": "Threat Detection Principle" },
		"tbl_r1_b_no": { "ru": "Только по базам (Сигнатурный)", "kg": "Базалар аркылуу гана (Сигнатуралык)", "en": "Database Only" },
		"tbl_r1_d_no": { "ru": "Ищет только старые, уже известные вирусы.", "kg": "Эски вирустарды издейт.", "en": "Known threats only." },
		"tbl_r1_b_yes": { "ru": "Проактивный нейросетевой TinyML", "kg": "Проактивдүү нейротармактык TinyML", "en": "Proactive TinyML Engine" },
		"tbl_r1_d_yes": { "ru": "Выявляет новые угрозы нулевого дня на лету.", "kg": "Жаңы коркунучтарды заматта аныктайт.", "en": "Detects zero-day threats." },
		"tbl_r2_c1": { "ru": "Защита от логических атак фрода", "kg": "Фрод логикалык чабуулдарынан коргоо", "en": "Fraud Logic Attack Safety" },
		"tbl_r2_b_no": { "ru": "Отсутствует полностью", "kg": "Толугу менен жок", "en": "Completely Missing" },
		"tbl_r2_d_no": { "ru": "Не видит манипуляции со смарт-контрактами.", "kg": "Манипуляцияны көрбөйт.", "en": "No protection." },
		"tbl_r2_b_yes": { "ru": "Блокировка за 0.002 секунды", "kg": "0.002 секундда бөгөттөө", "en": "0.002s Auto-Block" },
		"tbl_r2_d_yes": { "ru": "Останавливает Reentrancy и Flash-Loan атаки.", "kg": "Чабуулдарды токтотот.", "en": "Stops Reentrancy attacks." },
		"tbl_r3_c1": { "ru": "Защита оперативной памяти (RAM)", "kg": "Оперативдүү эсти (RAM) коргоо", "en": "Memory Isolation" },
		"tbl_r3_b_yes": { "ru": "Stealth-изоляция секторов RAM", "kg": "RAM секторлорун Stealth-изоляциялоо", "en": "Stealth RAM Isolation" },
		"tbl_r4_c1": { "ru": "Аппаратная верификация нод", "kg": "Түйүндөрдү аппараттык текшерүү", "en": "Hardware Verification" },
		"tbl_r4_d_no": { "ru": "Уязвим к подмене серверов (MitM).", "kg": "Алмаштыруусуна туруштук бербейт.", "en": "Vulnerable to MitM." },
		"tbl_r4_b_yes": { "ru": "Rust Device Fingerprinting", "kg": "Rust Device Fingerprinting технологиясы", "en": "Rust Fingerprinting" },
		"tbl_r4_d_yes": { "ru": "Жестко привязывается к ID процессора.", "kg": "Процессордун өзгөрүлбөс ID номерине байланат.", "en": "Locks code execution to CPU ID." },
		"tbl_r5_c1": { "ru": "Юридическая чистота (Без СКЗИ)", "kg": "Юридикалык тазалык (СКЗИсиз)", "en": "Legal Compliance" },
		"tbl_r5_d_no": { "ru": "Сложный комплаенс-контроль софта.", "kg": "Татаал комплаенс талап кылынат.", "en": "Complex compliance rules." },
		"tbl_r5_b_yes": { "ru": "100% Свободное обращение", "kg": "100% Эркин жүгүртүү", "en": "100% Globally Compliant" },
		"tbl_r5_d_yes": { "ru": "Не содержит СКЗИ, не требует лицензий.", "kg": "СКЗИ камтыбайт, лицензияларды талап кылбайт.", "en": "Allows rapid deploy." },
	}
	
	// Безопасный фолбэк для неописанных языков на русский язык
	res, ok := tableTexts[key][lang]
	if !ok || res == "" {
		res = tableTexts[key]["ru"]
	}
	return res
}
