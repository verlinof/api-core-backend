package usecase

import (
	"context"

	"payment-service/internal/modules/category/domain"

	"github.com/golangid/candi/tracer"
)

func (uc *categoryUsecaseImpl) GetDetailCategory(ctx context.Context, id int) (result domain.ResponseCategory, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "CategoryUsecase:GetDetailCategory")
	defer trace.Finish()

	repoFilter := domain.FilterCategory{ID: &id}
	data, err := uc.repoSQL.CategoryRepo().Find(ctx, &repoFilter)
	if err != nil {
		return result, err
	}

	result.Serialize(&data)
	return
}
