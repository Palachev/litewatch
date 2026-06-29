package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// SendTelegram отправляет сообщение в Telegram через Bot API (без внешних зависимостей)
func SendTelegram(cfg *Config, text string) {
	if cfg.TelegramToken == "" || cfg.TelegramChatID == "" {
		log.Println("⚠️ Telegram не настроен (нет токена или chat_id)")
		return
	}

	api := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", cfg.TelegramToken)
	data := url.Values{}
	data.Set("chat_id", cfg.TelegramChatID)
	data.Set("text", text)

	resp, err := http.PostForm(api, data)
	if err != nil {
		log.Printf("Ошибка отправки в Telegram: %v", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Telegram вернул статус: %s", strings.TrimSpace(resp.Status))
	}
}
