package usecase

import (
	"context"

	"payment-service/internal/modules/provider/domain"

	"github.com/golangid/candi/candishared"
	"github.com/golangid/candi/tracer"
)

func (uc *providerUsecaseImpl) UpdateProvider(ctx context.Context, data *domain.RequestProvider) (err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "ProviderUsecase:UpdateProvider")
	defer trace.Finish()

	repoFilter := domain.FilterProvider{ID: &data.ID}
	existing, err := uc.repoSQL.ProviderRepo().Find(ctx, &repoFilter)
	if err != nil {
		return err
	}
	existing.Name = data.Name
	existing.Code = data.Code
	existing.IsActive = data.IsActive
	err = uc.repoSQL.WithTransaction(ctx, func(ctx context.Context) error {
		return uc.repoSQL.ProviderRepo().Save(ctx, &existing, candishared.DBUpdateSetUpdatedFields("Name", "Code", "IsActive"))
	})
	return
}
