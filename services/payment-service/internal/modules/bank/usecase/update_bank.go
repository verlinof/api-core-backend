package usecase

import (
	"context"

	"payment-service/internal/modules/bank/domain"

	"github.com/golangid/candi/candishared"
	"github.com/golangid/candi/tracer"
)

func (uc *bankUsecaseImpl) UpdateBank(ctx context.Context, data *domain.RequestBank) (err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "BankUsecase:UpdateBank")
	defer trace.Finish()

	repoFilter := domain.FilterBank{ID: &data.ID}
	existing, err := uc.repoSQL.BankRepo().Find(ctx, &repoFilter)
	if err != nil {
		return err
	}
	existing.Name = data.Name
	existing.Code = data.Code
	existing.IsActive = data.IsActive
	err = uc.repoSQL.WithTransaction(ctx, func(ctx context.Context) error {
		return uc.repoSQL.BankRepo().Save(ctx, &existing, candishared.DBUpdateSetUpdatedFields("Name", "Code", "IsActive"))
	})
	return
}
