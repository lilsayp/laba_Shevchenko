CREATE TABLE IF NOT EXISTS groups (
    id   BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    year INT  NOT NULL
);

CREATE TABLE IF NOT EXISTS teachers (
    id        BIGSERIAL PRIMARY KEY,
    full_name TEXT NOT NULL,
    email     TEXT NOT NULL UNIQUE,
    degree    TEXT NOT NULL DEFAULT ''
);

CREATE TABLE IF NOT EXISTS subjects (
    id    BIGSERIAL PRIMARY KEY,
    title TEXT NOT NULL UNIQUE,
    hours INT  NOT NULL DEFAULT 0
);

CREATE TABLE IF NOT EXISTS rooms (
    id       BIGSERIAL PRIMARY KEY,
    number   TEXT NOT NULL UNIQUE,
    capacity INT  NOT NULL DEFAULT 0
);

CREATE TABLE IF NOT EXISTS lessons (
    id         BIGSERIAL PRIMARY KEY,
    group_id   BIGINT NOT NULL REFERENCES groups(id)   ON DELETE CASCADE,
    teacher_id BIGINT NOT NULL REFERENCES teachers(id) ON DELETE CASCADE,
    subject_id BIGINT NOT NULL REFERENCES subjects(id) ON DELETE CASCADE,
    room_id    BIGINT NOT NULL REFERENCES rooms(id)    ON DELETE CASCADE,
    starts_at  TIMESTAMPTZ NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX IF NOT EXISTS idx_lessons_group_id ON lessons(group_id);