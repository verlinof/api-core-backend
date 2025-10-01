package usecase

import (
	"context"

	"monorepo/services/order-service/internal/modules/order/domain"

	"github.com/golangid/candi/tracer"
)

func (uc *orderUsecaseImpl) GetDetailOrder(ctx context.Context, id int) (result domain.ResponseOrder, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "OrderUsecase:GetDetailOrder")
	defer trace.Finish()

	repoFilter := domain.FilterOrder{ID: &id}
	data, err := uc.repoSQL.OrderRepo().Find(ctx, &repoFilter)
	if err != nil {
		return result, err
	}

	result.Serialize(&data)
	return
}
