package controllers

import (
	"fmt"
	"midtrans-go/config"
	"midtrans-go/helper"
	"midtrans-go/model/web"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

type MidtransControllerImpl struct {
	Coreapi coreapi.Client
}

func NewMidtransController(coreapi coreapi.Client) *MidtransControllerImpl {
	return &MidtransControllerImpl{
		Coreapi: coreapi,
	}
}

func (c *MidtransControllerImpl) ChargeTransactions(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	chargeTransactionRequest := &web.ChargeTransactionsDto{}

	helper.ReadFromRequestBody(r, chargeTransactionRequest)

	// 1. Initiate coreapi client
	s := coreapi.Client{}
	s.New(config.NewConfig().Midtrans.ServerKey, midtrans.Sandbox)

	// 2. Initiate charge request
	chargeReq := &coreapi.ChargeReq{
		PaymentType: "gopay",
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  chargeTransactionRequest.TransactionDetails.OrderID,
			GrossAmt: int64(chargeTransactionRequest.TransactionDetails.GrossAmount),
		},
		// CreditCard: &coreapi.CreditCardDetails{
		// 	TokenID:        "YOUR-CC-TOKEN",
		// 	Authentication: true,
		// },
	}

	// 3. Request to Midtrans
	coreApiRes, _ := s.ChargeTransaction(chargeReq)
	fmt.Println("Response :", coreApiRes)

	helper.WriteToResponse(w, coreApiRes)

}
