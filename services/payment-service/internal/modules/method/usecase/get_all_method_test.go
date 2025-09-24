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

func Test_methodUsecaseImpl_GetAllMethod(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {
		methodRepo := &mockrepo.MethodRepository{}
		methodRepo.On("FetchAll", mock.Anything, mock.Anything, mock.Anything).Return([]shareddomain.PaymentMethod{}, nil)
		methodRepo.On("Count", mock.Anything, mock.Anything).Return(10)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("MethodRepo").Return(methodRepo)

		uc := methodUsecaseImpl{
			repoSQL: repoSQL,
		}

		_, err := uc.GetAllMethod(context.Background(), &domain.FilterMethod{})
		assert.NoError(t, err)
	})

	t.Run("Testcase #2: Negative", func(t *testing.T) {
		methodRepo := &mockrepo.MethodRepository{}
		methodRepo.On("FetchAll", mock.Anything, mock.Anything, mock.Anything).Return([]shareddomain.PaymentMethod{}, errors.New("Error"))
		methodRepo.On("Count", mock.Anything, mock.Anything).Return(10)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("MethodRepo").Return(methodRepo)

		uc := methodUsecaseImpl{
			repoSQL: repoSQL,
		}

		_, err := uc.GetAllMethod(context.Background(), &domain.FilterMethod{})
		assert.Error(t, err)
	})
}
