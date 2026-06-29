package main

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Thresholds struct {
	CPU    float64 `yaml:"cpu"`    // % загрузки CPU, выше которого тревога
	Memory float64 `yaml:"memory"` // % использования RAM
	Disk   float64 `yaml:"disk"`   // % заполнения диска
}

type Config struct {
	TelegramToken   string     `yaml:"telegram_token"`
	TelegramChatID  string     `yaml:"telegram_chat_id"`
	IntervalSeconds int        `yaml:"interval_seconds"`
	Thresholds      Thresholds `yaml:"thresholds"`
}

func LoadConfig() (*Config, error) {
	path := os.Getenv("LITEWATCH_CONFIG")
	if path == "" {
		path = "config.yaml"
	}

	cfg := &Config{
		IntervalSeconds: 30,
		Thresholds:      Thresholds{CPU: 90, Memory: 90, Disk: 90},
	}

	data, err := os.ReadFile(path)
	if err == nil {
		if err := yaml.Unmarshal(data, cfg); err != nil {
			return nil, err
		}
	}

	// ENV переопределяет конфиг (удобно для Docker)
	if v := os.Getenv("TELEGRAM_TOKEN"); v != "" {
		cfg.TelegramToken = v
	}
	if v := os.Getenv("TELEGRAM_CHAT_ID"); v != "" {
		cfg.TelegramChatID = v
	}

	return cfg, nil
}
