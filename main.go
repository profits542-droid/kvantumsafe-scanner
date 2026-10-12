package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

const defaultDomain = "bakai.kg"

func main() {
	fmt.Println("==================================================================")
	fmt.Println("KvantumSafe TLS Scanner — реальное подключение crypto/tls + net")
	fmt.Println("==================================================================")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", handleHome)
	http.HandleFunc("/report/mbank", handleReport)
	http.HandleFunc("/download", handleDownload)

	log.Printf("сервер слушает :%s  (/report/mbank?domain=%s , /download)", port, defaultDomain)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, pageShell(`
		<h1>TLS-сканер доменов</h1>
		<p class="muted">Честное TCP/TLS-подключение на порт 443 (таймаут 3 с). Без симуляции.</p>
		<form action="/report/mbank" method="GET" class="scan-form">
			<label for="domain">Домен</label>
			<input id="domain" name="domain" type="text" value="`+html.EscapeString(defaultDomain)+`" placeholder="example.com">
			<button type="submit">Сканировать</button>
		</form>
		<p><a href="/download">Модули ОсОО «Квантум Сейф»</a></p>
	`))
}

func handleReport(w http.ResponseWriter, r *http.Request) {
	domain := normalizeDomain(r.URL.Query().Get("domain"))
	if domain == "" {
		domain = defaultDomain
	}

	result := ScanDomainTLS(domain)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, renderReport(result))
}

func handleDownload(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, pageShell(`
		<h1>Модули ОсОО «Квантум Сейф»</h1>
		<p class="muted">Текстовое описание продуктовой линейки в разработке. PDF не формируется и никуда не отправляется.</p>
		<table>
			<thead>
				<tr><th>Модуль</th><th>Назначение</th><th>Статус</th></tr>
			</thead>
			<tbody>
				<tr>
					<td>AI P2P Referee</td>
					<td>Сервис-арбитр для разбора спорных P2P-операций: сверка сторон, журнал событий, выдача заключения оператору.</td>
					<td>В разработке</td>
				</tr>
				<tr>
					<td>AI-AntiHacker</td>
					<td>Контур мониторинга аномалий на периметре заказчика: сбор сигналов, правила реагирования, отчёты для ИБ.</td>
					<td>В разработке</td>
				</tr>
				<tr>
					<td>Hardware Dongle API</td>
					<td>Программный интерфейс к аппаратному ключу: проверка присутствия токена и вызов криптоопераций на устройстве.</td>
					<td>В разработке</td>
				</tr>
			</tbody>
		</table>
		<p>Это решения компании, которые проектируются и пишутся как отдельные модули платформы. Страница не является актом поставки, лицензией или подтверждением внедрения у клиента.</p>
		<p><a class="back" href="/">← Назад</a></p>
	`))
}

func renderReport(result ScanResult) string {
	esc := html.EscapeString
	status := "ОШИБКА"
	statusClass := "bad"
	if result.OK() {
		status = "ПОДКЛЮЧЕНО"
		statusClass = "ok"
	}

	notAfter := "—"
	validity := "—"
	if !result.NotAfter.IsZero() {
		notAfter = result.NotAfter.UTC().Format("2006-01-02 15:04:05 MST")
		left := time.Until(result.NotAfter)
		if left > 0 {
			validity = fmt.Sprintf("действителен ещё ~%d дн.", int(left.Hours()/24))
		} else {
			validity = "срок истёк"
		}
	}

	tlsVer := dash(result.TLSVersion)
	cipher := dash(result.CipherSuite)
	ip := dash(result.IP)
	subject := dash(result.Subject)
	issuer := dash(result.Issuer)
	errCell := "нет"
	if result.Error != "" {
		errCell = result.Error
	}

	var logBuf strings.Builder
	for _, line := range result.Logs {
		logBuf.WriteString(esc(line))
		logBuf.WriteByte('\n')
	}

	body := fmt.Sprintf(`
		<h1>TLS-отчёт: %s</h1>
		<p class="muted">Реальное сетевое подключение к порту 443. Таймаут 3 секунды. Пакеты <code>net</code> и <code>crypto/tls</code>.</p>
		<p>Статус: <span class="badge %s">%s</span></p>
		<table>
			<thead>
				<tr><th>Параметр</th><th>Значение</th></tr>
			</thead>
			<tbody>
				<tr><td>Домен</td><td>%s</td></tr>
				<tr><td>IP (DNS)</td><td>%s</td></tr>
				<tr><td>Порт</td><td>443</td></tr>
				<tr><td>Версия TLS</td><td>%s</td></tr>
				<tr><td>Cipher Suite</td><td>%s</td></tr>
				<tr><td>Сертификат (CN)</td><td>%s</td></tr>
				<tr><td>Издатель</td><td>%s</td></tr>
				<tr><td>Действителен до</td><td>%s</td></tr>
				<tr><td>Срок</td><td>%s</td></tr>
				<tr><td>Ошибка</td><td>%s</td></tr>
			</tbody>
		</table>
		<h2>Лог шагов</h2>
		<pre class="logs">%s</pre>
		<form action="/report/mbank" method="GET" class="scan-form">
			<label for="domain">Другой домен</label>
			<input id="domain" name="domain" type="text" value="%s">
			<button type="submit">Сканировать</button>
		</form>
		<p><a class="back" href="/">← Назад</a> · <a href="/download">Модули компании</a></p>
	`,
		esc(result.Domain),
		statusClass, esc(status),
		esc(result.Domain),
		esc(ip),
		esc(tlsVer),
		esc(cipher),
		esc(subject),
		esc(issuer),
		esc(notAfter),
		esc(validity),
		esc(errCell),
		logBuf.String(),
		esc(result.Domain),
	)
	return pageShell(body)
}

