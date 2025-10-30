package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"payment-service/internal/modules/payment/domain"
	pkg_midtrans "payment-service/pkg/helper/midtrans"
	"time"

	shareddomain "payment-service/pkg/shared/domain"

	"github.com/golangid/candi/tracer"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

func (uc *paymentUsecaseImpl) CreateTransaction(ctx context.Context, req *domain.CreateOrderRequest) (*domain.CreateTransactionResponse, error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "PaymentUsecase:CreateTransaction")
	defer trace.Finish()

	// 1. Create and save the order to the database with "PENDING" status.
	orderDataJSON, _ := json.Marshal(req)
	order := shareddomain.PaymentOrder{
		OrderID:   req.OrderID,
		Amount:    float64(req.Amount),
		Status:    "PENDING",
		Channel:   req.Channel,
		MethodID:  req.MethodID,
		OrderData: string(orderDataJSON),
	}

	// Map items from request to midtrans.ItemDetails
	var itemDetails []midtrans.ItemDetails
	for _, item := range req.Items {
		itemDetails = append(itemDetails, midtrans.ItemDetails{
			ID:    item.ID,
			Name:  item.Name,
			Price: item.Price,
			Qty:   item.Qty,
		})
	}

	snapReq := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  order.OrderID,
			GrossAmt: int64(order.Amount),
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: req.Customer.FirstName,
			LName: req.Customer.LastName,
			Email: req.Customer.Email,
			Phone: req.Customer.Phone,
		},
		Items: &itemDetails,
		Expiry: &snap.ExpiryDetails{
			StartTime: time.Now().Format("2006-01-02 15:04:05 -0700"),
			Unit:      "hour",
			Duration:  1,
		},
	}

	// hit midtrans
	snapResp, midtransErr := pkg_midtrans.SnapClient.CreateTransaction(snapReq)
	if midtransErr != nil {
		return nil, fmt.Errorf("payment error: %s", midtransErr.Message)
	}

	// Save DB order
	if err := uc.repoSQL.PaymentRepo().Save(ctx, &order); err != nil {
		return nil, fmt.Errorf("failed to save order: %w", err)
	}

	// Buat response yang terstruktur
	response := &domain.CreateTransactionResponse{
		Token:       snapResp.Token,
		RedirectURL: snapResp.RedirectURL,
		OrderID:     req.OrderID,
	}

	return response, nil
}
