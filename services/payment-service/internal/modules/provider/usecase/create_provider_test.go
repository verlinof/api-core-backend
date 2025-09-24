package usecase

import (
	"context"
	"testing"

	"payment-service/internal/modules/provider/domain"
	mockrepo "payment-service/pkg/mocks/modules/provider/repository"
	mocksharedrepo "payment-service/pkg/mocks/shared/repository"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_providerUsecaseImpl_CreateProvider(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {
		providerRepo := &mockrepo.ProviderRepository{}
		providerRepo.On("Save", mock.Anything, mock.Anything).Return(nil)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("ProviderRepo").Return(providerRepo)

		uc := providerUsecaseImpl{
			repoSQL: repoSQL,
		}

		_, err := uc.CreateProvider(context.Background(), &domain.RequestProvider{})
		assert.NoError(t, err)
	})
}
