package usecase

import (
	"context"

	"monorepo/services/order-service/internal/modules/order/domain"

	"github.com/golangid/candi/candishared"
	"github.com/golangid/candi/tracer"
)

func (uc *orderUsecaseImpl) UpdateOrder(ctx context.Context, data *domain.RequestOrder) (err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "OrderUsecase:UpdateOrder")
	defer trace.Finish()

	repoFilter := domain.FilterOrder{ID: &data.ID}
	existing, err := uc.repoSQL.OrderRepo().Find(ctx, &repoFilter)
	if err != nil {
		return err
	}
	existing.Field = data.Field
	err = uc.repoSQL.WithTransaction(ctx, func(ctx context.Context) error {
		return uc.repoSQL.OrderRepo().Save(ctx, &existing, candishared.DBUpdateSetUpdatedFields("Field"))
	})
	return
}
