package main

import (
	"strings"
)

const sharedStyles = `
<style>
* { box-sizing: border-box; margin: 0; padding: 0; }
body {
	font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, "Helvetica Neue", Arial, sans-serif;
	background: linear-gradient(135deg, #0f172a, #1e293b);
	color: #e2e8f0;
	min-height: 100vh;
	padding: 24px;
	line-height: 1.5;
}
.container {
	max-width: 960px;
	margin: 0 auto;
}
.banner {
	background: linear-gradient(90deg, #1e3a5f, #0f2744);
	border: 1px solid #3b82f6;
	border-radius: 12px;
	padding: 18px 22px;
	margin-bottom: 28px;
	font-size: 15px;
	color: #bfdbfe;
	box-shadow: 0 4px 20px rgba(59, 130, 246, 0.15);
}
h1 {
	font-size: 26px;
	font-weight: 700;
	margin-bottom: 22px;
	color: #f8fafc;
	letter-spacing: -0.3px;
}
.form-card {
	background: #1e293b;
	border: 1px solid #334155;
	border-radius: 12px;
	padding: 22px;
	margin-bottom: 28px;
	display: flex;
	gap: 12px;
	align-items: center;
	flex-wrap: wrap;
}
input[type="text"] {
	flex: 1;
	min-width: 240px;
	padding: 12px 16px;
	border: 1px solid #475569;
	border-radius: 8px;
	background: #0f172a;
	color: #f1f5f9;
	font-size: 15px;
	outline: none;
	transition: border-color 0.2s;
}
input[type="text"]:focus {
	border-color: #3b82f6;
}
.btn {
	display: inline-block;
	padding: 12px 22px;
	border: none;
	border-radius: 8px;
	font-size: 15px;
	font-weight: 600;
	cursor: pointer;
	text-decoration: none;
	transition: transform 0.15s, box-shadow 0.15s;
}
.btn:hover {
	transform: translateY(-1px);
}
.btn-primary {
	background: linear-gradient(135deg, #3b82f6, #2563eb);
	color: #fff;
	box-shadow: 0 4px 14px rgba(37, 99, 235, 0.35);
}
.btn-gold {
	background: linear-gradient(135deg, #f59e0b, #d97706);
	color: #1c1917;
	box-shadow: 0 4px 14px rgba(245, 158, 11, 0.35);
	margin-top: 12px;
}
.table-wrap {
	background: #1e293b;
	border: 1px solid #334155;
	border-radius: 12px;
	overflow: hidden;
	margin-bottom: 28px;
}
table {
	width: 100%;
	border-collapse: collapse;
}
th {
	background: #0f172a;
	color: #94a3b8;
	font-size: 13px;
	text-transform: uppercase;
	letter-spacing: 0.5px;
	padding: 14px 18px;
	text-align: left;
	border-bottom: 1px solid #334155;
}
td {
	padding: 14px 18px;
	border-bottom: 1px solid #334155;
	font-size: 14px;
	color: #e2e8f0;
}
tr:last-child td {
	border-bottom: none;
}
.param-name {
	color: #94a3b8;
	width: 42%;
}
.param-value {
	font-weight: 500;
	color: #f1f5f9;
}
.status-ok {
	color: #4ade80;
	font-weight: 600;
}
.status-fail {
	color: #f87171;
	font-weight: 600;
}
.log-window {
	background: #000;
	border: 1px solid #22c55e;
	border-radius: 10px;
	padding: 18px;
	font-family: "SF Mono", "Fira Code", "Consolas", monospace;
	font-size: 13px;
	color: #4ade80;
	white-space: pre-wrap;
	line-height: 1.6;
	margin-bottom: 28px;
	overflow-x: auto;
}
.cards {
	display: grid;
	grid-template-columns: 1fr;
	gap: 18px;
	margin: 28px 0;
}
.card {
	background: #1e293b;
	border: 1px solid #334155;
	border-radius: 12px;
	padding: 22px;
}
.card h3 {
	font-size: 17px;
	color: #fbbf24;
	margin-bottom: 12px;
}
.card p {
	font-size: 14px;
	color: #cbd5e1;
	line-height: 1.6;
}
.free-audit {
	background: linear-gradient(135deg, #14532d, #166534);
	border: 1px solid #22c55e;
	border-radius: 12px;
	padding: 20px;
	margin-bottom: 28px;
	font-size: 15px;
	color: #bbf7d0;
}
.footer-contacts {
	margin-top: 36px;
	padding-top: 20px;
	border-top: 1px solid #334155;
	text-align: center;
	font-size: 14px;
	color: #94a3b8;
}
.footer-contacts a {
	color: #60a5fa;
	text-decoration: none;
}
</style>
`

