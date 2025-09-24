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

func Test_providerUsecaseImpl_UpdateProvider(t *testing.T) {
	ctx := context.Background()
	t.Run("Testcase #1: Positive", func(t *testing.T) {
		providerRepo := &mockrepo.ProviderRepository{}
		providerRepo.On("Find", mock.Anything, mock.Anything).Return(shareddomain.PaymentProvider{}, nil)
		providerRepo.On("Save", mock.Anything, mock.Anything, mock.AnythingOfType("candishared.DBUpdateOptionFunc")).Return(nil)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("ProviderRepo").Return(providerRepo)
		repoSQL.On("WithTransaction", mock.Anything,
			mock.AnythingOfType("func(context.Context) error")).
			Return(nil).
			Run(func(args mock.Arguments) {
				arg := args.Get(1).(func(context.Context) error)
				arg(ctx)
			})
		uc := providerUsecaseImpl{
			repoSQL: repoSQL,
		}

		err := uc.UpdateProvider(ctx, &domain.RequestProvider{})
		assert.NoError(t, err)
	})

	t.Run("Testcase #2: Negative", func(t *testing.T) {
		providerRepo := &mockrepo.ProviderRepository{}
		providerRepo.On("Find", mock.Anything, mock.Anything).Return(shareddomain.PaymentProvider{}, errors.New("Error"))
		providerRepo.On("Save", mock.Anything, mock.Anything, mock.AnythingOfType("candishared.DBUpdateOptionFunc")).Return(nil)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("ProviderRepo").Return(providerRepo)
		repoSQL.On("WithTransaction", mock.Anything,
			mock.AnythingOfType("func(context.Context) error")).
			Return(nil).
			Run(func(args mock.Arguments) {
				arg := args.Get(1).(func(context.Context) error)
				arg(ctx)
			})
		uc := providerUsecaseImpl{
			repoSQL: repoSQL,
		}

		err := uc.UpdateProvider(ctx, &domain.RequestProvider{})
		assert.Error(t, err)
	})
}
