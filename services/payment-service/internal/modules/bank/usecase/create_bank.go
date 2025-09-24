package usecase

import (
	"context"

	"payment-service/internal/modules/bank/domain"

	"github.com/golangid/candi/tracer"
)

func (uc *bankUsecaseImpl) CreateBank(ctx context.Context, req *domain.RequestBank) (result domain.ResponseBank, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "BankUsecase:CreateBank")
	defer trace.Finish()

	data := req.Deserialize()
	err = uc.repoSQL.BankRepo().Save(ctx, &data)
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