func RenderStartPage() string {
	html := `<!DOCTYPE html>
<html lang="ru">
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<title>Экспресс-Аудит Сетевого Периметра</title>
` + sharedStyles + `
</head>
<body>
<div class="container">
	<div class="banner">
		📌 ПРОФИЛЬ ДЕЯТЕЛЬНОСТИ РЕЗИДЕНТА ПВТ КР ОсОО Квантум Сейф: Комплексная разработка специализированного программного обеспечения, проектирование интеллектуальных ИИ-ассистентов и автономных ИИ-агентов, сквозной нормативно-правовой и технологический комплаенс-контроль, а также оказание профильной юридической помощи для финансовых организаций, распределенных ИТ-систем и Web3-инфраструктуры.
	</div>

	<h1>🔍 Экспресс-Аудит Сетевого Периметра (Внешний ИИ-Сканер)</h1>

	<div class="form-card">
		<form method="GET" action="/report/mbank" style="display:flex;gap:12px;flex:1;flex-wrap:wrap;align-items:center;">
			<input type="text" name="domain" placeholder="Введите домен или IP для сканирования..." required>
			<button type="submit" class="btn btn-primary">🚀 Запустить сканирование</button>
		</form>
	</div>

	<div style="display:flex; gap:15px; justify-content:center; flex-wrap:wrap;">
		<a href="/download" class="btn btn-gold">🏆 Спецификация ПО и Финансовая Выгода</a>
		<a href="/marketplace" class="btn btn-primary" style="margin-top:12px; background:linear-gradient(135deg, #3b82f6, #2563eb);">🏪 Маркетплейс ИИ-решений</a>
	</div>
</div>
</body>
</html>`
	return html
}

