# Bybit Balance Checker

## Описание

Проект Bybit Balance Checker предназначен для получения инфомрации о кол-ве coin(ов) на бирже Bybit с использованием API. Он включает два модуля для получения информации с счетов: Funding и Unified. Для отправки запросов проект использует прокси.

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

3. Установка зависимостей:  вас должен быть установлен Go. Выполните команду для установки зависимостей:
   ```bash
   go mod tidy
   ```
4. **Конфигурация**: создайте или откройте файл `config.txt` в той же директории где проект, где указываете `apiKey`, `apiSecret`, `account` и `proxy` для каждого аккаунта Bybit.

   Пример `config.txt`:
   ```makefile
   apiKey=YOUR_API_KEY;apiSecret=YOUR_SECRET_KEY;account=example1@example.com;proxy=IP:PORT:LOGIN:PASSWORD
   apiKey=YOUR_API_KEY;apiSecret=YOUR_SECRET_KEY;account=example2@example.com;proxy=IP:PORT:LOGIN:PASSWORD
   и т.д.
   ```
5. **Сборка**: выполните следующие команды для сборки и запуска проекта:
   ```bash
   go build -o bybit-balance-checker main.go 
   ./bybit-balance-checker
   ```
   
#### Примечания
   - Убедитесь, что установили golang и зависимости для проекта;
   - Убедитесь, что у вас есть аккаунт на Bybit и весь доступ к API;
   - Проверьте, что ваш прокси сервер доступен и настроен правильно в конфигурационном файле;
   - Результаты будут сохранены в созданный файл output.txt в директории с скриптом.

