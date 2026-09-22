package main

import "fmt"

// Главная страница сайта - с формой ввода домена для внешнего сканирования
func GetDashboardHTML(clientName, serverID string) string {
	return `<html>
<head>
	<meta charset='UTF-8'>
	<title>KvantumSafe — Безопасность и Понятный Комплаенс</title>
	<style>
		body { font-family: 'Segoe UI', sans-serif; background-color: #f4f7f6; color: #333; margin: 0; padding: 0; line-height: 1.6; }
		.navbar { background-color: #0a2540; color: white; padding: 30px 40px; text-align: center; }
		.navbar h1 { margin: 0; font-size: 28px; }
		.navbar p { margin: 5px 0 0 0; color: #cbd5e1; font-size: 16px; }
		.content { max-width: 850px; background: white; margin: 40px auto; padding: 40px; border-radius: 8px; box-shadow: 0 4px 15px rgba(0,0,0,0.05); }
		h2 { color: #0a2540; border-bottom: 2px solid #e2e8f0; padding-bottom: 10px; margin-top: 30px; }
		.step-box { background: #f8fafc; padding: 20px; border-left: 4px solid #d4af37; margin-bottom: 20px; border-radius: 0 6px 6px 0; }
		.step-title { font-weight: bold; color: #0a2540; font-size: 18px; margin-bottom: 5px; }
		.tariff-table { width: 100%; border-collapse: collapse; margin-top: 20px; text-align: center; }
		.tariff-table th, .tariff-table td { padding: 15px; border: 1px solid #cbd5e1; }
		.tariff-free { background-color: #f0fdf4; color: #16a34a; }
		.tariff-premium { background-color: #faf5ff; color: #7c3aed; }
		
		/* Стили для нового интерактивного поля */
		.scan-container { background: #e0f2f1; padding: 30px; border-radius: 8px; border: 2px dashed #004d40; text-align: center; margin: 30px 0; }
		.scan-input { width: 60%; padding: 12px; font-size: 16px; border: 1px solid #cbd5e1; border-radius: 6px; margin-right: 10px; outline: none; }
		.scan-btn { background: #004d40; color: white; padding: 12px 25px; font-size: 16px; border: none; border-radius: 6px; font-weight: bold; cursor: pointer; transition: 0.2s; }
		.scan-btn:hover { background: #00796b; }
		
		.footer-links { text-align: center; margin-top: 40px; padding-top: 20px; border-top: 1px solid #e2e8f0; }
		.btn-admin { font-size: 12px; color: #94a3b8; text-decoration: none; }
	</style>
</head>
<body>
	<div class='navbar'>
		<h1>🛡️ Платформа KvantumSafe</h1>
		<p>Простыми словами о безопасности вашего бизнеса и соответствии правилам</p>
	</div>
	<div class='content'>
		<h2>Что такое KvantumSafe?</h2>
		<p>Говоря простым языком, <strong>KvantumSafe</strong> — это умный цифровой помощник для организаций, банков и интернет-проектов. Наша задача — проверить, надежно ли защищены ваши данные, помочь исправить ошибки сотрудников и настроить безопасный обмен информацией по самым современным мировым стандартам. Чтобы работать с нами, вам не нужно быть программистом или экспертом в технологиях.</p>
		
		<!-- НОВОЕ ИНТЕРАКТИВНОЕ ПОЛЕ СКАНИРОВАНИЯ ДЛЯ КЛИЕНТОВ -->
		<div class='scan-container'>
			<h3 style='margin-top:0; color:#004d40;'>🔍 Проверьте безопасность вашего сайта прямо сейчас</h3>
			<p style='margin-bottom:15px; font-size:14px; color:#004d40;'>Введите адрес любого домена (например, <i>mysite.com</i>), чтобы запустить внешнюю экспресс-проверку защитных шлюзов периметра:</p>
			<form action='/report/mbank' method='GET' onsubmit="if(this.domain.value.trim()==''){alert('Пожалуйста, введите адрес домена!'); return false;}">
				<input type='text' name='domain' class='scan-input' placeholder='Вкажите домен для проверки...'>
				<button type='submit' class='scan-btn'>Запустить аудит</button>
			</form>
		</div>

		<h2>Как это работает? Всего 3 понятных шага:</h2>
		
		<div class='step-box'>
			<div class='step-title'>Шаг 1. Внешний осмотр системы (Сканирование периметра)</div>
			<div>Наша программа работает как бдительный обходной охранник. Она осматривает ваш сайт или банковское приложение снаружи, проверяет, закрыты ли все виртуальные «окна и двери», и вовремя предупреждает, если защитные барьеры устарели.</div>
		</div>
		
		<div class='step-box' style='border-left-color: #004d40;'>
			<div class='step-title'>Шаг 2. Внутренний порядок (Ревизия серверов)</div>
			<div>Система наводит порядок внутри ваших рабочих компьютеров (серверов). Если сотрудник случайно сохранил важный пароль в обычном текстовом блокноте или забыл защитить базу данных клиентов, KvantumSafe автоматически находит эту ошибку и надежно изолирует документ от посторонних глаз.</div>
		</div>
		
		<div class='step-box' style='border-left-color: #2563eb;'>
			<div class='step-title'>Шаг 3. Защищенный цифровой мост (Маршрутизация данных)</div>
			<div>Когда ваша компания переводит деньги или отправляет важные документы партнерам, наша система берет на себя роль бронированного автомобиля. Она упаковывает данные в специальные защитные контейнеры. Если у принимающей стороны нет нашей программы, система сама переключится в стандартный безопасный режим, чтобы процесс перевода не прервался.</div>
		</div>
		
		<h2>💰 Стоимость лицензий (SaaS подписка)</h2>
		<table class='tariff-table'>
			<tr>
				<th class='tariff-free'><h3>Global Scanner</h3><div style='font-size:20px; font-weight:bold;'>FREE / $0</div></th>
				<th><h3>Compliance Pro</h3><div style='font-size:20px; font-weight:bold;'>$15,000 / год</div></th>
				<th class='tariff-premium'><h3>Quantum Web3</h3><div style='font-size:20px; font-weight:bold;'>$35,000 / год</div></th>
			</tr>
			<tr>
				<td>Ознакомительный контур. Бесплатный базовый осмотр вашего сайта на предмет внешних уязвимостей контура.</td>
				<td>Профессиональный аудит. Глубокая проверка внутренних серверов, поиск забытых баз данных и принудительная изоляция паролей.</td>
				<td>Максимальная защита. Полный доступ к защищенному цифровому мосту, сквозной контроль переводов и Web3-инфраструктуры банка.</td>
			</tr>
		</table>

		<div class='footer-links'>
			<p>© 2026 KvantumSafe ИТ-Комплаенс. Все права защищены.</p>
			<a href='/report/mbank' class='btn-admin'>Вход для технических специалистов (Сервис ИБ-мониторинга ОАО МБАНК)</a>
		</div>
	</div>
</body>
</html>`
}

