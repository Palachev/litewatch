package main

import (
	"log"
	"time"
)

func main() {
	cfg, err := LoadConfig()
	if err != nil {
		log.Fatalf("Ошибка загрузки конфига: %v", err)
	}

	log.Println("✅ LiteWatch запущен. Слежу за сервером...")
	SendTelegram(cfg, "🟢 LiteWatch запущен и следит за вашим сервером.")

	ticker := time.NewTicker(time.Duration(cfg.IntervalSeconds) * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		checkAll(cfg)
	}
}

// checkAll собирает метрики и шлёт человекопонятные алерты при превышении порогов
func checkAll(cfg *Config) {
	// CPU
	if cpu, err := GetCPU(); err == nil {
		if msg := HumanizeCPU(cpu, cfg.Thresholds.CPU); msg != "" {
			SendTelegram(cfg, msg)
		}
	}
	// RAM
	if mem, err := GetMemory(); err == nil {
		if msg := HumanizeMemory(mem, cfg.Thresholds.Memory); msg != "" {
			SendTelegram(cfg, msg)
		}
	}
	// Диск
	if disk, err := GetDisk("/"); err == nil {
		if msg := HumanizeDisk(disk, cfg.Thresholds.Disk); msg != "" {
			SendTelegram(cfg, msg)
		}
	}
}
