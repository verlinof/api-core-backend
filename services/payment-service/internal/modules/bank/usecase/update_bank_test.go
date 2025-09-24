package usecase

import (
	"context"
	"errors"
	"testing"

	"payment-service/internal/modules/bank/domain"
	mockrepo "payment-service/pkg/mocks/modules/bank/repository"
	mocksharedrepo "payment-service/pkg/mocks/shared/repository"
	shareddomain "payment-service/pkg/shared/domain"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_bankUsecaseImpl_UpdateBank(t *testing.T) {
	ctx := context.Background()
	t.Run("Testcase #1: Positive", func(t *testing.T) {
		bankRepo := &mockrepo.BankRepository{}
		bankRepo.On("Find", mock.Anything, mock.Anything).Return(shareddomain.PaymentBank{}, nil)
		bankRepo.On("Save", mock.Anything, mock.Anything, mock.AnythingOfType("candishared.DBUpdateOptionFunc")).Return(nil)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("BankRepo").Return(bankRepo)
		repoSQL.On("WithTransaction", mock.Anything,
			mock.AnythingOfType("func(context.Context) error")).
			Return(nil).
			Run(func(args mock.Arguments) {
				arg := args.Get(1).(func(context.Context) error)
				arg(ctx)
			})
		uc := bankUsecaseImpl{
			repoSQL: repoSQL,
		}

		err := uc.UpdateBank(ctx, &domain.RequestBank{})
		assert.NoError(t, err)
	})

	t.Run("Testcase #2: Negative", func(t *testing.T) {
		bankRepo := &mockrepo.BankRepository{}
		bankRepo.On("Find", mock.Anything, mock.Anything).Return(shareddomain.PaymentBank{}, errors.New("Error"))
		bankRepo.On("Save", mock.Anything, mock.Anything, mock.AnythingOfType("candishared.DBUpdateOptionFunc")).Return(nil)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("BankRepo").Return(bankRepo)
		repoSQL.On("WithTransaction", mock.Anything,
			mock.AnythingOfType("func(context.Context) error")).
			Return(nil).
			Run(func(args mock.Arguments) {
				arg := args.Get(1).(func(context.Context) error)
				arg(ctx)
			})
		uc := bankUsecaseImpl{
			repoSQL: repoSQL,
		}

		err := uc.UpdateBank(ctx, &domain.RequestBank{})
		assert.Error(t, err)
	})
}
