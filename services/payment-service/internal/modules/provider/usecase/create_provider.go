package usecase

import (
	"context"

	"payment-service/internal/modules/provider/domain"

	"github.com/golangid/candi/tracer"
)

func (uc *providerUsecaseImpl) CreateProvider(ctx context.Context, req *domain.RequestProvider) (result domain.ResponseProvider, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "ProviderUsecase:CreateProvider")
	defer trace.Finish()

	data := req.Deserialize()
	err = uc.repoSQL.ProviderRepo().Save(ctx, &data)
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
