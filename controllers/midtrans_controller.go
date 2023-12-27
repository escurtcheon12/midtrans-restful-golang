package controllers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type MidtransController interface {
	ChargeTransaction(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	CancelTransaction(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	RefundTransaction(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	GetTransactionStatus(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	VerifyPayment(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	Notification(w http.ResponseWriter, r *http.Request, p httprouter.Params)
}
