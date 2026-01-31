package todo

import (
	"time"

	"github.com/google/uuid"
)

type PriorityLevel int

const (
	PriorityLow PriorityLevel = iota
	PriorityMedium
	PriorityHigh
)

type Todo struct {
	ID          uuid.UUID     `json:"id"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	CreatedAt   time.Time     `json:"created_at"`
	Priority    PriorityLevel `json:"priority"`
}
