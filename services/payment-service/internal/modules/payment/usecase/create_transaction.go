package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"payment-service/internal/modules/payment/domain"
	pkg_midtrans "payment-service/pkg/helper/midtrans"
	"strconv"
	"time"

	shareddomain "payment-service/pkg/shared/domain"

	"github.com/golangid/candi/tracer"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

func (uc *paymentUsecaseImpl) CreateTransaction(ctx context.Context, req *domain.CreateOrderRequest) (*domain.CreateTransactionResponse, error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "PaymentUsecase:CreateTransaction")
	defer trace.Finish()

	// 1. Siapkan data Order
	orderDataJSON, _ := json.Marshal(req)
	order := shareddomain.PaymentOrder{
		OrderID:   req.OrderID,
		Amount:    float64(req.Amount),
		Status:    "PENDING",
		Channel:   req.Channel,
		MethodID:  req.MethodID,
		OrderData: string(orderDataJSON),
	}

	// 2. Siapkan data untuk request ke Midtrans
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

	if req.MethodID != nil {
		selectedPaymentType, ok := domain.PaymentMethodMap[*req.MethodID]
		if !ok {
			return nil, fmt.Errorf("invalid payment method ID: %d", *req.MethodID)
		}
		snapReq.EnabledPayments = []snap.SnapPaymentType{selectedPaymentType}
	}

	// 3. Panggil API Midtrans
	snapResp, midtransErr := pkg_midtrans.SnapClient.CreateTransaction(snapReq)
	if midtransErr != nil {
		// Log for Error
		go uc.createPaymentLog(context.Background(), &order, snapReq, midtransErr, snapResp)
		return nil, fmt.Errorf("midtrans error: %s", midtransErr.Message)
	}

	// Save DB order
	if err := uc.repoSQL.PaymentRepo().Save(ctx, &order); err != nil {
		return nil, fmt.Errorf("failed to save order: %w", err)
	}
	// Log if success
	go uc.createPaymentLog(context.Background(), &order, snapReq, snapResp, snapResp)

	response := &domain.CreateTransactionResponse{
		Token:       snapResp.Token,
		RedirectURL: snapResp.RedirectURL,
		OrderID:     req.OrderID,
	}
	return response, nil
}

// createPaymentLogWithRepo adalah inti logika pembuatan log yang menerima interface repository.
func (uc *paymentUsecaseImpl) createPaymentLog(ctx context.Context, order *shareddomain.PaymentOrder, reqPayload, respPayload interface{}, snapRes *snap.Response) error {
	reqJSON, _ := json.Marshal(reqPayload)
	respJSON, _ := json.Marshal(respPayload)

	statusCode, _ := strconv.Atoi(snapRes.StatusCode)

	logEntry := &shareddomain.PaymentLog{
		OrderID:    order.ID,
		StatusCode: statusCode,
		Request:    string(reqJSON),
		Response:   string(respJSON),
		PaymentURL: snapRes.RedirectURL,
	}

	return uc.repoSQL.PaymentRepo().SaveLog(ctx, logEntry, *order.MethodID)
}
