#!/bin/sh
set -e
echo "📦 Установка LiteWatch (systemd)..."

# 1. Скопировать бинарник (предполагается, что собран рядом: ./litewatch)
sudo cp ./litewatch /usr/local/bin/litewatch
sudo chmod +x /usr/local/bin/litewatch

# 2. Конфиг
sudo mkdir -p /etc/litewatch
if [ ! -f /etc/litewatch/config.yaml ]; then
  sudo cp ./config.yaml /etc/litewatch/config.yaml
  echo "✏️  Отредактируй /etc/litewatch/config.yaml (токен бота + chat_id)"
fi

# 3. systemd unit
sudo cp ./deploy/systemd/litewatch.service /etc/systemd/system/litewatch.service

# 4. Запуск + автозагрузка
sudo systemctl daemon-reload
sudo systemctl enable --now litewatch

echo "✅ LiteWatch установлен. Логи: journalctl -u litewatch -f"
