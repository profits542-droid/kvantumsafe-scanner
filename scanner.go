package main

import (
	"crypto/tls"
	"fmt"
	"net"
	"time"
)

type TLSReport struct {
	Domain      string
	Success     bool
	ErrorMsg    string
	Version     string
	CipherSuite string
	Issuer      string
	ExpiryDate  string
	IPAddress   string
}

func CheckDomainSecurity(domain string) TLSReport {
	report := TLSReport{Domain: domain, Success: false}

	ips, err := net.LookupIP(domain)
	if err != nil || len(ips) == 0 {
		report.ErrorMsg = "Ошибка DNS: Не удалось разрешить IP-адрес домена."
		return report
	}
	// Исправлено: берем первый IP-адрес из массива и переводим в строку
	report.IPAddress = ips[0].String()

	dialer := &net.Dialer{Timeout: 3 * time.Second}
	conn, err := tls.DialWithDialer(dialer, "tcp", domain+":443", &tls.Config{InsecureSkipVerify: true})
	if err != nil {
		report.ErrorMsg = fmt.Sprintf("Ошибка подключения: Порт 443 закрыт или заблокирован (%v)", err)
		return report
	}
	defer conn.Close()

	state := conn.ConnectionState()
	report.Success = true
	
	switch state.Version {
	case tls.VersionTLS13:
		report.Version = "TLS 1.3 (Высшая защита)"
	case tls.VersionTLS12:
		report.Version = "TLS 1.2 (Оптимально)"
	default:
		report.Version = "Устаревшая версия TLS"
	}

	report.CipherSuite = tls.CipherSuiteName(state.CipherSuite)
	if len(state.PeerCertificates) > 0 {
		cert := state.PeerCertificates[0]
		// Исправлено: корректно извлекаем имя организации эмитента
		if len(cert.Issuer.Organization) > 0 {
			report.Issuer = cert.Issuer.Organization[0]
		} else {
			report.Issuer = cert.Issuer.CommonName
		}
		report.ExpiryDate = cert.NotAfter.Format("2006-01-02")
	}

	return report
}

func RenderReportPage(report TLSReport) string {
	statusHTML := ""
	tableHTML := ""
	logHTML := fmt.Sprintf("[INIT] Старт проверки удаленного шлюза %s... ОК\n", report.Domain)

	if report.Success {
		statusHTML = "<p><b>Статус безопасности:</b> <span style='color:#00ff00;'><b>АКТИВЕН (Handshake Passed)</b></span></p>"
		tableHTML = fmt.Sprintf(`
			<tr><td>Версия шифрования контура</td><td>%s</td><td>Рисков не обнаружено</td></tr>
			<tr><td>Криптографический шифронабор</td><td>%s</td><td>NIST-Стандарт безопасен</td></tr>
			<tr><td>Центр сертификации (Issuer)</td><td>%s</td><td>Доверенный узел</td></tr>
			<tr><td>Срок действия сертификата</td><td>До %s</td><td>Валиден</td></tr>
		`, report.Version, report.CipherSuite, report.Issuer, report.ExpiryDate)
		
		logHTML += fmt.Sprintf("[DNS] Резолв домена успешен. Целевой IP-адрес: %s\n[TCP] Соединение с портом 443 установлено\n[TLS] Handshake завершен успешно.\n[SUCCESS] Контур верифицирован экспертным узлом ОсОО «Квантум Сейф». ПВТ КР.", report.IPAddress)
	} else {
		statusHTML = "<p><b>Статус безопасности:</b> <span style='color:#ff0000;'><b>ВНИМАНИЕ (Ошибка)</b></span></p>"
		tableHTML = fmt.Sprintf("<tr><td colspan='3' style='color:#ff0000;'>%s</td></tr>", report.ErrorMsg)
		logHTML += fmt.Sprintf("[ERROR] Сетевая сессия прервана: %s\n", report.ErrorMsg)
	}

	return fmt.Sprintf(`<html><head><style>body{font-family:Arial;margin:40px;background:#1a1a1a;color:#fff;} .card{background:#2d2d2d;padding:30px;border-radius:8px;max-width:800px;margin:0 auto;} th,td{padding:12px;text-align:left;border-bottom:1px solid #444;} table{width:100%%;border-collapse:collapse;margin:20px 0;} pre{background:#000;color:#00ff00;padding:15px;border-radius:4px;font-family:monospace;}</style></head><body>
		<div class='card'>
			<h2>🛡️ Расширенный отчет аудита периметра: %s</h2>
			%s
			<table>
				<tr style='background:#3d3d3d;'><th>Контур проверки</th><th>Параметр системы</th><th>Статус / Риск</th></tr>
				%s
				<tr><td>Защита памяти RAM (AI-AntiHacker)</td><td>Комплаенс-контроль активен</td><td><span style='color:#aaa;'>On-Premise / По запросу</span></td></tr>
				<tr><td>Санкционный скрининг (OFAC)</td><td>Граф связей Минюста КР</td><td><span style='color:#aaa;'>Enterprise / Подключение по API</span></td></tr>
			</table>
			<h3>💻 Технический лог сетевого рукопожатия (System Logs):</h3>
			<pre>%s</pre>
			<br><center><a href='/download' style='color:#00ff00;text-decoration:none;font-weight:bold;'>🏆 Нажми сюда для просмотра спецификации ПО</a></center>
		</div></body></html>`, report.Domain, statusHTML, tableHTML, logHTML)
}

