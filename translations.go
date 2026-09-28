package main

func getTranslation(key, lang string) string {
	translations := map[string]map[string]string{
		"nav_title": {
			"kg": "🛡️ KvantumSafe Pro Эл аралык ИТ-Платформасы",
			"ru": "🛡️ Международная ИТ-Платформа KvantumSafe Pro",
			"en": "🛡️ KvantumSafe Pro International IT Platform",
			"kz": "🛡️ KvantumSafe Pro Халықаралық ИТ-Платформасы",
			"ar": "🛡️ منصة كيفانتوم سيف برو الدولية لتكنولوجيا المعلومات",
		},
		"nav_sub": {
			"kg": "🚨 КӨҢҮЛ БУРУҢУЗ! Программное обеспечение «AI-AntiHacker» азыр периметрди чабуулдардан коргойт. KvantumSafe Pro Framework SDK (v2.6) каржылык маалыматтарды азыр уурдоодон жана кийин кванттык компьютерлердин бузуусунан коргойт.",
			"ru": "🚨 ВНИМАНИЕ! Программное обеспечение «AI-AntiHacker» необходимо всем СЕЙЧАС: защищает периметр от взломов и фрода на лету. Комплекс KvantumSafe Pro Framework SDK (версия 2.6) блокирует скачивание данных сейчас и гарантирует защиту от взлома квантовыми компьютерами потом.",
			"en": "🚨 ATTENTION! 'AI-AntiHacker' software is required by everyone NOW: it protects the perimeter from breaches and active fraud. KvantumSafe Pro Framework SDK (v2.6) stops data harvesting now and guarantees safety against quantum computer attacks later.",
			"kz": "🚨 НАЗАР АУДАРЫҢЫЗ! «AI-AntiHacker» қазір периметрді бұзудан қорғайды. KvantumSafe Pro Framework SDK (v2.6) деректерді қазір ұрлаудан жана кейін кванттық компьютерлердин бұзуынан қорғайды.",
			"ar": "🚨 انتباه! نظام مضاد الاختراق بالذكاء الاصطناعي مطلوب الآن لحماية الخوادم فوراً",
		},
		"scan_title": { 
			"ru": "🔍 Проверить безопасность сайта / Check Website Security", 
			"kg": "🔍 Сетевой периметрди реалдуу убакытта экспресс-аудиттөө", 
			"en": "🔍 Check Website Security / Verify Perimeter", 
			"kz": "🔍 Желілік қауіпсіздікті тексеру / Периметрді аудиттеу", 
			"ar": "🔍 تحقق من أمن الموقع / تدقيق محيط الشبكة",
		},
		"scan_btn": { 
			"ru": "Сканировать", "kg": "Аудитти баштоо", "en": "Scan Now", "kz": "Сканерлеу", "ar": "Фحص الآن",
		},
		"doc_btn": { 
			"ru": "📄 Скачать спецификацию софта и B2B-оффер (TXT)", "kg": "📄 Программалык камсыздоо спецификациясын жүктөө (TXT)", "en": "📄 Download Software Specification (TXT)", "kz": "📄  Бағдарламалық құралдың сипаттамасын жүктеу (TXT)", "ar": "📄 تنزيل مواصفات البرامج (TXT)",
		},
		"build_title": { 
			"ru": "⚙️ Верификация архитектуры и безопасности ядра софта:", "kg": "⚙️ Коддук базанын коопсуздук жана тазалык статусу:", "en": "⚙️ Core Codebase Verification Status:", "kz": "⚙️ Ядролық код базасының қауіпсіздігі:", "ar": "⚙️ حالة التحقق ونقاء قاعدة كود النواة الأساسية:",
		},
		"build_desc": {
			"ru": "Исходный программный код ядра полностью прошёл автоматическую верификацию облачной компиляцией Render. Наличие официального статуса Live подтверждает абсолютное отсутствие синтаксических ошибок, уязвимостей переполнения буфера памяти и гарантирует строгое соответствие кодовой базы стандартам безопасности.",
			"kg": "Программалык камсыздоонун ядросу автоматтык түрдө толук тексеруден өттү. Расмий Live компиляция статусу синтаксистик каталардын жоктугун тастыктайт.",
			"en": "The production source code of the core has successfully passed automated verification. The official Live compilation status verifies the absolute absence of syntax flaws.",
			"kz": "Бағдарламалық құралдың ядросу автоматты түрде толук тексеруден өттү.",
			"ar": "لقد اجتاز الكود البرمجي الأساسي لنواة النظام عملية التحقق التلقائي بنجاح.",
		},
		"render_badge": { 
			"ru": "Проверено безопасной сборкой Render Cloud (Passed)", "kg": "Render Cloud коопсуз жыйындысы тарабынан текшерилди (Passed)", "en": "Verified by Render Cloud Security (Passed)", "kz": "Render Cloud қауіпсіз жинағымен тексерілді (Passed)", "ar": "تم التحقق منه بواسطة أمان Render السحابي",
		},
		"saas_title": { 
			"ru": "💰 Лицензии и Стоимость софта:", "kg": "💰 Лицензиялар", "en": "💰 SaaS Corporate Subscription Licenses", "kz": "💰 SaaS корпоративтік жазылым лицензиялары", "ar": "💰 تراخيص الشركات لاشتрак SaaS",
		},
		"t1_free": { "ru": "FREE / $0", "kg": "Жылдык / $0", "en": "FREE / $0", "kz": "FREE / $0", "ar": "مجاني / $0" },
		"t2_pro": { "ru": "$19,000 / год", "kg": "$19,000 / жыл", "en": "$19,000 / year", "kz": "$19,000 / жыл", "ar": "$19,000 / سنوياً" },
		"t3_web3": { "ru": "$35,000 / год", "kg": "$35,000 / жыл", "en": "$35,000 / year", "kz": "$35,000 / жыл", "ar": "$35,000 / سنوياً" },
		"link_spec": { 
			"ru": "📄 Описание кнопки / Specification Info", "kg": "📄 Түйнөктүн сыпаттамасы / Specification Info", "en": "📄 Yellow Button Specification Info", "kz": "📄 Түйменің сипаттамасы / Specification Info", "ar": "📄 تفاصيل مواصفات الزر الأصفر",
		},
		"link_anti": { 
			"ru": "🤖 Софт «AI-AntiHacker» / AI-AntiHacker", "kg": "🤖 «AI-AntiHacker» софту / AI-AntiHacker", "en": "🤖 'AI-AntiHacker' Software Presentation", "kz": "🤖 «AI-AntiHacker» БҚ / AI-AntiHacker", "ar": "برمجيات 'AI-AntiHacker'",
		},
		"bot_welcome": { 
			"ru": "Здравствуйте! Я ИИ-консультант ОсОО «Квантум Сейф». Готов ответить на ваши вопросы по софту, NIST или GmSSL!", "kg": "Саламатсызбы! Мен ОсОО «Квантум Сейф» ИИ-консультантымын. Суроолоруңузду берсеңиз болот.", "en": "Hello! I am the AI Assistant of Quantum Safe LLC. Feel free to ask about our software, NIST or GmSSL.", "kz": "Сәлеметсіз бе! Мен «Квантум Сейф» ЖШС ИИ-консультантымын.", "ar": "مرحباً! أنا المستشار الذكي لشركة 'كوانتوم сэйф'.",
		},
		"bot_placeholder": { 
			"ru": "Задать вопрос ИИ...", "kg": "Текст жазыңыз...", "en": "Ask AI Assistant...", "kz": "Сұрақ қою...", "ar": "اسأل المستشار الذكي...", 
		},
		"bot_ans_nist": { 
			"ru": "Ядро KvantumSafe Pro оркестрирует постквантовые алгоритмы решеток стандарта NIST ML-KEM и суверенные азиатские криптопротоколы GmSSL (SM4-GCM). Система автоматически выбирает оптимальный маршрут данных, исключая риски дешифрования транзакций хакерами.", "kg": "KvantumSafe Pro ядросу NIST ML-KEM посткванттык алгоритмдерин оркестрациялайт.", "en": "The KvantumSafe Pro core orchestrates next-generation post-quantum lattice algorithms (NIST ML-KEM) and Asian sovereign protocols (GmSSL SM4-GCM).", "kz": "<b>NIST (ML-KEM)</b> — хакерлер мен болашақ кванттық компьютерлердің транзакцияларды дешифрлеуінен қорғайтын халықаралық посткванттық тор криптографиясының стандарты.", "ar": "تنسق نواة KvantumSafe Pro خوارزميات الشبكة لما بعد العصر الكمي.",
		},
		"bot_ans_meet": { 
			"ru": "Отличное решение! Наш Генеральный директор готов провести личную техническую презентацию контура безопасности. Оставьте ваши контакты.", "kg": "Биздин Башкы директорубуз коопсуздук контуру боюнча сизге жеке презентация өткровүүгө даяр. Байланыш маалымаңызды калтырыңыз.", "en": "Excellent choice! Our General Director is ready to conduct a personal technical presentation. Please leave your contact details.", "kz": "Өте жақсы шешім! Біждің Бас директорымыз қауіпсіздік контуру бойынша сізге жеке техникалық таныстырылым өткізуге дайын.", "ar": "قرار ممتاز! مديرنا العام مستعد لإجراء عرض فني شخصي لمنظومة الأمان.",
		},
		"bot_ans_default": { 
			"ru": "Platforma на языке Go обеспечивает автоматический комплаенс-контроль, сканирование внешнего периметра (порт 443) и фоновую изоляцию файлов под права POSIX 0600.", "kg": "Платформа автоматтык комплаенс-контролду, сыкткы периметрди сканерлөөнү (443-порт) жана файлдарды 0600 стандартына жашыруун которууну камсыз кылат.", "en": "The platform developed in Go provides automated compliance control, external perimeter audits (port 443), and stealth file isolation under strict POSIX 0600 tokens.",
			"kz": "«Квантум Сейф» ЖШС Go тілінде әзірлеген платформа автоматты комплаенс-бақылауды.",
			"ar": "توفر المنصة التي طورتها شركة 'كوانتوم сэйф' بلغة Go رقابة تلقائية.",
		},
	}
	res, ok := translations[key][lang]
	if !ok || res == "" {
		res = translations[key]["ru"]
	}
	return res
}
