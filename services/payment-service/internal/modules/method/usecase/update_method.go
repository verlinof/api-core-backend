package usecase

import (
	"context"

	"payment-service/internal/modules/method/domain"

	"github.com/golangid/candi/candishared"
	"github.com/golangid/candi/tracer"
)

func (uc *methodUsecaseImpl) UpdateMethod(ctx context.Context, data *domain.RequestMethod) (err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "MethodUsecase:UpdateMethod")
	defer trace.Finish()

	repoFilter := domain.FilterMethod{ID: &data.ID}
	existing, err := uc.repoSQL.MethodRepo().Find(ctx, &repoFilter)
	if err != nil {
		return err
	}
	existing.Code = data.Code
	existing.Name = data.Name
	existing.Description = data.Description
	existing.IconURL = data.IconURL
	existing.CategoryID = data.CategoryID
	existing.BankID = data.BankID
	existing.ProviderID = data.ProviderID
	existing.IsActive = data.IsActive
	err = uc.repoSQL.WithTransaction(ctx, func(ctx context.Context) error {
		return uc.repoSQL.MethodRepo().Save(ctx, &existing, candishared.DBUpdateSetUpdatedFields("Field"))
	})
	return
}
