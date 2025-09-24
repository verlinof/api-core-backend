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

func Test_categoryUsecaseImpl_UpdateCategory(t *testing.T) {
	ctx := context.Background()
	t.Run("Testcase #1: Positive", func(t *testing.T) {
		categoryRepo := &mockrepo.CategoryRepository{}
		categoryRepo.On("Find", mock.Anything, mock.Anything).Return(shareddomain.PaymentCategory{}, nil)
		categoryRepo.On("Save", mock.Anything, mock.Anything, mock.AnythingOfType("candishared.DBUpdateOptionFunc")).Return(nil)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("CategoryRepo").Return(categoryRepo)
		repoSQL.On("WithTransaction", mock.Anything,
			mock.AnythingOfType("func(context.Context) error")).
			Return(nil).
			Run(func(args mock.Arguments) {
				arg := args.Get(1).(func(context.Context) error)
				arg(ctx)
			})
		uc := categoryUsecaseImpl{
			repoSQL: repoSQL,
		}

		err := uc.UpdateCategory(ctx, &domain.RequestCategory{})
		assert.NoError(t, err)
	})

	t.Run("Testcase #2: Negative", func(t *testing.T) {
		categoryRepo := &mockrepo.CategoryRepository{}
		categoryRepo.On("Find", mock.Anything, mock.Anything).Return(shareddomain.PaymentCategory{}, errors.New("Error"))
		categoryRepo.On("Save", mock.Anything, mock.Anything, mock.AnythingOfType("candishared.DBUpdateOptionFunc")).Return(nil)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("CategoryRepo").Return(categoryRepo)
		repoSQL.On("WithTransaction", mock.Anything,
			mock.AnythingOfType("func(context.Context) error")).
			Return(nil).
			Run(func(args mock.Arguments) {
				arg := args.Get(1).(func(context.Context) error)
				arg(ctx)
			})
		uc := categoryUsecaseImpl{
			repoSQL: repoSQL,
		}

		err := uc.UpdateCategory(ctx, &domain.RequestCategory{})
		assert.Error(t, err)
	})
}
