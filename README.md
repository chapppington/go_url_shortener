# URL Shortener

Простой сервис для сокращения URL на Go с использованием GORM и PostgreSQL.
TODO добавить Dependency injection, Entity, и завязаться на абстракциях, чтобы добавить тесты

## Требования

- Go 1.25.3+
- PostgreSQL
- Docker и Docker Compose (опционально)

## Установка и запуск

### Локально

1. Установите зависимости:
```bash
go mod download
```

2. Создайте файл `.env`:
```env
API_PORT=8000
POSTGRES_HOST=localhost
POSTGRES_PORT=5432
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres
POSTGRES_DB=url_shortener
```

3. Запустите PostgreSQL и создайте базу данных

4. Запустите приложение:
```bash
make run
# или
go run ./cmd/main.go
```

### Docker

```bash
make all
```

## API

### Создать короткую ссылку

```bash
curl -X POST http://localhost:8000/api/v1/create \
  -H "Content-Type: application/json" \
  -d '{"url": "https://example.com"}'
```

Ответ:
```json
{
  "short_code": "abc12345",
  "long_url": "https://example.com"
}
```

### Получить оригинальный URL

```bash
curl http://localhost:8000/api/v1/abc12345
```

Ответ:
```json
{
  "long_url": "https://example.com"
}
```

### Редирект

```bash
curl -L http://localhost:8000/abc12345
```

Или откройте в браузере: `http://localhost:8000/abc12345`

## Команды

- `make run` - запустить приложение локально
- `make build` - собрать бинарник
- `make all` - запустить все сервисы в Docker
- `make all-down` - остановить все сервисы
- `make app-logs` - посмотреть логи приложения

