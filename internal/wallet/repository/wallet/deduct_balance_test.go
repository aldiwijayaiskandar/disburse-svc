package repository_test

import (
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/paper-assessment/internal/wallet/domain/models"
	repository "github.com/paper-assessment/internal/wallet/repository/wallet"
	"github.com/stretchr/testify/assert"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestWalletRepository_DeductBalance(t *testing.T) {
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

	t.Run("balance deducted", func(t *testing.T) {
		// expected variables
		deductionAmount := 5000.50
		expectedUserId := "expected_user_id"

		// mock expected query
		mock.ExpectBegin()
		mock.ExpectExec(`UPDATE "wallets" SET "balance"=balance - \$1 WHERE user_id = \$2`).
			WithArgs(deductionAmount, expectedUserId).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		// call deduct balance
		err := repo.DeductBalance(context.Background(), models.DeductBalanceRequest{
			UserId: expectedUserId,
			Amount: deductionAmount,
		})

		// assertion
		assert.NoError(t, err)
	})

	t.Run("database error", func(t *testing.T) {
		// variables
		deductionAmount := 100.0
		expectedUserId := "expected_user_id"

		// mocking query
		mock.ExpectBegin()
		mock.ExpectExec(`UPDATE "wallets" SET "balance"=balance - \$1 WHERE user_id = \$2`).
			WithArgs(deductionAmount, expectedUserId).
			WillReturnError(errors.New("database error"))
		mock.ExpectRollback()

		// call deduct balance
		err := repo.DeductBalance(context.TODO(), models.DeductBalanceRequest{
			UserId: expectedUserId,
			Amount: deductionAmount,
		})

		// assertion
		assert.Error(t, err)
		assert.EqualError(t, err, "database error")
	})
}
