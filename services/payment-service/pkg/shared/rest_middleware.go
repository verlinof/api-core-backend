package shared

import (
	"context"
	"net/http"
	"strings"

	"github.com/golangid/candi/candihelper"
	"github.com/golangid/candi/candishared"
	"github.com/golangid/candi/tracer"
	"github.com/golangid/candi/wrapper"
)

const (
	UnauthorizedUser = "Unauthorized user"
)

func HttpCustomBearerAuthMiddleware(tokenValidator interface {
	ValidateToken(ctx context.Context, token string) (*candishared.TokenClaim, error)
}) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			ctx := req.Context()

			// get auth
			authorization := req.Header.Get(candihelper.HeaderAuthorization)
			if authorization == "" {
				wrapper.NewHTTPResponse(http.StatusUnauthorized, UnauthorizedUser).JSON(rw)
				return
			}

			// get auth type
			authValues := strings.Split(authorization, " ")

			// validate value
			if len(authValues) != 2 || strings.ToLower(authValues[0]) != "bearer" {
				wrapper.NewHTTPResponse(http.StatusUnauthorized, "Invalid authorization type").JSON(rw)
				return
			}

			tokenClaim, err := tokenValidator.ValidateToken(ctx, authValues[1])
			if err != nil {
				tracer.SetError(ctx, err)
				wrapper.NewHTTPResponse(http.StatusUnauthorized, UnauthorizedUser).JSON(rw)
				return
			}
			tracer.Log(ctx, "token_claim", tokenClaim)
			ctx = candishared.SetToContext(ctx, candishared.ContextKeyTokenClaim, tokenClaim)

			next.ServeHTTP(rw, req.WithContext(ctx))
		})
	}
}

// HTTPCustomHTTPMultipleAuthFromCheckerMiddleware echo middleware
func HTTPCustomHTTPMultipleAuthFromCheckerMiddleware(multiAuthChecker interface {
	IsBasicAuthAllowed(ctx context.Context, username, password string) bool
	ValidateToken(ctx context.Context, token string) (*candishared.TokenClaim, error)
}) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			ctx := req.Context()

			// get auth
			authorization := req.Header.Get(candihelper.HeaderAuthorization)
			if authorization == "" {
				wrapper.NewHTTPResponse(http.StatusUnauthorized, UnauthorizedUser).JSON(rw)
				return
			}

			// get auth type
			authValues := strings.Split(authorization, " ")

			// validate value
			if len(authValues) != 2 {
				wrapper.NewHTTPResponse(http.StatusUnauthorized, "Invalid authorization type").JSON(rw)
				return
			}

			authType := strings.ToLower(authValues[0])
			if authType == "basic" {
				basicUser, basicPass, ok := req.BasicAuth()
				if !ok {
					wrapper.NewHTTPResponse(http.StatusUnauthorized, UnauthorizedUser).JSON(rw)
					return
				}
				if ok = multiAuthChecker.IsBasicAuthAllowed(ctx, basicUser, basicPass); !ok {
					wrapper.NewHTTPResponse(http.StatusUnauthorized, UnauthorizedUser).JSON(rw)
					return
				}
				ctx = candishared.SetToContext(ctx, candishared.ContextKeyTokenClaim, &candishared.TokenClaim{
					Role: "public",
				})
			} else if authType == "bearer" {
				tokenClaim, err := multiAuthChecker.ValidateToken(ctx, authValues[1])
				if err != nil {
					tracer.SetError(ctx, err)
					wrapper.NewHTTPResponse(http.StatusUnauthorized, UnauthorizedUser).JSON(rw)
					return
				}
				tracer.Log(ctx, "token_claim", tokenClaim)
				ctx = candishared.SetToContext(ctx, candishared.ContextKeyTokenClaim, tokenClaim)
			} else {
				wrapper.NewHTTPResponse(http.StatusUnauthorized, "Invalid authorization type").JSON(rw)
				return
			}

			next.ServeHTTP(rw, req.WithContext(ctx))
		})
	}
}
