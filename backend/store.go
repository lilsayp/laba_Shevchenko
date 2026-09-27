package main

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var ErrNotFound = errors.New("not found")

type ScheduleStore struct {
	db *pgxpool.Pool
}

func NewScheduleStore(db *pgxpool.Pool) *ScheduleStore {
	return &ScheduleStore{db: db}
}

// ---------- GROUPS ----------

func (s *ScheduleStore) ListGroups(ctx context.Context) ([]Group, error) {
	rows, err := s.db.Query(ctx, `SELECT id, name, year FROM groups ORDER BY id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := []Group{}
	for rows.Next() {
		var g Group
		if err := rows.Scan(&g.ID, &g.Name, &g.Year); err != nil {
			return nil, err
		}
		items = append(items, g)
	}
	return items, rows.Err()
}

func (s *ScheduleStore) GetGroup(ctx context.Context, id int64) (Group, error) {
	var g Group
	err := s.db.QueryRow(ctx,
		`SELECT id, name, year FROM groups WHERE id = $1`, id).
		Scan(&g.ID, &g.Name, &g.Year)
	if errors.Is(err, pgx.ErrNoRows) {
		return Group{}, ErrNotFound
	}
	if err != nil {
		return Group{}, err
	}
	return g, nil
}

func (s *ScheduleStore) CreateGroup(ctx context.Context, name string, year int) (Group, error) {
	var g Group
	err := s.db.QueryRow(ctx,
		`INSERT INTO groups (name, year) VALUES ($1, $2)
		 RETURNING id, name, year`, name, year).
		Scan(&g.ID, &g.Name, &g.Year)
	if err != nil {
		return Group{}, err
	}
	return g, nil
}

func (s *ScheduleStore) DeleteGroup(ctx context.Context, id int64) error {
	tag, err := s.db.Exec(ctx, `DELETE FROM groups WHERE id = $1`, id)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return ErrNotFound
	}
	return nil
}

// ---------- TEACHERS ----------

func (s *ScheduleStore) ListTeachers(ctx context.Context) ([]Teacher, error) {
	rows, err := s.db.Query(ctx,
		`SELECT id, full_name, email, degree FROM teachers ORDER BY id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := []Teacher{}
	for rows.Next() {
		var t Teacher
		if err := rows.Scan(&t.ID, &t.FullName, &t.Email, &t.Degree); err != nil {
			return nil, err
		}
		items = append(items, t)
	}
	return items, rows.Err()
}

func (s *ScheduleStore) GetTeacher(ctx context.Context, id int64) (Teacher, error) {
	var t Teacher
	err := s.db.QueryRow(ctx,
		`SELECT id, full_name, email, degree FROM teachers WHERE id = $1`, id).
		Scan(&t.ID, &t.FullName, &t.Email, &t.Degree)
	if errors.Is(err, pgx.ErrNoRows) {
		return Teacher{}, ErrNotFound
	}
	if err != nil {
		return Teacher{}, err
	}
	return t, nil
}

func (s *ScheduleStore) CreateTeacher(ctx context.Context, fullName, email, degree string) (Teacher, error) {
	var t Teacher
	err := s.db.QueryRow(ctx,
		`INSERT INTO teachers (full_name, email, degree) VALUES ($1, $2, $3)
		 RETURNING id, full_name, email, degree`, fullName, email, degree).
		Scan(&t.ID, &t.FullName, &t.Email, &t.Degree)
	if err != nil {
		return Teacher{}, err
	}
	return t, nil
}

func (s *ScheduleStore) DeleteTeacher(ctx context.Context, id int64) error {
	tag, err := s.db.Exec(ctx, `DELETE FROM teachers WHERE id = $1`, id)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return ErrNotFound
	}
	return nil
}

// ---------- SUBJECTS ----------

