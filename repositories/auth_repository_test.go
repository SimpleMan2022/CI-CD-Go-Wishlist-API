package repositories

import (
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"go-wishlist-api-2/entities"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
	"time"
)

func CreateGormDB(db *sql.DB) *gorm.DB {
	gormDB, err := gorm.Open(mysql.New(mysql.Config{Conn: db, SkipInitializeWithVersion: true}), &gorm.Config{})
	if err != nil {
		panic("failed to connect database:" + err.Error())
	}
	return gormDB
}
func TestAuthRepository(t *testing.T) {
	testCases := []struct {
		name      string
		setup     func(mock sqlmock.Sqlmock, repo authRepository)
		assertion func(t *testing.T, err error, user *entities.User)
	}{
		{
			name: "FindByEmail - success",
			setup: func(mock sqlmock.Sqlmock, repo authRepository) {
				user := &entities.User{
					Id:        1,
					Email:     "admin@example.com",
					Password:  "admin123",
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				}

				rows := sqlmock.
					NewRows([]string{"id", "email", "password", "created_at", "updated_at"}).
					AddRow(user.Id, user.Email, user.Password, user.CreatedAt, user.UpdatedAt)

				query := "SELECT * FROM `users` WHERE email = ? ORDER BY `users`.`id` LIMIT ?"
				mock.ExpectQuery(query).
					WithArgs("admin@example.com", 1).
					WillReturnRows(rows)
			},
			assertion: func(t *testing.T, err error, user *entities.User) {
				assert.NoError(t, err)
				assert.NotNil(t, user)
				assert.Equal(t, user.Email, "admin@example.com")
			},
		},
		{
			name: "FindByEmail - error",
			setup: func(mock sqlmock.Sqlmock, repo authRepository) {
				query := "SELECT * FROM `users` WHERE email = ? ORDER BY `users`.`id` LIMIT ?"
				mock.ExpectQuery(query).
					WithArgs("admin@example.com", 1).
					WillReturnError(fmt.Errorf("Failed find user by email"))
			},
			assertion: func(t *testing.T, err error, user *entities.User) {
				assert.Error(t, err)
				assert.Nil(t, user)
			},
		},
		{
			name: "Create - success",
			setup: func(mock sqlmock.Sqlmock, repo authRepository) {
				user := &entities.User{
					Id:        1,
					Email:     "admin@example.com",
					Password:  "admin123",
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				}

				mock.ExpectBegin()

				query := "INSERT INTO `users` (`email`,`password`,`created_at`,`updated_at`,`id`) VALUES (?,?,?,?,?)"
				mock.ExpectExec(query).
					WithArgs(user.Email, user.Password, user.CreatedAt, user.UpdatedAt, user.Id).
					WillReturnResult(sqlmock.NewResult(1, 1))

				mock.ExpectCommit()
			},
			assertion: func(t *testing.T, err error, user *entities.User) {
				assert.NoError(t, err)
				assert.NotNil(t, user)
				assert.Equal(t, user.Email, "admin@example.com")
			},
		},
		{
			name: "Create - error",
			setup: func(mock sqlmock.Sqlmock, repo authRepository) {
				user := &entities.User{
					Id:        1,
					Email:     "admin@example.com",
					Password:  "admin123",
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				}

				mock.ExpectBegin()

				query := "INSERT INTO `users` (`email`,`password`,`created_at`,`updated_at`,`id`) VALUES (?,?,?,?,?)"
				mock.ExpectExec(query).
					WithArgs(user.Email, user.Password, user.CreatedAt, user.UpdatedAt, user.Id).
					WillReturnError(fmt.Errorf("Failed create user"))

				mock.ExpectRollback()
			},
			assertion: func(t *testing.T, err error, user *entities.User) {
				assert.Error(t, err)
				assert.Nil(t, user)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			assert.NoError(t, err)
			defer db.Close()

			gormDB := CreateGormDB(db)
			repo := authRepository{db: gormDB}

			tc.setup(mock, repo)

			if tc.name == "FindByEmail - success" {
				user, err := repo.FindByEmail("admin@example.com")
				tc.assertion(t, err, user)
			} else if tc.name == "FindByEmail - error" {
				_, err := repo.FindByEmail("admin@example.com")
				tc.assertion(t, err, nil)
			} else if tc.name == "Create - success" {
				user := &entities.User{
					Id:        1,
					Email:     "admin@example.com",
					Password:  "admin123",
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				}
				got, err := repo.CreateUser(user)
				tc.assertion(t, err, got)
			} else if tc.name == "Create - error" {
				user := &entities.User{
					Id:        1,
					Email:     "admin@example.com",
					Password:  "admin123",
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				}
				_, err := repo.CreateUser(user)
				tc.assertion(t, err, nil)
			}
		})
	}
}
