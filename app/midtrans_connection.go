package app

import (
	"fmt"
	"midtrans-go/config"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

func NewMidtransServer() {
	// 1. Set you ServerKey with globally
	midtrans.ServerKey = config.NewConfig().Midtrans.ServerKey
	midtrans.Environment = midtrans.Sandbox

	// 2. Initiate Snap request
	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  "YOUR-ORDER-ID-12345",
			GrossAmt: 100000,
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
	}

	// 3. Request create Snap transaction to Midtrans
	snapResp, _ := snap.CreateTransaction(req)
	fmt.Println("Response :", snapResp)
}

func NewMidtransClient() {
	// 1. Initiate Snap client
	// midtrans.ClientKey = config.NewConfig().Midtrans.ClientKey

	var s = &snap.Client{}
	s.New(config.NewConfig().Midtrans.ServerKey, midtrans.Sandbox)

	// 2. Initiate Snap request
	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  "YOUR-ORDER-ID-12345",
			GrossAmt: 100000,
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
	}

	// 3. Request create Snap transaction to Midtrans
	snapResp, _ := s.CreateTransaction(req)
	fmt.Println("Response :", snapResp)
}
