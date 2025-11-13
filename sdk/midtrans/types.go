package midtrans

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"time"
)

// PaymentType type
type PaymentType string

const (
	VABankTransferType PaymentType = "bank_transfer"
	EchannelType       PaymentType = "echannel" // Mandiri Bill Payment
	PermataVaType      PaymentType = "permata_va"
	BcaVaType          PaymentType = "bca_va"
	BniVaType          PaymentType = "bni_va"
	BriVaType          PaymentType = "bri_va"
	CimbVaType         PaymentType = "cimb_va"
	DanamonVaType      PaymentType = "danamon_va"
	BsiVaType          PaymentType = "bsi_va"
	OtherVaType        PaymentType = "other_va"

	OtherQrisType PaymentType = "other_qris"
)

// TransactionDetails data structure
type TransactionDetails struct {
	OrderID     string `json:"order_id"`
	GrossAmount int    `json:"gross_amount"`
}

// ItemDetail data structure
type ItemDetails struct {
	ID       string `json:"id"`
	Price    int    `json:"price"`
	Quantity int    `json:"quantity"`
	Name     string `json:"name"`
	Type     string `json:"-"` // this type only for determining item, fee, or discount
}

// CustomerDetails data structure
type CustomerDetails struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone,omitempty"`
}

type Expiry struct {
	Unit     string `json:"unit,omitempty"`
	Duration int    `json:"duration,omitempty"`
}

type SnapRequest struct {
	TransactionDetails *TransactionDetails `json:"transaction_details,omitempty"`
	ItemDetails        []ItemDetails       `json:"item_details,omitempty"`
	CustomerDetails    *CustomerDetails    `json:"customer_details,omitempty"`
	EnabledPayments    []string            `json:"enabled_payments,omitempty"`
	CreditCard         *CreditCardSnap     `json:"credit_card,omitempty"`
	Expiry             *Expiry             `json:"expiry,omitempty"`
	VendorID           int                 `json:"vendorId,omitempty"`
	ServerKey          string              `json:"-"`
}

type Installment struct {
	Required bool   `json:"required,omitempty"`
	Terms    *Terms `json:"terms,omitempty"`
}

type Terms struct {
	Bca     []int `json:"bca,omitempty"`
	Bri     []int `json:"bri,omitempty"`
	Bni     []int `json:"bni,omitempty"`
	Cimb    []int `json:"cimb,omitempty"`
	Mandiri []int `json:"mandiri,omitempty"`
}

type BankRouting struct {
	Bin *Bin `json:"bin,omitempty"`
}

type Bin struct {
	Bca     []string `json:"bca,omitempty"`
	Bri     []string `json:"bri,omitempty"`
	Bni     []string `json:"bni,omitempty"`
	Cimb    []string `json:"cimb,omitempty"`
	Mandiri []string `json:"mandiri,omitempty"`
}

type CreditCardSnap struct {
	Secure          bool         `json:"secure,omitempty"`
	Bank            string       `json:"bank,omitempty"`
	Installment     *Installment `json:"installment,omitempty"`
	BankRouting     *BankRouting `json:"bank_routing,omitempty"`
	WhitelistBins   []string     `json:"whitelist_bins,omitempty"`
	InstallmentTerm int          `json:"installment_term,omitempty"`
}

type SnapResponse struct {
	Token             string   `json:"token,omitempty"`
	RedirectURL       string   `json:"redirect_url,omitempty"`
	ResponseTime      string   `json:"response_time"`
	TransactionStatus string   `json:"transaction_status"`
	FraudStatus       string   `json:"fraud_status"`
	StatusCode        int      `json:"status_code"`
	ErrorMessages     []string `json:"error_messages,omitempty"`
}

type VaNumbers struct {
	Bank     string `json:"bank"`
	VaNumber string `json:"va_number"`
}

type NotificationPayload struct {
	MaskedCard            string      `json:"masked_card"`
	ApprovalCode          string      `json:"approval_code"`
	Bank                  string      `json:"bank"`
	Eci                   string      `json:"eci"`
	TransactionTimeString string      `json:"transaction_time"`
	TransactionTime       time.Time   `json:"-"`
	SettlementTimeString  string      `json:"settlement_time"`
	SettlementTime        time.Time   `json:"-"`
	GrossAmount           string      `json:"gross_amount"`
	OrderID               string      `json:"order_id"`
	PaymentType           string      `json:"payment_type"`
	SignatureKey          string      `json:"signature_key"`
	StatusCode            string      `json:"status_code"`
	TransactionID         string      `json:"transaction_id"`
	TransactionStatus     string      `json:"transaction_status"`
	FraudStatus           string      `json:"fraud_status"`
	StatusMessage         string      `json:"status_message"`
	PermataVaNumber       string      `json:"permata_va_number"`
	VaNumbers             []VaNumbers `json:"va_numbers"`
}

func (n NotificationPayload) IsValidSignature(serverKey string) bool {

	input := fmt.Sprintf("%s%s%s%s", n.OrderID, n.StatusCode, n.GrossAmount, serverKey)

	mySha512 := sha512.New()
	mySha512.Write([]byte(input))

	newSignatureKey := hex.EncodeToString(mySha512.Sum(nil))

	return newSignatureKey == n.SignatureKey

}
