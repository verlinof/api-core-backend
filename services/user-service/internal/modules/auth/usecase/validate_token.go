package usecase

import (
	"context"
	"errors"
	"fmt"
	"monorepo/services/user-service/internal/modules/auth/domain"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/golangid/candi/tracer"
)

// Validate token
func (uc *authUsecaseImpl) ValidateToken(ctx context.Context, tokenString string) (claim *domain.CustomClaims, err error) {
	trace := tracer.StartTrace(ctx, "AuthUsecase:Validate")
	defer trace.Finish()
	ctx = trace.Context()

	tokenParse, err := jwt.ParseWithClaims(tokenString, &domain.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, domain.ErrUnexpectedSigningMethod
		}
		return []byte(uc.env.JWTSecret), nil
	})
	if err != nil {
		// v5 uses typed errors instead of ValidationError
		if errors.Is(err, jwt.ErrTokenMalformed) {
			return nil, domain.ErrTokenFormat
		} else if errors.Is(err, jwt.ErrTokenSignatureInvalid) {
			return nil, domain.ErrTokenSignatureInvalid
		} else if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, domain.ErrTokenExpired
		} else if errors.Is(err, jwt.ErrTokenNotValidYet) {
			return nil, domain.ErrTokenNotValidYet
		} else {
			return nil, domain.ErrTokenFormat
		}
	}

	// Check if token is valid
	claims, ok := tokenParse.Claims.(*domain.CustomClaims)
	if !ok || !tokenParse.Valid {
		return nil, domain.ErrTokenInvalid
	}

	// Check if token is expired
	if claims.ExpiresAt.Before(time.Now()) {
		return nil, domain.ErrTokenExpired
	}

	// Check not revoked
	redisKey := fmt.Sprintf("%s:%s", domain.RedisTokenConst, claims.UserID)
	_, errRedis := uc.cache.Get(ctx, redisKey)
	if errRedis != nil {
		return nil, domain.ErrTokenRevoked
	}

	return claims, nil

}
