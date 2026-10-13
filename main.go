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

	// 1. ГЛАВНАЯ СТРАНИЦА САЙТА-ВИТРИНЫ (СТРАТЕГИЯ FREEMIUM + UPSELL)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w, `<html><head><style>
			body{font-family:'Segoe UI',Arial,sans-serif;margin:0;padding:40px;background:#111111;color:#eeeeee;}
			.container{max-width:900px;margin:0 auto;background:#1a1a1a;padding:40px;border-radius:12px;box-shadow:0 10px 30px rgba(0,0,0,0.7);border:1px solid #2d2d2d;}
			h1{color:#ffffff;border-bottom:2px solid #b39212;padding-bottom:15px;margin-top:0;font-weight:600;}
			h2{color:#b39212;margin-top:30px;font-size:1.5em;border-left:4px solid #b39212;padding-left:10px;}
			p{line-height:1.6;color:#cccccc;}
			.product-card{background:#222222;padding:25px;border-radius:8px;margin:20px 0;border:1px solid #333333;transition:0.3s;}
			.product-card:hover{border-color:#b39212;transform:translateY(-2px);}
			.product-title{color:#ffffff;font-weight:bold;font-size:1.2em;margin-bottom:10px;display:flex;align-items:center;}
			.badge-free{background:#28a745;color:#fff;font-size:0.75em;padding:3px 8px;border-radius:4px;margin-left:10px;font-weight:bold;}
			.badge-gold{background:#b39212;color:#000;font-size:0.75em;padding:3px 8px;border-radius:4px;margin-left:10px;font-weight:bold;}
			.badge-blue{background:#0066cc;color:#fff;font-size:0.75em;padding:3px 8px;border-radius:4px;margin-left:10px;font-weight:bold;}
			ul{line-height:1.7;padding-left:20px;color:#cccccc;}
			li strong{color:#ffffff;}
			.gold-btn{display:inline-block;background:linear-gradient(135deg, #b39212, #ffd700);color:#000;padding:15px 35px;text-decoration:none;border-radius:6px;font-weight:bold;margin-top:30px;box-shadow:0 4px 15px rgba(179,146,18,0.4);transition:0.3s;text-transform:uppercase;letter-spacing:1px;}
			.gold-btn:hover{transform:translateY(-2px);box-shadow:0 6px 20px rgba(179,146,18,0.6);}
			.footer-text{color:#666666;font-size:0.85em;text-align:center;margin-top:40px;border-top:1px solid #2d2d2d;padding-top:20px;}
		</style></head><body>
			<div class='container'>
				<h1>🏛️ Экосистема Программных Комплексов «Quantum Safe»</h1>
				<p><b>Разработчик:</b> ОсОО «Квантум Сейф» (Официальный резидент Парка высоких технологий Кыргызской Республики).</p>
				<p>Мы поставляем передовые ИИ-решения для автоматизации комплаенса и обеспечения сквозной безопасности коммерческих банков, финтех-сервисов и Web3-инфраструктуры.</p>

				<!-- БЛОК 1: БЕСПЛАТНЫЕ ИНСТРУМЕНТЫ ДЛЯ ВХОДА В ДОВЕРИЕ -->
				<h2>🚀 Бесплатные Инструменты Продвижения (Open Source на GitHub):</h2>
				
				<div class='product-card'>
					<div class='product-title'>🔍 Quantum External AI Scanner & Internal Vulnerability Scanner <span class='badge-free'>FREE / OPEN-SOURCE</span></div>
					<p>Наши утилиты внешнего экспресс-аудита сетевого периметра и внутреннего сканирования уязвимостей АБС доступны для ИТ-сообщества абсолютно бесплатно.</p>
					<ul>
						<li><strong>Внешний сканер:</strong> Живой автоматический аудит TLS-рукопожатий, проверка Cipher Suites на соответствие стандартам NIST и контроль SSL-сертификатов.</li>
						<li><strong>Внутренний сканер:</strong> Экспресс-анализ кода на предмет SQL-инъекций и критических уязвимостей верхнего уровня (OWASP Top 10).</li>
					</ul>
				</div>

				<!-- БЛОК 2: НАШИ ГЛАВНЫЕ РЕШЕНИЯ ТЕКУЩИХ БОЛЕЙ РЫНКА -->
				<h2>🎯 Продукты Быстрого Реагирования (Решение болей банков и финтеха):</h2>

				<div class='product-card'>
					<div class='product-title'>📦 AI P2P Referee & Compliance Gateway <span class='badge-gold'>🔥 ТРЕБУЕТСЯ РЫНКУ</span></div>
					<p>Автономная мультимодальная ИИ-платформа для моментального разрешения споров в P2P-сделках и автоматической верификации транзакций. Спасает банки КР от операционного хаоса при массовых блокировках счетов физлиц.</p>
					<ul>
						<li><strong>ИИ-Vision OCR:</strong> Считывание параметров чека (Сумма, Время, ФИО, ID) за 5 секунд, полностью защищая банк от фальшивых чеков из Telegram-ботов.</li>
						<li><strong>Match Module:</strong> Прямая API-сверка данных скриншота с реальной выпиской АБС банка для мгновенного подтверждения легальности сделки комплаенс-офицеру.</li>
					</ul>
				</div>

				<div class='product-card'>
					<div class='product-title'>📦 Quantum Sanctions Screening & Graph AI Engine <span class='badge-gold'>🔥 АКТУАЛЬНО</span></div>
					<p>Интеллектуальная система сквозного скрининга международных SWIFT / ISO-20022 потоков для полной ликвидации рисков попадания банков КР под вторичные санкции.</p>
					<ul>
						<li><strong>Графовый ИИ (Graph AI):</strong> Построение цифровых карт связей и выявление скрытых подсанкционных бенефициаров до 4-го колена учредителей по базам Минюста.</li>
						<li><strong>Сквозной контроль:</strong> Онлайн-сверка контрагентов со списками OFAC (SDN-листы), ЕС, Великобритании и ООН.</li>
					</ul>
				</div>

				<!-- БЛОК 3: ТЯЖЕЛАЯ АРТИЛЛЕРИЯ ДЛЯ АПСЕЙЛА -->
				<h2>🛡️ Тяжелая Проактивная Артиллерия (Enterprise-Контур ПВТ):</h2>

				<div class='product-card'>
					<div class='product-title'>🔒 Quantum Web3 Orchestrator & Regulator <span class='badge-blue'>ENTERPRISE</span></div>
					<p>Программный комплекс, обеспечивающий комплексное оркестрирование распределенных нод, криптографическую инкапсуляцию банковских данных в защищенные контейнеры на базе гибридного постквантового шифрования и регуляторный комплаенс распределенных систем.</p>
				</div>

				<div class='product-card'>
					<div class='product-title'>🔒 AI-AntiHacker (Guard Engine) <span class='badge-blue'>ENTERPRISE</span></div>
					<p>Интеллектуальная система защиты оперативной памяти серверной инфраструктуры (RAM), осуществляющая проактивное пресечение эксплуатации уязвимостей нулевого дня и предотвращение атак типа переполнения буфера методом микросекундной изоляции скомпрометированных секторов (Stealth RAM Isolation).</p>
				</div>

				<center>
					<a href='/download?lang=ru' class='gold-btn'>🌟 Посмотреть Условия и Тарифы Внедрения</a>
				</center>

				<div class='footer-text'>
					© 2026 ОсОО «Квантум Сейф». Все права защищены. Продуктовая линейка депонирована в соответствии с регламентом Дирекции ПВТ КР.
				</div>
			</div>
		</body></html>`)
	})

	// 2. СТРАНИЦА ЗОЛОТОЙ КНОПКИ (ПОРТФЕЛЬ ТАРИФОВ БЕЗ ЦЕН НА ФЛАГМАНЫ)
	http.HandleFunc("/download", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w, `<html><head><style>
			body{font-family:'Segoe UI',Arial,sans-serif;margin:0;padding:40px;background:#111111;color:#eeeeee;text-align:center;}
			.card{background:#1a1a1a;padding:40px;border-radius:12px;max-width:650px;margin:0 auto;box-shadow:0 10px 30px rgba(0,0,0,0.7);border:1px solid #2d2d2d;}
			.price-box{background:#222222;padding:25px;border-radius:8px;margin:20px 0;text-align:left;border-left:4px solid #b39212;line-height:1.6;}
			h2{color:#ffffff;font-weight:600;}
			h3{color:#b39212;margin-top:0;}
			ul{padding-left:20px;color:#cccccc;}
			.btn-back{display:inline-block;background:#b39212;color:#000;padding:12px 30px;text-decoration:none;border-radius:4px;font-weight:bold;margin-top:25px;text-transform:uppercase;font-size:0.9em;letter-spacing:1px;transition:0.3s;}
			.btn-back:hover{background:#ffd700;}
		</style></head><body>
			<div class='card'>
				<h2>📥 Варианты Внедрения и Коммерческие Условия</h2>
				
				<div class='price-box'>
					<h3>🚀 Продукты Локомотивы (AI P2P Referee & Санкционный Скрининг):</h3>
					<ul>
						<li><strong>Транзакционная модель:</strong> $0.10 – $0.50 за каждый успешно верифицированный и разобранный ИИ-арбитром тикет.</li>
						<li><strong>Фиксированная SaaS-подписка:</strong> От $1,500/мес для финтех-стартапов до $15,000/мес для крупных Enterprise-платформ с адаптацией под дизайн документов заказчика.</li>
						<li><strong>Первичная интеграция (Setup Fee):</strong> $3,000 – $5,000 разово за кастомизацию шлюзов API.</li>
					</ul>
				</div>

				<div class='price-box' style='border-left-color: #0066cc;'>
					<h3>🛡️ Крупный Enterprise-Контур (Тяжелая Артиллерия):</h3>
					<p style='color: #cccccc;'>Стоимость развертывания систем <b>Quantum Web3 Orchestrator & Regulator</b> и <b>AI-AntiHacker (Guard Engine)</b> рассчитывается индивидуально под архитектуру ЦОД финансовой организации по On-Premise лицензиям.</p>
					<p style='color: #b39212; font-weight: bold; margin-bottom: 0;'>Поставка включает интеграцию Hardware Dongle API и привязку к аппаратным токенам защиты. Цена предоставляется строго по официальному B2B-запросу.</p>
				</div>
				
				<p style='color:#00ff00;font-weight:bold;font-size:1.1em;'>Запрос на получение полной коммерческой документации отправлен в отдел комплаенса!</p>
				<a href='/' class='btn-back'>← На Главную Страницу</a>
			</div>
		</body></html>`)
	})

	fmt.Printf("[SUCCESS] Сканер-сервер успешно развернут на порту: %s\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Printf("Ошибка запуска сервера: %v\n", err)
	}
}
