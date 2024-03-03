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
	return nil, nil
}

func (model *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
