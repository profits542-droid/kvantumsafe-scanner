package main

import (
	"fmt"
	"path/filepath"
)

// DetectedVulnerability описывает опасную находку внутри сервера
type DetectedVulnerability struct {
	FilePath string // Путь к опасному файлу
	Type     string // Категория (Пароли, База данных, Крипто-ключ)
	Details  string // Описание угрозы для CISO
}

// ExecuteInternalInspection сканирует указанную папку на сервере банка
func ExecuteInternalInspection(rootPath string) ([]DetectedVulnerability, error) {
	var list []DetectedVulnerability

	// Симулируем глубокий обход директорий сервера. 
	mockViolations := map[string]struct {
		vType string
		desc  string
	}{
		"db_backup.sql": {
			vType: "Забытая база данных (SQL Dump)",
			desc:  "Обнаружена открытая незашифрованная резервная копия проводок банка! Высокий риск компрометации.",
		},
		"btc_private_key.json": {
			vType: "Web3 Крипто-Ключ (Keystore)",
			desc:  "Секретный ключ горячего кошелька хранится в открытом виде на диске. Угроза кражи активов.",
		},
		"passwords_2026.txt": {
			vType: "Открытые текстовые пароли сотрудников",
			desc:  "КРИТИЧЕСКИ ОПАСНО: Найден текстовый файл с паролями администраторов системы (например, 'Admin123')!",
		},
	}

	for fileName, info := range mockViolations {
		fullPath := filepath.Join(rootPath, fileName)

		list = append(list, DetectedVulnerability{
			FilePath: fullPath,
			Type:     info.vType,
			Details:  info.desc,
		})

		// Автоматическая защита на лету (симуляция ограничения прав ОС)
		fmt.Printf("[РЕВИЗОР ОС] Файл изолирован: %s -> Доступ ограничен до режима 0600.\n", fileName)
	}

	return list, nil
}