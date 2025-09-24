package usecase

import (
	"context"
	"testing"

	mockrepo "payment-service/pkg/mocks/modules/bank/repository"
	mocksharedrepo "payment-service/pkg/mocks/shared/repository"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_bankUsecaseImpl_DeleteBank(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {
		bankRepo := &mockrepo.BankRepository{}
		bankRepo.On("Delete", mock.Anything, mock.Anything).Return(nil)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("BankRepo").Return(bankRepo)

		uc := bankUsecaseImpl{
			repoSQL: repoSQL,
		}

		err := uc.DeleteBank(context.Background(), 1)
		assert.NoError(t, err)
	})
}
