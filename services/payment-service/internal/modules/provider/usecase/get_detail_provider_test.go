package usecase

import (
	"context"
	"testing"

	mockrepo "payment-service/pkg/mocks/modules/provider/repository"
	mocksharedrepo "payment-service/pkg/mocks/shared/repository"
	shareddomain "payment-service/pkg/shared/domain"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_providerUsecaseImpl_GetDetailProvider(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {
		providerRepo := &mockrepo.ProviderRepository{}
		providerRepo.On("Find", mock.Anything, mock.Anything).Return(shareddomain.PaymentProvider{}, nil)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("ProviderRepo").Return(providerRepo)

		uc := providerUsecaseImpl{
			repoSQL: repoSQL,
		}

		_, err := uc.GetDetailProvider(context.Background(), 1)
		assert.NoError(t, err)
	})
}
