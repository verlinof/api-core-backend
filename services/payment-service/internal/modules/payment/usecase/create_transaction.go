package usecase

import (
	"context"
	"fmt"
	"payment-service/internal/modules/payment/domain"
	pkg_midtrans "payment-service/pkg/helper/midtrans"
	"time"

	"github.com/golangid/candi/tracer"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

func (uc *paymentUsecaseImpl) CreateTransaction(ctx context.Context, req *domain.CreateOrderRequest) (*domain.CreateTransactionResponse, error) {
	trace, _ := tracer.StartTraceWithContext(ctx, "PaymentUsecase:CreateTransaction")
	defer trace.Finish()

	// TODO: Di aplikasi nyata, kamu harus:
	// 1. Buat order di database-mu dengan status "PENDING".
	// 2. Gunakan OrderID dari database-mu (atau buat OrderID unik).
	// 3. Ambil total Amount dari databasemu, jangan percaya Amount dari frontend.

	snapReq := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  req.OrderID,
			GrossAmt: req.Amount,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: req.Customer.FirstName,
			Email: req.Customer.Email,
			Phone: req.Customer.Phone,
		},
		Expiry: &snap.ExpiryDetails{
			StartTime: time.Now().Format(time.RFC3339),
			Unit:      "hour",
			Duration:  24,
		},
	}

	snapResp, midtransErr := pkg_midtrans.SnapClient.CreateTransaction(snapReq)
	if midtransErr != nil {
		return nil, fmt.Errorf("payment error: %s", midtransErr.Message)
	}

	// Buat response yang terstruktur
	response := &domain.CreateTransactionResponse{
		Token:       snapResp.Token,
		RedirectURL: snapResp.RedirectURL,
		OrderID:     req.OrderID,
	}

	return response, nil
}
