package web

type OrdersRequestDto struct {
	Status           string `validate:"required" json:"status"`
	MidtransResponse string `validate:"required" json:"midtrans_response"`
}
