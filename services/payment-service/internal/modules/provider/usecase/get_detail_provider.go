package usecase

import (
	"context"

	"payment-service/internal/modules/provider/domain"

	"github.com/golangid/candi/tracer"
)

func (uc *providerUsecaseImpl) GetDetailProvider(ctx context.Context, id int) (result domain.ResponseProvider, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "ProviderUsecase:GetDetailProvider")
	defer trace.Finish()

	repoFilter := domain.FilterProvider{ID: &id}
	data, err := uc.repoSQL.ProviderRepo().Find(ctx, &repoFilter)
	if err != nil {
		return result, err
	}

	result.Serialize(&data)
	return
}
