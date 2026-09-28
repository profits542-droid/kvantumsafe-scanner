
```go
package main

func GetSpecificationPageHTML(lang string) string {
	title := "Технический регламент и функционал Золотой Кнопки"
	desc := "Настоящий веб-интерфейс реализует функцию динамической генерации и безопасного депонирования технической документации. При активации триггера (нажатии на кнопку) ядро сервера на Go на лету определяет активную языковую локализацию клиентской сессии и формирует официальный защищенный файл спецификации (TXT) для предоставления ИТ-департаментам и комплаенс-контролю финансовых организаций.<br><br>Сгенерированный документ содержит полные архитектурные параметры платформы KvantumSafe Pro Framework, подробную разбивку коммерческой тарифной сетки (Global Scanner, ПО АнтиХакер AI, Quantum Web3) и легальное обоснование юридической чистоты софта перед государственными регуляторами. Файл принудительно сохраняется в локальное хранилище (папку Загрузки) ноутбука или персонального компьютера пользователя по протоколу контентной диспетчеризации (Content-Disposition)."
	back := "← Назад / Back"

	if lang == "kg" {
		title = "Алтын Түйнөктүн Техникалык Регламенти"
		desc = "Бул веб-интерфейс программалык камсыздоонун техникалык документтерин динамикалык түрдө генерациялоо жана криптографиялык коопсуз депонирлөө функциясын аткарат. Түймени басканда, Go тилиндеги сервердин ядросу кардардын сессиясынын тилин аныктайт жана өзгөчөлөнгөн тексттик документти (TXT) түзөт."
		back = "← Артка"
	}
	return "<html><head><meta charset='UTF-8'><title>Specification</title></head><body style='font-family:sans-serif;padding:40px;line-height:1.6;max-width:850px;margin:auto;background:#f4f7f6;'><p><a href='/?lang=" + lang + "' style='font-weight:bold;text-decoration:none;color:#0a2540;'>" + back + "</a></p><h2 style='color:#0a2540; border-bottom:2px solid #cbd5e1; padding-bottom:10px;'>🛡️ " + title + "</h2><p style='font-size:16px; color:#334155; text-align:justify;'>" + desc + "</p></body></html>"
}

func GetAntiHackerPageHTML(lang string) string {
	title := "KvantumSafe AI-AntiHacker (Guard Engine)"
	desc := "<b>Статус разработки:</b> Автономный оборонный программный комплекс нового поколения от ОсОО «Квантум Сейф».<br><br>Уникальное ИТ-решение на гибридном стеке <b>Go + Rust</b> с привлечением нейросетевых ИИ-моделей TinyML. Комплекс осуществляет фоновый контроль оперативной памяти, считывает уникальный аппаратный ID процессоров (Device Fingerprinting) и на лету блокирует хакерские логические атаки Reentrancy и Flash-Loan фрода транзакций финансовых организаций."
	back := "← Назад / Back"

	if lang == "kg" {
		title = "KvantumSafe AI-AntiHacker (Guard Engine)"
		desc = "<b>Иштеп чыгуу статусу:</b> «Квантум Сейф» ОсООсунан жаңы муундагы өзүнчө автономдуу коргонуу программалык комплекси."
		back = "← Артка"
	}
	return "<html><head><meta charset='UTF-8'><title>AI-AntiHacker</title></head><body style='font-family:sans-serif;padding:40px;line-height:1.6;max-width:850px;margin:auto;background:#0a2540;color:white;'><p><a href='/?lang=" + lang + "' style='color:#cbd5e1;text-decoration:none;font-weight:bold;'> " + back + "</a></p><h2 style='color:#d4af37;'>🤖 " + title + "</h2><p style='font-size:16px; text-align:justify; line-height:1.7; color:#f1f5f9;'>" + desc + "</p></body></html>"
}

// УЛЬТРА-ЧИСТАЯ ПОЛНОСТЬЮ ЛОКАЛИЗОВАННАЯ МАТРИЦА СРАВНЕНИЯ ТЕХНОЛОГИЙ БЕЗ ОБРЕЗАНИЙ КАВЫЧЕК
func GetComparisonTableHTML(lang string) string {
	return `<h3 class='comp-heading'>` + getTranslation("tbl_main_title", lang) + `</h3>
		<table class='comp-table'>
			<tr>
				<th style='width:40%;'>` + getTranslation("tbl_h1", lang) + `</th>
				<th style='width:30%; background-color:#475569;'>` + getTranslation("tbl_h2", lang) + `</th>
				<th style='width:30%; background-color:#1e3a8a;'>` + getTranslation("tbl_h3", lang) + `</th>
			</tr>
			<tr class='comp-row'>
				<td><b>` + getTranslation("tbl_c1_t", lang) + `</b></td>
				<td><span class='badge-no'>` + getTranslation("tbl_c1_no", lang) + `</span><br><small style='color:#64748b;'>` + getTranslation("tbl_c1_no_sub", lang) + `</small></td>
				<td><b><span class='badge-yes'>` + getTranslation("tbl_c1_yes", lang) + `</span></b><br><small style='color:#64748b;'>` + getTranslation("tbl_c1_yes_sub", lang) + `</small></td>
			</tr>
			<tr class='comp-row'>
				<td><b>` + getTranslation("tbl_c2_t", lang) + `</b></td>
				<td><span class='badge-no'>` + getTranslation("tbl_c2_no", lang) + `</span><br><small style='color:#64748b;'>` + getTranslation("tbl_c2_no_sub", lang) + `</small></td>
				<td><b><span class='badge-yes'>` + getTranslation("tbl_c2_yes", lang) + `</span></b><br><small style='color:#64748b;'>` + getTranslation("tbl_c2_yes_sub", lang) + `</small></td>
			</tr>
			<tr class='comp-row'>
				<td><b>` + getTranslation("tbl_c3_t", lang) + `</b></td>
				<td><span class='badge-no'>` + getTranslation("tbl_c3_no", lang) + `</span><br><small style='color:#64748b;'>` + getTranslation("tbl_c3_no_sub", lang) + `</small></td>
				<td><b><span class='badge-yes'>` + getTranslation("tbl_c3_yes", lang) + `</span></b><br><small style='color:#64748b;'>` + getTranslation("tbl_c3_yes_sub", lang) + `</small></td>
			</tr>
			<tr class='comp-row'>
				<td><b>` + getTranslation("tbl_c4_t", lang) + `</b></td>
				<td><span class='badge-no'>` + getTranslation("tbl_c4_no", lang) + `</span><br><small style='color:#64748b;'>` + getTranslation("tbl_c4_no_sub", lang) + `</small></td>
				<td><b><span class='badge-yes'>` + getTranslation("tbl_c4_yes", lang) + `</span></b><br><small style='color:#64748b;'>` + getTranslation("tbl_c4_yes_sub", lang) + `</small></td>
			</tr>
			<tr class='comp-row'>
				<td><b>` + getTranslation("tbl_c5_t", lang) + `</b></td>
				<td><span class='badge-no'>` + getTranslation("tbl_c5_no", lang) + `</span><br><small style='color:#64748b;'>` + getTranslation("tbl_c5_no_sub", lang) + `</small></td>
				<td><b><span class='badge-yes'>` + getTranslation("tbl_c5_yes", lang) + `</span></b><br><small style='color:#64748b;'>` + getTranslation("tbl_c5_yes_sub", lang) + `</small></td>
			</tr>
		</table>`
}

func GetFullDownloadOffer(lang string) string {
	return `ОФИЦИАЛЬНАЯ ТЕХНИЧЕСКАЯ СПЕЦИФИКАЦИЯ И ТАРИФНАЯ СЕТКА КОРПОРАТИВНОГО B2B-ОФФЕРА