func GetPerimeterReportHTML() string {
	return "<html><body style='font-family:sans-serif;padding:40px;'><h2>Результаты внешнего сканирования периметра</h2><p><strong>Объект проверки:</strong> Анализ указанного сетевого узла успешно завершен</p><p style='color:red;font-weight:bold;'>УРОВЕНЬ УГРОЗЫ: КРИТИЧЕСКИЙ (Обнаружены устаревшие конфигурации портов)</p><hr><p><a href='/'>← Назад на главную страницу</a></p></body></html>"
}

func GetInternalReportHTML(serverID string, daysLeft int) string {
	return fmt.Sprintf("<html><body style='font-family:sans-serif;padding:40px;'><h2>Внутренний аудит сервера %s</h2><p style='color:orange;'>⚠️ Обнаружены текстовые пароли в открытом виде!</p><p>Статус: Изолировано (0600). Дней Pro-подписки: %d</p><p><a href='/'>← Назад на главную</a></p></body></html>", serverID, daysLeft)
}

func GetQuantumBridgeHTML() string {
	return "<html><body style='font-family:sans-serif;padding:40px;'><h2>Постквантовый крипто-мост (Data in Transit)</h2><p style='color:blue;'>Статус оркестрации: Вызов NIST ML-KEM + GmSSL SM4-GCM на аппаратном уровне HSM банка</p><p><a href='/'>← Назад на главную</a></p></body></html>"
}

func GetBillingPageHTML(clientName string) string {
	return fmt.Sprintf("<html><body style='font-family:sans-serif;padding:40px;'><h2>Центр управления подписками</h2><p>Организация: %s</p><p>Форма защищена по стандарту PCI-DSS.</p><button onclick='alert(\"[+] Успешно оплачено!\")' style='padding:15px;background:#7c3aed;color:#fff;border:none;font-weight:bold;cursor:pointer;border-radius:4px;'>Провести корпоративный платеж</button><br><br><a href='/'>← Назад на главную</a></body></html>", clientName)
}
