package usecase

import (
	"context"
	"testing"

	mockrepo "payment-service/pkg/mocks/modules/category/repository"
	mocksharedrepo "payment-service/pkg/mocks/shared/repository"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_categoryUsecaseImpl_DeleteCategory(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {
		categoryRepo := &mockrepo.CategoryRepository{}
		categoryRepo.On("Delete", mock.Anything, mock.Anything).Return(nil)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("CategoryRepo").Return(categoryRepo)

		uc := categoryUsecaseImpl{
			repoSQL: repoSQL,
		}

		err := uc.DeleteCategory(context.Background(), 1)
		assert.NoError(t, err)
	})
}
