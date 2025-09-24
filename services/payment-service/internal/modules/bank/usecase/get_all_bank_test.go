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

func Test_bankUsecaseImpl_GetAllBank(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {
		bankRepo := &mockrepo.BankRepository{}
		bankRepo.On("FetchAll", mock.Anything, mock.Anything, mock.Anything).Return([]shareddomain.PaymentBank{}, nil)
		bankRepo.On("Count", mock.Anything, mock.Anything).Return(10)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("BankRepo").Return(bankRepo)

		uc := bankUsecaseImpl{
			repoSQL: repoSQL,
		}

		_, err := uc.GetAllBank(context.Background(), &domain.FilterBank{})
		assert.NoError(t, err)
	})

	t.Run("Testcase #2: Negative", func(t *testing.T) {
		bankRepo := &mockrepo.BankRepository{}
		bankRepo.On("FetchAll", mock.Anything, mock.Anything, mock.Anything).Return([]shareddomain.PaymentBank{}, errors.New("Error"))
		bankRepo.On("Count", mock.Anything, mock.Anything).Return(10)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("BankRepo").Return(bankRepo)

		uc := bankUsecaseImpl{
			repoSQL: repoSQL,
		}

		_, err := uc.GetAllBank(context.Background(), &domain.FilterBank{})
		assert.Error(t, err)
	})
}
