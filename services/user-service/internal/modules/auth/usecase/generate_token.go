package usecase

import (
	"context"
	"time"

	"monorepo/services/user-service/internal/modules/auth/domain"

	"github.com/golang-jwt/jwt/v5"
	"github.com/golangid/candi/tracer"
)

func (uc *authUsecaseImpl) GenerateToken(ctx context.Context, userID string, ttl time.Duration) (tokenString string, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "AuthUsecase:GenerateToken")
	defer trace.Finish()

	claim := domain.CustomClaims{}
	claim.UserID = userID
	claim.IssuedAt = jwt.NewNumericDate(time.Now())
	claim.ExpiresAt = jwt.NewNumericDate(time.Now().Add(ttl))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err = token.SignedString([]byte(uc.env.JWTSecret))
	if err != nil {
		return tokenString, err
	}

	return tokenString, nil
}
