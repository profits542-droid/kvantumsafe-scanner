package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Главная страница: теперь показывает чистую форму ввода без результатов
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w, RenderStartPage())
	})

	// Страница генерации отчета: срабатывает строго ПОСЛЕ нажатия кнопки сканирования
	http.HandleFunc("/report/mbank", func(w http.ResponseWriter, r *http.Request) {
		domain := r.URL.Query().Get("domain")
		if domain == "" {
			// Если нажали пустую кнопку, возвращаем на главную
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		report := CheckDomainSecurity(domain)
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w, RenderReportPage(report))
	})

	// Страница спецификации ПО, тарифов и NDA соглашения
	http.HandleFunc("/download", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w, RenderDownloadPage())
	})

	fmt.Printf("[SUCCESS] Квантум Сейф Веб-Ядро запущено на порту: %s\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Printf("Ошибка запуска сервера: %v\n", err)
	}
}
