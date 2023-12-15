package midtranscontroller

import (
	"fmt"
	"midtrans-go/config"
	"midtrans-go/helper"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type MidtransControllerImpl struct {
	Client snap.Client
}

func NewMidtransController(client snap.Client) *MidtransControllerImpl {
	return &MidtransControllerImpl{
		Client: client,
	}
}

func (c *MidtransControllerImpl) ChargeTransactions(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// c.Client.New(config.NewConfig().Midtrans.ServerKey, midtrans.Sandbox)

	var s = c.Client
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

	helper.WriteToResponse(w, snapResp)

	// chargeReq := &c.Client.CreateTransaction{
	// 	PaymentType: midtrans.SourceCreditCard,
	// 	TransactionDetails: midtrans.TransactionDetails{
	// 		OrderID:  "12345",
	// 		GrossAmt: 200000,
	// 	},
	// 	CreditCard: &coreapi.CreditCardDetails{
	// 		TokenID:        "YOUR-CC-TOKEN",
	// 		Authentication: true,
	// 	},
	// 	Items: &[]midtrans.ItemDetail{
	// 		coreapi.ItemDetail{
	// 			ID:    "ITEM1",
	// 			Price: 200000,
	// 			Qty:   1,
	// 			Name:  "Someitem",
	// 		},
	// 	},
	// }

	// chargeReq := &c.ChargeReq{
	// 	PaymentType: midtrans.SourceCreditCard,
	// 	TransactionDetails: midtrans.TransactionDetails{
	// 		OrderID:  "12345",
	// 		GrossAmt: 200000,
	// 	},
	// 	CreditCard: &c.CreditCardDetails{
	// 		TokenID:        "YOUR-CC-TOKEN",
	// 		Authentication: true,
	// 	},
	// 	Items: &[]midtrans.ItemDetail{
	// 		c.ItemDetail{
	// 			ID:    "ITEM1",
	// 			Price: 200000,
	// 			Qty:   1,
	// 			Name:  "Someitem",
	// 		},
	// 	},
	// }
}
