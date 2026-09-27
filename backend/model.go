package main

import "time"

type Group struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Year int    `json:"year"`
}

type Teacher struct {
	ID       int64  `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Degree   string `json:"degree"`
}

type Subject struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
	Hours int    `json:"hours"`
}

type Room struct {
	ID       int64  `json:"id"`
	Number   string `json:"number"`
	Capacity int    `json:"capacity"`
}

type Lesson struct {
	ID        int64     `json:"id"`
	GroupID   int64     `json:"group_id"`
	TeacherID int64     `json:"teacher_id"`
	SubjectID int64     `json:"subject_id"`
	RoomID    int64     `json:"room_id"`
	StartsAt  time.Time `json:"starts_at"`
	CreatedAt time.Time `json:"created_at"`
}
