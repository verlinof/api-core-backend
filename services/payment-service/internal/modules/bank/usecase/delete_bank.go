package usecase

import (
	"context"
	
	"payment-service/internal/modules/bank/domain"

	"github.com/golangid/candi/tracer"
)

func (uc *bankUsecaseImpl) DeleteBank(ctx context.Context, id int) (err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "BankUsecase:DeleteBank")
	defer trace.Finish()

	repoFilter := domain.FilterBank{ID: &id}
	return uc.repoSQL.BankRepo().Delete(ctx, &repoFilter)
}
