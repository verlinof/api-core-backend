package usecase

import (
	"context"
	"testing"

	mockrepo "payment-service/pkg/mocks/modules/category/repository"
	mocksharedrepo "payment-service/pkg/mocks/shared/repository"
	shareddomain "payment-service/pkg/shared/domain"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_categoryUsecaseImpl_GetDetailCategory(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {
		categoryRepo := &mockrepo.CategoryRepository{}
		categoryRepo.On("Find", mock.Anything, mock.Anything).Return(shareddomain.PaymentCategory{}, nil)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("CategoryRepo").Return(categoryRepo)

		uc := categoryUsecaseImpl{
			repoSQL: repoSQL,
		}

		_, err := uc.GetDetailCategory(context.Background(), 1)
		assert.NoError(t, err)
	})
}
