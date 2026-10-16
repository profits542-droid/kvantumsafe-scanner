package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(RenderStartPage()))
	})
	http.HandleFunc("/report/mbank", func(w http.ResponseWriter, r *http.Request) {
		domain := r.URL.Query().Get("domain")
		if domain == "" {
			http.Error(w, "Введите домен или IP", http.StatusBadRequest)
			return
		}
		report := ScanDomain(domain)
		w.Write([]byte(RenderReportPage(report)))
	})
	http.HandleFunc("/download", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(RenderDownloadPage()))
	})
	http.HandleFunc("/marketplace", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(RenderDownloadPage()))
	})
	log.Println("[SUCCESS] Квантум Сейф Веб-Ядро запущено на порту: 8080")
	http.ListenAndServe(":8080", nil)
}