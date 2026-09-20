package main

import (
	"encoding/base64"
	"fmt"
	"html"
	"net/http"
	"time"

	"github.com/skip2/go-qrcode"
)

func main() {
	isPremiumActive := true

	qrcodeBytes, err := qrcode.Encode("https://kvantumsafe.tech", qrcode.Medium, 256)
	if err != nil {
		fmt.Printf("[-] Ошибка генерации QR-кода: %v\n", err)
		return
	}
	qrcodeSrc := "data:image/png;base64," + base64.StdEncoding.EncodeToString(qrcodeBytes)

	_ = isPremiumActive
	_ = qrcodeSrc

	cssStyles := "<style>" +
		"body { font-family: 'Segoe UI', Arial, sans-serif; background-color: #f4f7f6; color: #333; margin: 0; padding: 0; }" +
		".navbar { background-color: #0a2540; color: white; padding: 20px 40px; display: flex; align-items: center; gap: 20px; box-shadow: 0 4px 10px rgba(0,0,0,0.1); }" +
		".logo-img { width: 70px; height: 70px; border: 2px solid #d4af37; background: white; object-fit: cover; }" +
		".navbar h1 { margin: 0; font-size: 24px; letter-spacing: 0.5px; }" +
		".content-area { max-width: 900px; background: white; margin: 40px auto; padding: 40px; border-radius: 8px; box-shadow: 0 4px 15px rgba(0,0,0,0.05); }" +
		"h2 { color: #0a2540; margin-top: 0; font-size: 22px; border-bottom: 2px solid #e2e8f0; padding-bottom: 10px; }" +
		".btn-action { display: inline-block; background-color: #004d40; color: white; padding: 14px 28px; border-radius: 6px; text-decoration: none; font-weight: bold; margin-top: 20px; box-shadow: 0 2px 5px rgba(0,0,0,0.15); transition: 0.2s; }" +
		".btn-action:hover { background-color: #00796b; }" +
		".btn-internal { display: inline-block; background-color: #d97706; color: white; padding: 14px 28px; border-radius: 6px; text-decoration: none; font-weight: bold; margin-top: 20px; box-shadow: 0 2px 5px rgba(0,0,0,0.15); transition: 0.2s; }" +
		".btn-internal:hover { background-color: #b45309; }" +
		".btn-bridge { display: inline-block; background-color: #2563eb; color: white; padding: 14px 28px; border-radius: 6px; text-decoration: none; font-weight: bold; margin-top: 20px; box-shadow: 0 2px 5px rgba(0,0,0,0.15); transition: 0.2s; }" +
		".btn-bridge:hover { background-color: #1d4ed8; }" +
		".meta-grid { display: grid; grid-template-columns: 1fr 1fr; gap: 15px; background: #f8fafc; padding: 20px; border-radius: 6px; margin-bottom: 30px; border-left: 4px solid #d4af37; font-size: 14px; }" +
		".table-vulnerabilities { width: 100%; border-collapse: collapse; margin-top: 20px; }" +
		".table-vulnerabilities th { background-color: #f1f5f9; text-align: left; padding: 12px; border: 1px solid #cbd5e1; font-weight: bold; }" +
		".table-vulnerabilities td { padding: 12px; border: 1px solid #cbd5e1; font-size: 14px; }" +
		".danger-row { background-color: #fff5f5; }" +
		".warning-row { background-color: #fffbeb; }" +
		".bridge-row { background-color: #f0fdf4; }" +
		".status-badge { display: block; margin-top: 5px; color: #b71c1c; font-weight: bold; font-size: 12px; }" +
		".conclusion-box { background-color: #e0f2f1; border-left: 4px solid #004d40; padding: 20px; border-radius: 6px; margin-top: 30px; font-size: 14px; line-height: 1.6; }" +
		"</style>"

	// 1. РОУТЕР ГЛАВНОЙ СТРАНИЦЫ (DASHBOARD)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		statusText := "<span style=\"color:#16a34a; font-weight:bold;\">● КОРПОРАТИВНАЯ ЛИЦЕНЗИЯ АКТИВНА</span>"
		htmlPage := "<!DOCTYPE html><html><head><meta charset=\"UTF-8\"><title>KvantumSafe Dashboard</title>" + cssStyles + "</head><body><div class=\"navbar\"><img class=\"logo-img\" src=\"/static/logo.jpg\" alt=\"Logo\"><h1>KvantumSafe Corporate Dashboard</h1></div><div class=\"content-area\"><h2>Мониторинг банковских приложений</h2><div class=\"meta-grid\"><div><strong>Организация:</strong> ОАО Коммерческий банк КЫРГЫЗСТАН (MBANK)</div><div><strong>Статус подписки:</strong> " + statusText + "</div><div><strong>Ядро платформы:</strong> Go Engine v2.6 Pro</div><div><strong>Зона сканирования:</strong> Внешний периметр</div></div><div style=\"text-align: center; margin-top: 40px; display: flex; justify-content: center; gap: 20px;\"><a class=\"btn-action\" href=\"/report/mbank\">Стартовать внешний аудит</a><a class=\"btn-bridge\" href=\"/report/quantum-bridge\">Запустить крипто-мост (NIST / GmSSL)</a></div></div></body></html>"
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		_, _ = w.Write([]byte(htmlPage))
	})

	// 2. РОУТЕР ВНЕШНЕГО ОТЧЕТА MBANK
	http.HandleFunc("/report/mbank", func(w http.ResponseWriter, r *http.Request) {
		currentTime := time.Now().Format("2006-01-02 15:04")
		reportData := UniversalIdentifySoftware("95.46.154.19")
		var rowsHtml string
		for _, app := range reportData.DetectedApps {
			rowsHtml += "<tr class=\"danger-row\"><td><strong>• " + html.EscapeString(app.Name) + "</strong><span class=\"status-badge\">🛑 Уязвимо перед дешифрованием</span></td><td><span style=\"background:#e2e8f0; padding:4px 8px; border-radius:4px; font-weight:bold;\">" + fmt.Sprintf("%d", app.Port) + "</span> (Версия: " + html.EscapeString(app.Version) + ")</td></tr>"
		}
		htmlPage := "<!DOCTYPE html><html><head><meta charset=\"UTF-8\"><title>KvantumSafe — Внешний Отчет</title>" + cssStyles + "</head><body><div class=\"navbar\"><img class=\"logo-img\" src=\"/static/logo.jpg\" alt=\"Logo\"><h1>KvantumSafe — Внешний Отчет</h1></div><div class=\"content-area\"><p><a href=\"/\" style=\"color: #004d40; text-decoration: none; font-weight: bold;\">← Назад к панели</a></p><h2>Результаты внешнего инфраструктурного сканирования</h2><div class=\"meta-grid\"><div><strong>Объект проверки:</strong> mbank.kg (95.46.154.19)</div><div><strong>Уровень угрозы:</strong> <span style=\"color:#b71c1c; font-weight:bold;\">КРИТИЧЕСКИЙ</span></div><div><strong>Дата и время:</strong> " + currentTime + "</div><div><strong>Статус подписки:</strong> <span style=\"color:#16a34a; font-weight:bold;\">● КОРПОРАТИВНАЯ ЛИЦЕНЗИЯ АКТИВНА</span></div></div><table class=\"table-vulnerabilities\"><thead><tr><th>Приложение</th><th>Порт / Версия</th></tr></thead><tbody>" + rowsHtml + "</tbody></table><div class=\"conclusion-box\"><h3>Рекомендательное заключение:</h3><ol><li>Срочно перевести инфраструктуру в Stealth Mode.</li><li>Интегрировать стандарты NIST (ML-KEM) и китайские криптоалгоритмы GmSSL (SM4-GCM / SM9).</li></ol></div><div style=\"text-align: center; margin-top: 30px;\"><a class=\"btn-internal\" href=\"/report/internal\">Запустить внутренний аудит безопасности серверов банка</a></div></div></body></html>"
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		_, _ = w.Write([]byte(htmlPage))
	})

	// 3. РОУТЕР ДЛЯ ВНУТРЕННЕГО ОТЧЕТА (ТАБЛИЦА С ПАРОЛЯМИ ИЗ INTERNAL_FS.GO)
	http.HandleFunc("/report/internal", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		currentTime := time.Now().Format("2006-01-02 15:04")
		internalList, _ := ExecuteInternalInspection("C:\\BankServer\\Secret")
		var rowsHtml string
		for _, vuln := range internalList {
			rowsHtml += "<tr class=\"warning-row\"><td><strong style=\"color:#b45309;\">⚠️ " + html.EscapeString(vuln.Type) + "</strong><br><span style=\"font-size:12px; color:#64748b;\">Путь: " + html.EscapeString(vuln.FilePath) + "</span></td><td>" + html.EscapeString(vuln.Details) + "</td><td style=\"color:#16a34a; font-weight:bold; text-align:center;\">🛡️ Изолирован (0600)</td></tr>"
		}
		htmlPage := "<!DOCTYPE html><html><head><meta charset=\"UTF-8\"><title>KvantumSafe — Внутренний Аудит</title>" + cssStyles + "</head><body><div class=\"navbar\"><img class=\"logo-img\" src=\"/static/logo.jpg\" alt=\"Logo\"><h1>KvantumSafe — Закрытый Внутренний Аудит</h1></div><div class=\"content-area\"><p><a href=\"/report/mbank\" style=\"color: #004d40; text-decoration: none; font-weight: bold;\">← Вернуться к внешнему аудиту</a></p><h2>Отчет ревизии данных внутри сервера (Data at Rest)</h2><div class=\"meta-grid\" style=\"border-left-color: #d97706;\"><div><strong>Сервер проверки:</strong> Core-Node-01</div><div><strong>Статус комплаенса:</strong> <span style=\"color:#b45309; font-weight:bold;\">⚠️ ОБНАРУЖЕНЫ НАРУШЕНИЯ</span></div><div><strong>Время ревизии:</strong> " + currentTime + "</div><div><strong>Статус подписки:</strong> <span style=\"color:#16a34a; font-weight:bold;\">● КОРПОРАТИВНАЯ ЛИЦЕНЗИЯ АКТИВНА</span></div></div><table class=\"table-vulnerabilities\"><thead><tr><th>Категория / Расположение</th><th>Описание угрозы</th><th>Статус защиты</th></tr></thead><tbody>" + rowsHtml + "</tbody></table></div></body></html>"
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		_, _ = w.Write([]byte(htmlPage))
	})

	// 4. РОУТЕР ДЛЯ ТРЕТЬЕГО ПРОДУКТА (КВАНТОВЫЙ МОСТ ИЗ QUANTUM_BRIDGE.GO)
	http.HandleFunc("/report/quantum-bridge", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		currentTime := time.Now().Format("2006-01-02 15:04")
		encryptedFiles := ProcessPostQuantumEncryption("C:\\BankServer\\CriticalData")
		var rowsHtml string
		for _, file := range encryptedFiles {
			rowsHtml += "<tr class=\"bridge-row\"><td><strong>📄 " + html.EscapeString(file.OriginalName) + "</strong></td><td><span style=\"color:#16a34a; font-weight:bold;\">🛡️ " + html.EscapeString(file.LibraryUsed) + "</span></td><td><span style=\"background:#e2e8f0; padding:4px 8px; border-radius:4px; font-weight:bold; font-size:12px;\">" + html.EscapeString(file.Algorithm) + "</span></td><td><code>" + html.EscapeString(file.CipherText) + "</code></td></tr>"
		}
		htmlPage := "<!DOCTYPE html><html><head><meta charset=\"UTF-8\"><title>KvantumSafe — Квантовый Мост</title>" + cssStyles + "</head><body><div class=\"navbar\"><img class=\"logo-img\" src=\"/static/logo.jpg\" alt=\"Logo\"><h1>KvantumSafe — Постквантовый Мост</h1></div><div class=\"content-area\"><p><a href=\"/\" style=\"color: #004d40; text-decoration: none; font-weight: bold;\">← Назад к панели</a></p><h2>Активное сквозное шифрование критических потоков данных (Data in Transit)</h2><div class=\"meta-grid\" style=\"border-left-color: #2563eb;\"><div><strong>Операция:</strong> Постквантовый крипто-мост (NIST / GmSSL)</div><div><strong>Крипто-защита:</strong> NIST ML-KEM + GmSSL SM4-GCM (Суверенный стандарт КНР, стойкий к квантовым атакам)</div><div><strong>Время оркестрации:</strong> " + currentTime + "</div><div><strong>Лицензия:</strong> <span style=\"color:#16a34a; font-weight:bold;\">● КОРПОРАТИВНАЯ ЛИЦЕНЗИЯ АКТИВНА</span></div></div><table class=\"table-vulnerabilities\"><thead><tr><th>Исходный объект передачи</th><th>Библиотека ядра</th><th>Постквантовый алгоритм</th><th>Квантово-стойкий шифротекст</th></tr></thead><tbody>" + rowsHtml + "</tbody></table></div></body></html>"
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		_, _ = w.Write([]byte(htmlPage))
	})

	http.HandleFunc("/static/logo.jpg", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/logo.jpg")
	})

	fmt.Println(">>> Глобальная平台 KvantumSafe Pro успешно запущена <<<")
	fmt.Println("Сервис отчёта запущен на http://localhost:8080/")

	server := &http.Server{
		Addr:              ":8080",
		ReadHeaderTimeout: 3 * time.Second,
		ReadTimeout:       5 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       30 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil {
		fmt.Printf("[-] Ошибка запуска порта: %v\n", err)
	}
}
