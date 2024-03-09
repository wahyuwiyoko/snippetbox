package mysql

import (
	"reflect"
	"testing"
	"time"

	"github.com/wahyuwiyoko/snippetbox/pkg/models"
)

func TestUserModelGet(t *testing.T) {
	if testing.Short() {
		t.Skip("mysql: skipping integration test")
	}

	tests := []struct {
		name      string
		userID    int
		wantUser  *models.User
		wantError error
	}{
		{
			name:   "Valid ID",
			userID: 1,
			wantUser: &models.User{
				ID:      1,
				Name:    "Alice Jones",
				Email:   "alice@example.com",
				Created: time.Date(2018, 12, 23, 17, 25, 22, 0, time.UTC),
			},
			wantError: nil,
		},
		{
			name:      "Zero ID",
			userID:    0,
			wantUser:  nil,
			wantError: models.ErrNoRecord,
		},
		{
			name:      "Non-existent ID",
			userID:    2,
			wantUser:  nil,
			wantError: models.ErrNoRecord,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			db, teardown := newTestDB(t)

			defer teardown()

			model := UserModel{db}

			user, err := model.Get(test.userID)

			if err != test.wantError {
				t.Errorf("want %v; got %s", test.wantError, err)
			}

			if !reflect.DeepEqual(user, test.wantUser) {
				t.Errorf("want %v; got %v", test.wantUser, user)
			}
		})
	}
}
