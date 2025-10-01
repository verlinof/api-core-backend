package usecase

import (
	"context"
	
	"monorepo/services/order-service/internal/modules/order/domain"

	"github.com/golangid/candi/tracer"
)

func (uc *orderUsecaseImpl) DeleteOrder(ctx context.Context, id int) (err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "OrderUsecase:DeleteOrder")
	defer trace.Finish()

	repoFilter := domain.FilterOrder{ID: &id}
	return uc.repoSQL.OrderRepo().Delete(ctx, &repoFilter)
}
