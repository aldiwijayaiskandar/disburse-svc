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

func TestWalletRepository_Update(t *testing.T) {
	// Create sqlmock database connection and mock
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Create GORM DB from sqlmock connection
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	assert.NoError(t, err)

	repo := repository.NewWalletRepository(gormDB)

	t.Run("wallet found", func(t *testing.T) {
		expectedUserId := "test_user"
		expectedBalance := 100.50
		// mock query
		mock.ExpectQuery(`SELECT \* FROM "wallets" WHERE user_id = \$1`).
			WithArgs(expectedUserId).
			WillReturnRows(sqlmock.NewRows([]string{"user_id", "balance"}).
				AddRow(expectedUserId, expectedBalance))

		// Call the method to be tested
		ctx := context.TODO()
		wallet, err := repo.Get(ctx, expectedUserId)

		// Assertions
		assert.NoError(t, err)
		assert.NotNil(t, wallet)
		assert.Equal(t, expectedUserId, wallet.UserId)
		assert.Equal(t, expectedBalance, wallet.Balance)

		// Ensure all expectations were met
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("wallet not found found", func(t *testing.T) {
		nonExistantUser := "not_exist_user"

		// mock query
		mock.ExpectQuery(`SELECT \* FROM "wallets" WHERE user_id = \$1`).
			WithArgs(nonExistantUser).
			WillReturnError(gorm.ErrRecordNotFound)

		// Call the method to be tested
		ctx := context.TODO()
		wallet, err := repo.Get(ctx, nonExistantUser)

		// Assertions
		assert.NoError(t, err)
		assert.Nil(nil, wallet)
	})

	t.Run("Database Error", func(t *testing.T) {
		expectedUserId := "test_user"

		// mock query
		mock.ExpectQuery(`SELECT \* FROM "wallets" WHERE user_id = \$1`).
			WithArgs(expectedUserId).
			WillReturnError(errors.New("database error"))

		// Call the function with expected user ID
		wallet, err := repo.Get(context.Background(), expectedUserId)

		// Assert expectations
		assert.Error(t, err)
		assert.EqualError(t, err, "database error")
		assert.Nil(t, wallet)
	})
}
