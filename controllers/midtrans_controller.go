package controllers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type MidtransController interface {
	ChargeTransactions(w http.ResponseWriter, r *http.Request, p httprouter.Params)
}
