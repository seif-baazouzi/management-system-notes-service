package models

import (
	"database/sql"
	"notes-service/src/db"
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
