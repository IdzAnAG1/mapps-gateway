# Gateway Routes

**BASE URL:** `http://130.49.148.135`

---

## Auth

| Method | Path | Body |
|--------|------|------|
| `POST` | `/api/mobile/v1/auth/register` | `email`, `password`, `username`, `nickname` |
| `POST` | `/api/mobile/v1/auth/login` | `email`, `password` |

### Register
```json
{
  "email": "user@example.com",
  "password": "secret123",
  "username": "john_doe",
  "nickname": "John"
}
```
Response:
```json
{
  "user_id": "uuid",
  "access_token": "jwt_token"
}
```

### Login
```json
{
  "email": "user@example.com",
  "password": "secret123"
}
```
Response:
```json
{
  "user_id": "uuid",
  "access_token": "jwt_token"
}
```

---

## Products

| Method | Path | Params / Body |
|--------|------|---------------|
| `GET` | `/api/v1/mobile/products` | `?category`, `?name`, `?price`, `?page`, `?page_size` |
| `GET` | `/api/v1/mobile/products/{product_id}` | — |
| `POST` | `/api/v1/mobile/products` | `name`, `description`, `price`, `category`, `model_id`, `virtual_image_id` |
| `PUT` | `/api/v1/mobile/products/{product_id}` | `name`, `description`, `price`, `category`, `model_id`, `virtual_image_id` |

### List Products
```
GET /api/v1/mobile/products?category=sofas&page=1&page_size=10
```

### Get Product
```
GET /api/v1/mobile/products/{product_id}
```

### Create Product
```json
{
  "name": "Диван угловой",
  "description": "Мягкий угловой диван",
  "price": 49999.99,
  "category": "sofas",
  "model_id": "edb83f76-9e72-48c9-91b0-3644c34550a6",
  "virtual_image_id": ""
}
```

### Update Product
```json
{
  "name": "Диван обновлённый",
  "description": "Обновлённое описание",
  "price": 44999.99,
  "category": "sofas",
  "model_id": "edb83f76-9e72-48c9-91b0-3644c34550a6",
  "virtual_image_id": ""
}
```

---

## Assets

| Method | Path | Body |
|--------|------|------|
| `POST` | `/api/v1/assets/models/upload-url` | `name`, `format`, `mime_type` |
| `GET` | `/api/v1/assets/models/{model_id}` | — |
| `POST` | `/api/v1/assets/textures/upload-url` | `name`, `mime_type` |
| `GET` | `/api/v1/assets/textures/{asset_id}` | — |

### Get Model Upload URL
```json
{
  "name": "sofa",
  "format": "glb",
  "mime_type": "model/gltf-binary"
}
```
Response:
```json
{
  "upload_url": "http://...",
  "model_id": "uuid"
}
```

### Upload Model (прямо в S3 по presigned URL)
```
PUT <upload_url>
Content-Type: model/gltf-binary
x-amz-meta-name: sofa
x-amz-meta-format: glb

Body: бинарный .glb файл
```

### Get Model
```
GET /api/v1/assets/models/{model_id}
```
Response:
```json
{
  "model": {
    "id": "uuid",
    "name": "sofa",
    "format": "glb",
    "url": "http://... (presigned, 15 min)"
  }
}
```

### Get Texture Upload URL
```json
{
  "name": "wood_texture",
  "mime_type": "image/png"
}
```
Response:
```json
{
  "upload_url": "http://...",
  "asset_id": "uuid"
}
```

---

## Viability

| Method | Path | Description |
|--------|------|-------------|
| `GET` | `/api/v1/viability/health` | Статус gateway |
| `GET` | `/api/v1/viability/ready` | Статус всех сервисов (auth, product, asset_manager) |

---

## Asset Upload Flow

```
1. POST /api/v1/assets/models/upload-url  →  получить upload_url + model_id
2. PUT <upload_url> с файлом             →  загрузить файл напрямую в S3
3. POST /api/v1/mobile/products          →  создать продукт с model_id
4. GET /api/v1/assets/models/{model_id}  →  получить download url для рендера AR
```
