package task

import (
	"fmt"
	"time"
)

type Status string

const (
	StatusPending Status = "pending"
	StatusDone    Status = "done"
)

type Task struct {
	ID        int        `json:"id"`
	Title     string     `json:"title"`
	Status    Status     `json:"status"`
	CreatedAt time.Time  `json:"created_at"`
	DoneAt    *time.Time `json:"done_at,omitempty"`
}

func NewTask(id int, title string) Task {
	return Task{
		ID:        id,
		Title:     title,
		Status:    StatusPending,
		CreatedAt: time.Now(),
	}
}

func (t *Task) IsPending() bool {
	return t.Status == StatusPending
}

func (t *Task) Completed() {
	now := time.Now()
	t.Status = StatusDone
	t.DoneAt = &now
}

func (t Task) String() string {
	icon := "[]"

	if t.Status == StatusDone {
		icon = "[✓]"
	}

	return fmt.Sprintf("%s %d: %s", icon, t.ID, t.Title)
}
