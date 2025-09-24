package usecase

import (
	"context"
	"testing"

	mockrepo "payment-service/pkg/mocks/modules/method/repository"
	mocksharedrepo "payment-service/pkg/mocks/shared/repository"
	shareddomain "payment-service/pkg/shared/domain"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_methodUsecaseImpl_GetDetailMethod(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {
		methodRepo := &mockrepo.MethodRepository{}
		methodRepo.On("Find", mock.Anything, mock.Anything).Return(shareddomain.PaymentMethod{}, nil)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("MethodRepo").Return(methodRepo)

		uc := methodUsecaseImpl{
			repoSQL: repoSQL,
		}

		_, err := uc.GetDetailMethod(context.Background(), 1)
		assert.NoError(t, err)
	})
}
