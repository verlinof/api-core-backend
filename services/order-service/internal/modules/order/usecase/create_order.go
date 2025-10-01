package usecase

import (
	"context"

	"monorepo/services/order-service/internal/modules/order/domain"

	"github.com/golangid/candi/tracer"
)

func (uc *orderUsecaseImpl) CreateOrder(ctx context.Context, req *domain.RequestOrder) (result domain.ResponseOrder, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "OrderUsecase:CreateOrder")
	defer trace.Finish()

	data := req.Deserialize()
	err = uc.repoSQL.OrderRepo().Save(ctx, &data)
	result.Serialize(&data)

	// Sample using broker publisher
	// uc.deps.GetBroker(types.Kafka). // get registered broker type (sample Kafka)
	// 				GetPublisher().
	// 				PublishMessage(ctx, &candishared.PublisherArgument{
	// 		Topic:   "[topic]",
	// 		Key:     "[key]",
	// 		Message: candihelper.ToBytes("[message]"),
	// 	})
	return
}
