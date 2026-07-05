<div align="center">

# 🪶 LiteWatch

### Лёгкий и быстрый агент мониторинга серверов с понятными уведомлениями в Telegram

*В мире тяжёлых и сложных систем мониторинга — LiteWatch делает наоборот.*

[![Go](https://img.shields.io/badge/Go-1.22-00ADD8?logo=go&logoColor=white)](https://go.dev)
[![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)
[![Platform](https://img.shields.io/badge/platform-Linux-333?logo=linux&logoColor=white)](#)
[![Docker](https://img.shields.io/badge/Docker-ready-2496ED?logo=docker&logoColor=white)](#-вариант-2--docker)
[![Telegram](https://img.shields.io/badge/Telegram-alerts-26A5E4?logo=telegram&logoColor=white)](#)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](#-вклад)
[![Stars](https://img.shields.io/github/stars/ТЫ/litewatch?style=social)](https://github.com/ТЫ/litewatch)

</div>

---

## 💡 Зачем это

Большие системы мониторинга (Prometheus, Zabbix, Netdata) — мощные, но **тяжёлые и сложные**:
дни настройки, конфиги, экспортеры, графики, в которых тонешь.

Если у тебя 1–2 VPS, домашний сервер или маленький проект — тебе не нужен дашборд с 500 графиками.
Тебе нужно одно: **чтобы по-человечески сказали, когда что-то не так.**

**LiteWatch — это про простоту:** один маленький бинарник, ноль конфигов, понятный язык.

---

## ✨ Чем отличается

| | LiteWatch | Тяжёлые системы |
|---|:---:|:---:|
| 🪶 Вес | один бинарник | сервер + БД + экспортеры |
| ⚡ Установка | 30 секунд | часы / дни |
| 🗣 Алерты | человеческим языком | `CPU 94%` |
| 🇷🇺 Русский | да | редко |
| 🧠 Подсказки «что делать» | да | нет |

---

## 📲 Пример уведомления

Вместо сухого `disk usage: 96%` ты получаешь:

> 🔴 **Диск почти полон (96% занято).**
> Скоро сервис может остановиться. Почисти логи: `du -sh /var/log/* | sort -h`.

> ⚠️ **Процессор загружен на 94% уже несколько минут — сервер может тормозить.**
> Что проверить: тяжёлые процессы (`top` или `htop`).

```
  ┌─────────────────────────────────────────┐
  │  🟢 LiteWatch запущен и следит за сервером │
  │                                           │
  │  🔴 Диск почти полон (96% занято).        │
  │     Почисти логи: du -sh /var/log/*       │
  └─────────────────────────────────────────┘
```


## 🚀 Установка

### Вариант 1 — systemd (рекомендуется, легче всего)
```bash
git clone https://github.com/ТЫ/litewatch && cd litewatch/agent
go mod tidy && go build -o ../litewatch . && cd ..
# отредактируй config.yaml — впиши токен бота и chat_id
sudo sh deploy/systemd/install.sh
```

### Вариант 2 — Docker
```bash
export TELEGRAM_TOKEN=xxx TELEGRAM_CHAT_ID=yyy
docker compose up -d
```

---

## ⚙️ Настройка

`config.yaml`:
```yaml
telegram_token: "токен_от_BotFather"
telegram_chat_id: "твой_chat_id"
interval_seconds: 30
thresholds:
  cpu: 90      # тревога если CPU выше 90%
  memory: 90   # тревога если RAM выше 90%
  disk: 90     # тревога если диск выше 90%
```

**Где взять данные:**
- 🤖 Токен бота — у [@BotFather](https://t.me/BotFather)
- 🆔 Chat ID — у [@userinfobot](https://t.me/userinfobot)

---

## 📊 Что мониторит (MVP)

- ✅ CPU — загрузка
- ✅ RAM — использование
- ✅ Диск — заполнение
- ✅ Heartbeat — «сервер запущен»

## 🗺 Дорожная карта

- [ ] Мониторинг сервисов (порты, docker-контейнеры)
- [ ] Проверка SSL-сертификатов и доменов
- [ ] Несколько серверов в одном боте
- [ ] Веб-панель
- [ ] Облачная версия (hosted)

---

## 🤝 Вклад

PR и идеи приветствуются! Если LiteWatch оказался полезным — **поставь ⭐, это лучшая поддержка проекта.**

## 📄 Лицензия

[MIT](LICENSE) — используй свободно.

---

<div align="center">
<sub>Сделано с 🪶 для тех, кто не хочет разбираться в сложном мониторинге.</sub>
</div>
