package usecase

import (
	"context"

	"payment-service/internal/modules/method/domain"

	"github.com/golangid/candi/candishared"
	"github.com/golangid/candi/tracer"
)

func (uc *methodUsecaseImpl) GetAllMethod(ctx context.Context, filter *domain.FilterMethod) (result domain.ResponseMethodList, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "MethodUsecase:GetAllMethod")
	defer trace.Finish()

	data, err := uc.repoSQL.MethodRepo().FetchAll(ctx, filter)
	if err != nil {
		return result, err
	}
	count := uc.repoSQL.MethodRepo().Count(ctx, filter)
	result.Meta = candishared.NewMeta(filter.Page, filter.Limit, count)

	result.Data = make([]domain.ResponseMethod, len(data))
	for i, detail := range data {
		result.Data[i].Serialize(&detail)
	}

	return
}
