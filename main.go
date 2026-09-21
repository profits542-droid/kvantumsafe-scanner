package main

import (
	"encoding/base64"
	"fmt"
	"html"
	"net/http"
	"strings"
	"time"
"github.com/skip2/go-qrcode"
)

type LicenseInfo struct {
	ClientName string
	ExpireDate time.Time
	ServerID   string
	IsValid    bool
	DaysLeft   int
}

func VerifyLicenseKey(encodedKey string) LicenseInfo {
	info := LicenseInfo{IsValid: false, DaysLeft: 0}
	decodedBytes, err := base64.StdEncoding.DecodeString(encodedKey)
	if err != nil {
		return info
	}
	parts := strings.Split(string(decodedBytes), "|")
	if len(parts) != 3 {
		return info
	}
	info.ClientName = parts[0]
	info.ServerID = parts[2]
	expTime, err := time.Parse("2006-01-02", parts[1])
	if err != nil {
		return info
	}
	info.ExpireDate = expTime
	currentTime := time.Now()
	if currentTime.Before(info.ExpireDate) {
		info.IsValid = true
		duration := info.ExpireDate.Sub(currentTime)
		info.DaysLeft = int(duration.Hours() / 24)
	}
	return info
}

var cssStyles = "<style>body{font-family:'Segoe UI',sans-serif;background-color:#f4f7f6;color:#333;margin:0;padding:0}.navbar{background-color:#0a2540;color:#fff;padding:20px 40px;display:flex;align-items:center;gap:20px;box-shadow:0 4px 10px rgba(0,0,0,.1)}.logo-img{width:70px;height:70px;border:2px solid #d4af37;background:#fff;object-fit:cover}.navbar h1{margin:0;font-size:24px;letter-spacing:.5px}.content-area{max-width:900px;background:#fff;margin:40px auto;padding:40px;border-radius:8px;box-shadow:0 4px 15px rgba(0,0,0,.05)}h2{color:#0a2540;margin-top:0;font-size:22px;border-bottom:2px solid #e2e8f0;padding-bottom:10px}.btn-action{display:inline-block;background-color:#004d40;color:#fff;padding:14px 28px;border-radius:6px;text-decoration:none;font-weight:bold;margin-top:20px;box-shadow:0 2px 5px rgba(0,0,0,.15);transition:.2s}.btn-action:hover{background-color:#00796b}.btn-bridge{display:inline-block;background-color:#2563eb;color:#fff;padding:14px 28px;border-radius:6px;text-decoration:none;font-weight:bold;margin-top:20px;box-shadow:0 2px 5px rgba(0,0,0,.15);transition:.2s}.btn-bridge:hover{background-color:#1d4ed8}.btn-billing{display:inline-block;background-color:#7c3aed;color:#fff;padding:14px 28px;border-radius:6px;text-decoration:none;font-weight:bold;margin-top:20px;box-shadow:0 2px 5px rgba(0,0,0,.15);transition:.2s}.btn-billing:hover{background-color:#6d28d9}.meta-grid{display:grid;grid-template-columns:1fr 1fr;gap:15px;background:#f8fafc;padding:20px;border-radius:6px;margin-bottom:30px;border-left:4px solid #d4af37;font-size:14px}.billing-grid{display:grid;grid-template-columns:1fr 1fr 1fr;gap:20px;margin-top:20px}.tier-card{border:1px solid #e2e8f0;padding:20px;border-radius:8px;text-align:center;background:#fff;box-shadow:0 2px 5px rgba(0,0,0,.02)}.tier-card.free-tier{border:2px solid #16a34a;background:#f0fdf4}.tier-card.premium{border:2px solid #7c3aed;background:#faf5ff}.tier-price{font-size:24px;font-weight:bold;color:#0a2540;margin:15px 0}.input-field{width:100%;padding:10px;margin:10px 0;border:1px solid #cbd5e1;border-radius:4px;box-sizing:border-box}.table-vulnerabilities{width:100%;border-collapse:collapse;margin-top:20px}.table-vulnerabilities th{background-color:#f1f5f9;text-align:left;padding:12px;border:1px solid #cbd5e1;font-weight:bold}.table-vulnerabilities td{padding:12px;border:1px solid #cbd5e1;font-size:14px}.danger-row{background-color:#fff5f5}.warning-row{background-color:#fffbeb}.bridge-row{background-color:#f0fdf4}.status-badge{display:block;margin-top:5px;color:#b71c1c;font-weight:bold;font-size:12px}.conclusion-box{background-color:#e0f2f1;border-left:4px solid #004d40;padding:20px;border-radius:6px;margin-top:30px;font-size:14px;line-height:1.6}</style>"

