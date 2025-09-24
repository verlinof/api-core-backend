package usecase

import (
	"context"
	"errors"
	"testing"

	"payment-service/internal/modules/category/domain"
	mockrepo "payment-service/pkg/mocks/modules/category/repository"
	mocksharedrepo "payment-service/pkg/mocks/shared/repository"
	shareddomain "payment-service/pkg/shared/domain"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_categoryUsecaseImpl_GetAllCategory(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {
		categoryRepo := &mockrepo.CategoryRepository{}
		categoryRepo.On("FetchAll", mock.Anything, mock.Anything, mock.Anything).Return([]shareddomain.PaymentCategory{}, nil)
		categoryRepo.On("Count", mock.Anything, mock.Anything).Return(10)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("CategoryRepo").Return(categoryRepo)

		uc := categoryUsecaseImpl{
			repoSQL: repoSQL,
		}

		_, err := uc.GetAllCategory(context.Background(), &domain.FilterCategory{})
		assert.NoError(t, err)
	})

	t.Run("Testcase #2: Negative", func(t *testing.T) {
		categoryRepo := &mockrepo.CategoryRepository{}
		categoryRepo.On("FetchAll", mock.Anything, mock.Anything, mock.Anything).Return([]shareddomain.PaymentCategory{}, errors.New("Error"))
		categoryRepo.On("Count", mock.Anything, mock.Anything).Return(10)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("CategoryRepo").Return(categoryRepo)

		uc := categoryUsecaseImpl{
			repoSQL: repoSQL,
		}

		_, err := uc.GetAllCategory(context.Background(), &domain.FilterCategory{})
		assert.Error(t, err)
	})
}
