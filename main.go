package main

import (
	"fmt"
	"time"
)

// Информация о подписке банка
type LicenseInfo struct {
	OrganizationName string
	ValidUntil       time.Time
}

// Модуль «Кибер-Айкидо» — растворение кода хакера в бесконечной ловушке
func AikidoTrap(attackerIP string) {
	fmt.Printf("\n[КИБЕР-АЙКИДО] Хакер с IP %s обнаружен 'по запаху'!\n", attackerIP)
	fmt.Println("[*] Не блокируем атаку, чтобы не выдать себя. Затягиваем хакера в невидимую ловушку...")

	for i := 1; i <= 3; i++ {
		fmt.Printf("[ЛОВУШКА] Этап %d: Генерируем ложный квантовый ключ. Скрипт хакера увязает в расчетах...\n", i)
		time.Sleep(1 * time.Second)
	}
	fmt.Printf("[УСПЕХ] Сила атаки с IP %s полностью исчерпана. Вредоносный код растворился в западне.\n", attackerIP)
}

// Контролер подписки (Биллинг) — проверяет, заплатил ли банк за SaaS
func CheckSaaSLicense(lic LicenseInfo) bool {
	if time.Now().After(lic.ValidUntil) {
		fmt.Printf("\n[ВНИМАНИЕ] Подписка для банка '%s' ИСТЕКЛА %s!\n", lic.OrganizationName, lic.ValidUntil.Format("2006-01-02"))
		fmt.Println("[-] КРИПТО-РЕГУЛИРОВЩИК ОТКЛЮЧЕН. Банк снова стал видимым для хакеров.")
		return false
	}
	fmt.Printf("[БИЛЛИНГ] Подписка банка '%s' активна. Режим невидимки работает штатно.\n", lic.OrganizationName)
	return true
}

func main() {
	fmt.Println(">>> Универсальный Крипто-Оркестратор v2.0 (Режим: Кибер-Айкидо) запущен <<<")

	// Симулируем проверку подписки (для теста создадим лицензию, которая истекла вчера)
	expiredLicense := LicenseInfo{
		OrganizationName: "Финтех Банк Бишкек",
		ValidUntil:       time.Now().Add(-24 * time.Hour),
	}

	// Запускаем проверку оплаты
	if !CheckSaaSLicense(expiredLicense) {
		fmt.Println("[СИСТЕМА] Работа приостановлена. Ожидание оплаты от руководства банка...")
	}

	// Симулируем, что на банк напал хакер, и включается наше Кибер-Айкидо
	fmt.Println("\n--- Имитация сетевой атаки на банк ---")
	hackerIP := "185.220.101.5"
	AikidoTrap(hackerIP)

	// ==========================================================================
	// 🔥 СВЯЗУЮЩИЙ МОСТ: ТЕПЕРЬ ОРКЕСТРАТОР ВЫЗЫВАЕТ ВНЕШНИЙ СКАНЕР-ПРИМАНКУ!
	// ==========================================================================
	fmt.Println("\n--- Активация Внешнего Сканера (Модуль-Приманка) ---")
	testIP := "127.0.0.1" // Имитируем сканирование локального тестового сервера

	// Вызываем функцию сканирования софта из файла scanner.go
	report := UniversalIdentifySoftware(testIP)

	// Вызываем функцию печати брендированного отчета с логотипом и QR-кодом из файла scanner.go
	PrintRecommendationReport(report)
}
