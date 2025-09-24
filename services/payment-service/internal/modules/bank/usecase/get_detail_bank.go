package usecase

import (
	"context"

	"payment-service/internal/modules/bank/domain"

	"github.com/golangid/candi/tracer"
)

func (uc *bankUsecaseImpl) GetDetailBank(ctx context.Context, id int) (result domain.ResponseBank, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "BankUsecase:GetDetailBank")
	defer trace.Finish()

	repoFilter := domain.FilterBank{ID: &id}
	data, err := uc.repoSQL.BankRepo().Find(ctx, &repoFilter)
	if err != nil {
		return result, err
	}

	result.Serialize(&data)
	return
}
