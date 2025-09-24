package usecase

import (
	"context"
	"errors"
	"testing"

	"payment-service/internal/modules/method/domain"
	mockrepo "payment-service/pkg/mocks/modules/method/repository"
	mocksharedrepo "payment-service/pkg/mocks/shared/repository"
	shareddomain "payment-service/pkg/shared/domain"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_methodUsecaseImpl_UpdateMethod(t *testing.T) {
	ctx := context.Background()
	t.Run("Testcase #1: Positive", func(t *testing.T) {
		methodRepo := &mockrepo.MethodRepository{}
		methodRepo.On("Find", mock.Anything, mock.Anything).Return(shareddomain.PaymentMethod{}, nil)
		methodRepo.On("Save", mock.Anything, mock.Anything, mock.AnythingOfType("candishared.DBUpdateOptionFunc")).Return(nil)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("MethodRepo").Return(methodRepo)
		repoSQL.On("WithTransaction", mock.Anything,
			mock.AnythingOfType("func(context.Context) error")).
			Return(nil).
			Run(func(args mock.Arguments) {
				arg := args.Get(1).(func(context.Context) error)
				arg(ctx)
			})
		uc := methodUsecaseImpl{
			repoSQL: repoSQL,
		}

		err := uc.UpdateMethod(ctx, &domain.RequestMethod{})
		assert.NoError(t, err)
	})

	t.Run("Testcase #2: Negative", func(t *testing.T) {
		methodRepo := &mockrepo.MethodRepository{}
		methodRepo.On("Find", mock.Anything, mock.Anything).Return(shareddomain.Method{}, errors.New("Error"))
		methodRepo.On("Save", mock.Anything, mock.Anything, mock.AnythingOfType("candishared.DBUpdateOptionFunc")).Return(nil)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("MethodRepo").Return(methodRepo)
		repoSQL.On("WithTransaction", mock.Anything,
			mock.AnythingOfType("func(context.Context) error")).
			Return(nil).
			Run(func(args mock.Arguments) {
				arg := args.Get(1).(func(context.Context) error)
				arg(ctx)
			})
		uc := methodUsecaseImpl{
			repoSQL: repoSQL,
		}

		err := uc.UpdateMethod(ctx, &domain.RequestMethod{})
		assert.Error(t, err)
	})
}
