# OTUS | Микросервисная архитектура | 2025

Автор: Дубровин Е.Н.

Содержит в себе ответы по ДЗ (начиная с ДЗ2) для курса "Микросервисная архитектура" до тех пор, пока это имеет смысл вести в одном репозитории

## ДЗ 2 | Приложение в docker-образ и запушить его на Dockerhub

**Цель:** В этом ДЗ вы сможете обернуть приложение в docker-образ и запушить его на Dockerhub.

**Вариант 1 (С КОДОМ)**

Шаг 1. Создать минимальный сервис, который

```
    отвечает на порту 8000
    имеет http-метод:
        GET /health/
        RESPONSE: {"status": "OK"}
```

Шаг 2. Cобрать локально образ приложения в докер контейнер под архитектуру AMD64.

Запушить образ в dockerhub

### Локальный запуск проекта

```bash
# Запустить локальный Postgres
docker compose up

# Запустить приложение
go run ./src
```

### Сборка и пуш контейнера

```bash
docker build --platform linux/amd64,linux/arm64 . --push -t twobrowin/2025-msa-project:hw4.0.1
```

## ДЗ 3 | Основы работы с Kubernetes

**Цель:** В этом ДЗ вы научитесь создавать минимальный сервис.

**Вариант 1 (С КОДОМ)**

Написать манифесты для деплоя в k8s для этого сервиса.


Манифесты должны описывать сущности: Deployment, Service, Ingress.

В Deployment могут быть указаны Liveness, Readiness пробы.

Количество реплик должно быть не меньше 2. Image контейнера должен быть указан с Dockerhub.


Хост в ингрессе должен быть arch.homework. В итоге после применения манифестов GET запрос на `http://arch.homework/health `должен отдавать `{“status”: “OK”}`.


*Задание со звездой:*

В Ingress-е должно быть правило, которое форвардит все запросы с `/otusapp/{student name}/*` на сервис с rewrite-ом пути. Где `{student name}` - это имя студента.

Например: `curl arch.homework/otusapp/aeugene/health -> рерайт пути на arch.homework/health`

### Развёртывание сервиса в k8s

```bash
kubectl apply -f k8s-manifests
```

## ДЗ 4 | Работа с Helm-ом

**Цель:** В этом ДЗ вы создадите простейший RESTful CRUD.

**Вариант 1 (С КОДОМ)**

Сделать простейший RESTful CRUD по созданию, удалению, просмотру и обновлению пользователей.

Пример API - https://app.swaggerhub.com/apis/otus55/users/1.0.0

Добавить базу данных для приложения.

Конфигурация приложения должна хранится в Configmaps.

Доступы к БД должны храниться в Secrets.

Первоначальные миграции должны быть оформлены в качестве Job-ы, если это требуется.

Ingress-ы должны также вести на url arch.homework/ (как и в прошлом задании)

На выходе должны быть предоставлена
1. ссылка на директорию в github, где находится директория с манифестами кубернетеса (в виде pull request).
2. инструкция по запуску приложения.
    * команда установки БД из helm, вместе с файлом values.yaml.
    * команда применения первоначальных миграций
    * команда kubectl apply -f, которая запускает в правильном порядке манифесты кубернетеса
3. Postman коллекция, в которой будут представлены примеры запросов к сервису на создание, получение, изменение и удаление пользователя. Важно: в postman коллекции использовать базовый url - arch.homework.
4. Проверить корректность работы приложения используя созданную коллекцию `newman run коллекция_постман` и приложить скриншот/вывод исполнения корректной работы

*Задание со звездочкой:*

Добавить шаблонизацию приложения в helm чартах

### Конфигурирование приложения

Для конфигурирования приложения следует задать переменную окружения `CONFIG_PATH` - путь к файлу конфигурации

В случае, если переменная окружения не будет задана будет выполнена попытка получить конфигурацию из файла `.env` (см файл `.env.example` за примерами)

Поддерживаемые форматы конфигурационных файлов:
* YAML (`.yaml`, `.yml`)
* JSON (`.json`)
* TOML (`.toml`)
* EDN (`.edn`)
* ENV (`.env`)

Все параметры файла конфигурации можно переписать при помощи переменных окружения, см. описания полей в файле `src/config/config.go`

### Выполнение миграций приложения

Миграции выполняются отдельным Go-приложением, которое ведётся в директории `migrate`

Приложение имеет cli, помощь можно получить выполнив команду `go run ./migrate db --help`

#### Подготовка файлов миграции

1. Следует создать файл миграции (`migration_name` должно содержать описание вносимых изменений, например, `create_users_table`)

```bash
# Создать файл Go-миграции
go run ./migrate db create_go migration_name

# Создать файл SQL-миграции
go run ./migrate db create_sql migration_name
```

2. Заполнить функции прямой и обратной миграций, например, создание и удаление таблицы

#### Ручные миграции

```bash
# Инициализировать миграции
go run ./migrate db init

# Выполнить миграции
go run ./migrate db migrate

# Откатить миграции
go run ./migrate db rollback
```

#### Docker-образ миграций

```bash
docker build --platform linux/amd64,linux/arm64 . -f Dockerfile.migrate --push -t twobrowin/2025-msa-project:migrate-20260110134623_create_users_table
```

### Развертывание приложения

#### 1. Подготовка

```bash
# Создать namespace
kubectl create namespace otus-2025-msa

# Создать секрет с паролями для БД PostgreSQL
kubectl create secret -n otus-2025-msa generic pg-dev --from-literal='postgres-password=<ADMIN_PASSWORD>' --from-literal='password=<USER_PASSWORD>' --from-literal='replication-password=<REPLICATION_PASSWORD>'
```

#### 2. Развёртывание БД Postgresql

```bash
# Развёртывание БД
helm upgrade --install -n otus-2025-msa pg-dev oci://registry-1.docker.io/bitnamicharts/postgresql -f ./db-charts/pg-values-dev.yaml

# Port-forward для подключения к БД
kubectl port-forward --namespace otus-2025-msa svc/pg-dev-postgresql 5432:5432
```

#### 3. Развёртывание приложения с проведением миграций

```bash
helm upgrade --install -n otus-2025-msa otus-2025-msa-project ./charts -f ./charts/values_dev.yaml
```