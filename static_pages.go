```go
package main

func GetSpecificationPageHTML(lang string) string {
	title := "Технический регламент и функционал Золотой Кнопки"
	desc := "Настоящий веб-интерфейс реализует функцию динамической генерации и безопасного депонирования технической документации. При активации триггера (нажатии на кнопку) ядро сервера на Go на лету определяет активную языковую локализацию клиентской сессии и формирует официальный защищенный файл спецификации (TXT) для предоставления ИТ-департаментам и комплаенс-контролю финансовых организаций.<br><br>Сгенерированный документ содержит полные архитектурные параметры платформы KvantumSafe Pro Framework, подробную разбивку коммерческой тарифной сетки (Global Scanner, ПО АнтиХакер AI, Quantum Web3) и легальное обоснование юридической чистоты софта перед государственными регуляторами. Файл принудительно сохраняется в локальное хранилище (папку Загрузки) ноутбука или персонального компьютера пользователя по протоколу контентной диспетчеризации (Content-Disposition)."
	back := "← Назад / Back"

	if lang == "kg" {
		title = "Алтын Түйнөктүн Техникалык Регламенти"
		desc = "Бул веб-интерфейс программалык камсыздоонун техникалык документтерин динамикалык түрдө генерациялоо жана криптографиялык коопсуз депонирлөө функциясын аткарат. Түймени басканда, Go тилиндеги serverдин ядросу кардардын сессиясынын тилин аныктайт жана өзгөчөлөнгөн тексттик документти (TXT) түзөт."
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

// ВЫСОКОИНТЕЛЛЕКТУАЛЬНАЯ ДИНАМИЧЕСКАЯ ТАБЛИЦА СРАВНЕНИЯ НА 5 ЯЗЫКАХ
func GetComparisonTableHTML(lang string) string {
	// Безопасный фолбэк для арабского и казахского на английский язык внутри технической сетки таблицы
	cleanLang := lang
	if cleanLang != "ru" && cleanLang != "kg" && cleanLang != "en" && cleanLang != "kz" && cleanLang != "ar" {
		cleanLang = "en"
	}

	return `<h3 class='comp-heading'>` + getTranslation("tbl_main_title", cleanLang) + `</h3>
		<table class='comp-table'>
			<tr>
				<th style='width:40%;'>` + getTranslation("tbl_h1", cleanLang) + `</th>
				<th style='width:30%; background-color:#475569;'>` + getTranslation("tbl_h2", cleanLang) + `</th>
				<th style='width:30%; background-color:#1e3a8a;'>` + getTranslation("tbl_h3", cleanLang) + `</th>
			</tr>
			<tr class='comp-row'>
				<td><b>` + getTranslation("tbl_r1_c1", cleanLang) + `</b></td>
				<td><span class='badge-no'>` + getTranslation("tbl_r1_b_no", cleanLang) + `</span><br><small style='color:#64748b;'>` + getTranslation("tbl_r1_d_no", cleanLang) + `</small></td>
				<td><b><span class='badge-yes'>` + getTranslation("tbl_r1_b_yes", cleanLang) + `</span></b><br><small style='color:#64748b;'>` + getTranslation("tbl_r1_d_yes", cleanLang) + `</small></td>
			</tr>
			<tr class='comp-row'>
				<td><b>` + getTranslation("tbl_r2_c1", cleanLang) + `</b></td>
				<td><span class='badge-no'>` + getTranslation("tbl_r2_b_no", cleanLang) + `</span><br><small style='color:#64748b;'>` + getTranslation("tbl_r2_d_no", cleanLang) + `</small></td>
				<td><b><span class='badge-yes'>` + getTranslation("tbl_r2_b_yes", cleanLang) + `</span></b><br><small style='color:#64748b;'>` + getTranslation("tbl_r2_d_yes", cleanLang) + `</small></td>
			</tr>
			<tr class='comp-row'>
				<td><b>` + getTranslation("tbl_r3_c1", cleanLang) + `</b></td>
				<td><span class='badge-no'>` + getTranslation("tbl_r2_b_no", cleanLang) + `</span><br><small style='color:#64748b;'>` + getTranslation("tbl_r3_d_no", cleanLang) + `</small></td>
				<td><b><span class='badge-yes'>` + getTranslation("tbl_r3_b_yes", cleanLang) + `</span></b><br><small style='color:#64748b;'>` + getTranslation("tbl_r3_d_yes", cleanLang) + `</small></td>
			</tr>
			<tr class='comp-row'>
				<td><b>` + getTranslation("tbl_r4_c1", cleanLang) + `</b></td>
				<td><span class='badge-no'>` + getTranslation("tbl_r4_d_no", cleanLang) + `</span><br><small style='color:#64748b;'>` + getTranslation("tbl_r4_d_no", cleanLang) + `</small></td>
				<td><b><span class='badge-yes'>` + getTranslation("tbl_r4_b_yes", cleanLang) + `</span></b><br><small style='color:#64748b;'>` + getTranslation("tbl_r4_d_yes", cleanLang) + `</small></td>
			</tr>
			<tr class='comp-row'>
				<td><b>` + getTranslation("tbl_r5_c1", cleanLang) + `</b></td>
				<td><span class='badge-no'>` + getTranslation("tbl_r5_d_no", cleanLang) + `</span><br><small style='color:#64748b;'>` + getTranslation("tbl_r5_d_no", cleanLang) + `</small></td>
				<td><b><span class='badge-yes'>` + getTranslation("tbl_r5_b_yes", cleanLang) + `</span></b><br><small style='color:#64748b;'>` + getTranslation("tbl_r5_d_yes", cleanLang) + `</small></td>
			</tr>
		</table>`
}

func GetFullDownloadOffer(lang string) string {
	return `ОФИЦИАЛЬНОЕ КОММЕРЧЕСКОЕ ПРЕДЛОЖЕНИЕ И ТАРИФНАЯ СЕТКА ДЛЯ РУКОВОДСТВА БАНКА
