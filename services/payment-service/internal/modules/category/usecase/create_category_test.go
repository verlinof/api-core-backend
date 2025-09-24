package usecase

import (
	"context"
	"testing"

	"payment-service/internal/modules/category/domain"
	mockrepo "payment-service/pkg/mocks/modules/category/repository"
	mocksharedrepo "payment-service/pkg/mocks/shared/repository"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_categoryUsecaseImpl_CreateCategory(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {
		categoryRepo := &mockrepo.CategoryRepository{}
		categoryRepo.On("Save", mock.Anything, mock.Anything).Return(nil)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("CategoryRepo").Return(categoryRepo)

		uc := categoryUsecaseImpl{
			repoSQL: repoSQL,
		}

		_, err := uc.CreateCategory(context.Background(), &domain.RequestCategory{})
		assert.NoError(t, err)
	})
}
