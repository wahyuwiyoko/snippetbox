package models

import (
	"errors"
	"time"
)

var (
	ErrNoRecord           = errors.New("Models: no matching record found")
	ErrInvalidCredentials = errors.New("Models: invalid credentials")
	ErrDuplicateEmail     = errors.New("Models: duplicate email")
)

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

type User struct {
	ID             int
	Name           string
	Email          string
	HashedPassword []byte
	Created        time.Time
}
