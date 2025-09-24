package usecase

import (
	"context"
	"testing"

	mockrepo "payment-service/pkg/mocks/modules/provider/repository"
	mocksharedrepo "payment-service/pkg/mocks/shared/repository"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_providerUsecaseImpl_DeleteProvider(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {
		providerRepo := &mockrepo.ProviderRepository{}
		providerRepo.On("Delete", mock.Anything, mock.Anything).Return(nil)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("ProviderRepo").Return(providerRepo)

		uc := providerUsecaseImpl{
			repoSQL: repoSQL,
		}

		err := uc.DeleteProvider(context.Background(), 1)
		assert.NoError(t, err)
	})
}
