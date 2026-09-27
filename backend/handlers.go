package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"
)

type ScheduleHandler struct {
	store *ScheduleStore
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func writeError(w http.ResponseWriter, status int, msg string) {
	writeJSON(w, status, map[string]string{"error": msg})
}

func (h *ScheduleHandler) listGroups(w http.ResponseWriter, r *http.Request) {
	items, err := h.store.ListGroups(r.Context())
	if err != nil {
		writeError(w, http.StatusInternalServerError, "database error")
		return
	}
	writeJSON(w, http.StatusOK, items)
}

func (h *ScheduleHandler) getGroup(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		writeError(w, http.StatusBadRequest, "id must be a number")
		return
	}
	g, err := h.store.GetGroup(r.Context(), id)
	if errors.Is(err, ErrNotFound) {
		writeError(w, http.StatusNotFound, "group not found")
		return
	}
	if err != nil {
		writeError(w, http.StatusInternalServerError, "database error")
		return
	}
	writeJSON(w, http.StatusOK, g)
}

type groupRequest struct {
	Name string `json:"name"`
	Year int    `json:"year"`
}

func (h *ScheduleHandler) createGroup(w http.ResponseWriter, r *http.Request) {
	var req groupRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON")
		return
	}
	req.Name = strings.TrimSpace(req.Name)
	if req.Name == "" {
		writeError(w, http.StatusUnprocessableEntity, "name is required")
		return
	}
	g, err := h.store.CreateGroup(r.Context(), req.Name, req.Year)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "database error")
		return
	}
	writeJSON(w, http.StatusCreated, g)
}

func (h *ScheduleHandler) deleteGroup(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		writeError(w, http.StatusBadRequest, "id must be a number")
		return
	}
	err = h.store.DeleteGroup(r.Context(), id)
	if errors.Is(err, ErrNotFound) {
		writeError(w, http.StatusNotFound, "group not found")
		return
	}
	if err != nil {
		writeError(w, http.StatusInternalServerError, "database error")
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *ScheduleHandler) listTeachers(w http.ResponseWriter, r *http.Request) {
	items, err := h.store.ListTeachers(r.Context())
	if err != nil {
		writeError(w, http.StatusInternalServerError, "database error")
		return
	}
	writeJSON(w, http.StatusOK, items)
}

func (h *ScheduleHandler) getTeacher(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		writeError(w, http.StatusBadRequest, "id must be a number")
		return
	}
	t, err := h.store.GetTeacher(r.Context(), id)
	if errors.Is(err, ErrNotFound) {
		writeError(w, http.StatusNotFound, "teacher not found")
		return
	}
	if err != nil {
		writeError(w, http.StatusInternalServerError, "database error")
		return
	}
	writeJSON(w, http.StatusOK, t)
}

type teacherRequest struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Degree   string `json:"degree"`
}

func (h *ScheduleHandler) createTeacher(w http.ResponseWriter, r *http.Request) {
	var req teacherRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON")
		return
	}
	req.FullName = strings.TrimSpace(req.FullName)
	req.Email = strings.TrimSpace(req.Email)
	if req.FullName == "" || req.Email == "" {
		writeError(w, http.StatusUnprocessableEntity, "full_name and email are required")
		return
	}
	t, err := h.store.CreateTeacher(r.Context(), req.FullName, req.Email, req.Degree)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "database error")
		return
	}
	writeJSON(w, http.StatusCreated, t)
}

func (h *ScheduleHandler) deleteTeacher(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		writeError(w, http.StatusBadRequest, "id must be a number")
		return
	}
	err = h.store.DeleteTeacher(r.Context(), id)
	if errors.Is(err, ErrNotFound) {
		writeError(w, http.StatusNotFound, "teacher not found")
		return
	}
	if err != nil {
		writeError(w, http.StatusInternalServerError, "database error")
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *ScheduleHandler) listSubjects(w http.ResponseWriter, r *http.Request) {
	items, err := h.store.ListSubjects(r.Context())
	if err != nil {
		writeError(w, http.StatusInternalServerError, "database error")
		return
	}
	writeJSON(w, http.StatusOK, items)
}

func (h *ScheduleHandler) getSubject(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		writeError(w, http.StatusBadRequest, "id must be a number")
		return
	}
	sub, err := h.store.GetSubject(r.Context(), id)
	if errors.Is(err, ErrNotFound) {
		writeError(w, http.StatusNotFound, "subject not found")
		return
	}
	if err != nil {
		writeError(w, http.StatusInternalServerError, "database error")
		return
	}
	writeJSON(w, http.StatusOK, sub)
}

type subjectRequest struct {
	Title string `json:"title"`
	Hours int    `json:"hours"`
}

