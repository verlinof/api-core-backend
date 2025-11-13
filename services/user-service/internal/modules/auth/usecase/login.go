package usecase

import (
	"context"
	"fmt"
	"monorepo/services/user-service/internal/modules/auth/domain"
	user_domain "monorepo/services/user-service/internal/modules/user/domain"
	"strconv"

	"github.com/golangid/candi/tracer"
)

func (uc *authUsecaseImpl) Login(ctx context.Context, email, password string) (resp domain.ResponseLogin, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "AuthUsecase:Login")
	defer trace.Finish()

	// check email
	userFilter := user_domain.FilterUser{Email: &email}
	user, err := uc.repoSQL.UserRepo().Find(ctx, &userFilter)
	if err != nil {
		return resp, domain.ErrUserNotFound
	}

	// compare password
	if !user.CheckPassword(password) {
		return resp, domain.ErrInvalidCredentials
	}

	// Generate JWT token
	accessTokenString, err := uc.GenerateToken(ctx, strconv.Itoa(user.ID), uc.env.JWTAccessTTL)
	if err != nil {
		return resp, err
	}

	refreshTokenString, err := uc.GenerateToken(ctx, strconv.Itoa(user.ID), uc.env.JWTRefreshTTL)
	if err != nil {
		return resp, err
	}
	resp.RefreshToken = refreshTokenString
	resp.UserID = strconv.Itoa(user.ID)
	resp.Token = accessTokenString

	redisKey := fmt.Sprintf("%s:%d", domain.RedisTokenConst, user.ID)
	uc.cache.Set(ctx, redisKey, refreshTokenString, uc.env.JWTRefreshTTL)

	return
}
