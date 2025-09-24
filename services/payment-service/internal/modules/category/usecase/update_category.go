package usecase

import (
	"context"

	"payment-service/internal/modules/category/domain"

	"github.com/golangid/candi/candishared"
	"github.com/golangid/candi/tracer"
)

func (uc *categoryUsecaseImpl) UpdateCategory(ctx context.Context, data *domain.RequestCategory) (err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "CategoryUsecase:UpdateCategory")
	defer trace.Finish()

	repoFilter := domain.FilterCategory{ID: &data.ID}
	existing, err := uc.repoSQL.CategoryRepo().Find(ctx, &repoFilter)
	if err != nil {
		return err
	}
	existing.Name = data.Name
	existing.Code = data.Code
	existing.IsActive = data.IsActive
	err = uc.repoSQL.WithTransaction(ctx, func(ctx context.Context) error {
		return uc.repoSQL.CategoryRepo().Save(ctx, &existing, candishared.DBUpdateSetUpdatedFields("Name", "Code", "IsActive"))
	})
	return
}