Правообладатель: Общество с ограниченной ответственностью «Квантум Сейф» (ОсОО «Квантум Сейф», г. Ош, КР)

ДОСТУПНЫЕ КОРПОРАТИВНЫЕ ЛИЦЕНЗИИ И СТОИМОСТЬ ПОДПИСКИ:

1. ТАРИФ «GLOBAL SCANNER» — БЕСПЛАТНО / USD 0
Базовый инструмент для оперативного экспресс-анализа внешних сетевых шлюзов ИТ-инфраструктуры организации.

2. ТАРИФ «ПО АНТИХАКЕР AI» — 19 000 долларов США / год
Автономный оборонный программный комплекс СЕЙЧАС. Предназначен для развертывания внутри закрытого серверного контура финансовой организации. Включает в себя:
 - Модуль Hardware Device Fingerprinting (Стек Rust): опрашивает регистры процессоров, считывает неизменяемый аппаратный ID и предотвращает атаки типа 'Человек по средине' (MitM).
 - Модуль TinyML Аномалий (Стек Go): на лету анализирует тайминги транзакционных пакетов XML/ISO-20022 и блокирует логические атаки классов Reentrancy и Flash-Loan фрода за 0.002 секунды до списания средств.
 - Модуль Изоляции Памяти: осуществляет непрерывный фоновый контроль оперативной памяти (RAM) серверов банка и блокирует хакерские инъекции вредоносного кода.

3. ТАРИФ «QUANTUM WEB3» (МАКСИМАЛЬНАЯ ЗАЩИТА / КВАНТОВОЕ БЕССМЕРТИЕ) — 35 000 долларов США / год
Максимальный оборонный комплекс на базе ядра KvantumSafe Pro Framework SDK (версия 2.6). Включает в себя полный пакет защитных модулей:
 - Все модули тарифа 'ПО АнтиХакер AI' для принудительного отражения текущих хакерских взломов и фрода в оперативной памяти СЕЙЧАС.
 - Модуль внутреннего сканера комплаенс-контроля (Stealth-сейф) для перманентного маскирования и принудительной изоляции резервных копий баз данных SQL под строгие права доступа стандарта POSIX 0600.
 - Интеллектуальный транзитный диспетчер и гибридный постквантовый оркестратор трансграничных платежей. Наше программное ядро осуществляет сквозную контейнеризацию и безопасную маршрутизацию потоков данных (SWIFT, XML) как в международных криптоконтейнерах на базе решеток нового поколения по стандартам NIST (ML-KEM), так и в суверенных азиатских шлюзах по государственным стандартам Китая GmSSL (алгоритмы семейства SM4-GCM). Гарантирует абсолютную защиту от скачивания данных сейчас и их последующего взлома сторонними вычислительными системами ПОТОМ.

Поставляется на защищенном физическом Flash-носителе с жесткой привязкой к Hardware ID главного сервера Вашей организации.`
}
