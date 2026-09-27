///// Делали через PostgreSQL, там на паре Вы говорили про Docker, но он у нас не скачивается и прошлую вашу штучку делали тоже через него же, всё работало(как мы поняли, файл docker-compose.yml получается не нужным, но оставил, schema +)/// 

1. Создать базу `app`:

       psql -U postgres -c "CREATE DATABASE app;"

2. Из папки проекта применить схему и тестовые данные:

       psql -U postgres -d app -f schema.sql
       psql -U postgres -d app -f seed.sql

3. В `main.go` в переменной `dsn` указать свой пароль PostgreSQL.

4. Запустить:

       go run .

5. Проверить:

       curl localhost:8080/health
       # {"status":"ok"}

Схема БД

Пять таблиц. lessons — основная, ссылается на четыре справочника.

groups — учебные группы
- id — BIGSERIAL, PRIMARY KEY
- name — TEXT, NOT NULL, UNIQUE
- year — INT, NOT NULL

teachers — преподаватели
- id — BIGSERIAL, PRIMARY KEY
- full_name — TEXT, NOT NULL
- email — TEXT, NOT NULL, UNIQUE
- degree — TEXT, NOT NULL, DEFAULT ''

subjects — предметы
- id — BIGSERIAL, PRIMARY KEY
- title — TEXT, NOT NULL, UNIQUE
- hours — INT, NOT NULL, DEFAULT 0

rooms — аудитории
- id — BIGSERIAL, PRIMARY KEY
- number — TEXT, NOT NULL, UNIQUE
- capacity — INT, NOT NULL, DEFAULT 0

lessons — занятия
- id — BIGSERIAL, PRIMARY KEY
- group_id — BIGINT, NOT NULL, FK → groups(id) ON DELETE CASCADE
- teacher_id — BIGINT, NOT NULL, FK → teachers(id) ON DELETE CASCADE
- subject_id — BIGINT, NOT NULL, FK → subjects(id) ON DELETE CASCADE
- room_id — BIGINT, NOT NULL, FK → rooms(id) ON DELETE CASCADE
- starts_at — TIMESTAMPTZ, NOT NULL
- created_at — TIMESTAMPTZ, NOT NULL, DEFAULT now()

Связи: у одной группы, преподавателя, предмета и аудитории — много занятий.
При удалении справочной записи её занятия удаляются (ON DELETE CASCADE).

Эндпоинты

Для каждой сущности — CRUD: список, один по id, создание, удаление.

GET /health — 200 {"status":"ok"}

Groups
- GET /groups — 200, массив
- GET /groups/{id} — 200, объект / 400, 404, 500
- POST /groups — тело {"name":"ПИ-21","year":2021} — 201, объект / 400, 422, 500
- DELETE /groups/{id} — 204 / 400, 404, 500

Teachers
- GET /teachers — 200, массив
- GET /teachers/{id} — 200, объект / 400, 404, 500
- POST /teachers — тело {"full_name":"Иванов И.И.","email":"i@e.com","degree":"к.т.н."} — 201, объект / 400, 422, 500
- DELETE /teachers/{id} — 204 / 400, 404, 500

Subjects
- GET /subjects — 200, массив
- GET /subjects/{id} — 200, объект / 400, 404, 500
- POST /subjects — тело {"title":"Матан","hours":144} — 201, объект / 400, 422, 500
- DELETE /subjects/{id} — 204 / 400, 404, 500

Rooms
- GET /rooms — 200, массив
- GET /rooms/{id} — 200, объект / 400, 404, 500
- POST /rooms — тело {"number":"А-301","capacity":30} — 201, объект / 400, 422, 500
- DELETE /rooms/{id} — 204 / 400, 404, 500

Lessons
- GET /lessons — 200, массив
- GET /lessons/{id} — 200, объект / 400, 404, 500
- POST /lessons — тело {"group_id":1,"teacher_id":1,"subject_id":1,"room_id":1,"starts_at":"2026-09-28T09:00:00+05:00"} — 201, объект / 400, 422, 500
- DELETE /lessons/{id} — 204 / 400, 404, 500

Коды ответов

200 — успешный GET
201 — успешный POST
204 — успешный DELETE
400 — id не число или тело не JSON
404 — запись не найдена
422 — пустое обязательное поле
500 — ошибка базы