package usecase

import (
	"context"

	"payment-service/internal/modules/checkout/domain"

	"github.com/golangid/candi/candishared"
	"github.com/golangid/candi/tracer"
)

func (uc *checkoutUsecaseImpl) GetListMethod(ctx context.Context, filter *domain.FilterCheckout) (result domain.ResponseMethodList, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "CheckoutUsecase:GetListMethod")
	defer trace.Finish()

	data, err := uc.repoSQL.CheckoutRepo().FetchMethodGroupByCategory(ctx, filter)
	if err != nil {
		return result, err
	}
	count := uc.repoSQL.CheckoutRepo().CountMethod(ctx, filter)
	result.Meta = candishared.NewMeta(filter.Page, filter.Limit, count)

	result.Data = make([]domain.ResponseMethod, len(data))
	for i, detail := range data {
		result.Data[i].Serialize(&detail)
	}

	return
}
