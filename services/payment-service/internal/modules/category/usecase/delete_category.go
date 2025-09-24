package usecase

import (
	"context"
	
	"payment-service/internal/modules/category/domain"

	"github.com/golangid/candi/tracer"
)

func (uc *categoryUsecaseImpl) DeleteCategory(ctx context.Context, id int) (err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "CategoryUsecase:DeleteCategory")
	defer trace.Finish()

	repoFilter := domain.FilterCategory{ID: &id}
	return uc.repoSQL.CategoryRepo().Delete(ctx, &repoFilter)
}
