# Bybit Balance Checker

## Описание

Проект Bybit Balance Checker предназначен для получения данных о балансе на счетах Bybit с использованием их API. Этот инструмент позволяет проверять балансы нескольких аккаунтов и записывать результаты в файл.

## Использование

### Шаги для запуска

1. **Установка зависимостей**: перед началом работы убедитесь, что у вас установлен Go (версия >= 1.13).

2. **Конфигурация**: измените или создайте файл `config.txt` в формате JSON, где указываете `apiKey`, `apiSecret` и `account` для каждого аккаунта Bybit.

   Пример `config.txt`:
   ```json
   [
     {"apiKey": "YOUR_API_KEY", "apiSecret": "YOUR_SECRET_KEY", "account": "example1@example.com"},
     {"apiKey": "YOUR_API_KEY", "apiSecret": "YOUR_SECRET_KEY", "account": "example2@example.com"}
   ]
   
3. **Запуск**: выполните следующие команды для сборки и запуска проекта:
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

### Инструкции по сборке для разных ОС

#### Для Windows:

Для сборки под Windows:

```bash
GOOS=windows GOARCH=amd64 go build -o bybit-balance-checker.exe main.go
```

### Для Linux:

Для сборки под Linux:
```bash
GOOS=linux GOARCH=amd64 go build -o bybit-balance-checker main.go
```

### Для macOS:

Для сборки под macOS:
```bash
GOOS=darwin GOARCH=amd64 go build -o bybit-balance-checker main.go
```