func (h *ScheduleHandler) createSubject(w http.ResponseWriter, r *http.Request) {
	var req subjectRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON")
		return
	}
	req.Title = strings.TrimSpace(req.Title)
	if req.Title == "" {
		writeError(w, http.StatusUnprocessableEntity, "title is required")
		return
	}
	sub, err := h.store.CreateSubject(r.Context(), req.Title, req.Hours)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "database error")
		return
	}
	writeJSON(w, http.StatusCreated, sub)
}

func (h *ScheduleHandler) deleteSubject(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		writeError(w, http.StatusBadRequest, "id must be a number")
		return
	}
	err = h.store.DeleteSubject(r.Context(), id)
	if errors.Is(err, ErrNotFound) {
		writeError(w, http.StatusNotFound, "subject not found")
		return
	}
	if err != nil {
		writeError(w, http.StatusInternalServerError, "database error")
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *ScheduleHandler) listRooms(w http.ResponseWriter, r *http.Request) {
	items, err := h.store.ListRooms(r.Context())
	if err != nil {
		writeError(w, http.StatusInternalServerError, "database error")
		return
	}
	writeJSON(w, http.StatusOK, items)
}

func (h *ScheduleHandler) getRoom(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		writeError(w, http.StatusBadRequest, "id must be a number")
		return
	}
	rm, err := h.store.GetRoom(r.Context(), id)
	if errors.Is(err, ErrNotFound) {
		writeError(w, http.StatusNotFound, "room not found")
		return
	}
	if err != nil {
		writeError(w, http.StatusInternalServerError, "database error")
		return
	}
	writeJSON(w, http.StatusOK, rm)
}

type roomRequest struct {
	Number   string `json:"number"`
	Capacity int    `json:"capacity"`
}

func (h *ScheduleHandler) createRoom(w http.ResponseWriter, r *http.Request) {
	var req roomRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON")
		return
	}
	req.Number = strings.TrimSpace(req.Number)
	if req.Number == "" {
		writeError(w, http.StatusUnprocessableEntity, "number is required")
		return
	}
	rm, err := h.store.CreateRoom(r.Context(), req.Number, req.Capacity)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "database error")
		return
	}
	writeJSON(w, http.StatusCreated, rm)
}

func (h *ScheduleHandler) deleteRoom(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		writeError(w, http.StatusBadRequest, "id must be a number")
		return
	}
	err = h.store.DeleteRoom(r.Context(), id)
	if errors.Is(err, ErrNotFound) {
		writeError(w, http.StatusNotFound, "room not found")
		return
	}
	if err != nil {
		writeError(w, http.StatusInternalServerError, "database error")
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *ScheduleHandler) listLessons(w http.ResponseWriter, r *http.Request) {
	items, err := h.store.ListLessons(r.Context())
	if err != nil {
		writeError(w, http.StatusInternalServerError, "database error")
		return
	}
	writeJSON(w, http.StatusOK, items)
}

func (h *ScheduleHandler) getLesson(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		writeError(w, http.StatusBadRequest, "id must be a number")
		return
	}
	l, err := h.store.GetLesson(r.Context(), id)
	if errors.Is(err, ErrNotFound) {
		writeError(w, http.StatusNotFound, "lesson not found")
		return
	}
	if err != nil {
		writeError(w, http.StatusInternalServerError, "database error")
		return
	}
	writeJSON(w, http.StatusOK, l)
}

type lessonRequest struct {
	GroupID   int64  `json:"group_id"`
	TeacherID int64  `json:"teacher_id"`
	SubjectID int64  `json:"subject_id"`
	RoomID    int64  `json:"room_id"`
	StartsAt  string `json:"starts_at"`
}

func (h *ScheduleHandler) createLesson(w http.ResponseWriter, r *http.Request) {
	var req lessonRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON")
		return
	}
	if req.GroupID == 0 || req.TeacherID == 0 || req.SubjectID == 0 || req.RoomID == 0 || req.StartsAt == "" {
		writeError(w, http.StatusUnprocessableEntity, "all fields are required")
		return
	}
	l, err := h.store.CreateLesson(r.Context(),
		req.GroupID, req.TeacherID, req.SubjectID, req.RoomID, req.StartsAt)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "database error")
		return
	}
	writeJSON(w, http.StatusCreated, l)
}

func (h *ScheduleHandler) deleteLesson(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		writeError(w, http.StatusBadRequest, "id must be a number")
		return
	}
	err = h.store.DeleteLesson(r.Context(), id)
	if errors.Is(err, ErrNotFound) {
		writeError(w, http.StatusNotFound, "lesson not found")
		return
	}
	if err != nil {
		writeError(w, http.StatusInternalServerError, "database error")
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
