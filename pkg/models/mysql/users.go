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
	var id int
	var hashedPassword []byte

	stmt := `SELECT	id, hashed_password FROM users WHERE email = ?`

	row := model.DB.QueryRow(stmt, email)
	err := row.Scan(&id, &hashedPassword)

	if err == sql.ErrNoRows {
		return 0, models.ErrInvalidCredentials
	} else if err != nil {
		return 0, err
	}

	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))

	if err == bcrypt.ErrMismatchedHashAndPassword {
		return 0, models.ErrInvalidCredentials
	} else if err != nil {
		return 0, err
	}

	return id, nil
}

func (model *UserModel) Get(id int) (*models.User, error) {
	user := &models.User{}

	stmt := `SELECT id, name, email, created FROM users WHERE id = ?`
	err := model.DB.QueryRow(stmt, id).Scan(&user.ID, &user.Name, &user.Email, &user.Created)

	if err == sql.ErrNoRows {
		return nil, models.ErrNoRecord
	} else if err != nil {
		return nil, err
	}

	return user, nil
}
