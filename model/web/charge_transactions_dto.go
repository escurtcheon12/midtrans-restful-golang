package web

type ChargeTransactionsDto struct {
	PaymentType        string             `json:"payment_type"`
	TransactionDetails TransactionDetails `json:"transaction_details"`
	CustomerDetails    CustomerDetails    `json:"customer_details"`
	CustomField1       string             `json:"custom_field1"`
	CustomField2       string             `json:"custom_field2"`
	CustomField3       string             `json:"custom_field3"`
	CustomExpiry       CustomExpiry       `json:"custom_expiry"`
	Metadata           map[string]string  `json:"metadata"`
}

type TransactionDetails struct {
	OrderID     string `json:"order_id"`
	GrossAmount int    `json:"gross_amount"`
}

type CustomerDetails struct {
	FirstName                     string   `json:"first_name"`
	LastName                      string   `json:"last_name"`
	Email                         string   `json:"email"`
	Phone                         string   `json:"phone"`
	CustomerDetailsRequiredFields []string `json:"customer_details_required_fields"`
}

type CustomExpiry struct {
	ExpiryDuration int    `json:"expiry_duration"`
	Unit           string `json:"unit"`
}
