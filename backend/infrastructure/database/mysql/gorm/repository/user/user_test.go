package user

import (
	"database/sql"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/posiposi/project/backend/infrastructure/database/mysql/gorm/model"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupMockDB(t *testing.T) (*gorm.DB, sqlmock.Sqlmock, *sql.DB) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	d := mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	})

	gdb, err := gorm.Open(d, &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a gorm database connection", err)
	}

	return gdb, mock, db
}

func TestGetList(t *testing.T) {
	t.Run("users exist in the database", func(t *testing.T) {
		gdb, mock, db := SetupMockDB(t)
		defer db.Close()

		r, _ := NewUserRepository(gdb)

		user := model.User{
			Id:        "user_id",
			Name:      "user_name",
			Email:     "user_email",
			Password:  "user_password",
			CreatedAt: time.Now(),
		}
		user_2 := model.User{
			Id:        "user_id_2",
			Name:      "user_name_2",
			Email:     "user_email_2",
			Password:  "user_password_2",
			CreatedAt: time.Now(),
		}

		rows := sqlmock.NewRows([]string{"id", "name", "email", "password", "created_at"}).
			AddRow(user.Id, user.Name, user.Email, user.Password, user.CreatedAt).
			AddRow(user_2.Id, user_2.Name, user_2.Email, user_2.Password, user_2.CreatedAt)
		mock.ExpectQuery("^SELECT \\* FROM `users`$").WillReturnRows(rows)

		users, err := r.GetList()
		assert.NoError(t, err)
		assert.Len(t, users, 2)
		user_1_id := users[0].Id()
		user_1_name := users[0].Name()
		user_2_name := users[1].Name()
		assert.Equal(t, user.Id, user_1_id.Value())
		assert.Equal(t, user.Name, user_1_name.Value())
		assert.Equal(t, user_2.Name, user_2_name.Value())

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})

	t.Run("users do not exist in the database", func(t *testing.T) {
		gdb, mock, db := SetupMockDB(t)
		defer db.Close()

		r, _ := NewUserRepository(gdb)

		rows := sqlmock.NewRows([]string{"id", "name", "email", "password", "created_at"})
		mock.ExpectQuery("^SELECT \\* FROM `users`$").WillReturnRows(rows)

		users, err := r.GetList()
		assert.NoError(t, err)
		assert.Len(t, users, 0)
	})
}
