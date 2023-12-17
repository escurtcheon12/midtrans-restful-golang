package controllers

import (
	"fmt"
	"midtrans-go/config"
	"midtrans-go/helper"
	"midtrans-go/model/web"
	"midtrans-go/service"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

type MidtransControllerImpl struct {
	Coreapi         coreapi.Client
	MidtransService service.MidtransService
}

func NewMidtransController(coreapi coreapi.Client, midtransService service.MidtransService) *MidtransControllerImpl {
	return &MidtransControllerImpl{
		Coreapi:         coreapi,
		MidtransService: midtransService,
	}
}

func (c *MidtransControllerImpl) ChargeTransaction(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	s := coreapi.Client{}
	s.New(config.NewConfig().Midtrans.ServerKey, midtrans.Sandbox)

	chargeTransactionRequest := &coreapi.ChargeReq{}

	helper.ReadFromRequestBody(r, chargeTransactionRequest)

	coreApiRes, _ := s.ChargeTransaction(chargeTransactionRequest)

	helper.WriteToResponse(w, coreApiRes)
}

func (c *MidtransControllerImpl) CancelTransaction(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	cancelTransactionRequest := &web.CancelTransactionDto{}

	helper.ReadFromRequestBody(r, cancelTransactionRequest)

	s := coreapi.Client{}
	s.New(config.NewConfig().Midtrans.ServerKey, midtrans.Sandbox)

	coreApiRes, _ := s.CancelTransaction(cancelTransactionRequest.OrderId)

	helper.WriteToResponse(w, coreApiRes)
}

func (c *MidtransControllerImpl) RefundTransaction(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	orderId := r.URL.Query().Get("order_id")
	refundTransactionRequest := &web.RefundTransactionDto{}

	helper.ReadFromRequestBody(r, refundTransactionRequest)

	s := coreapi.Client{}
	s.New(config.NewConfig().Midtrans.ServerKey, midtrans.Sandbox)

	refundReq := &coreapi.RefundReq{
		RefundKey: refundTransactionRequest.RefundKey,
		Amount:    int64(refundTransactionRequest.Amount),
		Reason:    refundTransactionRequest.Reason,
	}

	coreApiRes, _ := s.RefundTransaction(orderId, refundReq)

	helper.WriteToResponse(w, coreApiRes)
}

func (c *MidtransControllerImpl) GetTransactionStatus(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	getStatusTransactionRequest := &web.GetStatusTransactionDto{
		OrderId: r.URL.Query().Get("order_id"),
	}

	fmt.Println(getStatusTransactionRequest.OrderId)

	s := coreapi.Client{}
	s.New(config.NewConfig().Midtrans.ServerKey, midtrans.Sandbox)

	coreApiRes, _ := s.CancelTransaction(getStatusTransactionRequest.OrderId)

	helper.WriteToResponse(w, coreApiRes)
}

func (c *MidtransControllerImpl) VerifyPayment(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	s := coreapi.Client{}
	s.New(config.NewConfig().Midtrans.ServerKey, midtrans.Sandbox)

	captureTransactionRequest := web.CaptureTransactionDto{}
	helper.ReadFromRequestBody(r, &captureTransactionRequest)

	// 4. Check transaction to Midtrans with param orderId
	transactionStatusResp, e := s.CheckTransaction(captureTransactionRequest.OrderId)
	fmt.Println(transactionStatusResp)
	if e != nil {
		http.Error(w, e.GetMessage(), http.StatusInternalServerError)
		return
	} else {
		if transactionStatusResp != nil {
			// 5. Do set transaction status based on response from check transaction status
			if transactionStatusResp.TransactionStatus == "capture" {
				if transactionStatusResp.FraudStatus == "challenge" {
					// TODO set transaction status on your database to 'challenge'
					// e.g: 'Payment status challenged. Please take action on your Merchant Administration Portal
				} else if transactionStatusResp.FraudStatus == "accept" {
					// TODO set transaction status on your database to 'success'
				}
			} else if transactionStatusResp.TransactionStatus == "settlement" {
				// TODO set transaction status on your databaase to 'success'
			} else if transactionStatusResp.TransactionStatus == "deny" {
				// TODO you can ignore 'deny', because most of the time it allows payment retries
				// and later can become success
			} else if transactionStatusResp.TransactionStatus == "cancel" || transactionStatusResp.TransactionStatus == "expire" {
				// TODO set transaction status on your databaase to 'failure'
			} else if transactionStatusResp.TransactionStatus == "pending" {
				// TODO set transaction status on your databaase to 'pending' / waiting payment
			}
		}
	}

	helper.WriteToResponse(w, transactionStatusResp)
}