Разработчик и правообладатель: ОсОО «Квантум Сейф» (г. Ош, Кыргызская Республика)

Уважаемый Председатель Правления / Директор по безопасности! 
Ниже представлена простая и понятная расшифровка защитных программных пакетов, которые оградят ваш банк от финансовых потерь, кражи денег со счетов и репутационных рисков.

ДОСТУПНЫЕ ВАРИАНТЫ КОРПОРАТИВНЫХ ЛИЦЕНЗИИ И СТОИМОСТЬ:

1. ТАРИФ «GLOBAL SCANNER» — БЕСПЛАТНО / USD 0
- ЧТО ЭТО ТАКОЕ: Базовый поверхностный инструмент («экспресс-термометр»).
- КАК ПОМОГАЕТ БАНКУ: Проверяет внешние двери вашего сайта в реальном времени. Показывает, какие шлюзы забыли закрыть ваши ИТ-специалисты и где хакеры могут начать прощупывать систему.

2. ТАРИФ «ПО АНТИХАКЕР AI» — 19 000 долларов США / год
- ЧТО ЭТО ТАКОЕ: Мощный невидимый щит, который ставится внутрь серверов банка для защиты от взломов и воровства денег ПРЯМО СЕЙЧАС.
- КАК ПОМОГАЕТ БАНКУ:
  - Цифровой паспорт компьютерного железа: Программа намертво привязывается к заводскому номеру процессора вашего банка. Хакер не сможет подменить сервер или перехватить управление транзакциями снаружи.
  - Скоростная блокировка умного мошенничества: Искусственный интеллект проверяет платежи и за 0.002 секунды автоматически останавливает хитрые хакерские схемы по выводу балансов со счетов до того, как деньги уйдут.
  - Невидимый охранник оперативной памяти: Круглосуточно следит за памятью серверов и блокирует любые попытки злоумышленников внедрить вирусный код прямо в работающие процессы банка.

3. ТАРИФ «QUANTUM WEB3» (МАКСИМАЛЬНАЯ ЗАЩИТА И БЕЗОПАСНОСТЬ) — 35 000 долларов США / год
- ЧТО ЭТО ТАКОЕ: Абсолютная защита высшего уровня на базе системы KvantumSafe Pro Framework SDK (версия 2.6). Включает в себя вообще все наши технологии в один клик.
- ЧТО ВХОДИТ В ПАКЕТ:
  - Все защитные модули от взломов и кражи денег из тарифа «ПО АнтиХакер AI» (Защита ПРЯМО СЕЙЧАС).
  - Умный комплаенс-сейф (Внутренний сканер контроля баз данных): Автоматический фоновый ревизор. Он на лету находит важные резервные копии баз данных SQL, мгновенно стирает их из общей видимости и прячет под строжайшие права доступа стандарта POSIX 0600. Если хакер проникнет внутрь сети, он физически не сможет скопировать или стереть ваши бэкапы клиентов.
  - Защитный бронеконтейнер против суперкомпьютеров будущего (Защита ПОТОМ): Упаковывает международные переводы (SWIFT, XML) в новые защитные криптоконтейнеры по строгим стандартам США (NIST ML-KEM) и суверенным стандартам Китая GmSSL (SM4-GCM). Это гарантирует, что хакеры не смогут украсть ваши данные сейчас, чтобы расшифровать их потом на квантовых компьютерах.

Программный комплекс поставляется на физической защищенной Flash-карте с жесткой привязкой к оборудованию Вашего банка.`
}
