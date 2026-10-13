package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	// Считываем порт из окружения Render
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// 1. Главная страница сайта-витрины (Внешний, Внутренний сканеры и Оркестратор)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w, `<html><head><style>
			body{font-family:Arial,sans-serif;margin:0;padding:40px;background:#131313;color:#fff;}
			.container{max-width:850px;margin:0 auto;background:#1e1e1e;padding:40px;border-radius:12px;box-shadow:0 4px 15px rgba(0,0,0,0.5);}
			h1{color:#ffffff;border-bottom:2px solid #b39212;padding-bottom:10px;}
			h3{color:#b39212;margin-top:25px;}
			.tech-box{background:#262626;padding:20px;border-left:4px solid #b39212;border-radius:4px;margin:20px 0;line-height:1.6;}
			ul{line-height:1.7;padding-left:20px;}
			.gold-btn{display:inline-block;background:linear-gradient(135deg, #b39212, #ffd700);color:#000;padding:15px 30px;text-decoration:none;border-radius:6px;font-weight:bold;margin-top:25px;box-shadow:0 4px 10px rgba(179,146,18,0.3);transition:0.3s;}
			.gold-btn:hover{transform:translateY(-2px);box-shadow:0 6px 15px rgba(179,146,18,0.5);}
			hr{border:0;border-top:1px solid #333;margin:30px 0;}
		</style></head><body>
			<div class='container'>
				<h1>🏛️ Экосистема Программных Комплексов «Quantum Safe»</h1>
				<p><b>Разработчик:</b> ОсОО «Квантум Сейф» (Официальный резидент Парка высоких технологий Кыргызской Республики).</p>
				
				<div class='tech-box'>
					<b>Глобальная миссия:</b> Сквозная технологическая и правовая защита банковской ИТ-инфраструктуры, автоматизация комплаенс-контроля финансовых ISO-20022 потоков данных и полное пресечение рисков вторичных санкций.
				</div>

				<hr>
				<h3>1. Внешний ИИ-Сканер Сетевого Периметра (Quantum External AI Scanner)</h3>
				<p>Интеллектуальный программный комплекс для непрерывного внешнего аудита шлюзов финансовых организаций. Осуществляет мониторинг доступности портов, верифицирует Cipher Suites на соответствие жестким стандартам кибербезопасности NIST и контролирует сроки действия и легитимность цифровых SSL/TLS сертификатов банковского контура в автоматическом режиме.</p>

				<h3>2. Внутренний ИИ-Сканер Уязвимостей Ядра (Quantum Internal AI Security Scanner)</h3>
				<p>Глубокий анализатор ИТ-периметра, развертываемый непосредственно внутри локального закрытого контура банка (On-Premise ЦОД). Предназначен для постоянного статического и динамического аудита исходного кода автоматизированных банковских систем (АБС), выявления SQL-инъекций, уязвимостей нулевого дня (OWASP Top 10) и поиска логических девиаций в смарт-контрактах.</p>

				<h3>3. Постквантовый Оркестратор-Регулятор (Quantum Web3 Orchestrator & Regulator)</h3>
				<p>Главное управляющее ядро экосистемы. Отвечает за координацию защищенных нод, криптографическую инкапсуляцию транзакций в контейнеры на базе гибридного постквантового шифрования (NIST ML-KEM + GmSSL) и аппаратную привязку лицензий через Hardware Dongle API к физическим USB-токенам (FEITIAN / Рутокен 3.0).</p>

				<hr>
				<h3>⚙️ Встроенные прикладные модули безопасности:</h3>
				<ul>
					<li><b>AI P2P Referee & Anti-Fraud Module:</b> ИИ-арбитраж сделок, моментальное OCR-распознавание параметров чеков и Match-сверка с выпиской АБС для снижения нагрузки на саппорт на 80%%.</li>
					<li><b>Quantum Sanctions Screening Engine:</b> Графовый ИИ-скрининг скрытых цепочек бенефициаров до 4-го колена по мировым спискам (OFAC, SDN, ЕС) для полной ликвидации рисков санкций.</li>
					<li><b>AI-AntiHacker (Guard Engine):</b> Микросекундная защита оперативной памяти серверов (RAM) от переполнения буфера по методу Stealth RAM Isolation.</li>
				</ul>

				<center>
					<a href='/download?lang=ru' class='gold-btn'>🌟 Спецификация ПО и Техническая документация</a>
				</center>
			</div>
		</body></html>`)
	})

	// 2. Страница золотой кнопки (Подготовка документации)
	http.HandleFunc("/download", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w, `<html><head><style>
			body{font-family:Arial,sans-serif;margin:0;padding:40px;background:#131313;color:#fff;text-align:center;}
			.card{background:#1e1e1e;padding:40px;border-radius:12px;max-width:600px;margin:0 auto;box-shadow:0 4px 15px rgba(0,0,0,0.5);}
			.status-box{background:#262626;padding:20px;border-radius:6px;margin:20px 0;font-family:monospace;text-align:left;line-height:1.5;}
			.btn-back{display:inline-block;background:#b39212;color:#000;padding:12px 25px;text-decoration:none;border-radius:4px;font-weight:bold;margin-top:20px;}
		</style></head><body>
			<div class='card'>
				<h2>📥 Скачивание технической документации Quantum Safe</h2>
				<p style='color:#ccc;'>Подготовка пакета документов: <b>Specification_AI_Compliance_Gateway_RU.pdf</b></p>
				
				<div class='status-box'>
					[STATUS] Формирование выгрузки спецификации... ОК<br>
					[INFO] Экспертный узел ПВТ КР: ОсОО «Квантум Сейф»<br>
					[INFO] Контур защиты: Лицензия ACTIVE [Enterprise]
				</div>
				
				<p style='color:#00ff00;font-weight:bold;'>Запрос на загрузку спецификации программного комплекса успешно отправлен!</p>
				<a href='/' class='btn-back'>← Вернуться на главную страницу</a>
			</div>
		</body></html>`)
	})

	fmt.Printf("[SUCCESS] Витрина успешно запущена на порту: %s\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Printf("Ошибка сервера: %v\n", err)
	}
}
