# Bybit Balance Checker

## Описание

Проект Bybit Balance Checker предназначен для получения инфомрации о кол-ве coin(ов) на бирже Bybit с использованием API. Он включает два модуля для получения информации с счетов: Funding и Unified. Для отправки запросов проект использует протокол SOCKS5 для проксирования.

## Использование

### Шаги для запуска

1. **Клонирование репозитория**: сначала склонируйте репозиторий на ваш локальный компьютер:
   ```bash
   git clone https://github.com/Ravierin/bybit-balance-checker.git
   cd bybit-balance-checker
   ```
2. **Установка зависимостей**:  вас должен быть установлен Golang версии 1.22.4. Выполните команду для установки зависимостей:
   ```bash
   go mod tidy
   ```
3. **Конфигурация**: создайте файл с названием `config.txt`(важно) в той же директории где проект и укажите `apiKey`, `apiSecret`, `account` и `proxy` для каждого аккаунта Bybit в том формате что указан ниже(важно).

   Пример `config.txt`:
   ```makefile
   apiKey=YOUR_API_KEY;apiSecret=YOUR_SECRET_KEY;account=example1@example.com;proxy=IP:PORT:LOGIN:PASSWORD
   apiKey=YOUR_API_KEY;apiSecret=YOUR_SECRET_KEY;account=example2@example.com;proxy=IP:PORT:LOGIN:PASSWORD
   и т.д.
   ```
4. **Сборка**: выполните следующие команды для сборки в исполняемый файл:

   #### Для Windows:
   ```bash
   go build -o bybit-balance-checker.exe main.go
   ```
   #### Для macOS:
   ```bash
   go build -o bybit-balance-checker main.go
   ```
   #### Для Linux:
   ```bash
   go build -o bybit-balance-checker main.go 
   ```
   После этого для запуска будет достаточно исользовать исполняемый файл и config.txt, поэтому для удобства можно будет перенести это в отдельную директорию.
   
#### Примечания
   - Убедитесь, что установили golang, git и зависимости для проекта;
   - Убедитесь, что у вас есть аккаунт на Bybit и весь доступ к API;
   - Проверьте, что ваш прокси сервер доступен и настроен правильно в конфигурационном файле;
   - Результаты будут сохранены в созданный файл output.txt в директории с скриптом.
   - Файл output.txt не обязательно удалять, при каждом запуске скрипт будет перезаписывать его.
   - Монеты отображаются только больше одной и также на акаунтах, где на Funding и Unified нету монет больше одной то в output.txt, отображаться не будет

