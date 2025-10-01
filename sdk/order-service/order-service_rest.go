package orderservice

import (
	"net/http"
	"time"

	"github.com/golangid/candi/candiutils"
)

type orderserviceRESTImpl struct {
	host    string
	authKey string
	httpReq candiutils.HTTPRequest
}

// NewOrderserviceServiceREST constructor
func NewOrderserviceServiceREST(host string, authKey string) Orderservice {

	return &orderserviceRESTImpl{
		host:    host,
		authKey: authKey,
		httpReq: candiutils.NewHTTPRequest(
			candiutils.HTTPRequestSetRetries(5),
			candiutils.HTTPRequestSetSleepBetweenRetry(500*time.Millisecond),
			candiutils.HTTPRequestSetHTTPErrorCodeThreshold(http.StatusBadRequest),
			candiutils.HTTPRequestSetBreakerName("orderservice"),
		),
	}
}
