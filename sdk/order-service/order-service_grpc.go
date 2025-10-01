package orderservice

import (
	"net/url"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/backoff"
)

type orderserviceGRPCImpl struct {
	host    string
	authKey string
	conn    *grpc.ClientConn
}

// NewOrderserviceServiceGRPC constructor
func NewOrderserviceServiceGRPC(host string, authKey string) Orderservice {

	if u, _ := url.Parse(host); u.Host != "" {
		host = u.Host
	}
	conn, err := grpc.Dial(host, grpc.WithInsecure(), grpc.WithConnectParams(grpc.ConnectParams{
		Backoff: backoff.Config{
			BaseDelay:  50 * time.Millisecond,
			Multiplier: 5,
			MaxDelay:   50 * time.Millisecond,
		},
		MinConnectTimeout: 1 * time.Second,
	}))
	if err != nil {
		panic(err)
	}

	return &orderserviceGRPCImpl{
		host:    host,
		authKey: authKey,
		conn:    conn,
	}
}