func (s *ScheduleStore) ListSubjects(ctx context.Context) ([]Subject, error) {
	rows, err := s.db.Query(ctx, `SELECT id, title, hours FROM subjects ORDER BY id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := []Subject{}
	for rows.Next() {
		var sub Subject
		if err := rows.Scan(&sub.ID, &sub.Title, &sub.Hours); err != nil {
			return nil, err
		}
		items = append(items, sub)
	}
	return items, rows.Err()
}

func (s *ScheduleStore) GetSubject(ctx context.Context, id int64) (Subject, error) {
	var sub Subject
	err := s.db.QueryRow(ctx,
		`SELECT id, title, hours FROM subjects WHERE id = $1`, id).
		Scan(&sub.ID, &sub.Title, &sub.Hours)
	if errors.Is(err, pgx.ErrNoRows) {
		return Subject{}, ErrNotFound
	}
	if err != nil {
		return Subject{}, err
	}
	return sub, nil
}

func (s *ScheduleStore) CreateSubject(ctx context.Context, title string, hours int) (Subject, error) {
	var sub Subject
	err := s.db.QueryRow(ctx,
		`INSERT INTO subjects (title, hours) VALUES ($1, $2)
		 RETURNING id, title, hours`, title, hours).
		Scan(&sub.ID, &sub.Title, &sub.Hours)
	if err != nil {
		return Subject{}, err
	}
	return sub, nil
}

func (s *ScheduleStore) DeleteSubject(ctx context.Context, id int64) error {
	tag, err := s.db.Exec(ctx, `DELETE FROM subjects WHERE id = $1`, id)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return ErrNotFound
	}
	return nil
}

// ---------- ROOMS ----------

func (s *ScheduleStore) ListRooms(ctx context.Context) ([]Room, error) {
	rows, err := s.db.Query(ctx, `SELECT id, number, capacity FROM rooms ORDER BY id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := []Room{}
	for rows.Next() {
		var r Room
		if err := rows.Scan(&r.ID, &r.Number, &r.Capacity); err != nil {
			return nil, err
		}
		items = append(items, r)
	}
	return items, rows.Err()
}

func (s *ScheduleStore) GetRoom(ctx context.Context, id int64) (Room, error) {
	var r Room
	err := s.db.QueryRow(ctx,
		`SELECT id, number, capacity FROM rooms WHERE id = $1`, id).
		Scan(&r.ID, &r.Number, &r.Capacity)
	if errors.Is(err, pgx.ErrNoRows) {
		return Room{}, ErrNotFound
	}
	if err != nil {
		return Room{}, err
	}
	return r, nil
}

func (s *ScheduleStore) CreateRoom(ctx context.Context, number string, capacity int) (Room, error) {
	var r Room
	err := s.db.QueryRow(ctx,
		`INSERT INTO rooms (number, capacity) VALUES ($1, $2)
		 RETURNING id, number, capacity`, number, capacity).
		Scan(&r.ID, &r.Number, &r.Capacity)
	if err != nil {
		return Room{}, err
	}
	return r, nil
}

func (s *ScheduleStore) DeleteRoom(ctx context.Context, id int64) error {
	tag, err := s.db.Exec(ctx, `DELETE FROM rooms WHERE id = $1`, id)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return ErrNotFound
	}
	return nil
}

// ---------- LESSONS ----------

func (s *ScheduleStore) ListLessons(ctx context.Context) ([]Lesson, error) {
	rows, err := s.db.Query(ctx,
		`SELECT id, group_id, teacher_id, subject_id, room_id, starts_at, created_at
		 FROM lessons ORDER BY starts_at`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := []Lesson{}
	for rows.Next() {
		var l Lesson
		if err := rows.Scan(&l.ID, &l.GroupID, &l.TeacherID, &l.SubjectID,
			&l.RoomID, &l.StartsAt, &l.CreatedAt); err != nil {
			return nil, err
		}
		items = append(items, l)
	}
	return items, rows.Err()
}

func (s *ScheduleStore) GetLesson(ctx context.Context, id int64) (Lesson, error) {
	var l Lesson
	err := s.db.QueryRow(ctx,
		`SELECT id, group_id, teacher_id, subject_id, room_id, starts_at, created_at
		 FROM lessons WHERE id = $1`, id).
		Scan(&l.ID, &l.GroupID, &l.TeacherID, &l.SubjectID,
			&l.RoomID, &l.StartsAt, &l.CreatedAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return Lesson{}, ErrNotFound
	}
	if err != nil {
		return Lesson{}, err
	}
	return l, nil
}

func (s *ScheduleStore) CreateLesson(ctx context.Context,
	groupID, teacherID, subjectID, roomID int64, startsAt string) (Lesson, error) {

	var l Lesson
	err := s.db.QueryRow(ctx,
		`INSERT INTO lessons (group_id, teacher_id, subject_id, room_id, starts_at)
		 VALUES ($1, $2, $3, $4, $5::timestamptz)
		 RETURNING id, group_id, teacher_id, subject_id, room_id, starts_at, created_at`,
		groupID, teacherID, subjectID, roomID, startsAt).
		Scan(&l.ID, &l.GroupID, &l.TeacherID, &l.SubjectID,
			&l.RoomID, &l.StartsAt, &l.CreatedAt)
	if err != nil {
		return Lesson{}, err
	}
	return l, nil
}

func (s *ScheduleStore) DeleteLesson(ctx context.Context, id int64) error {
	tag, err := s.db.Exec(ctx, `DELETE FROM lessons WHERE id = $1`, id)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return ErrNotFound
	}
	return nil
}
