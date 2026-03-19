# mApps Gateway

API Gateway для мобильного приложения mApps.
Принимает HTTP и gRPC запросы от клиентов и проксирует их в downstream микросервисы: **auth**, **product_service**, **asset_manager**.

---

## Архитектура

```
Mobile Client
      │
      ▼
┌─────────────────────────────────┐
│         mApps Gateway           │
│                                 │
│  HTTP :8000   │   gRPC :9000    │
│               │                 │
│  ┌─────────┐  │  ┌───────────┐  │
│  │  HTTP   │  │  │   gRPC    │  │
│  │ Server  │  │  │  Server   │  │
│  └────┬────┘  │  └─────┬─────┘  │
│       │       │        │        │
│       └───────┴────────┘        │
│               │                 │
│        ┌──────▼──────┐          │
│        │   Service   │          │
│        │    Layer    │          │
│        └──────┬──────┘          │
│               │                 │
│        ┌──────▼──────┐          │
│        │  Data Layer │          │
│        │ gRPC clients│          │
│        └──────┬──────┘          │
└───────────────┼─────────────────┘
                │
    ┌───────────┼───────────┐
    ▼           ▼           ▼
 auth       product     asset_manager
:49780      :49782        :49783
```

**Фреймворк:** [Kratos v2](https://github.com/go-kratos/kratos)
**DI:** [Google Wire](https://github.com/google/wire)
**Кодогенерация контрактов:** [Buf](https://buf.build)

---

## Структура проекта

```
gateway/
├── api/
│   ├── generated/                  # Сгенерированный Go код из proto-контрактов
│   │   └── proto/
│   │       ├── auth/v1/
│   │       ├── products/v1/
│   │       └── asset_manager/v1/
│   └── viability/                  # Health/Readiness proto
├── cmd/mApps_gateway/
│   ├── main.go                     # Точка входа
│   ├── wire.go                     # Wire DI граф
│   └── wire_gen.go                 # Сгенерированный DI код
├── configs/
│   └── config.yaml                 # Конфигурация
├── configs_example/
│   └── configs.yaml.example        # Пример конфигурации
├── internal/
│   ├── conf/
│   │   └── conf.go                 # Конфигурационные структуры
│   ├── data/
│   │   └── data.go                 # gRPC клиенты downstream сервисов
│   ├── server/
│   │   ├── http.go                 # HTTP сервер
│   │   ├── grpc.go                 # gRPC сервер
│   │   └── server.go               # Wire ProviderSet
│   └── service/
│       ├── health_service.go       # Health/Readiness
│       ├── auth_service.go         # Прокси → auth
│       ├── product_service.go      # Прокси → product_service
│       └── asset_manager_service.go# Прокси → asset_manager
├── buf.gen.yaml                    # Конфигурация генерации proto
├── Makefile
└── Dockerfile
```

---

## API Endpoints

### Viability

| Method | Path | Описание |
|--------|------|----------|
| `GET` | `/api/v1/viability/health` | Статус и uptime gateway |
| `GET` | `/api/v1/viability/ready` | Готовность всех downstream сервисов |

### Auth

| Method | Path | Описание |
|--------|------|----------|
| `POST` | `/api/v1/mobile/auth/register` | Регистрация пользователя |
| `POST` | `/api/v1/mobile/auth/login` | Авторизация пользователя |

### Products

| Method | Path | Описание |
|--------|------|----------|
| `GET` | `/api/v1/mobile/products` | Список продуктов |
| `GET` | `/api/v1/mobile/products/{product_id}` | Получить продукт по ID |
| `POST` | `/api/v1/mobile/products` | Создать продукт |
| `PUT` | `/api/v1/mobile/products/{product_id}` | Обновить продукт |

### Asset Manager

| Method | Path | Описание |
|--------|------|----------|
| `POST` | `/api/v1/assets/models/upload-url` | Получить URL для загрузки 3D модели |
| `GET` | `/api/v1/assets/models/{model_id}` | Получить 3D модель по ID |
| `POST` | `/api/v1/assets/textures/upload-url` | Получить URL для загрузки текстуры |
| `GET` | `/api/v1/assets/textures/{asset_id}` | Получить текстуру по ID |

---

## Конфигурация

Конфигурация загружается из `configs/config.yaml`.
Пример — в `configs_example/configs.yaml.example`.

```yaml
server:
  http:
    addr: 0.0.0.0:8000
    timeout: 1s
  grpc:
    addr: 0.0.0.0:9000
    timeout: 1s
data:
  redis:
    addr: 127.0.0.1:6379
    read_timeout: 0.2s
    write_timeout: 0.2s
  auth:
    addr: 127.0.0.1:49780
    timeout: 0.2s
  product:
    addr: 127.0.0.1:49782
    timeout: 0.2s
  asset_manager:
    addr: 127.0.0.1:49783
    timeout: 0.2s
```

Конфиг также можно переопределить через переменные окружения с префиксом `KRATOS_`.

---

## Запуск

### Локально

```bash
# Установить зависимости для кодогенерации
make init

# Сгенерировать proto-контракты из remote репозитория
make generate_contracts   # auth + products + asset_manager

# Пересобрать DI (если менялись зависимости)
make wire

# Запустить
make local
```

### Сборка бинарника

```bash
make build
# бинарник: ./bin/mApps_gateway
./bin/mApps_gateway -conf ./configs
```

### Docker

```bash
docker build -t mapps-gateway .
docker run -p 8000:8000 -p 9000:9000 mapps-gateway
```

---

## Разработка

### Кодогенерация

Контракты хранятся в отдельном репозитории. Генерация через `buf`:

```bash
# Все контракты сразу
make generate_contracts

# Или по отдельности
make auth
make products
make asset_manager
```

Сгенерированный код попадает в `api/generated/proto/`.

### Обновление DI после изменения зависимостей

```bash
make wire
```

### Полный цикл пересборки

```bash
make all   # config + api + generate
```

---

## Зависимости

| Пакет | Версия | Назначение |
|-------|--------|------------|
| `github.com/go-kratos/kratos/v2` | v2.9.2 | HTTP/gRPC фреймворк |
| `github.com/google/wire` | v0.7.0 | Dependency Injection |
| `google.golang.org/grpc` | v1.79.1 | gRPC транспорт |
| `google.golang.org/protobuf` | v1.36.11 | Protobuf runtime |
| `go.uber.org/automaxprocs` | v1.6.0 | Авто-настройка GOMAXPROCS |
