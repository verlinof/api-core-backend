package usecase

import (
	"context"
	"testing"

	"payment-service/internal/modules/bank/domain"
	mockrepo "payment-service/pkg/mocks/modules/bank/repository"
	mocksharedrepo "payment-service/pkg/mocks/shared/repository"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_bankUsecaseImpl_CreateBank(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {
		bankRepo := &mockrepo.BankRepository{}
		bankRepo.On("Save", mock.Anything, mock.Anything).Return(nil)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("BankRepo").Return(bankRepo)

		uc := bankUsecaseImpl{
			repoSQL: repoSQL,
		}

		_, err := uc.CreateBank(context.Background(), &domain.RequestBank{})
		assert.NoError(t, err)
	})
}