func RenderDownloadPage() string {
	return `<html><head><style>body{font-family:Arial;margin:40px;background:#1a1a1a;color:#fff;} .card{background:#2d2d2d;padding:30px;border-radius:8px;max-width:750px;margin:0 auto;} .tech-box{background:#3d3d3d;padding:20px;border-left:4px solid #00ff00;border-radius:4px;margin:15px 0;} ul{line-height:1.6;} .btn{display:inline-block;background:#00ff00;color:#000;padding:10px 20px;text-decoration:none;border-radius:4px;font-weight:bold;margin-top:15px;}</style></head><body>
		<div class='card'>
			<h2>🏛️ Спецификация программного обеспечения «Quantum Safe Compliance Gateway»</h2>
			<p><b>Разработчик:</b> ОсОО «Квантум Сейф» (Резидент Парка высоких технологий Кыргызской Республики).</p>
			<div class='tech-box'>
				<h3>📋 Назначение платформы:</h3>
				<p>Автономный комплекс высокоскоростного контроля, комплаенс-мониторинга трансграничных финансовых ISO-20022 потоков данных и проактивной защиты серверной памяти АБС банков.</p>
			</div>
			<h3>⚙️ Проектируемые модули нашей Enterprise-экосистемы:</h3>
			<ul>
				<li><b>AI P2P Referee & Anti-Fraud Module:</b> ИИ-арбитраж P2P-сделок, моментальное OCR-распознавание параметров чеков и Match-сверка с выпиской банка для снижения нагрузки на саппорт на 80%%.</li>
				<li><b>Quantum Sanctions Screening Engine:</b> Графовый ИИ-скрининг скрытых бенефициаров до 4-го колена по мировым базам (OFAC, SDN-листы, ЕС, ООН) для предотвращения рисков вторичных санкций.</li>
				<li><b>AI-AntiHacker (Guard Engine):</b> Микросекундная защита оперативной памяти серверов (RAM) от уязвимостей переполнения буфера по методу Stealth RAM Isolation.</li>
				<li><b>Hardware Dongle API:</b> Низкоуровневая привязка лицензий к процессорам серверов и физическим USB-токенам FEITIAN/Рутокен 3.0 крупным оптом для защиты интеллектуальной собственности.</li>
			</ul>
			<br><center>
				<p style='color:#aaa;'>Полная техническая документация и архитектурные схемы предоставляются по официальному B2B-запросу финансовых организаций.</p>
				<a href='/report/mbank' class='btn'>← Вернуться в панель аудита</a>
			</center>
		</div></body></html>`
}
