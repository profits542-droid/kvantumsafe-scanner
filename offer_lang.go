package main

// Функция связи перенаправляет золотую кнопку на прямые стабильные файлы спецификаций
func getOfferTranslationFromB2B(lang string) string {
	return "https://githubusercontent.com_" + lang + ".txt"
}

// Легкая функция переводов для страницы АнтиХакера, без тяжелых текстовых массивов
func getAntiHackerPageText(key, lang string) string {
	texts := map[string]map[string]string{
		"ah_title": {
			"ru": "KvantumSafe AI-AntiHacker (Guard Engine)",
			"kg": "KvantumSafe AI-AntiHacker (Guard Engine)",
			"en": "KvantumSafe AI-AntiHacker (Guard Engine)",
			"kz": "KvantumSafe AI-AntiHacker (Guard Engine)",
			"ar": "KvantumSafe AI-AntiHacker (Guard Engine)",
		},
		"ah_status": {
			"ru": "<b>Статус разработки:</b> Действующий автономный программный комплекс нового поколения.",
			"kg": "<b>Иштеп чыгуу статусу:</b> Жаңы муундагы өзүнчө автономдуу коргонуу программалык комплекси.",
			"en": "<b>Development Status:</b> Active next-generation autonomous cyber-defense software suite.",
			"kz": "<b>Әзірлеу мәртебесі:</b> Жаңа буындағы автономды қорғаныс бағдарламалық кешені.",
			"ar": "<b>حالة التطوير:</b> منظومة برمجية دفاعية مستقلة ونشطة من الجيل الجديد.",
		},
		"ah_sub": {
			"ru": "Программный комплекс <b>«AI-AntiHacker»</b> идеально встает на защиту <b>Web3-проектов, криптобирж, необанков, смарт-контрактов и блокчейн-транзакций!</b> 🚀🛡️",
			"kg": "<b>«AI-AntiHacker»</b> программалык комплекси <b>Web3-долбоорлорун, криптобиржаларды, необанктарды жана блокчейн-транзакцияларды</b> коргоо үчүн иштелип чыккан! 🚀🛡️",
			"en": "The <b>'AI-AntiHacker'</b> software suite perfectly fits the defense of <b>Web3 projects, crypto exchanges, neobanks, and blockchain transactions!</b> 🚀🛡️",
			"kz": "<b>«AI-AntiHacker»</b> кешені <b>Web3 жобаларын, криптобиржаларды, необанктерді және блокчейн-транзакцияларды</b> қорғауға арналған! 🚀🛡️",
			"ar": "تم تصميم المنظومة البرمجية <b>'AI-AntiHacker'</b> كحل أمني شامل يناسب حماية <b>مشاريع Web3، منصات العملات الرقمية، البنوك الرقمية، ومعاملات блокчейн!</b> 🚀🛡️",
		},
		"ah_p1": { "ru": "🔹 <b>Защита Web3:</b> Пресекает атаки Reentrancy на смарт-контракты за 0.002 сек.", "kg": "🔹 <b>Web3 Коргоо:</b> Смарт-контракттардагы Reentrancy чабуулдарын 0.002 секундда бөгөттөө.", "en": "🔹 <b>Web3 Safety:</b> Stops smart contract Reentrancy attacks within 0.002 seconds.", "kz": "🔹 <b>Web3 Қорғау:</b> Смарт-келісімшарттардағы Reentrancy шабуылдарын 0.002 секундта тоқтату.", "ar": "🔹 <b>حماية Web3:</b> تمنع هجمات Reentrancy على العقود الذكية خلال 0.002 ثانية." },
		"ah_p2": { "ru": "🔹 <b>Криптобиржи:</b> Верификация и блокировка Flash-Loan фрода пулов ликвидности.", "kg": "🔹 <b>Криптобиржалар:</b> Ири ликвиддүүлүк пулдары менен аномалдуу фродду бөгөттөө.", "en": "🔹 <b>Crypto Exchanges:</b> Real-time Flash-Loan fraud detection and ledger isolation.", "kz": "🔹 <b>Криптобиржалар:</b> Ликвидтілік пулдарымен аномальды фродты заматта бұғаттау.", "ar": "🔹 <b>منصات التداول:</b> كشف ومنع احتيال القروض الخاطفة Flash-Loan فوراً." },
		"ah_p3": { "ru": "🔹 <b>Необанки:</b> Скрытый контроль и Stealth-изоляция секторов оперативной памяти (RAM).", "kg": "🔹 <b>Необанктар:</b> Серверлердин ыкчам эс тутумун (Stealth RAM Isolation) үзгүлтүксүз сканерлөө.", "en": "🔹 <b>Neobanks:</b> Low-level background daemons execute non-stop Stealth RAM Isolation.", "kz": "🔹 <b>Необанктер:</b> Жедел жадты үздіксіз сканерлеу (Stealth RAM Isolation).", "ar": "🔹 <b>البنوك الرقمية:</b> فحص مستمر لذاكرة الوصول العشوائي وعزل القطاعات المخترقة." },
		"ah_p4": { "ru": "🔹 <b>Привязка к железу:</b> Жесткая синхронизация с неизменяемым заводским ID процессора.", "kg": "🔹 <b>Аппараттык байлоо:</b> Процессордун өзгөрүлбөс заводдук ID номерине бекем байланат.", "en": "🔹 <b>Hardware Binding:</b> Locks server execution to the unalterable CPU core factory ID.", "kz": "🔹 <b>Аппараттық байланыс:</b> Процессордың зауыттық ID нөміріне қатаң байланыстыру.", "ar": "🔹 <b>الارتباط بالأجهزة:</b> ترتبط البرمجيات بشكل صارم بالرقم المصнعي لمعالج الخادم." },
	}
	return texts[key][lang]
}
