package usecase

import (
	"context"
	
	"payment-service/internal/modules/provider/domain"

	"github.com/golangid/candi/tracer"
)

func (uc *providerUsecaseImpl) DeleteProvider(ctx context.Context, id int) (err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "ProviderUsecase:DeleteProvider")
	defer trace.Finish()

	repoFilter := domain.FilterProvider{ID: &id}
	return uc.repoSQL.ProviderRepo().Delete(ctx, &repoFilter)
}
