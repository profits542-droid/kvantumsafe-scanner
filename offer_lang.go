package main

func GetFullDownloadOffer(lang string) string {
	if lang == "kg" {
		return `ПРОГРАММАЛЫК КАМСЫЗДООНУН ТЕХНИКАЛЫК СПЕЦИФИКАЦИЯСЫ ЖАНА КОРПОРАТИВДИК B2B СУНУШУ
Укук ээси: «Квантум Сейф» ОсООсу (Ош ш., Кыргыз Республикасы)

Урматтуу Башкармалыктын Төрагасы / Коопсуздук боюнча директор!
Төмөндө банкты каржылык жоготуулардан жана хакердик чабуулдардан коргой турган программалык пакеттердин жөнөкөй жана түшүнүктүү сыпаттамасы келтирилген.

ЖЕТКИЛИКТҮҮ КОРПОРАТИВДИК ЛИЦЕНЗИЯЛАР ЖАНА БААСЫ:

1. ТАРИФ «GLOBAL SCANNER» — АКЫСЫЗ / USD 0
- БУЛ ЭМНЕ: Сырткы тармактык периметрди тез арада экспресс-талдоо куралы.
- БАНККА КАНДАЙ ЖАРДАМ БЕРЕТ: Сайтыңыздын коопсуздугун реалдуу убакытта текшерет. ИТ-адистериңиз кайсы шлюздарды жабууну унутуп калганын көрсөтөт.

2. ТАРИФ «AI-AntiHacker» — жылына 19 000 АКШ доллары
- БУЛ ЭМНЕ: Хакердик чабуулдардан жана акча уурдоодон коргоо үчүн банктын серверлеринин ичине орнотулуучу кубаттуу коргоочу калкан (АЗЫР КОРГОО).
- БАНККА КАНДАЙ ЖАРДАМ БЕРЕТ:
  - Компьютердик жабдуунун санариптик паспорту: Программа сервериңиздин процессорунун заводдук номерине бекем байланат. Хакер сырттан башкарууну колго ала албайт.
  - Алдамчылыкты тез арада бөгөттөө: Жасалма интеллект төлөмдөрдү текшерип, 0.002 секунддун ичинде акчаны уурдоо схемаларын автоматтык түрдө токтотот.
  - Оперативдүү эстин коргоочусу: Серверлердин эс тутумун (RAM) күнү-түнү көзөмөлдөп, зыяндуу коддорду киргизүү аракеттерин бөгөттөйт.

3. ТАРИФ «QUANTUM WEB3» (МАКСИМАЛДУУ КОРГОО ЖАНА КООПСУЗДУК) — жылына 35 000 АКШ доллары
- БУЛ ЭМНЕ: KvantumSafe Pro Framework SDK (v2.6) тутумунун базасындагы эң жогорку деңгээлдеги коргоо пакети. 
- ПАКЕТКЕ ЭМНЕЛЕР КИРЕТ:
  - «AI-AntiHacker» тарифиндеги бардык коргоо модулдары (АЗЫР КОРГОО).
  - Акплиенс-сейф (Маалымат базаларын ички сканерлөө): Автоматтык фонодук ревизор. Ал SQL маалымат базаларынын камдык көчүрмөлөрүн (бэкап) таап, аларды POSIX 0600 стандартынын алдында жашырат. Хакер ички тармакка кирсе дагы, кардарлардын маалыматтарын көчүрө же өчүрө албайт.
  - Суперкомпьютерлерге каршы коргоочу бронеконтейнер (КИЙИН КОРГОО): Эл аралык которууларды (SWIFT, XML) АКШ (NIST ML-KEM) и Кытайдын (GmSSL SM4-GCM) жаңы посткванттык криптоконтейнерлерине таңгактайт. Бул маалыматтарды кийин кванттык компьютерлерде бузуудан коргойт.`
	}

	if lang == "en" || lang == "ar" || lang == "kz" {
		return `OFFICIAL B2B COMMERCIAL OFFER FOR THE BANK MANAGEMENT
Developer and Trademark Owner: Quantum Safe LLC (Osh, Kyrgyz Republic)

Dear Chairman of the Board / Chief Information Security Officer!
Below is a clear, non-technical breakdown of our security software packages engineered to defend your financial institution against critical losses, logic transaction fraud, and cyber risks.

AVAILABLE SUBSCRIPTION LICENSES & PRICING:

1. "GLOBAL SCANNER" LICENSE — FREE / USD 0
- DESCRIPTION: Basic external network perimeter express audio tool.
- HOW IT HELPS THE BANK: Real-time validation of your website gates. Shows which cloud ports were left open by your IT team and where hackers can attempt to initiate an intrusion.

2. "AI-ANTIHACKER" LICENSE — USD 19,000 / year
- DESCRIPTION: Powerful proactive shield deployed inside the server environment of the bank to block ongoing exploits and fund theft NOW.
- HOW IT HELPS THE BANK:
  - Hardware Device Fingerprinting: Securely binds software execution to the unique server CPU factory IDs. Prevents Middle-man spoofing attacks.
  - Active Fraud Interception: Artificial intelligence analyzes transaction timing parameters and automatically blocks sophisticated logic exploits within 0.002 seconds, preventing money from leaving the bank.
  - Core Memory Guard: Provides background tracking of operational server RAM, instantly blocking hacker code injections into running processes.

3. "QUANTUM WEB3" LICENSE (MAXIMUM DEFENSE AND QUANTUM IMMORTALITY) — USD 35,000 / year
- DESCRIPTION: Supreme enterprise protection package running on KvantumSafe Pro Framework SDK (v2.6) core. Includes absolutely all features in one click.
- WHAT IS INCLUDED IN THE PACKAGE:
  - All defensive active monitors from the "AI-AntiHacker" license (Protection NOW).
  - Stealth Compliance Vault (Internal Database Scanner): Automated background auditor that searches for SQL backup copies, erases them from common visibility, and hides them under strict POSIX 0600 access tokens. If an intruder breaks in, they cannot harvest or wipe your customer records.
  - Anti-Supercomputer Shield (Protection LATER): Encapsulates incoming and outgoing SWIFT / XML data batches inside next-generation quantum-resistant lattice crypto-containers complying with US standards (NIST ML-KEM) and Chinese state regulations GmSSL (SM4-GCM). Guarantees that data harvested now cannot be cracked by quantum systems later.`
	}

	return `ОФИЦИАЛЬНОЕ КОММЕРЧЕСКОЕ ПРЕДЛОЖЕНИЕ И ТАРИФНАЯ СЕТКА ДЛЯ РУКОВОДСТВА БАНКА
Разработчик и правообладатель: ОсОО «Квантум Сейф» (г. Ош, Кыргызская Республика)

Уважаемый Председатель Правления / Директор по безопасности! 
Ниже представлена простая и понятная расшифровка защитных программных пакетов, которые оградят ваш банк от финансовых потерь, кражи денег со счетов и репутационных рисков.

ДОСТУПНЫЕ ВАРИАНТЫ КОРПОРАТИВНЫХ ЛИЦЕНЗИЙ И СТОИМОСТЬ:

1. ТАРИФ «GLOBAL SCANNER» — БЕСПЛАТНО / USD 0
- ЧТО ЭТО ТАКОЕ: Базовый поверхностный инструмент («экспресс-термометр»).
- КАК ПОМОГАЕТ БАНКУ: Проверяет внешние двери вашего сайта в реальном времени. Показывает, какие шлюзы забыли закрыть ваши ИТ-специалисты и где хакеры могут начать прощупывать систему.

2. ТАРИФ «AI-AntiHacker» — 19 000 долларов США / год
- ЧТО ЭТО ТАКОЕ: Мощный невидимый щит, который ставится внутрь серверов банка для защиты от взломов и воровства денег ПРЯМО СЕЙЧАС.
- КАК ПОМОГАЕТ БАНКУ:
  - Цифровой паспорт компьютерного железа: Программа намертво привязывается к заводскому номеру процессора вашего банка. Хакер не сможет подменить сервер или перехватить управление транзакциями снаружи.
  - Скоростной блокировщик умного мошенничества: Искусственный интеллект проверяет платежи и за 0.002 секунды автоматически останавливает хитрые хакерские схемы по выводу балансов со счетов до того, как деньги уйдут.
  - Невидимый охранник оперативной памяти: Круглосуточно следит за памятью серверов и блокирует любые попытки злоумышленников внедрить вирусный код прямо в работающие процессы банка.

3. ТАРИФ «QUANTUM WEB3» (МАКСИМАЛЬНАЯ ЗАЩИТА И БЕЗОПАСНОСТЬ) — 35 000 долларов США / год
- ЧТО ЭТО ТАКОЕ: Абсолютная защита высшего уровня на базе системы KvantumSafe Pro Framework SDK (версия 2.6). Включает в себя вообще все наши технологии в один клик.
- ЧТО ВХОДИТ В ПАКЕТ:
  - Все защитные модули от взломов и кражи денег из тарифа «AI-AntiHacker» (Защита ПРЯМО СЕЙЧАС).
  - Умный комплаенс-сейф (Внутренний сканер контроля баз данных): Автоматический фоновый ревизор. Он на лету находит важные резервные копии баз данных SQL, мгновенно стирает их из общей видимости и прячет под строжайшие права доступа стандарта POSIX 0600. Если хакер проникнет внутрь сети, он физически не сможет скопировать или стереть ваши бэкапы клиентов.
  - Защитный бронеконтейнер против суперкомпьютеров будущего (Защита ПОТОМ): Упаковывает международные переводы (SWIFT, XML) в новые защитные криптоконтейнеры по строгим стандартам США (NIST ML-KEM) и суверенным стандартам Китая GmSSL (SM4-GCM). Это гарантирует, что хакеры не смогут украсть ваши данные сейчас, чтобы расшифровать их потом на квантовых компьютерах.

Программный комплекс поставляется на физической защищенной Flash-карте с жесткой привязкой к оборудованию Вашего банка.`
}
