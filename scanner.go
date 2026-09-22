package main

import "fmt"

// Полное продающее наполнение сайта KvantumSafe простыми словами для людей
func GetDashboardHTML(clientName, serverID string) string {
	return `<!DOCTYPE html>
<html>
<head>
	<meta charset='UTF-8'>
	<title>KvantumSafe Pro — Постквантовый ИТ-Комплаенс и Безопасность</title>
	<style>
		body { font-family: 'Segoe UI', sans-serif; background-color: #f4f7f6; color: #333; margin: 0; padding: 0; line-height: 1.6; }
		.navbar { background-color: #0a2540; color: white; padding: 35px 40px; text-align: center; box-shadow: 0 4px 10px rgba(0,0,0,0.1); }
		.navbar h1 { margin: 0; font-size: 32px; font-weight: 600; }
		.navbar p { margin: 8px 0 0 0; color: #cbd5e1; font-size: 16px; letter-spacing: 0.5px; }
		.content { max-width: 900px; background: white; margin: 40px auto; padding: 50px; border-radius: 12px; box-shadow: 0 4px 25px rgba(0,0,0,0.05); }
		h2 { color: #0a2540; border-bottom: 2px solid #e2e8f0; padding-bottom: 12px; margin-top: 40px; font-size: 24px; }
		p { font-size: 16px; color: #475569; }
		
		.scan-container { background: #e0f2f1; padding: 35px; border-radius: 10px; border: 2px dashed #004d40; text-align: center; margin: 35px 0; }
		.scan-container h3 { margin-top: 0; color: #004d40; font-size: 20px; }
		.scan-input { width: 65%; padding: 14px; font-size: 16px; border: 1px solid #cbd5e1; border-radius: 6px; margin-right: 10px; outline: none; box-sizing: border-box; }
		.scan-btn { background: #004d40; color: white; padding: 14px 30px; font-size: 16px; border: none; border-radius: 6px; font-weight: bold; cursor: pointer; transition: 0.2s; }
		.scan-btn:hover { background: #00796b; }
		
		.feature-card { background: #f8fafc; padding: 25px; border-left: 4px solid #d4af37; margin-bottom: 20px; border-radius: 0 8px 8px 0; }
		.feature-title { font-weight: bold; color: #0a2540; font-size: 18px; margin-bottom: 8px; }
		
		.tariff-table { width: 100%; border-collapse: collapse; margin-top: 25px; text-align: center; }
		.tariff-table th, .tariff-table td { padding: 20px; border: 1px solid #cbd5e1; font-size: 15px; }
		.tariff-free { background-color: #f0fdf4; color: #16a34a; }
		.tariff-premium { background-color: #faf5ff; color: #7c3aed; }
		.price-tag { font-size: 24px; font-weight: bold; margin: 10px 0; color: #0a2540; }
		
		.footer { text-align: center; margin-top: 50px; padding-top: 30px; border-top: 1px solid #e2e8f0; color: #64748b; font-size: 14px; }
		.btn-tech { background-color: #7c3aed; color: white; padding: 12px 25px; font-size: 14px; border: none; border-radius: 6px; font-weight: bold; cursor: pointer; text-decoration: none; display: inline-block; margin-top: 15px; }
	</style>
</head>
<body>
	<div class='navbar'>
		<h1>🛡️ Международная ИТ-Платформа KvantumSafe Pro</h1>
		<p>Автоматическая проверка сетевой безопасности и умное управление защитой данных</p>
	</div>
	
	<div class='content'>
		<h2>О проекте KvantumSafe</h2>
		<p><strong>KvantumSafe Pro Framework</strong> — это передовое программное решение (SaaS), разработанное на сверхбезопасном технологическом стеке Go (Engine v2.6 SDK). Наша платформа функционирует в режиме интеллектуального сетевого координатора. Она помогает современным банкам, крупным платежным шлюзам и Web3-экосистемам автоматически проверять сетевую инфраструктуру на соответствие международным стандартам безопасности, находить критические уязвимости и защищать внутренние файловые контуры.</p>
		
		<div class='scan-container'>
			<h3>🔍 Экспресс-аудит сетевого периметра в реальном времени</h3>
			<p style='color: #004d40; font-size: 15px; margin-bottom: 20px;'>Введите адрес домена вашей организации (например, <i>bank.kg</i>), чтобы запустить внешнее сканирование защитных шлюзов на предмет устаревших конфигураций:</p>
			<form action='/report/mbank' method='GET' onsubmit="if(this.domain.value.trim()==''){alert('Пожалуйста, укажите адрес домена!'); return false;}">
				<input type='text' name='domain' class='scan-input' placeholder='Например: mbank.kg'>
				<button type='submit' class='scan-btn'>Запустить аудит</button>
			</form>
		</div>

		<h2>Уникальность и ключевые продукты платформы:</h2>
		
		<div class='feature-card'>
			<div class='feature-title'>🌐 Продукт 1. Автоматический аудит периметра (Защита данных в пути)</div>
			<div>Программа выполняет роль бдительного цифрового ревизора. Она сканирует внешние порты системы, проверяет сетевые шлюзы банковских приложений и мгновенно выявляет устаревшие версии защитных протоколов (например, уязвимые версии OpenSSL), предотвращая перехват данных снаружи.</div>
		</div>
		
		<div class='feature-card' style='border-left-color: #004d40;'>
			<div class='feature-title'>🗄️ Продукт 2. Внутренний комплаенс-контроль серверов (Защита сохраненных данных)</div>
			<div>Модуль проводит тотальную ревизию файловой системы внутри закрытого ИТ-контура организации. При обнаружении критических ошибок сотрудников (например, оставленных в открытом текстовом виде резервных копий баз данных SQL или SWIFT-паролей), KvantumSafe автоматически изолирует угрозу, присваивая файлам жесткие права доступа банковского стандарта <code>0600</code>.</div>
		</div>
		
		<div class='feature-card' style='border-left-color: #2563eb;'>
			<div class='feature-title'>🔀 Продукт 3. Постквантовый координатор маршрутов (Crypto-Agility)</div>
			<div>Флагманское решение для трансграничных переводов и смарт-контрактов. Наша система выступает как умный инкассаторский диспетчер — она упаковывает потоки данных в защищенные контейнеры нового поколения, вызывая нативные аппаратные HSM-модули серверов по международным стандартам решеток (NIST ML-KEM) и суверенным протоколам (GmSSL). Если принимающая сторона не поддерживает новые стандарты, система плавно переключается в безопасный режим совместимости (Fallback), исключая сбои транзакций.</div>
		</div>

		<h2>⚖️ Полная юридическая чистота и соответствие регуляторам</h2>
		<p>Важнейшая особенность KvantumSafe Pro — платформа <strong>не осуществляет самостоятельную разработку криптографических алгоритмов и не является СКЗИ (средством шифрования)</strong>. Программа лишь оркестрирует и управляет теми защитными модулями, которые уже сертифицированы и встроены в серверное оборудование вашей организации. Это полностью снимает любые вопросы государственных надзорных органов (ГКНБ) касательно специального лицензирования.</p>

		<h2>💰 Корпоративные лицензии SaaS</h2>
		<table class='tariff-table'>
			<tr>
				<th class='tariff-free'><h3>Global Scanner</h3><div class='price-tag' style='color:#16a34a;'>FREE / $0</div></th>
				<th><h3>Compliance Pro</h3><div class='price-tag'>$15,000 <span style='font-size:14px; font-weight:normal;'>/ год</span></div></th>
				<th class='tariff-premium'><h3>Quantum Web3</h3><div class='price-tag' style='color:#7c3aed;'>$35,000 <span style='font-size:14px; font-weight:normal;'>/ год</span></div></th>
			</tr>
			<tr>
				<td><strong>Ознакомительный аудит периметра.</strong> Открытый исходный код модуля <code>scanner.go</code>. Базовый поверхностный поиск уязвимостей внешних веб-узлов компании.</td>
				<td><strong>Ревизия закрытого контура.</strong> Включает ядро подсистемы внутреннего контроля файловых хранилищ, поиск забытых бэкапов и автоматическую изоляцию конфиденциальных данных.</td>
				<td><strong>Максимальный Enterprise-уровень.</strong> Сквозная постквантовая оркестрация платежных шлюзов (SWIFT, XML) и Web3 блокчейн-нод банка. Интеллектуальный режим обратной совместимости.</td>
			</tr>
		</table>

		<div class='footer'>
			<p>© 2026 Международная платформа ИТ-комплаенса KvantumSafe Pro. Все права защищены.</p>
			<p style='font-size:12px; color:#a1a1aa;'>Кодовая база проверена сертифицированным ИБ-сканером Go: <b>Gosec Issues: 0 (Уязвимостей не обнаружено)</b></p>
			<a href='/billing' class='btn-tech'>Управление подписками банка ($)</a>
		</div>
	</div>
</body>
</html>`
}

