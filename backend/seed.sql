INSERT INTO groups (name, year) VALUES
    ('ПИ-21', 2021), ('ПИ-22', 2022)
ON CONFLICT (name) DO NOTHING;

INSERT INTO teachers (full_name, email, degree) VALUES
    ('Иванов Иван Иванович',   'ivanov@example.com',  'к.т.н.'),
    ('Петрова Анна Сергеевна', 'petrova@example.com', '')
ON CONFLICT (email) DO NOTHING;

INSERT INTO subjects (title, hours) VALUES
    ('Математический анализ', 144),
    ('Программирование на Go', 108)
ON CONFLICT (title) DO NOTHING;

INSERT INTO rooms (number, capacity) VALUES
    ('А-301', 30), ('Б-105', 25)
ON CONFLICT (number) DO NOTHING;

INSERT INTO lessons (group_id, teacher_id, subject_id, room_id, starts_at) VALUES
    (1, 1, 1, 1, '2026-09-28 09:00:00+05'),
    (1, 2, 2, 2, '2026-09-28 10:45:00+05'),
    (2, 1, 1, 1, '2026-09-29 12:30:00+05');