package usecase

import (
	"context"

	"payment-service/internal/modules/method/domain"

	"github.com/golangid/candi/tracer"
)

func (uc *methodUsecaseImpl) CreateMethod(ctx context.Context, req *domain.RequestMethod) (result domain.ResponseMethod, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "MethodUsecase:CreateMethod")
	defer trace.Finish()

	data := req.Deserialize()
	err = uc.repoSQL.MethodRepo().Save(ctx, &data)
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
