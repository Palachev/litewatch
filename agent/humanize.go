package main

import "fmt"

// HumanizeCPU превращает загрузку CPU в понятное сообщение.
// Возвращает "" если всё в норме.
func HumanizeCPU(value, threshold float64) string {
	if value < threshold {
		return ""
	}
	return fmt.Sprintf(
		"⚠️ Процессор загружен на %.0f%% — сервер может тормозить.\n"+
			"Что проверить: тяжёлые процессы (команда: top или htop).",
		value,
	)
}

// HumanizeMemory превращает использование RAM в понятное сообщение.
func HumanizeMemory(value, threshold float64) string {
	if value < threshold {
		return ""
	}
	return fmt.Sprintf(
		"🟡 Память почти закончилась (%.0f%% занято).\n"+
			"Возможна утечка в одном из процессов. Проверь: free -h, ps aux --sort=-%%mem | head.",
		value,
	)
}

// HumanizeDisk превращает заполнение диска в понятное сообщение.
func HumanizeDisk(value, threshold float64) string {
	if value < threshold {
		return ""
	}
	return fmt.Sprintf(
		"🔴 Диск почти полон (%.0f%% занято).\n"+
			"Скоро сервис может остановиться. Почисти логи: du -sh /var/log/* | sort -h.",
		value,
	)
}
