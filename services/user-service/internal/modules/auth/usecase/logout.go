package usecase

import (
	"context"
	"fmt"
	"monorepo/services/user-service/internal/modules/auth/domain"

	"github.com/golang-jwt/jwt/v5"
)

func (uc *authUsecaseImpl) Logout(ctx context.Context, token string) error {
	uc.ValidateToken(ctx, token)
	tokenParse, err := jwt.ParseWithClaims(token, &domain.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, domain.ErrUnexpectedSigningMethod
		}
		return []byte(uc.env.JWTSecret), nil
	})

	if err != nil {
		return err
	}

	claims, ok := tokenParse.Claims.(*domain.CustomClaims)
	if !ok || !tokenParse.Valid {
		return domain.ErrTokenInvalid
	}

	redisKey := fmt.Sprintf("%s:%s", domain.RedisTokenConst, claims.UserID)
	err = uc.cache.Delete(ctx, redisKey)
	if err != nil {
		return err
	}

	return nil
}
