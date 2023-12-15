package authcontroller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type AuthController interface {
	Login(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	CreateJWT(w http.ResponseWriter, r *http.Request, p httprouter.Params)
}
