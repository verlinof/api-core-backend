package usecase

import (
	"context"
	"testing"

	mockrepo "payment-service/pkg/mocks/modules/method/repository"
	mocksharedrepo "payment-service/pkg/mocks/shared/repository"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_methodUsecaseImpl_DeleteMethod(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {
		methodRepo := &mockrepo.MethodRepository{}
		methodRepo.On("Delete", mock.Anything, mock.Anything).Return(nil)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("MethodRepo").Return(methodRepo)

		uc := methodUsecaseImpl{
			repoSQL: repoSQL,
		}

		err := uc.DeleteMethod(context.Background(), 1)
		assert.NoError(t, err)
	})
}
