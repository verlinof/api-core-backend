package usecase

import (
	"context"

	"payment-service/internal/modules/provider/domain"

	"github.com/golangid/candi/candishared"
	"github.com/golangid/candi/tracer"
)

func (uc *providerUsecaseImpl) GetAllProvider(ctx context.Context, filter *domain.FilterProvider) (result domain.ResponseProviderList, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "ProviderUsecase:GetAllProvider")
	defer trace.Finish()

	data, err := uc.repoSQL.ProviderRepo().FetchAll(ctx, filter)
	if err != nil {
		return result, err
	}
	count := uc.repoSQL.ProviderRepo().Count(ctx, filter)
	result.Meta = candishared.NewMeta(filter.Page, filter.Limit, count)

	result.Data = make([]domain.ResponseProvider, len(data))
	for i, detail := range data {
		result.Data[i].Serialize(&detail)
	}

	return
}
