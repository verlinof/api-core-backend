package userservice

import (
	"context"
	"encoding/json"
	"fmt"
	"monorepo/globalshared"
	"net/http"
	"reflect"
	"time"

	"github.com/golangid/candi/candihelper"
	"github.com/golangid/candi/candiutils"
	"github.com/golangid/candi/tracer"
	"github.com/golangid/candi/wrapper"
)

const (
	serviceName = "userservice"
)

type userserviceRESTImpl struct {
	host    string
	authKey string
	httpReq candiutils.HTTPRequest
}

// NewUserserviceServiceREST constructor
func NewUserserviceServiceREST(host string) Userservice {

	return &userserviceRESTImpl{
		host: host,
		httpReq: candiutils.NewHTTPRequest(
			candiutils.HTTPRequestSetRetries(5),
			candiutils.HTTPRequestSetSleepBetweenRetry(500*time.Millisecond),
			candiutils.HTTPRequestSetHTTPErrorCodeThreshold(http.StatusBadRequest),
			candiutils.HTTPRequestSetBreakerName("userservice"),
		),
	}
}

func (s *userserviceRESTImpl) VerifyToken(ctx context.Context, token string) (resp ClaimResponse, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "Userservice:VerifyToken")
	defer func() {
		trace.Finish(tracer.FinishWithError(err))
	}()
	url := fmt.Sprintf("%s/v1/auth/token/verify", s.host)
	headers := map[string]string{
		candihelper.HeaderContentType: candihelper.HeaderMIMEApplicationJSON,
	}
	reqBody := map[string]string{
		"token": token,
	}
	respBody, _, err := s.httpReq.Do(ctx, http.MethodPost, url, candihelper.ToBytes(reqBody), headers)
	if err != nil {
		return resp, &globalshared.ErrorResponse{
			Code:       http.StatusInternalServerError,
			Message:    fmt.Sprintf("Something error when get data from %s", serviceName),
			MultiError: candihelper.NewMultiError().Append(serviceName, err),
		}
	}

	var respService wrapper.HTTPResponse
	respService.Data = new(ClaimResponse)
	if err := json.Unmarshal(respBody, &respService); err != nil {
		return resp, err
	}

	resp = reflect.ValueOf(respService.Data).Elem().Interface().(ClaimResponse)

	return
}
