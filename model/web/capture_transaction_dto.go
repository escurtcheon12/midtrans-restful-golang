package web

type CaptureTransactionDto struct {
	OrderId       string  `json:"order_id"`
	TransactionId string  `json:"transaction_id"`
	GrossAmount   float64 `json:"gross_amount"`
}
