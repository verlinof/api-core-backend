package usecase

import (
	"context"
	"testing"

	mockrepo "monorepo/services/order-service/pkg/mocks/modules/order/repository"
	mocksharedrepo "monorepo/services/order-service/pkg/mocks/shared/repository"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_orderUsecaseImpl_DeleteOrder(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {
		orderRepo := &mockrepo.OrderRepository{}
		orderRepo.On("Delete", mock.Anything, mock.Anything).Return(nil)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("OrderRepo").Return(orderRepo)

		uc := orderUsecaseImpl{
			repoSQL: repoSQL,
		}

		err := uc.DeleteOrder(context.Background(), 1)
		assert.NoError(t, err)
	})
}
