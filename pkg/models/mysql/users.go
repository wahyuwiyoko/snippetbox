package mysql

import (
	"database/sql"

	"github.com/wahyuwiyoko/snippetbox/pkg/models"
)

type UserModel struct {
	DB *sql.DB
}

func (model *UserModel) Insert(name, email, password string) error {
	return nil
}

func (model *UserModel) Authenticate(email, password string) (int, error) {
	return 0, nil
}

func (model *UserModel) Get(id int) (*models.User, error) {
	return nil, nil
}
