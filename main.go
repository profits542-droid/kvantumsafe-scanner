package main

import (
	"log"
	"net/http"
)

func main() {
	// 1. Главная страница
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(RenderStartPage()))
	})

	// 2. Страница результатов ИИ-сканирования
	http.HandleFunc("/report/mbank", func(w http.ResponseWriter, r *http.Request) {
		domain := r.URL.Query().Get("domain")
		if domain == "" {
			w.Write([]byte(RenderStartPage()))
			return
		}
		report := ScanDomain(domain)
		w.Write([]byte(RenderReportPage(report)))
	})

	// 3. Страница спецификации софта ПВТ КР
	http.HandleFunc("/download", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(RenderDownloadPage()))
	})

	// 4. НАШ ОФИЦИАЛЬНЫЙ СВЯЗАННЫЙ МАРШРУТ ДЛЯ МАРКЕТПЛЕЙСА ИИ-РЕШЕНИЙ
	http.HandleFunc("/marketplace", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(RenderMarketplacePage()))
	})

	log.Println("[SUCCESS] Квантум Сейф Веб-Ядро запущено на порту: 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Ошибка запуска сервера: ", err)
	}
}