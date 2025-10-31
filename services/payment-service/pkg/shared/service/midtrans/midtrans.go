package mid_trans

import (
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/midtrans/midtrans-go/snap"
)

type midtransImpl struct {
	Env       string
	ServerKey string
}

type Midtrans interface {
	SnapCreateTransaction(req *snap.Request) (*snap.Response, *midtrans.Error)
	SnapCreateTransactionToken(req *snap.Request) (string, *midtrans.Error)
	SnapCreateTransactionUrl(req *snap.Request) (string, *midtrans.Error)
	CoreCheckTransaction(orderID string) (*coreapi.TransactionStatusResponse, *midtrans.Error)
}

func NewMidtransService(env string, serverKey string) Midtrans {
	return &midtransImpl{
		Env:       env,
		ServerKey: serverKey,
	}
}

func (m *midtransImpl) SnapCreateTransaction(req *snap.Request) (*snap.Response, *midtrans.Error) {
	snapClient := snap.Client{}
	if m.Env == "production" {
		snapClient.New(m.ServerKey, midtrans.Production)
	} else {
		snapClient.New(m.ServerKey, midtrans.Sandbox)
	}
	return snapClient.CreateTransaction(req)
}

func (m *midtransImpl) SnapCreateTransactionToken(req *snap.Request) (string, *midtrans.Error) {
	snapClient := snap.Client{}
	if m.Env == "production" {
		snapClient.New(m.ServerKey, midtrans.Production)
	} else {
		snapClient.New(m.ServerKey, midtrans.Sandbox)
	}
	return snapClient.CreateTransactionToken(req)
}

func (m *midtransImpl) SnapCreateTransactionUrl(req *snap.Request) (string, *midtrans.Error) {
	snapClient := snap.Client{}
	if m.Env == "production" {
		snapClient.New(m.ServerKey, midtrans.Production)
	} else {
		snapClient.New(m.ServerKey, midtrans.Sandbox)
	}
	return snapClient.CreateTransactionUrl(req)
}

func (m *midtransImpl) CoreCheckTransaction(orderID string) (*coreapi.TransactionStatusResponse, *midtrans.Error) {
	coreClient := coreapi.Client{}
	if m.Env == "production" {
		coreClient.New(m.ServerKey, midtrans.Production)
	} else {
		coreClient.New(m.ServerKey, midtrans.Sandbox)
	}
	return coreClient.CheckTransaction(orderID)
}