func dash(s string) string {
	if s == "" {
		return "—"
	}
	return s
}

func pageShell(body string) string {
	return `<!DOCTYPE html>
<html lang="ru">
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<title>KvantumSafe TLS Scanner</title>
<style>
  :root { --bg:#0b1220; --panel:#111827; --line:#1f2937; --text:#e5e7eb; --muted:#94a3b8; --gold:#d4af37; --ok:#22c55e; --bad:#f87171; }
  * { box-sizing: border-box; }
  body { margin:0; font-family:"Segoe UI", Arial, sans-serif; background:var(--bg); color:var(--text); }
  .wrap { max-width:920px; margin:40px auto; padding:32px; background:var(--panel); border:1px solid var(--line); border-radius:10px; box-shadow:0 12px 40px rgba(0,0,0,.35); }
  h1,h2 { margin:0 0 12px; font-weight:600; letter-spacing:.02em; }
  h1 { font-size:22px; color:#f8fafc; }
  h2 { font-size:16px; color:var(--gold); margin-top:28px; }
  p { line-height:1.55; color:var(--text); }
  .muted { color:var(--muted); font-size:14px; }
  table { width:100%; border-collapse:collapse; margin:18px 0 8px; font-size:14px; }
  th, td { padding:12px 14px; text-align:left; border-bottom:1px solid var(--line); vertical-align:top; }
  thead th { background:#0a1628; color:var(--gold); font-size:12px; text-transform:uppercase; letter-spacing:.08em; border-bottom:1px solid #334155; }
  tbody tr:nth-child(even) { background:#0d1726; }
  tbody td:first-child { color:var(--muted); width:28%; }
  .badge { display:inline-block; padding:4px 10px; border-radius:4px; font-weight:700; font-size:12px; }
  .ok { background:#052e16; color:var(--ok); }
  .bad { background:#3f1212; color:var(--bad); }
  pre.logs { background:#05080f; color:#86efac; padding:16px; border-radius:6px; overflow:auto; font:13px/1.5 Consolas, "Courier New", monospace; border:1px solid var(--line); min-height:120px; white-space:pre-wrap; }
  .scan-form { display:flex; flex-wrap:wrap; gap:10px; align-items:flex-end; margin:22px 0; }
  label { display:block; font-size:12px; color:var(--muted); margin-bottom:6px; width:100%; }
  input[type=text] { flex:1; min-width:220px; padding:10px 12px; border-radius:6px; border:1px solid #334155; background:#0b1220; color:var(--text); }
  button { background:var(--gold); color:#0b1220; border:0; padding:11px 18px; border-radius:6px; font-weight:700; cursor:pointer; }
  button:hover { filter:brightness(1.08); }
  a { color:#93c5fd; }
  a.back { color:var(--gold); font-weight:600; text-decoration:none; }
  code { color:#fde68a; }
</style>
</head>
<body><div class="wrap">` + body + `</div></body>
</html>`
}
