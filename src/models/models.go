package models

import (
	"time"

	"github.com/google/uuid"
)

type Note struct {
	NoteID      uuid.UUID `json:"todoID"`
	Title       string    `json:"title"`
	Body        string    `json:"body"`
	UserID      uuid.UUID `json:"userID"`
	WorkspaceID uuid.UUID `json:"workspaceID"`
	CreatedAt   time.Time `json:"createdAt"`
}

type NoteBody struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}
