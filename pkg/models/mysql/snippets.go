package mysql

import (
	"database/sql"

	"github.com/wahyuwiyoko/snippetbox/pkg/models"
)

type SnippetModel struct {
	DB *sql.DB
}

func (model *SnippetModel) Insert(title, content, expires string) (int, error) {
	stmt := `INSERT INTO snippets (title, content, created, expires)
	VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	result, err := model.DB.Exec(stmt, title, content, expires)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (model *SnippetModel) Get(id int) (*models.Snippet, error) {
	stmt := `SELECT id, title, content, created, expires FROM snippets
	WHERE expires > UTC_TIMESTAMP() AND id = ?`

	row := model.DB.QueryRow(stmt, id)

	snippet := &models.Snippet{}

	err := row.Scan(&snippet.ID, &snippet.Title, &snippet.Content, &snippet.Created, &snippet.Expires)

	if err == sql.ErrNoRows {
		return nil, models.ErrNoRecord
	} else if err != nil {
		return nil, err
	}

	return snippet, nil
}

func (model *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
