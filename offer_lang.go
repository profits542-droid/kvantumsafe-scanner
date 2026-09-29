package main

func getOfferTranslationFromB2B(lang string) string {
	if lang == "kg" {
		return `ПРОГРАММАЛЫК КАМСЫЗДООНУН ТЕХНИКАЛЫК СПЕЦИФИКАЦИЯСЫ ЖАНА КОРПОРАТИВДИК B2B СУНУШУ
Укук ээси: «Квантум Сейф» ОсООсу (Ош ш., Кыргыз Республикасы)

1. Global Scanner — АКЫСЫЗ / USD 0
2. AI-AntiHacker — USD 19,000 / жыл (Периметрди жана оперативдүү эсти азыр коргоо)
3. Quantum Web3 (SDK v2.6) — USD 35,000 / жыл (Ички сканер комплаенс-сейф POSIX 0600 + NIST ML-KEM жана GmSSL)`
	}

	if lang == "en" || lang == "ar" || lang == "kz" {
		return `OFFICIAL B2B COMMERCIAL OFFER FOR THE BANK MANAGEMENT
Developer and Trademark Owner: Quantum Safe LLC (Osh, Kyrgyz Republic)

Dear Chairman of the Board / Chief Information Security Officer!
Below is a clear, non-technical breakdown of our security software packages engineered to defend your financial institution against critical losses, logic transaction fraud, and cyber risks.

AVAILABLE SUBSCRIPTION LICENSES & PRICING:

1. "GLOBAL SCANNER" LICENSE — FREE / USD 0
2. "AI-ANTIHACKER" LICENSE — USD 19,000 / year
3. "QUANTUM WEB3" LICENSE (MAXIMUM DEFENSE AND QUANTUM IMMORTALITY) — USD 35,000 / year`
	}

	return `ОФИЦИАЛЬНОЕ КОММЕРЧЕСКОЕ ПРЕДЛОЖЕНИЕ И ТАРИФНАЯ СЕТКА ДЛЯ РУКОВОДСТВА БАНКА
Разработчик и правообладатель: ОсОО «Квантум Сейф» (г. Ош, Кыргызская Республика)

Уважаемый Председатель Правления / Директор по безопасности! 
ДОСТУПНЫЕ ВАРИАНТЫ КОРПОРАТИВНЫХ ЛИЦЕНЗИЙ И СТОИМОСТЬ:

1. ТАРИФ «GLOBAL SCANNER» — БЕСПЛАТНО / USD 0
2. ТАРИФ «AI-AntiHacker» — 19 000 долларов США / год
3. ТАРИФ «QUANTUM WEB3» (МАКСИМАЛЬНАЯ ЗАЩИТА И БЕЗОПАСНОСТЬ) — 35 000 долларов США / год`
}

