package web

type RefundTransactionDto struct {
	RefundKey string `json:"refund_key"`
	Amount    int    `json:"amount"`
	Reason    string `json:"reason"`
}
