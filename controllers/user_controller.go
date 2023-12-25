package controllers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type UserController interface {
	Create(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	GetById(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	Get(w http.ResponseWriter, r *http.Request, p httprouter.Params)
}
