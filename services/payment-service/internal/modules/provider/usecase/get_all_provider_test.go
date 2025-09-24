package usecase

import (
	"context"
	"errors"
	"testing"

	"payment-service/internal/modules/provider/domain"
	mockrepo "payment-service/pkg/mocks/modules/provider/repository"
	mocksharedrepo "payment-service/pkg/mocks/shared/repository"
	shareddomain "payment-service/pkg/shared/domain"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_providerUsecaseImpl_GetAllProvider(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {
		providerRepo := &mockrepo.ProviderRepository{}
		providerRepo.On("FetchAll", mock.Anything, mock.Anything, mock.Anything).Return([]shareddomain.PaymentProvider{}, nil)
		providerRepo.On("Count", mock.Anything, mock.Anything).Return(10)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("ProviderRepo").Return(providerRepo)

		uc := providerUsecaseImpl{
			repoSQL: repoSQL,
		}

		_, err := uc.GetAllProvider(context.Background(), &domain.FilterProvider{})
		assert.NoError(t, err)
	})

	t.Run("Testcase #2: Negative", func(t *testing.T) {
		providerRepo := &mockrepo.ProviderRepository{}
		providerRepo.On("FetchAll", mock.Anything, mock.Anything, mock.Anything).Return([]shareddomain.PaymentProvider{}, errors.New("Error"))
		providerRepo.On("Count", mock.Anything, mock.Anything).Return(10)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("ProviderRepo").Return(providerRepo)

		uc := providerUsecaseImpl{
			repoSQL: repoSQL,
		}

		_, err := uc.GetAllProvider(context.Background(), &domain.FilterProvider{})
		assert.Error(t, err)
	})
}
