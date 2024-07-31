package repository_test

import (
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	repository "github.com/paper-assessment/internal/user/repository/user"
	"github.com/stretchr/testify/assert"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestUserRepository_Get(t *testing.T) {
	// create sql mock connection
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// create gorm connection
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	assert.NoError(t, err)

	repo := repository.NewUserRepository(gormDB)

	t.Run("wallet found", func(t *testing.T) {
		// expected variables
		expectedUserId := "test_user"
		expectedName := "user"

		// mock query
		mock.ExpectQuery(`SELECT \* FROM "users" WHERE id = \$1`).
			WithArgs(expectedUserId).
			WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).
				AddRow(expectedUserId, expectedName))

		// call get user
		ctx := context.TODO()
		user, err := repo.Get(ctx, expectedUserId)

		// assertion
		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, expectedUserId, user.Id)
		assert.Equal(t, expectedName, user.Name)

		// ensure no error
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("user not found found", func(t *testing.T) {
		// expected variables
		nonExistantUser := "not_exist_user"

		// mock query
		mock.ExpectQuery(`SELECT \* FROM "users" WHERE id = \$1`).
			WithArgs(nonExistantUser).
			WillReturnError(gorm.ErrRecordNotFound)

		// call get user
		ctx := context.TODO()
		user, err := repo.Get(ctx, nonExistantUser)

		// assertion
		assert.NoError(t, err)
		assert.Nil(nil, user)
	})

	t.Run("database error", func(t *testing.T) {
		// expected variables
		expectedUserId := "test_user"

		// mock query
		mock.ExpectQuery(`SELECT \* FROM "users" WHERE id = \$1`).
			WithArgs(expectedUserId).
			WillReturnError(errors.New("database error"))

		// call repository
		user, err := repo.Get(context.Background(), expectedUserId)

		// assertion
		assert.Error(t, err)
		assert.EqualError(t, err, "database error")
		assert.Nil(t, user)
	})
}
