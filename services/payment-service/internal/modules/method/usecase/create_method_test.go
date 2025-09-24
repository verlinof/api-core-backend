package usecase

import (
	"context"
	"testing"

	"payment-service/internal/modules/method/domain"
	mockrepo "payment-service/pkg/mocks/modules/method/repository"
	mocksharedrepo "payment-service/pkg/mocks/shared/repository"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_methodUsecaseImpl_CreateMethod(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {
		methodRepo := &mockrepo.MethodRepository{}
		methodRepo.On("Save", mock.Anything, mock.Anything).Return(nil)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("MethodRepo").Return(methodRepo)

		uc := methodUsecaseImpl{
			repoSQL: repoSQL,
		}

		_, err := uc.CreateMethod(context.Background(), &domain.RequestMethod{})
		assert.NoError(t, err)
	})
}