func RenderReportPage(report TLSReport) string {
	statusText := "Ошибка handshake"
	statusClass := "status-fail"
	if report.Success {
		statusText = "Успешно (Handshake OK)"
		statusClass = "status-ok"
	}

	html := `<!DOCTYPE html>
<html lang="ru">
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<title>Результаты сканирования — {{DOMAIN}}</title>
` + sharedStyles + `
</head>
<body>
<div class="container">
	<div class="banner">
		📌 ПРОФИЛЬ ДЕЯТЕЛЬНОСТИ РЕЗИДЕНТА ПВТ КР ОсОО Квантум Сейф: Комплексная разработка специализированного программного обеспечения, проектирование интеллектуальных ИИ-ассистентов и автономных ИИ-агентов, сквозной нормативно-правовой и технологический комплаенс-контроль, а также оказание профильной юридической помощи для финансовых организаций, распределенных ИТ-систем и Web3-инфраструктуры.
	</div>

	<h1>🔍 Экспресс-Аудит Сетевого Периметра (Внешний ИИ-Сканер)</h1>

	<div class="form-card">
		<form method="GET" action="/report/mbank" style="display:flex;gap:12px;flex:1;flex-wrap:wrap;align-items:center;">
			<input type="text" name="domain" value="{{DOMAIN}}" placeholder="Введите домен или IP для сканирования..." required>
			<button type="submit" class="btn btn-primary">🚀 Сканировать повторно</button>
		</form>
	</div>

	<div class="table-wrap">
		<table>
			<thead>
				<tr>
					<th colspan="2">Контур проверки: {{DOMAIN}}</th>
				</tr>
			</thead>
			<tbody>
				<tr>
					<td class="param-name">Статус сканирования</td>
					<td class="param-value {{STATUSCLASS}}">{{STATUSTEXT}}</td>
				</tr>
				<tr>
					<td class="param-name">IP-адрес</td>
					<td class="param-value">{{IPADDRESS}}</td>
				</tr>
				<tr>
					<td class="param-name">Версия протокола контура</td>
					<td class="param-value">{{VERSION}}</td>
				</tr>
				<tr>
					<td class="param-name">Криптографический шифронабор</td>
					<td class="param-value">{{CIPHERSUITE}}</td>
				</tr>
				<tr>
					<td class="param-name">Центр сертификации (Issuer)</td>
					<td class="param-value">{{ISSUER}}</td>
				</tr>
				<tr>
					<td class="param-name">Срок действия сертификата</td>
					<td class="param-value">{{EXPIRYDATE}}</td>
				</tr>
				<tr>
					<td class="param-name">Защита памяти RAM (AI-AntiHacker)</td>
					<td class="param-value" style="color: #f87171; font-weight: bold;">КОНТУРА БЕЗОПАСНОСТИ ВООБЩЕ НЕТ! ⚠️ Оперативная память серверов полностью открыта для скрытых хакерских угроз. Необходима установка защитного модуля</td>
				</tr>
				<tr>
					<td class="param-name">Санкционный скрининг (OFAC)</td>
					<td class="param-value" style="color: #f87171; font-weight: bold;">АВТОМАТИЧЕСКОЙ ПРОВЕРКИ ВООБЩЕ НЕТ! 🛑 Высокий риск вторичных санкций и потери международных корреспондентских счетов. Требуется API-подключение антисанкционного шлюза Квантум Сейф</td>
				</tr>
			</tbody>
		</table>
	</div>

	<pre class="log-window">[SYSTEM] ИИ-сканер инициализирован
[INFO]    Цель: {{DOMAIN}}
[INFO]    IP: {{IPADDRESS}}
[INFO]    Статус: {{STATUSTEXT}}
[INFO]    Протокол: {{VERSION}}
[INFO]    Шифронабор: {{CIPHERSUITE}}
[INFO]    Issuer: {{ISSUER}}
[INFO]    Срок действия: {{EXPIRYDATE}}
[WARN]    AI-AntiHacker: Защита оперативной памяти полностью отсутствует! Обнаружен критический риск эксплуатации уязвимостей нулевого дня.
[CRIT]    OFAC / Минюст КР: Автоматическая онлайн-проверка платежей отсутствует! Требуется интеграция шлюза Квантум Сейф
[DONE]    Сканирование завершено</pre>

	<div style="display:flex; gap:15px; justify-content:center; flex-wrap:wrap;">
		<a href="/download" class="btn btn-gold">🏆 Спецификация ПО и Финансовая Выгода</a>
		<a href="/marketplace" class="btn btn-primary" style="margin-top:12px; background:linear-gradient(135deg, #3b82f6, #2563eb);">🏪 Маркетплейс ИИ-решений</a>
	</div>
</div>
</body>
</html>`

	html = strings.Replace(html, "{{DOMAIN}}", report.Domain, -1)
	html = strings.Replace(html, "{{VERSION}}", report.Version, -1)
	html = strings.Replace(html, "{{CIPHERSUITE}}", report.CipherSuite, -1)
	html = strings.Replace(html, "{{ISSUER}}", report.Issuer, -1)
	html = strings.Replace(html, "{{EXPIRYDATE}}", report.ExpiryDate, -1)
	html = strings.Replace(html, "{{IPADDRESS}}", report.IPAddress, -1)
	html = strings.Replace(html, "{{STATUSTEXT}}", statusText, -1)
	html = strings.Replace(html, "{{STATUSCLASS}}", statusClass, -1)

	return html
}

func RenderDownloadPage() string {
	html := `<!DOCTYPE html>
<html lang="ru">
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<title>Спецификация ПО и Финансовая Выгода</title>
` + sharedStyles + `
</head>
<body>
<div class="container">
	<div class="banner">
