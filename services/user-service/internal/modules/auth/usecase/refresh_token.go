package usecase

import (
	"context"
	"fmt"
	"monorepo/services/user-service/internal/modules/auth/domain"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/golangid/candi/tracer"
)

func (uc *authUsecaseImpl) RefreshToken(ctx context.Context, token, refreshToken string) (resp *domain.ResponseGenerateToken, err error) {
	trace := tracer.StartTrace(ctx, "AuthUsecase:RefreshToken")
	defer trace.Finish()
	ctx = trace.Context()

	now := time.Now()
	refreshTokenParse, err := jwt.ParseWithClaims(refreshToken, &domain.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, domain.ErrUnexpectedSigningMethod
		}
		return []byte(uc.env.JWTSecret), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := refreshTokenParse.Claims.(*domain.CustomClaims)
	if !ok || !refreshTokenParse.Valid {
		return nil, domain.ErrTokenInvalid
	}

	if claims.ExpiresAt.Before(now) {
		return nil, domain.ErrTokenExpired
	}

	// generate new token
	newToken, err := uc.GenerateToken(ctx, claims.UserID, uc.env.JWTAccessTTL)
	if err != nil {
		return nil, err
	}

	redisKey := fmt.Sprintf("%s:%s", domain.RedisTokenConst, claims.UserID)
	uc.cache.Set(ctx, redisKey, newToken, uc.env.JWTAccessTTL)

	resp = &domain.ResponseGenerateToken{
		Token:        newToken,
		RefreshToken: refreshToken,
	}

	return resp, nil
}
