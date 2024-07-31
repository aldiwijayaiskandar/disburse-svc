package repository_test

import (
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	repository "github.com/paper-assessment/internal/wallet/repository/wallet"
	"github.com/stretchr/testify/assert"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestWalletRepository_Get(t *testing.T) {
	// create sql mock connection
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// create gorm connection
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	assert.NoError(t, err)

	repo := repository.NewWalletRepository(gormDB)

	t.Run("wallet found", func(t *testing.T) {
		// expected variables
		expectedUserId := "test_user"
		expectedBalance := 100.50

		// mock query
		mock.ExpectQuery(`SELECT \* FROM "wallets" WHERE user_id = \$1`).
			WithArgs(expectedUserId).
			WillReturnRows(sqlmock.NewRows([]string{"user_id", "balance"}).
				AddRow(expectedUserId, expectedBalance))

		// call get wallet
		ctx := context.TODO()
		wallet, err := repo.Get(ctx, expectedUserId)

		// assertion
		assert.NoError(t, err)
		assert.NotNil(t, wallet)
		assert.Equal(t, expectedUserId, wallet.UserId)
		assert.Equal(t, expectedBalance, wallet.Balance)

		// ensure no error
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("wallet not found found", func(t *testing.T) {
		// expected variables
		nonExistantUser := "not_exist_user"

		// mock query
		mock.ExpectQuery(`SELECT \* FROM "wallets" WHERE user_id = \$1`).
			WithArgs(nonExistantUser).
			WillReturnRows(sqlmock.NewRows([]string{"user_id", "balance"}))

		// call get wallet
		ctx := context.TODO()
		wallet, err := repo.Get(ctx, nonExistantUser)

		// assertion
		assert.NoError(t, err)
		assert.Nil(nil, wallet)
	})

	t.Run("database error", func(t *testing.T) {
		// expected variables
		expectedUserId := "test_user"

		// mock query
		mock.ExpectQuery(`SELECT \* FROM "wallets" WHERE user_id = \$1`).
			WithArgs(expectedUserId).
			WillReturnError(errors.New("database error"))

		// call repository
		wallet, err := repo.Get(context.Background(), expectedUserId)

		// assertion
		assert.Error(t, err)
		assert.EqualError(t, err, "database error")
		assert.Nil(t, wallet)
	})
}
