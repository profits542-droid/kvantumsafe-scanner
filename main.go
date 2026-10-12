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

	// Маршрут реального B2B-аудита
	http.HandleFunc("/report/mbank", func(w http.ResponseWriter, r *http.Request) {
		domain := r.URL.Query().Get("domain")
		if domain == "" {
			domain = "bakai.kg"
		}

		report := CheckDomainSecurity(domain)
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w, RenderReportPage(report))
	})

	// Маршрут честной B2B-витрины (Спецификация)
	http.HandleFunc("/download", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w, RenderDownloadPage())
	})

	fmt.Printf("[SUCCESS] KvantumSafe TLS Scanner запущен на порту: %s\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Printf("Ошибка запуска сервера: %v\n", err)
	}
}
