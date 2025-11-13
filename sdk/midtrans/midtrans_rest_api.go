package midtrans

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"monorepo/globalshared"
	"net/http"
	"reflect"

	"github.com/golangid/candi/candihelper"
	"github.com/golangid/candi/candiutils"
	"github.com/golangid/candi/wrapper"
)

const (
	serviceName = "midtrans"
)

// MidTransRestApi data structure for REST API service
type midtransRestApi struct {
	host        string
	clientKey   string
	serverKey   string
	httpRequest candiutils.HTTPRequest
}

func NewMidtransRestApi(host string, serverKey string, clientKey string, options ...candiutils.HTTPRequestOption) Midtrans {
	options = append(options, candiutils.HTTPRequestSetBreakerName(serviceName))
	return &midtransRestApi{
		host:        host,
		clientKey:   clientKey,
		serverKey:   serverKey,
		httpRequest: candiutils.NewHTTPRequest(options...),
	}
}

func (m *midtransRestApi) ChargeSnap(ctx context.Context, req *SnapRequest) (resp SnapResponse, err error) {
	// implement charge snap via REST API if needed
	token := "Basic " + base64.StdEncoding.EncodeToString([]byte(m.serverKey+":"))
	headers := map[string]string{
		candihelper.HeaderContentType:   candihelper.HeaderMIMEApplicationJSON,
		candihelper.HeaderAuthorization: token,
	}
	respBody, code, err := m.httpRequest.Do(ctx, http.MethodPost, m.host+"/snap/v1/transactions", candihelper.ToBytes(req), headers)

	if err != nil {
		return resp, &globalshared.ErrorResponse{
			Code:       code,
			Message:    err.Error(),
			MultiError: candihelper.NewMultiError().Append(serviceName, err),
		}
	}
	var result wrapper.HTTPResponse
	result.Data = new(SnapResponse)
	if err := json.Unmarshal(respBody, &result); err != nil {
		return resp, err
	}
	resp = reflect.ValueOf(result.Data).Elem().Interface().(SnapResponse)
	return
}

func (m *midtransRestApi) GetStatus(ctx context.Context, orderID string) (resp NotificationPayload, err error) {
	token := "Basic " + base64.StdEncoding.EncodeToString([]byte(m.serverKey+":"))
	headers := map[string]string{
		candihelper.HeaderContentType:   candihelper.HeaderMIMEApplicationJSON,
		candihelper.HeaderAuthorization: token,
	}
	path := m.host + fmt.Sprintf("/v2/%s/status", orderID)
	respBody, code, err := m.httpRequest.Do(ctx, http.MethodGet, path, nil, headers)
	if err != nil {
		err = &globalshared.ErrorResponse{
			Code:       code,
			Message:    err.Error(),
			MultiError: candihelper.NewMultiError().Append(serviceName, err),
		}
		var result wrapper.HTTPResponse
		result.Data = new(SnapResponse)
		if err := json.Unmarshal(respBody, &result); err != nil {
			return resp, err
		}
		// resp = reflect.ValueOf(result.Data).Elem().Interface().(SnapResponse)
		// return resp
	}
	return
}