func main() {
	premiumLicenseKey := "T0FPIEtvbW1lcmNoZXNreWkgYmFuayBLWVJHWVpTVEFOIChNQkFOSyl8MjAyNy0wOS0yMHxDb3JlLU5vZGUtMDE="
	license := VerifyLicenseKey(premiumLicenseKey)

	// 1. DASHBOARD
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		statusText := "<span style=\"color:#b71c1c; font-weight:bold;\">● ЛИЦЕНЗИЯ НЕ НАЙДЕНА</span>"
		if license.IsValid {
			statusText = "<span style=\"color:#16a34a; font-weight:bold;\">● КОРПОРАТИВНАЯ ЛИЦЕНЗИЯ АКТИВНА</span>"
		}
		htmlPage := "<!DOCTYPE html><html><head><meta charset=\"UTF-8\"><title>KvantumSafe Dashboard</title>" + cssStyles + "</head><body><div class=\"navbar\"><img class=\"logo-img\" src=\"/static/logo.jpg\"><h1>KvantumSafe Corporate Dashboard</h1></div><div class=\"content-area\"><h2>Мониторинг банковских приложений</h2><div class=\"meta-grid\"><div><strong>Организация:</strong> " + html.EscapeString(license.ClientName) + "</div><div><strong>Статус подписки:</strong> " + statusText + "</div><div><strong>Ядро платформы:</strong> Go Engine v2.6 SDK</div><div><strong>Привязка к серверу:</strong> " + html.EscapeString(license.ServerID) + "</div></div><div style=\"text-align: center; margin-top: 40px; display: flex; justify-content: center; gap: 20px;\"><a class=\"btn-action\" href=\"/report/mbank\">Запустить внешний аудит</a><a class=\"btn-bridge\" href=\"/report/quantum-bridge\">Запустить крипто-мост (NIST / GmSSL)</a><a class=\"btn-billing\" href=\"/billing\">Управление биллингом ($)</a></div></div></body></html>"
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		_, _ = w.Write([]byte(htmlPage))
	})

	// 2. ВНЕШНИЙ ОТЧЕТ
	http.HandleFunc("/report/mbank", func(w http.ResponseWriter, r *http.Request) {
		currentTime := time.Now().Format("2006-01-02 15:04")
		reportData := UniversalIdentifySoftware("95.46.154.19")
		var rowsHtml string
		for _, app := range reportData.DetectedApps {
			rowsHtml += "<tr class=\"danger-row\"><td><strong>• " + html.EscapeString(app.Name) + "</strong><span class=\"status-badge\">🛑 Несоответствие регламентам сетевой защиты</span></td><td><span style=\"background:#e2e8f0; padding:4px 8px; border-radius:4px; font-weight:bold;\">" + fmt.Sprintf("%d", app.Port) + "</span> (Версия: " + html.EscapeString(app.Version) + ")</td></tr>"
		}
		htmlPage := "<!DOCTYPE html><html><head><meta charset=\"UTF-8\"><title>KvantumSafe — Внешний Отчет</title>" + cssStyles + "</head><body><div class=\"navbar\"><img class=\"logo-img\" src=\"/static/logo.jpg\"><h1>KvantumSafe — Внешний Отчет</h1></div><div class=\"content-area\"><p><a href=\"/\" style=\"color: #004d40; text-decoration: none; font-weight: bold;\">← Назад к панели</a></p><h2>Результаты внешнего инфраструктурного сканирования</h2><div class=\"meta-grid\"><div><strong>Объект проверки:</strong> mbank.kg (95.46.154.19)</div><div><strong>Уровень угрозы:</strong> <span style=\"color:#b71c1c; font-weight:bold;\">КРИТИЧЕСКИЙ</span></div><div><strong>Дата и время:</strong> " + currentTime + "</div><div><strong>Лицензия сканера:</strong> <span style=\"color:#16a34a; font-weight:bold;\">● FREE (Ознакомительная)</span></div></div><table class=\"table-vulnerabilities\"><thead><tr><th>Приложение</th><th>Порт / Версия</th></tr></thead><tbody>" + rowsHtml + "</tbody></table><div class=\"conclusion-box\"><h3>Рекомендательное заключение:</h3><ol><li>Срочно перевести инфраструктуру в Stealth Mode.</li><li>Интегрировать стандарты NIST (ML-KEM) и суверенные протоколы оркестрации GmSSL (SM4-GCM / SM9).</li></ol></div><div style=\"text-align: center; margin-top: 30px;\"><a class=\"btn-internal\" href=\"/report/internal\">Запустить внутренний аудит безопасности серверов банка</a></div></div></body></html>"
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		_, _ = w.Write([]byte(htmlPage))
	})

	// 3. ВНУТРЕННИЙ АУДИТ
	http.HandleFunc("/report/internal", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		currentTime := time.Now().Format("2006-01-02 15:04")
		htmlPage := "<!DOCTYPE html><html><head><meta charset=\"UTF-8\"><title>KvantumSafe — Внутренний Аудит</title>" + cssStyles + "</head><body><div class=\"navbar\"><img class=\"logo-img\" src=\"/static/logo.jpg\"><h1>KvantumSafe — Закрытый Внутренний Аудит</h1></div><div class=\"content-area\"><p><a href=\"/report/mbank\" style=\"color: #004d40; text-decoration: none; font-weight: bold;\">← Вернуться к внешнему аудиту</a></p><h2>Отчет ревизии данных внутри сервера (Data at Rest)</h2><div class=\"meta-grid\" style=\"border-left-color: #d97706;\"><div><strong>Сервер проверки:</strong> " + html.EscapeString(license.ServerID) + "</div><div><strong>Статус комплаенса:</strong> <span style=\"color:#b45309; font-weight:bold;\">⚠️ ОБНАРУЖЕНЫ НАРУШЕНИЯ</span></div><div><strong>Время ревизии:</strong> " + currentTime + "</div><div><strong>Оставшийся срок Pro-подписки:</strong> <span style=\"color:#16a34a; font-weight:bold;\">" + fmt.Sprintf("%d дней", license.DaysLeft) + "</span></div></div><table class=\"table-vulnerabilities\"><thead><tr><th>Категория / Расположение</th><th>Описание угрозы</th><th>Статус защиты</th></tr></thead><tbody><tr class=\"warning-row\"><td><strong>⚠️ SQL Dump Backup</strong><br><span style=\"font-size:12px; color:#64748b;\">Путь: C:\\BankServer\\Secret\\backup.sql</span></td><td>Обнаружена открытая база данных без маскирования чувствительных полей</td><td style=\"color:#16a34a; font-weight:bold; text-align:center;\">🛡️ Изолирован (0600)</td></tr><tr class=\"warning-row\"><td><strong>⚠️ Hardcoded Credentials</strong><br><span style=\"font-size:12px; color:#64748b;\">Путь: C:\\BankServer\\Secret\\config.json</span></td><td>Обнаружен открытый текстовый пароль доступа к шлюзу SWIFT контура</td><td style=\"color:#16a34a; font-weight:bold; text-align:center;\">🛡️ Изолирован (0600)</td></tr></tbody></table></div></body></html>"
		_, _ = w.Write([]byte(htmlPage))
	})

	// 4. КВАНТОВЫЙ МОСТ
	http.HandleFunc("/report/quantum-bridge", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		currentTime := time.Now().Format("2006-01-02 15:04")
