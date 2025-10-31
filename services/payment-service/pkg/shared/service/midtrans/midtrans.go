package mid_trans

import (
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type midtransImpl struct {
	Env       string
	ServerKey string
}

type Midtrans interface {
	SnapCreateTransaction(req *snap.Request) (*snap.Response, error)
	SnapCreateTransactionToken(req *snap.Request) (string, error)
	SnapCreateTransactionUrl(req *snap.Request) (string, error)
}

func NewMidtransService(env string, serverKey string) Midtrans {
	return &midtransImpl{
		Env:       env,
		ServerKey: serverKey,
	}
}

func (m *midtransImpl) SnapCreateTransaction(req *snap.Request) (*snap.Response, error) {
	snapClient := snap.Client{}
	if m.Env == "production" {
		snapClient.New(m.ServerKey, midtrans.Production)
	} else {
		snapClient.New(m.ServerKey, midtrans.Sandbox)
	}
	return snapClient.CreateTransaction(req)
}

func (m *midtransImpl) SnapCreateTransactionToken(req *snap.Request) (string, error) {
	snapClient := snap.Client{}
	if m.Env == "production" {
		snapClient.New(m.ServerKey, midtrans.Production)
	} else {
		snapClient.New(m.ServerKey, midtrans.Sandbox)
	}
	return snapClient.CreateTransactionToken(req)
}

func (m *midtransImpl) SnapCreateTransactionUrl(req *snap.Request) (string, error) {
	snapClient := snap.Client{}
	if m.Env == "production" {
		snapClient.New(m.ServerKey, midtrans.Production)
	} else {
		snapClient.New(m.ServerKey, midtrans.Sandbox)
	}
	return snapClient.CreateTransactionUrl(req)
}
