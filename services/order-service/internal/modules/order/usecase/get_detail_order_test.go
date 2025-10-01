package usecase

import (
	"context"
	"testing"

	mockrepo "monorepo/services/order-service/pkg/mocks/modules/order/repository"
	mocksharedrepo "monorepo/services/order-service/pkg/mocks/shared/repository"
	shareddomain "monorepo/services/order-service/pkg/shared/domain"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_orderUsecaseImpl_GetDetailOrder(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {
		orderRepo := &mockrepo.OrderRepository{}
		orderRepo.On("Find", mock.Anything, mock.Anything).Return(shareddomain.Order{}, nil)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("OrderRepo").Return(orderRepo)

		uc := orderUsecaseImpl{
			repoSQL: repoSQL,
		}

		_, err := uc.GetDetailOrder(context.Background(), 1)
		assert.NoError(t, err)
	})
}
