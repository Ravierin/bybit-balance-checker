# Bybit Balance Checker

## Описание

Проект Bybit Balance Checker предназначен для получения данных о балансе на счетах Bybit с использованием их API. Этот инструмент позволяет проверять балансы нескольких аккаунтов и записывать результаты в файл.

## Использование

### Шаги для запуска

1. **Клонирование репозитория**: сначала склонируйте репозиторий на ваш локальный компьютер:
   ```bash
   git clone https://github.com/Ravierin/bybit-balance-checker.git
   cd bybit-balance-checker
   ```
2. **Установка Go**: Убедитесь, что у вас установлен Go. Если Go не установлен, следуйте инструкциям ниже для вашей операционной системы.

   ### Linux:

   Установите Go через пакетный менеджер вашей системы. Например, для Ubuntu:
   ```bash
   sudo apt install golang
   ```
   ### Для macOS:

   Установите Go через Homebrew:
   ```bash
   brew install go
   ```
   #### Для Windows:

   Скачайте установочный файл с [официального сайта Go](https://go.dev/dl/) и следуйте инструкциям для установки.
   
3. **Конфигурация**: измените или создайте файл `config.txt` в формате JSON, где указываете `apiKey`, `apiSecret` и `account` для каждого аккаунта Bybit.

   Пример `config.txt`:
   ```json
   [
     {"apiKey": "YOUR_API_KEY", "apiSecret": "YOUR_SECRET_KEY", "account": "example1@example.com"},
     {"apiKey": "YOUR_API_KEY", "apiSecret": "YOUR_SECRET_KEY", "account": "example2@example.com"}
   ]
   
4. **Запуск**: выполните следующие команды для сборки и запуска проекта:
   ```bash
   go build -o bybit-balance-checker main.go 
   ./bybit-balance-checker
   ```
#### Запись результатов

Результаты проверки балансов будут сохранены в файл `output.txt`, который будет создан после выполнения программы.
Каждый аккаунт имеет свою секцию в формате:

Account: "example1@example.com":
   - BTC: 1.2345
   - ETH: 5.6789

