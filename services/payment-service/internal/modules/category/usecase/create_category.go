package usecase

import (
	"context"

	"payment-service/internal/modules/category/domain"

	"github.com/golangid/candi/tracer"
)

func (uc *categoryUsecaseImpl) CreateCategory(ctx context.Context, req *domain.RequestCategory) (result domain.ResponseCategory, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "CategoryUsecase:CreateCategory")
	defer trace.Finish()

	data := req.Deserialize()
	err = uc.repoSQL.CategoryRepo().Save(ctx, &data)
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
