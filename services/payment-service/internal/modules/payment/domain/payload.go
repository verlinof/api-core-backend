package domain

type MidtransNotification struct {
	// Common fields for all payment methods
	TransactionTime   string `json:"transaction_time,omitempty"`
	TransactionStatus string `json:"transaction_status,omitempty"`
	TransactionID     string `json:"transaction_id,omitempty"`
	StatusMessage     string `json:"status_message,omitempty"`
	StatusCode        string `json:"status_code,omitempty"`
	SignatureKey      string `json:"signature_key,omitempty"`
	PaymentType       string `json:"payment_type,omitempty"`
	OrderID           string `json:"order_id,omitempty"`
	MerchantID        string `json:"merchant_id,omitempty"`
	GrossAmount       string `json:"gross_amount,omitempty"`
	FraudStatus       string `json:"fraud_status,omitempty"`
	Currency          string `json:"currency,omitempty"`
	SettlementTime    string `json:"settlement_time,omitempty"`
	ExpiryTime        string `json:"expiry_time,omitempty"`

	// Credit Card specific fields
	MaskedCard             string `json:"masked_card,omitempty"`
	ApprovalCode           string `json:"approval_code,omitempty"`
	ECI                    string `json:"eci,omitempty"`
	ChannelResponseCode    string `json:"channel_response_code,omitempty"`
	ChannelResponseMessage string `json:"channel_response_message,omitempty"`
	CardType               string `json:"card_type,omitempty"`
	Bank                   string `json:"bank,omitempty"`
	SavedTokenID           string `json:"saved_token_id,omitempty"`
	SavedTokenIDExpiredAt  string `json:"saved_token_id_expired_at,omitempty"`
}

const (
	TransactionStatusCapture    = "capture"
	TransactionStatusSettlement = "settlement"
	TransactionStatusPending    = "pending"
	TransactionStatusDeny       = "deny"
	TransactionStatusCancel     = "cancel"
	TransactionStatusExpire     = "expire"
	TransactionStatusFailure    = "failure"

	FraudStatusAccept    = "accept"
	FraudStatusChallenge = "challenge"
	FraudStatusDeny      = "deny"
)
