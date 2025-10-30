package service

import (
	"payment-service/pkg/shared"
	mid_trans "payment-service/pkg/shared/service/midtrans"
)

const KeyExternalService string = "midtrans-external-service"

type serviceImpl struct {
	midtransService mid_trans.Midtrans
}

type Service interface {
	Midtrans() mid_trans.Midtrans
}

func New() Service {
	return &serviceImpl{
		midtransService: mid_trans.NewMidtransService(
			shared.GetEnv().MidtransEnv, shared.GetEnv().MidtransServerKey,
		),
	}
}

func (s *serviceImpl) Midtrans() mid_trans.Midtrans {
	return s.midtransService
}
