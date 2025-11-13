package globalshared

// this file only for example

import (
	"context"

	"github.com/golangid/candi/candishared"
	"github.com/golangid/candi/logger"
	"github.com/golangid/candi/tracer"
)

// DefaultMiddleware for middleware validator example
type DefaultMiddleware struct {
}

// ValidateToken implement TokenValidator
func (DefaultMiddleware) ValidateToken(ctx context.Context, token string) (*candishared.TokenClaim, error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "DefaultMiddleware:ValidateToken")
	defer trace.Finish()

	var tokenClaim candishared.TokenClaim
	tokenClaim.Subject = "USER_ID"
	// tokenClaim.Additional = CustomUserTokenClaim{}

	logger.LogI("validate token: allowed")
	return &tokenClaim, nil
}

// CheckPermission implement interfaces.ACLPermissionChecker
func (DefaultMiddleware) CheckPermission(ctx context.Context, userID string, permissionCode string) (role string, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "DefaultMiddleware:CheckPermission")
	defer trace.Finish()

	/* add check allow permission for user access (is given "userID" can access "permissionCode" ?)
	if !contains(getAllPermissionFromUser(userID), permissionCode) {
		return role, errors.New("Forbidden")
	}
	*/
	logger.LogIf("check permission: users with id '%s' can access resource with permission code '%s' (return role for this user is 'superadmin')\n", userID, permissionCode)
	return "superadmin", nil
}
