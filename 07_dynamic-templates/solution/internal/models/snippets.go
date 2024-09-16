package models

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

type SnippetModel struct {
	DB *sql.DB
}

// Insert inserts a snippet into the database, and returns the snippet's id.
func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	stmt := `
		INSERT INTO snippets (title, content, created, expires)
		VALUES (?, ?, CURRENT_TIMESTAMP, ?)
	`
	expiresParam := fmt.Sprintf("DATETIME(CURRENT_TIMESTAMP, '+%v days')", expires)
	result, err := m.DB.Exec(stmt, title, content, expiresParam)
	if err != nil {
		return 0, fmt.Errorf("inserting row into database: %w", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("getting last insert ID: %w", err)
	}
	return int(id), nil
}

func (m *SnippetModel) Get(id int) (*Snippet, error) {
	stmt := `
		SELECT id, title, content, created, expires
		FROM snippets
		WHERE expires > CURRENT_TIMESTAMP AND id = ?
	`
	row := m.DB.QueryRow(stmt, id)
	var s Snippet
	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		}
		return nil, err
	}
	return &s, nil
}

func (m *SnippetModel) Latest() ([]*Snippet, error) {
	stmt := `
		SELECT id, title, content, created, expires
		FROM snippets
		WHERE expires > CURRENT_TIMESTAMP
		ORDER BY id
		LIMIT 10
	`
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var snippets []*Snippet
	for rows.Next() {
		var s Snippet
		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}
		snippets = append(snippets, &s)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return snippets, nil
}
