package mysql

import (
	"database/sql"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/wahyuwiyoko/snippetbox/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

type UserModel struct {
	DB *sql.DB
}

func (model *UserModel) Insert(name, email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)

	if err != nil {
		return err
	}

	stmt := `INSERT INTO users (name, email, hashed_password, created)
	VALUES(?, ?, ?, UTC_TIMESTAMP())`

	_, err = model.DB.Exec(stmt, name, email, string(hashedPassword))

	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			if mysqlErr.Number == 1062 && strings.Contains(mysqlErr.Message, "users_uc_email") {
				return models.ErrDuplicateEmail
			}
		}
	}

	return err
}

func (model *UserModel) Authenticate(email, password string) (int, error) {
	return 0, nil
}

func (model *UserModel) Get(id int) (*models.User, error) {
	return nil, nil
}
