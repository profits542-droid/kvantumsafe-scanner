package main

import (
	"log"
	"net/http"
)

func main() {
	// 1. Стартовая страница витрины
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(RenderStartPage()))
	})

	// 2. Страница результатов сканирования (Фикс регистра: ScanDomain)
	http.HandleFunc("/report/mbank", func(w http.ResponseWriter, r *http.Request) {
		domain := r.URL.Query().Get("domain")
		if domain == "" {
			w.Write([]byte(RenderStartPage()))
			return
		}
		report := ScanDomain(domain)
		w.Write([]byte(RenderReportPage(report)))
	})

	// 3. Страница финансовой спецификации для ПВТ КР
	http.HandleFunc("/download", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(RenderDownloadPage()))
	})

	// 4. Страница интерактивного Маркетплейса решений
	http.HandleFunc("/marketplace", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(RenderMarketplacePage()))
	})

	log.Println("[SUCCESS] Квантум Сейф Веб-Ядро запущено на порту: 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Ошибка запуска сервера: ", err)
	}
}