// НАДЕЖНО ЗАКРЫВАЕМ ОПИСАНИЕ СТРАНИЦЫ ЗДЕСЬ, БУФЕР GITHUB ЕГО БОЛЬШЕ НЕ ОБРУБИТ!
func getAntiHackerPageText(key, lang string) string {
	texts := map[string]map[string]string{
		"ah_title": { "ru": "KvantumSafe AI-AntiHacker (Guard Engine)", "kg": "KvantumSafe AI-AntiHacker (Guard Engine)", "en": "KvantumSafe AI-AntiHacker (Guard Engine)", "kz": "KvantumSafe AI-AntiHacker (Guard Engine)", "ar": "KvantumSafe AI-AntiHacker (Guard Engine)" },
		"ah_status": {
			"ru": "<b>Статус разработки:</b> Действующий автономный оборонный программный комплекс нового поколения от ОсОО «Квантум Сейф».",
			"kg": "<b>Иштеп чыгуу статусу:</b> «Квантум Сейф» ОсООсунан жаңы муундагы өзүнчө автономдуу коргонуу программалык комплекси.",
			"en": "<b>Development Status:</b> Active next-generation autonomous cyber-defense software suite owned by Quantum Safe LLC.",
			"kz": "<b>Әзірлеу мәртебесі:</b> «Квантум Сейф» ЖШС-нің жаңа буындағы автономды қорғаныс бағдарламалық кешені.",
			"ar": "<b>حالة التطوير:</b> منظومة برمجية دفاعية مستقلة من الجيل الجديد مملوكة لشركة كوانتوم سيف.",
		},
		"ah_sub": {
			"ru": "Программный комплекс <b>«AI-AntiHacker» (Guard Engine)</b> спроектирован как универсальное сквозное решение кибербезопасности, которое идеально встает на защиту <b>Web3-проектов, криптобирж, необанков, смарт-контрактов и любых блокчейн-транзакций!</b> 🚀🛡️<br><br>Уникальное ИТ-решение на гибридном стеке <b>Go + Rust</b> с привлечением нейросетевых ИИ-моделей TinyML осуществляет превентивную защиту по следующим критическим направлениям:",
			"kg": "<b>«AI-AntiHacker» (Guard Engine)</b> программалык комплекси <b>Web3-долбоорлорун, криптобиржаларды, необанктарды, смарт-контракттарды жана каалаган блокчейн-транзакцияларды</b> коргоо үчүн универсалдуу коопсуздук чечими катары иштелип чыккан! 🚀🛡️<br><br><b>Go + Rust</b> гибриддик стегиндеги уникалдуу ИТ-чечим TinyML жасалма интеллект моделдерин колдонуу менен төмөнкү багыттар боюнча алдын ала коргоону камсыз кылат:",
			"en": "The <b>'AI-AntiHacker' (Guard Engine)</b> software suite is engineered as a universal end-to-end cybersecurity solution that perfectly fits the defense of <b>Web3 projects, crypto exchanges, neobanks, smart contracts, and any blockchain transactions!</b> 🚀🛡️<br><br>Built on a hybrid <b>Go + Rust</b> stack powered by TinyML neural models, it delivers preemptive protection across critical vectors:",
			"kz": "<b>«AI-AntiHacker» (Guard Engine)</b> бағдарламалық кешені <b>Web3 жобаларын, криптобиржаларды, необанктерді, смарт-келісімшарттарды және кез келген блокчейн-транзакцияларды</b> қорғауға арналған әмбебап қауіпсіздік шешімі ретінде жасалған! 🚀🛡️<br><br>TinyML нейрожелілік ИИ моделдері бар <b>Go + Rust</b> гибридті стегіндегі бірегей ИТ-шешім келесі бағыттар бойынша алдын ала қорғанысты жүзеге асырады:",
			"ar": "تم تصميم المنظومة البرمجية <b>'AI-AntiHacker' (Guard Engine)</b> كحل أمني شامل يناسب حماية <b>مشاريع Web3، منصات العملات الرقمية، البنوك الرقمية، العقود الذكية، وأي معاملات بلوكشين!</b> 🚀🛡️<br><br>يعتمد هذا الحل التقني الفريد على تكنولوجيا هجينة من <b>Go + Rust</b> مدعومة بنماذج الذكاء الاصطناعي TinyML لحماية الخوادم وفق المسارات التالية:",
		},
		"ah_p1": {
			"ru": "🔹 <b>Защита Web3 и смарт-контрактов (Атаки Reentrancy):</b> Встроенные TinyML алгоритмы в реальном времени анализируют тайминги транзакционных пакетов и пресекают попытки циклического вывода активов из DeFi-протоколов за 0.002 секунды до фактического списания баланса.",
			"kg": "🔹 <b>Web3 жана смарт-контракттарды коргоо (Reentrancy чабуулдары):</b> Орнотулган TinyML алгоритмдери транзакциялардын убактысын реалдуу убакытта талдап, DeFi протоколдорунан активдерди уурдоо аракеттерин 0.002 секундда токтотот.",
			"en": "🔹 <b>Web3 & Smart Contract Safety (Reentrancy Exploits):</b> Core TinyML engines analyze transaction micro-timings in real time, preventing automated balance draining from DeFi protocols 0.002 seconds before asset extraction.",
			"kz": "🔹 <b>Web3 және смарт-келісімшарттарды қорғау (Reentrancy шабуылдары):</b> Кірістірілген TinyML алгоритмдері транзакция пакеттерінің таймингтерін нақты уақытта талдап, DeFi хаттамаларынан активдерді циклдік шығару әрекеттерін 0.002 секундта тоқтатады.",
			"ar": "🔹 <b>حماية Web3 والعقود الذكية (هجمات Reentrancy):</b> تحلل خوارزميات TinyML المدمجة توقيت المعاملات في الوقت الفعلي وتمنع محركات سحب الأموال التلقائي من بروتوكولات DeFi خلال 0.002 ثانية فقط.",
		},
		"ah_p2": {
			"ru": "🔹 <b>Пресечение Flash-Loan фрода на Криптобиржах:</b> Алгоритмы ядра на Go на лету верифицируют аномальные манипуляции со сверхобъёмными пулами ликвидности на торговых площадках, защищая балансы от логических манипуляций хакеров.",
			"kg": "🔹 <b>Flash-Loan фродун алдын алуу:</b> Go тилиндеги сервердин ядросу криптобиржалардагы ири ликвиддүүлүк пулдары менен аномалдуу манипуляцияларды заматта аныктап, логикалык чабуулдардан коргойт.",
			"en": "🔹 <b>Flash-Loan Fraud Mitigation on Crypto Exchanges:</b> High-speed Go core algorithms detect and isolate anomalous liquidity pool manipulations on token ledgers, protecting corporate asset sheets from logic exploits.",
			"kz": "🔹 <b>Flash-Loan фродының алдын алу:</b> Go-дағы ядро алгоритмдері криптобиржалардағы аса ірі ликвидтілік пулдарымен аномальды манипуляцияларды заматта тексеріп, баланстарды хакерлердің логикалық шабуылдарынан қорғайды.",
			"ar": "🔹 <b>منع احتيال القروض الخاطفة في منصات التداول:</b> تتحقق خوارزميات النواة بلغة Go فوراً من التلاعبات غير الطبيعية بسيولة المنصات، مما يحمي الأصول من هجمات التلاعب بالأسعار.",
		},
		"ah_p3": {
			"ru": "🔹 <b>Безопасность Необанков и RAM-памяти:</b> Фоновый низкоуровневый демон осуществляет непрерывное сканирование оперативной памяти серверов (Stealth RAM Isolation). Любые попытки инъекции вредоносного кода или уязвимости переполнения буфера мгновенно блокируются аппаратной изоляцией секторов ядра.",
			"kg": "🔹 <b>Необанктардын жана RAM-эстин коопсуздугу:</b> Төмөнкү деңгээлдеги фондук демон серверлердин ыкчам эс тутумун (Stealth RAM Isolation) үзгүлтүксүз сканерлейт. Буфердин толуп кетиши же зыяндуу код киргизүү аракеттери заматта бөгөттөлөт.",
			"en": "🔹 <b>Neobank buffers & Runtime Memory (RAM) Isolation:</b> Low-level background daemons execute non-stop pointer verification (Stealth RAM Isolation), shielding operational memory matrix from byte overflows and binary injections.",
			"kz": "🔹 <b>Необанктер мен RAM-жад қауіпсіздігі:</b> Фондық төменгі деңгейдегі демон серверлердің жедел жадын үздіксіз сканерлеуді жүзеге асырады (Stealth RAM Isolation). Кез келген зиянды код енгізу әрекеттері ядро секторларын оқшаулаумен бұғатталады.",
			"ar": "🔹 <b>أمن البنوك الرقمية وذاكرة RAM:</b> يقوم برنامج خلفي منخفض المستوى بفحص مستمر لذاكرة الوصول العشوائي وعزل القطاعات المخترقة (Stealth RAM Isolation) لحظر أي برمجيات خبيثة فوراً.",
		},
		"ah_p4": {
			"ru": "🔹 <b>Аппаратная привязка ноды (Device Fingerprinting):</b> Софт жестко синхронизируется с неизменяемым ID процессоров физического оборудования, полностью исключая риски подмены серверов или перехвата управления (MitM).",
			"kg": "🔹 <b>Аппараттык байлоо (Device Fingerprinting):</b> Программалык камсыздоо сервердин процессорунун өзгөрүлбөс заводдук ID номерине бекем байланат. Бул серверди алмаштыруу же трафикти уурдоо (MitM) тобокелдиктерин жокко чыгарат.",
			"en": "🔹 <b>Hardware Binding (Device Fingerprinting):</b> The software layer explicitly synchronizes with the unalterable CPU core factory identifier of the node, eliminating server spoofing and Man-in-the-Middle (MitM) interception vectors.",
			"kz": "🔹 <b>Аппараттық байланыстыру (Device Fingerprinting):</b> Бағдарламалық құрал серверлік торап процессорларының өзгермейтін зауыттық ID-імен қатаң синхрондалады, бұл серверлерді ауыстыру немесе басқаруды басып алу (MitM) қауіптерін толық жояды.",
			"ar": "🔹 <b>الارتباط بالأجهزة (Device Fingerprinting):</b> ترتبط البرمجيات بشكل صارم بالرقم المصنعي غير القابل للتغيير لمعالج الخادم لمنع أي محاولات لتزوير هوية الخوادم (MitM).",
		},
	}
	return texts[key][lang]
}
