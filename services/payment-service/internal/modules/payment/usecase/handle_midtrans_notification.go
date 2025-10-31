package usecase

import (
	"context"
	"encoding/json"
	"log"
	"payment-service/internal/modules/payment/domain"
	shareddomain "payment-service/pkg/shared/domain"
	"strconv"

	"github.com/golangid/candi/tracer"
)

func (uc *paymentUsecaseImpl) HandleMidtransNotification(ctx context.Context, notif domain.MidtransNotification) error {
	trace, ctx := tracer.StartTraceWithContext(ctx, "PaymentUsecase:HandleMidtransNotification")
	defer trace.Finish()
	log.Printf("Received notification for Order ID: %s, Status: %s",
		notif.OrderID, notif.TransactionStatus)

	// Verify signature
	res, midtransErr := uc.service.Midtrans().CoreCheckTransaction(notif.OrderID)
	if midtransErr != nil {
		log.Printf("[Webhook] Error checking transaction status: %v\n", midtransErr)
		return nil
	}

	// Cari order di database
	order, err := uc.repoSQL.PaymentRepo().FindByOrderID(ctx, res.OrderID)
	if err != nil {
		log.Printf("[Webhook] Error finding order %s: %v\n", res.OrderID, err)
		return nil
	}

	var orderStatus string

	switch res.TransactionStatus {
	case domain.TransactionStatusSettlement:
		orderStatus = "paid"
	case domain.TransactionStatusPending:
		orderStatus = "pending"
	case domain.TransactionStatusDeny:
		orderStatus = "denied"
	case domain.TransactionStatusExpire:
		orderStatus = "expired"
	case domain.TransactionStatusCancel:
		orderStatus = "cancelled"
	case domain.TransactionStatusFailure:
		orderStatus = "failed"
	default:
		return nil
	}

	// Update status order di database
	order.Status = orderStatus
	order.TransactionStatus = res.TransactionStatus
	order.FraudStatus = res.FraudStatus
	if err := uc.repoSQL.PaymentRepo().Save(ctx, order); err != nil {
		log.Printf("[Webhook] Error updating order status for OrderID %s: %v\n", order.OrderID, err)
		return nil
	}

	// Simpan log notifikasi
	reqJSON, _ := json.Marshal(notif)
	respJSON, _ := json.Marshal(res)
	statusCode, _ := strconv.Atoi(res.StatusCode)

	logEntry := &shareddomain.PaymentLog{
		OrderID:    order.ID,
		StatusCode: statusCode,
		Request:    string(reqJSON),
		Response:   string(respJSON),
	}

	if err := uc.repoSQL.PaymentRepo().SaveLog(ctx, logEntry, *order.MethodID); err != nil {
		log.Printf("[Webhook] Error saving notification log for OrderID %s: %v\n", order.OrderID, err)
	}

	log.Printf("Order %s status updated to: %s", notif.OrderID, orderStatus)

	return nil
}
