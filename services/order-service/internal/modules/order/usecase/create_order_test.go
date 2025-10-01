package usecase

import (
	"context"
	"testing"

	"monorepo/services/order-service/internal/modules/order/domain"
	mockrepo "monorepo/services/order-service/pkg/mocks/modules/order/repository"
	mocksharedrepo "monorepo/services/order-service/pkg/mocks/shared/repository"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_orderUsecaseImpl_CreateOrder(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {
		orderRepo := &mockrepo.OrderRepository{}
		orderRepo.On("Save", mock.Anything, mock.Anything).Return(nil)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("OrderRepo").Return(orderRepo)

		uc := orderUsecaseImpl{
			repoSQL: repoSQL,
		}

		_, err := uc.CreateOrder(context.Background(), &domain.RequestOrder{})
		assert.NoError(t, err)
	})
}
