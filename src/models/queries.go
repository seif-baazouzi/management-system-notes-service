package models

import (
	"database/sql"
	"notes-service/src/db"

	"github.com/google/uuid"
)

func GetNotes(workspaceID string) ([]Note, error) {
	conn := db.GetPool()
	defer db.ClosePool(conn)

	var notes []Note

	var rows *sql.Rows
	var err error

	rows, err = conn.Query(
		`SELECT noteID, title, body, userID, workspaceID, createdAt FROM notes WHERE workspaceID = $1 ORDER BY createdAt`,
		workspaceID,
	)

	if err != nil {
		return notes, err
	}

	defer rows.Close()

	for rows.Next() {
		var note Note
		rows.Scan(&note.NoteID, &note.Title, &note.Body, &note.UserID, &note.WorkspaceID, &note.CreatedAt)

		notes = append(notes, note)
	}

	return notes, nil
}

func GetSingleNote(noteID string, userID string) (bool, Note, error) {
	conn := db.GetPool()
	defer db.ClosePool(conn)

	var note Note

	var rows *sql.Rows
	var err error

	rows, err = conn.Query(
		`SELECT noteID, title, body, userID, workspaceID, createdAt FROM notes WHERE userID = $1 AND noteID = $2 ORDER BY createdAt`,
		userID,
		noteID,
	)

	if err != nil {
		return false, note, err
	}

	defer rows.Close()

	if rows.Next() {
		rows.Scan(&note.NoteID, &note.Title, &note.Body, &note.UserID, &note.WorkspaceID, &note.CreatedAt)

		return true, note, nil
	}

	return false, note, nil
}

func CreateNote(note NoteBody, workspaceID string, userID string) (string, error) {
	conn := db.GetPool()
	defer db.ClosePool(conn)

	noteID := uuid.New()

	_, err := conn.Exec(
		"INSERT INTO notes (noteID, title, body, workspaceID, userID) VALUES ($1, $2, $3, $4, $5)",
		noteID,
		note.Title,
		note.Body,
		workspaceID,
		userID,
	)

	if err != nil {
		return "", err
	}

	return noteID.String(), nil
}

func UpdateNote(note NoteBody, noteID string, userID string) error {
	conn := db.GetPool()
	defer db.ClosePool(conn)

	_, err := conn.Exec(
		"UPDATE notes SET title = $1, body = $2 WHERE noteID = $3 AND userID = $4",
		note.Title,
		note.Body,
		noteID,
		userID,
	)

	if err != nil {
		return err
	}

	return nil
}

func DeleteNote(noteID string, userID string) error {
	conn := db.GetPool()
	defer db.ClosePool(conn)

	_, err := conn.Exec(
		"DELETE FROM notes WHERE noteID = $1 AND userID = $2",
		noteID,
		userID,
	)

	if err != nil {
		return err
	}

	return nil
}

func DeleteWorkspaceNotes(workspaceID string) error {
	conn := db.GetPool()
	defer db.ClosePool(conn)

	_, err := conn.Exec(
		"DELETE FROM notes WHERE workspaceID = $1",
		workspaceID,
	)

	if err != nil {
		return err
	}

	return nil
}
