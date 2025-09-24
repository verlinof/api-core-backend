package usecase

import (
	"context"
	"testing"

	mockrepo "payment-service/pkg/mocks/modules/bank/repository"
	mocksharedrepo "payment-service/pkg/mocks/shared/repository"
	shareddomain "payment-service/pkg/shared/domain"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_bankUsecaseImpl_GetDetailBank(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {
		bankRepo := &mockrepo.BankRepository{}
		bankRepo.On("Find", mock.Anything, mock.Anything).Return(shareddomain.PaymentBank{}, nil)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("BankRepo").Return(bankRepo)

		uc := bankUsecaseImpl{
			repoSQL: repoSQL,
		}

		_, err := uc.GetDetailBank(context.Background(), 1)
		assert.NoError(t, err)
	})
}
