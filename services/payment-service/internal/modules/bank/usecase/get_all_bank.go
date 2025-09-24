package usecase

import (
	"context"

	"payment-service/internal/modules/bank/domain"

	"github.com/golangid/candi/candishared"
	"github.com/golangid/candi/tracer"
)

func (uc *bankUsecaseImpl) GetAllBank(ctx context.Context, filter *domain.FilterBank) (result domain.ResponseBankList, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "BankUsecase:GetAllBank")
	defer trace.Finish()

	data, err := uc.repoSQL.BankRepo().FetchAll(ctx, filter)
	if err != nil {
		return result, err
	}
	count := uc.repoSQL.BankRepo().Count(ctx, filter)
	result.Meta = candishared.NewMeta(filter.Page, filter.Limit, count)

	result.Data = make([]domain.ResponseBank, len(data))
	for i, detail := range data {
		result.Data[i].Serialize(&detail)
	}

	return
}
