package service

import (
	"payment-service/pkg/shared"
	mid_trans "payment-service/pkg/shared/service/midtrans"
	"payment-service/pkg/shared/service/notifier"
)

const KeyExternalService string = "midtrans-external-service"

type serviceImpl struct {
	midtransService mid_trans.Midtrans
	notifierService notifier.Notifier
}

type Service interface {
	Midtrans() mid_trans.Midtrans
	Notifier() notifier.Notifier
}

func New() Service {
	return &serviceImpl{
		midtransService: mid_trans.NewMidtransService(
			shared.GetEnv().MidtransEnv, shared.GetEnv().MidtransServerKey,
		),
		notifierService: notifier.NewNotifierService(),
	}
}

func (s *serviceImpl) Midtrans() mid_trans.Midtrans {
	return s.midtransService
}

func (s *serviceImpl) Notifier() notifier.Notifier {
	return s.notifierService
}
