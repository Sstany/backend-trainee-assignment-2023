# Segmenty

Тестовое задание для стажера Backend Avito 2023

## Run the app

По дефолту сервис будет запущен на порту 8090

```bash
docker compose up -d

```

## Stop the app

```bash
docker compose down
```

## Stop the app with DB deletion

Если требуется удалить базу данных

```bash
docker compose down --volumes

```

# REST API

Описание REST API в doc/openapi.yaml

На порту 3000 стартует swagger-ui, в котором описаны все запросы

В файле doc/insomnia_collection.json содержится коллекция запросов, которую можно импортировать в Insomnia.

### Request

```bash
`GET /api/v1/users/1`
```

```bash
curl --request GET \
  --url http://localhost:8090/api/v1/users/1

```

### Response


HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Wed, 30 Aug 2023 23:08:48 GMT
Content-Length: 30

{"userId":1,"username":"Test"}

## Get non-existent user

### Request

```bash
`GET /api/v1/users/9999 `
```

```bash
curl --request GET \
  --url http://localhost:8090/api/v1/users/9999

```

### Response


HTTP/1.1 404 Not Found
Content-Type: text/plain; charset=utf-8
Date: Thu, 31 Aug 2023 00:46:06 GMT
Content-Length: 30

User with id '9999' is not found

### Questions

Могут ли быть одинаковые slug у  сегментов?

- Нет
