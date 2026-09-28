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

func GetComparisonTableHTML(lang string) string {
	cleanLang := lang
	if cleanLang != "ru" && cleanLang != "kg" && cleanLang != "en" && cleanLang != "kz" && cleanLang != "ar" {
		cleanLang = "en"
	}

	return `<h3 class='comp-heading'>` + getTableTranslation("tbl_main_title", cleanLang) + `</h3>
		<table class='comp-table'>
			<tr>
				<th style='width:40%;'>` + getTableTranslation("tbl_h1", cleanLang) + `</th>
				<th style='width:30%; background-color:#475569;'>` + getTableTranslation("tbl_h2", cleanLang) + `</th>
				<th style='width:30%; background-color:#1e3a8a;'>` + getTableTranslation("tbl_h3", cleanLang) + `</th>
			</tr>
			<tr class='comp-row'>
				<td><b>` + getTableTranslation("tbl_r1_c1", cleanLang) + `</b></td>
				<td><span class='badge-no'>` + getTableTranslation("tbl_r1_b_no", cleanLang) + `</span><br><small style='color:#64748b;'>` + getTableTranslation("tbl_r1_d_no", cleanLang) + `</small></td>
				<td><b><span class='badge-yes'>` + getTableTranslation("tbl_r1_b_yes", cleanLang) + `</span></b><br><small style='color:#64748b;'>` + getTableTranslation("tbl_r1_d_yes", cleanLang) + `</small></td>
			</tr>
			<tr class='comp-row'>
				<td><b>` + getTableTranslation("tbl_r2_c1", cleanLang) + `</b></td>
				<td><span class='badge-no'>` + getTableTranslation("tbl_r2_b_no", cleanLang) + `</span><br><small style='color:#64748b;'>` + getTableTranslation("tbl_r2_d_no", cleanLang) + `</small></td>
				<td><b><span class='badge-yes'>` + getTableTranslation("tbl_r2_b_yes", cleanLang) + `</span></b><br><small style='color:#64748b;'>` + getTableTranslation("tbl_r2_d_yes", cleanLang) + `</small></td>
			</tr>
			<tr class='comp-row'>
				<td><b>` + getTableTranslation("tbl_r3_c1", cleanLang) + `</b></td>
				<td><span class='badge-no'>` + getTableTranslation("tbl_r2_b_no", cleanLang) + `</span><br><small style='color:#64748b;'>` + getTableTranslation("tbl_r3_d_no", cleanLang) + `</small></td>
				<td><b><span class='badge-yes'>` + getTableTranslation("tbl_r3_b_yes", cleanLang) + `</span></b><br><small style='color:#64748b;'>` + getTableTranslation("tbl_r3_d_yes", cleanLang) + `</small></td>
			</tr>
			<tr class='comp-row'>
				<td><b>` + getTableTranslation("tbl_r4_c1", cleanLang) + `</b></td>
				<td><span class='badge-no'>` + getTableTranslation("tbl_r4_d_no", cleanLang) + `</span><br><small style='color:#64748b;'>` + getTableTranslation("tbl_r4_d_no", cleanLang) + `</small></td>
				<td><b><span class='badge-yes'>` + getTableTranslation("tbl_r4_b_yes", cleanLang) + `</span></b><br><small style='color:#64748b;'>` + getTableTranslation("tbl_r4_d_yes", cleanLang) + `</small></td>
			</tr>
			<tr class='comp-row'>
				<td><b>` + getTableTranslation("tbl_r5_c1", cleanLang) + `</b></td>
				<td><span class='badge-no'>` + getTableTranslation("tbl_r5_d_no", cleanLang) + `</span><br><small style='color:#64748b;'>` + getTableTranslation("tbl_r5_d_no", cleanLang) + `</small></td>
				<td><b><span class='badge-yes'>` + getTableTranslation("tbl_r5_b_yes", cleanLang) + `</span></b><br><small style='color:#64748b;'>` + getTableTranslation("tbl_r5_d_yes", cleanLang) + `</small></td>
			</tr>
		</table>`
}

func GetFullDownloadOffer(lang string) string {
	return getTranslation("download_b2b_text", lang)
}
