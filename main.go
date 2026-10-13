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

	// 1. ГЛАВНАЯ СТРАНИЦА САЙТА-ВИТРИНЫ (ЮРИДИЧЕСКИ ВЫВЕРЕННЫЙ ENTERPRISE-КОНТУР)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w, `<html><head><style>
			body{font-family:'Segoe UI',Arial,sans-serif;margin:0;padding:40px;background:#111111;color:#eeeeee;}
			.container{max-width:900px;margin:0 auto;background:#1a1a1a;padding:40px;border-radius:12px;box-shadow:0 10px 30px rgba(0,0,0,0.7);border:1px solid #2d2d2d;}
			.pvt-header{background:#262626;padding:20px;border-radius:8px;border-left:4px solid #b39212;margin-bottom:30px;line-height:1.6;}
			h1{color:#ffffff;border-bottom:2px solid #b39212;padding-bottom:15px;margin-top:0;font-weight:600;}
			h2{color:#b39212;margin-top:30px;font-size:1.5em;border-left:4px solid #b39212;padding-left:10px;}
			p{line-height:1.6;color:#cccccc;}
			.product-card{background:#222222;padding:25px;border-radius:8px;margin:20px 0;border:1px solid #333333;transition:0.3s;}
			.product-card:hover{border-color:#b39212;transform:translateY(-2px);}
			.product-title{color:#ffffff;font-weight:bold;font-size:1.2em;margin-bottom:10px;display:flex;align-items:center;}
			.badge-free{background:#dc3545;color:#fff;font-size:0.75em;padding:3px 8px;border-radius:4px;margin-left:10px;font-weight:bold;}
			.badge-gold{background:#b39212;color:#000;font-size:0.75em;padding:3px 8px;border-radius:4px;margin-left:10px;font-weight:bold;}
			.badge-blue{background:#0066cc;color:#fff;font-size:0.75em;padding:3px 8px;border-radius:4px;margin-left:10px;font-weight:bold;}
			ul{line-height:1.7;padding-left:20px;color:#cccccc;}
			li strong{color:#ffffff;}
			.gold-btn{display:inline-block;background:linear-gradient(135deg, #b39212, #ffd700);color:#000;padding:15px 35px;text-decoration:none;border-radius:6px;font-weight:bold;margin-top:30px;box-shadow:0 4px 15px rgba(179,146,18,0.4);transition:0.3s;text-transform:uppercase;letter-spacing:1px;}
			.gold-btn:hover{transform:translateY(-2px);box-shadow:0 6px 20px rgba(179,146,18,0.6);}
			.contact-box{background:#222222;padding:20px;border-radius:8px;margin-top:30px;border:1px solid #2d2d2d;}
			.footer-text{color:#666666;font-size:0.85em;text-align:center;margin-top:40px;border-top:1px solid #2d2d2d;padding-top:20px;}
		</style></head><body>
			<div class='container'>
				
				<!-- НАШЕ ОФИЦИАЛЬНОЕ НАПРАВЛЕНИЕ НА САМОМ ВИДНОМ МЕСТЕ -->
				<div class='pvt-header'>
					<b>📌 ОФИЦИАЛЬНЫЙ ПРОФИЛЬ ДЕЯТЕЛЬНОСТИ РЕЗИДЕНТА ПВТ КР:</b><br>
					Комплексная разработка специализированного программного обеспечения, проектирование интеллектуальных ИИ-ассистентов и автономных ИИ-агентов, сквозной нормативно-правовой и технологический комплаенс-контроль, а также оказание профильной юридической помощи для финансовых организаций, распределенных ИТ-систем и Web3-инфраструктуры.
				</div>

				<h1>🏛️ Экосистема Программных Комплексов «Квантум Сейф»</h1>
				<p><b>Разработчик:</b> ОсОО «Квантум Сейф» (Официальный резидент Парка высоких технологий Кыргызской Республики).</p>

				<!-- БЛОК 1: БЕСПЛАТНЫЙ ДИАГНОСТИЧЕСКИЙ АУДИТ -->
				<h2>🔍 Бесплатный Диагностический Экспресс-Аудит ИТ-Периметра:</h2>
				
				<div class='product-card' style='border-left: 4px solid #dc3545;'>
					<div class='product-title'>🛡️ Сканирование уязвимостей внешнего и внутреннего контура <span class='badge-free'>БЕСПЛАТНАЯ ДИАГНОСТИКА</span></div>
					<p>Мы бесплатно выявляем критические проблемы, скрытые угрозы и технологические девиации вашей ИТ-инфраструктуры. <strong>Устранение обнаруженных проблем, перевод систем под защиту и сквозная автоматизация комплаенса осуществляются в рамках платных прикладных модулей:</strong></p>
					<ul>
						<li><strong>Внешний экспресс-анализ:</strong> Проверка защищенных шлюзов на соответствие международным спецификациям NIST и китайскому государственному стандарту GM/T (алгоритмы SM2/SM3/SM4).</li>
						<li><strong>Внутренний аудит кода АБС:</strong> Автоматический поиск уязвимостей верхнего уровня (OWASP Top 10) и критических SQL-инъекций ядра.</li>
					</ul>
				</div>

				<!-- БЛОК 2: НАШИ ГЛАВНЫЕ РЕШЕНИЯ ТЕКУЩИХ БОЛЕЙ РЫНКА -->
				<h2>🎯 Продукты Быстрого Реагирования (Решение выявленных проблем):</h2>

				<div class='product-card'>
					<div class='product-title'>📦 AI P2P Referee & Compliance Gateway <span class='badge-gold'>РЕШЕНИЕ ДЛЯ ФИНТЕХА</span></div>
					<p>Автономная мультимодальная ИИ-платформа для моментального разрешения споров в P2P-сделках и автоматической верификации транзакций. Развертывание интеллектуальных ИИ-ассистентов и ИИ-агентов полностью снимает операционную нагрузку с комплаенс-отделов банков при блокировках счетов.</p>
					<ul>
						<li><strong>ИИ-Vision OCR:</strong> Автоматическое распознавание параметров чека (Сумма, Время, ФИО, ID) за 5 секунд, полностью защищая банк от фальшивых чеков из Telegram-ботов.</li>
						<li><strong>Match Module:</strong> Прямая API-сверка данных скриншота с реальной выпиской АБС банка для мгновенного подтверждения легальности сделки комплаенс-офицеру.</li>
					</ul>
				</div>

				<div class='product-card'>
					<div class='product-title'>📦 Quantum Sanctions Screening & Graph AI Engine <span class='badge-gold'>АНТИСАНКЦИОННЫЙ ШЛЮЗ</span></div>
					<p>Интеллектуальная система сквозного скрининга международных SWIFT / ISO-20022 потоков для полной ликвидации рисков попадания банков КР под вторичные санкции.</p>
					<ul>
						<li><strong>Графовый ИИ (Graph AI):</strong> Построение цифровых карт связей и выявление скрытых подсанкционных бенефициаров до 4-го колена учредителей по базам Минюста.</li>
						<li><strong>Сквозной контроль:</strong> Онлайн-сверка контрагентов со списками OFAC (SDN-листы), ЕС, Великобритании и ООН.</li>
					</ul>
				</div>

				<!-- БЛОК 3: ТЯЖЕЛАЯ АРТИЛЛЕРИЯ ДЛЯ АПСЕЙЛА (БЕЗ СЛОВА ШИФРОВАНИЕ) -->
				<h2>🛡️ Тяжелая Проактивная Артиллерия (Enterprise-Контур ПВТ):</h2>

				<div class='product-card'>
					<div class='product-title'>🔒 Quantum Web3 Orchestrator & Regulator <span class='badge-blue'>ENTERPRISE КОНТУР</span></div>
					<p>Программный комплекс, обеспечивающий комплексное оркестрирование распределенных нод, криптографическое контейнирование и высокотехнологичную инкапсуляцию банковских данных в защищенные контейнеры на базе гибридных постквантовых алгоритмов (NIST ML-KEM + китайский стандарт защиты GM/T SM4-GCM).</p>
				</div>

				<div class='product-card'>
					<div class='product-title'>🔒 AI-AntiHacker (Guard Engine) <span class='badge-blue'>ENTERPRISE ЗАЩИТА</span></div>
					<p>Интеллектуальная система защиты оперативной памяти серверной инфраструктуры (RAM), осуществляющая проактивное пресечение эксплуатации уязвимостей нулевого дня методом микросекундной изоляции скомпрометированных секторов (Stealth RAM Isolation).</p>
				</div>

				<center>
					<a href='/download?lang=ru' class='gold-btn'>🌟 Посмотреть Условия Внедрения и Тарифы</a>
				</center>

				<!-- НАШИ ОФИЦИАЛЬНЫЕ КОНТАКТЫ ДЛЯ КЛИЕНТОВ И БАНКОВ -->
				<div class='contact-box'>
					<h3 style='color:#b39212;margin-top:0;'>📞 Юридический отдел и Департамент B2B-внедрения:</h3>
					<p style='margin:5px 0;'><b>Адрес:</b> Кыргызская Республика, г. Ош, ул. Ленина / г. Бишкек, Проспект Чуй, 265/1 (ПВТ КР)</p>
					<p style='margin:5px 0;'><b>Официальный Email:</b> corporate@kvantumsafe.tech</p>
					<p style='margin:5px 0;'><b>Телефон / Telegram для связи:</b> +996 (700) 00-XX-XX / @QuantumSafe_Compliance</p>
				</div>

				<div class='footer-text'>
					© 2026 ОсОО «Квантум Сейф». Все права защищены. Продуктовая линейка депонирована в соответствии с регламентом Дирекции ПВТ КР. Оказание юридической помощи и ИТ-комплаенса.
				</div>
			</div>
		</body></html>`)
	})

	// 2. СТРАНИЦА ЗОЛОТОЙ КНОПКИ (ПОРТФЕЛЬ ТАРИФОВ)
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
				<h2>📥 Варианты Решения Выявленных Проблем</h2>
				
				<div class='price-box'>
					<h3>🚀 Ликвидация P2P-Фрода и Автоматизация Санкционного Скрининга:</h3>
					<ul>
						<li><strong>Транзакционная модель:</strong> $0.10 – $0.50 за каждый успешно верифицированный и защищенный ИИ-арбитром тикет.</li>
						<li><strong>Фиксированная SaaS-подписка (Устранение рисков):</strong> От $1,500/мес для финтех-стартапов до $15,000/мес для крупных Enterprise-платформ с дообучением ИИ.</li>
						<li><strong>Первичная интеграция API шлюзов (Setup Fee):</strong> $3,000 – $5,000 разово.</li>
					</ul>
				</div>

				<div class='price-box' style='border-left-color: #0066cc;'>
					<h3>🛡️ Комплексный Глубокий Контур (Тяжелая Артиллерия):</h3>
					<p style='color: #cccccc;'>Развертывание систем <b>Quantum Web3 Orchestrator & Regulator</b> и <b>AI-AntiHacker (Guard Engine)</b> внутри вашего ЦОД (On-Premise) для полной ликвидации системных уязвимостей.</p>
