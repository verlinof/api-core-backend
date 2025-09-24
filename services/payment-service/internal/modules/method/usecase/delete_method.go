package usecase

import (
	"context"
	
	"payment-service/internal/modules/method/domain"

	"github.com/golangid/candi/tracer"
)

func (uc *methodUsecaseImpl) DeleteMethod(ctx context.Context, id int) (err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "MethodUsecase:DeleteMethod")
	defer trace.Finish()

	repoFilter := domain.FilterMethod{ID: &id}
	return uc.repoSQL.MethodRepo().Delete(ctx, &repoFilter)
}
