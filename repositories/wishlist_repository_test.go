package repositories

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"go-wishlist-api-2/entities"
	"testing"
	"time"
)

func TestWishlistRepository(t *testing.T) {
	testCases := []struct {
		name      string
		setup     func(mock sqlmock.Sqlmock, repo WishlistRepository)
		assertion func(t *testing.T, err error, wishlists []*entities.Wishlist)
	}{
		{
			name: "GetAll - success",
			setup: func(mock sqlmock.Sqlmock, repo WishlistRepository) {
				wishlists := []*entities.Wishlist{
					{ID: 1, Title: "Wishlist 1", IsAchieved: false},
					{ID: 2, Title: "Wishlist 2", IsAchieved: true},
				}

				rows := sqlmock.
					NewRows([]string{"id", "title", "is_achieved"}).
					AddRow(wishlists[0].ID, wishlists[0].Title, wishlists[0].IsAchieved).
					AddRow(wishlists[1].ID, wishlists[1].Title, wishlists[1].IsAchieved)

				query := "SELECT * FROM `wishlists` WHERE `wishlists`.`deleted_at` IS NULL"
				mock.ExpectQuery(query).
					WillReturnRows(rows)
			},
			assertion: func(t *testing.T, err error, wishlists []*entities.Wishlist) {
				assert.NoError(t, err)
				assert.NotNil(t, wishlists)
				assert.Len(t, wishlists, 2)
				assert.Equal(t, wishlists[0].Title, "Wishlist 1")
				assert.Equal(t, wishlists[1].Title, "Wishlist 2")
			},
		},
		{
			name: "GetAll - error",
			setup: func(mock sqlmock.Sqlmock, repo WishlistRepository) {
				query := "SELECT * FROM `wishlists` WHERE `wishlists`.`deleted_at` IS NULL"
				mock.ExpectQuery(query).
					WillReturnError(fmt.Errorf("Failed to get wishlists"))
			},
			assertion: func(t *testing.T, err error, wishlists []*entities.Wishlist) {
				assert.Error(t, err)
				assert.Nil(t, wishlists)
			},
		},
		{
			name: "Create - success",
			setup: func(mock sqlmock.Sqlmock, repo WishlistRepository) {
				wishlist := &entities.Wishlist{
					ID:         1,
					Title:      "New Wishlist",
					IsAchieved: false,
					CreatedAt:  time.Now(),
					UpdatedAt:  time.Now(),
				}

				mock.ExpectBegin()
				query := "INSERT INTO `wishlists` (`title`,`is_achieved`,`created_at`,`updated_at`,`deleted_at`,`id`) VALUES (?,?,?,?,?,?)"
				mock.ExpectExec(query).
					WithArgs(wishlist.Title, wishlist.IsAchieved, wishlist.CreatedAt, wishlist.UpdatedAt, nil, 1).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
			assertion: func(t *testing.T, err error, wishlists []*entities.Wishlist) {
				
			},
		},
		{
			name: "Create - error",
			setup: func(mock sqlmock.Sqlmock, repo WishlistRepository) {
				wishlist := &entities.Wishlist{
					ID:         1,
					Title:      "New Wishlist",
					IsAchieved: false,
					CreatedAt:  time.Now(),
					UpdatedAt:  time.Now(),
				}
				mock.ExpectBegin()
				query := "INSERT INTO `wishlists` (`title`,`is_achieved`,`created_at`,`updated_at`,`deleted_at`,`id`) VALUES (?,?,?,?,?,?)"
				mock.ExpectExec(query).
					WithArgs(wishlist.Title, wishlist.IsAchieved, wishlist.CreatedAt, wishlist.UpdatedAt, nil, 1).
					WillReturnError(fmt.Errorf("Failed to create wishlist"))
				mock.ExpectRollback()
			},
			assertion: func(t *testing.T, err error, wishlists []*entities.Wishlist) {
				
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			assert.NoError(t, err)
			defer func() {
				db.Close()
			}()

			gormDB := CreateGormDB(db)
			repo := NewWishlistRepository(gormDB)

			tc.setup(mock, repo)

			if tc.name == "GetAll - success" {
				wishlists, err := repo.GetAll()
				tc.assertion(t, err, wishlists)
			} else if tc.name == "GetAll - error" {
				_, err := repo.GetAll()
				tc.assertion(t, err, nil)
			} else if tc.name == "Create - success" {
				wishlist := &entities.Wishlist{
					ID:         1,
					Title:      "New Wishlist",
					IsAchieved: false,
					CreatedAt:  time.Now(),
					UpdatedAt:  time.Now(),
				}
				got, err := repo.CreateWishlist(wishlist)
				tc.assertion(t, err, []*entities.Wishlist{got})
			} else if tc.name == "Create - error" {
				wishlist := &entities.Wishlist{
					ID:         1,
					Title:      "New Wishlist",
					IsAchieved: false,
					CreatedAt:  time.Now(),
					UpdatedAt:  time.Now(),
				}
				_, err := repo.CreateWishlist(wishlist)
				tc.assertion(t, err, nil)
			}
		})
	}
}
