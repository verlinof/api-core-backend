package usecase

import (
	"context"
	"errors"
	"testing"

	"monorepo/services/order-service/internal/modules/order/domain"
	mockrepo "monorepo/services/order-service/pkg/mocks/modules/order/repository"
	mocksharedrepo "monorepo/services/order-service/pkg/mocks/shared/repository"
	shareddomain "monorepo/services/order-service/pkg/shared/domain"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_orderUsecaseImpl_UpdateOrder(t *testing.T) {
	ctx := context.Background()
	t.Run("Testcase #1: Positive", func(t *testing.T) {
		orderRepo := &mockrepo.OrderRepository{}
		orderRepo.On("Find", mock.Anything, mock.Anything).Return(shareddomain.Order{}, nil)
		orderRepo.On("Save", mock.Anything, mock.Anything, mock.AnythingOfType("candishared.DBUpdateOptionFunc")).Return(nil)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("OrderRepo").Return(orderRepo)
		repoSQL.On("WithTransaction", mock.Anything,
			mock.AnythingOfType("func(context.Context) error")).
			Return(nil).
			Run(func(args mock.Arguments) {
				arg := args.Get(1).(func(context.Context) error)
				arg(ctx)
			})
		uc := orderUsecaseImpl{
			repoSQL: repoSQL,
		}

		err := uc.UpdateOrder(ctx, &domain.RequestOrder{})
		assert.NoError(t, err)
	})

	t.Run("Testcase #2: Negative", func(t *testing.T) {
		orderRepo := &mockrepo.OrderRepository{}
		orderRepo.On("Find", mock.Anything, mock.Anything).Return(shareddomain.Order{}, errors.New("Error"))
		orderRepo.On("Save", mock.Anything, mock.Anything, mock.AnythingOfType("candishared.DBUpdateOptionFunc")).Return(nil)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("OrderRepo").Return(orderRepo)
		repoSQL.On("WithTransaction", mock.Anything,
			mock.AnythingOfType("func(context.Context) error")).
			Return(nil).
			Run(func(args mock.Arguments) {
				arg := args.Get(1).(func(context.Context) error)
				arg(ctx)
			})
		uc := orderUsecaseImpl{
			repoSQL: repoSQL,
		}

		err := uc.UpdateOrder(ctx, &domain.RequestOrder{})
		assert.Error(t, err)
	})
}
