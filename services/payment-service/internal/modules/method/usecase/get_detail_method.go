package usecase

import (
	"context"

	"payment-service/internal/modules/method/domain"

	"github.com/golangid/candi/tracer"
)

func (uc *methodUsecaseImpl) GetDetailMethod(ctx context.Context, id int) (result domain.ResponseMethod, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "MethodUsecase:GetDetailMethod")
	defer trace.Finish()

	repoFilter := domain.FilterMethod{ID: &id}
	data, err := uc.repoSQL.MethodRepo().Find(ctx, &repoFilter)
	if err != nil {
		return result, err
	}

	result.Serialize(&data)
	return
}
