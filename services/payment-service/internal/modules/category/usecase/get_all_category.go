package usecase

import (
	"context"

	"payment-service/internal/modules/category/domain"

	"github.com/golangid/candi/candishared"
	"github.com/golangid/candi/tracer"
)

func (uc *categoryUsecaseImpl) GetAllCategory(ctx context.Context, filter *domain.FilterCategory) (result domain.ResponseCategoryList, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "CategoryUsecase:GetAllCategory")
	defer trace.Finish()

	data, err := uc.repoSQL.CategoryRepo().FetchAll(ctx, filter)
	if err != nil {
		return result, err
	}
	count := uc.repoSQL.CategoryRepo().Count(ctx, filter)
	result.Meta = candishared.NewMeta(filter.Page, filter.Limit, count)

	result.Data = make([]domain.ResponseCategory, len(data))
	for i, detail := range data {
		result.Data[i].Serialize(&detail)
	}

	return
}