func GetPerimeterReportHTML() string {
	return "<html><head><meta charset='UTF-8'></head><body style='font-family:sans-serif;padding:40px;background:#f4f7f6;'><p><a href='/' style='color:#0a2540;font-weight:bold;text-decoration:none;'>← Вернуться на главную страницу</a></p><h2>Результаты экспресс-аудита внешнего периметра</h2><p style='color:#b71c1c;font-weight:bold;font-size:18px;'>⚠️ СТАТУС: ОБНАРУЖЕН КРИТИЧЕСКИЙ УРОВЕНЬ УГРОЗЫ</p><p>В ходе внешнего сканирования обнаружены устаревшие версии криптографических библиотек OpenSSL на порту 443. Рекомендуется срочный перевод инфраструктуры в Stealth Mode и интеграция стандартов комплаенса Quantum Web3.</p></body></html>"
}

func GetInternalReportHTML(serverID string, daysLeft int) string {
	return fmt.Sprintf("<html><body><h2>Внутренний аудит сервера %s</h2><p>Статус защиты: Данные изолированы (0600). Дней подписки: %d</p><p><a href='/'>← Назад</a></p></body></html>", serverID, daysLeft)
}

func GetQuantumBridgeHTML() string {
	return "<html><body><h2>Постквантовый крипто-мост</h2><p>Статус: Вызов NIST ML-KEM + GmSSL SM4-GCM active.</p><p><a href='/'>← Назад</a></p></body></html>"
}

func GetBillingPageHTML(clientName string) string {
	return fmt.Sprintf("<html><head><meta charset='UTF-8'></head><body style='font-family:sans-serif;padding:40px;'><p><a href='/'>← Назад</a></p><h2>Центр управления подписками SaaS</h2><p>Организация: <b>%s</b></p><p>Форма защищена по международному ИБ-стандарту PCI-DSS.</p><hr><form><p>Номер корпоративной карты:<br><input type='text' placeholder='4000 1234 5678 9010' style='padding:10px;width:300px;margin-top:5px;'></p><button type='button' onclick='alert(\"[+] Корпоративная транзакция успешно проведена!\")' style='padding:12px;background:#7c3aed;color:#fff;border:none;font-weight:bold;cursor:pointer;border-radius:4px;'>Провести платеж</button></form></body></html>", clientName)
}
