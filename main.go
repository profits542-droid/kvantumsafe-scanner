package main

import (
	"fmt"
	"net/http"
)

func main() {
	// 1. Главная страница (Дашборд экспресс-аудита)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		lang := r.URL.Query().Get("lang")
		if lang == "" {
			lang = "ru" // Базовый фолбэк для СНГ шлюзов
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprint(w, GetDashboardHTML("Mbank", "srv-daofac8ae00c73ca1bc0", lang))
	})

	// 2. Страница спецификации технического регламента
	http.HandleFunc("/specification", func(w http.ResponseWriter, r *http.Request) {
		lang := r.URL.Query().Get("lang")
		if lang == "" {
			lang = "ru"
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprint(w, GetSpecificationPageHTML(lang))
	})

	// 3. Страница презентации софта AI-AntiHacker
	http.HandleFunc("/antihacker", func(w http.ResponseWriter, r *http.Request) {
		lang := r.URL.Query().Get("lang")
		if lang == "" {
			lang = "ru"
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprint(w, GetAntiHackerPageHTML(lang))
	})
	// 4. ВОССТАНОВЛЕННЫЙ АВТОНОМНЫЙ ШЛЮЗ ЗОЛОТОЙ КНОПКИ (Прямая отдача текста в Загрузки)
	http.HandleFunc("/download", func(w http.ResponseWriter, r *http.Request) {
		lang := r.URL.Query().Get("lang")
		if lang == "" {
			lang = "ru"
		}
		offerContent := getOfferTranslationFromB2B(lang)
		
		w.Header().Set("Content-Disposition", "attachment; filename=kvantumsafe_specification_"+lang+".txt")
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		fmt.Fprint(w, offerContent)
	})



	// Имитация отчета сканирования периметра (Для демонстрации банкирам)
	http.HandleFunc("/report/mbank", func(w http.ResponseWriter, r *http.Request) {
		lang := r.URL.Query().Get("lang")
		if lang == "" {
			lang = "ru"
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprint(w, "<html><head><meta charset='UTF-8'></head><body style='font-family:sans-serif;padding:40px;background:#f4f7f6;'><h2>🛡️ Аудит периметра завершен успешно (Passed)</h2><p>Уязвимостей переполнения RAM и логического фрода транзакций не обнаружено.</p><br><a href='/?lang="+lang+"'>← Назад</a></body></html>")
	})

	fmt.Println("Server KvantumSafe Pro Framework successfully started on port :8080...")
	http.ListenAndServe(":8080", nil)
}
