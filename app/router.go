package app

import (
	"fmt"
	authcontroller "midtrans-go/controllers/auth_controller"
	midtranscontroller "midtrans-go/controllers/midtrans_controller"
	"midtrans-go/exception"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(authController authcontroller.AuthController, midtransController midtranscontroller.MidtransController) *httprouter.Router {
	router := httprouter.New()

	router.POST("/api/midtrans/create-transactions", validateJWT(midtransController.ChargeTransactions))

	router.GET("/api/jwt/generate", authController.CreateJWT)
	router.POST("/api/login", authController.Login)

	router.PanicHandler = exception.ErrorHandler

	return router
}

func validateJWT(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		if r.Header["Token"] != nil {
			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {

					return nil, fmt.Errorf("Invalid signing method")
				}
				return []byte("secret-token"), nil
			})

			fmt.Println(token)

			if err != nil || !token.Valid {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Not authorized"))
				return
			}

			next(w, r, p)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Not authorized"))
		}
	}
}
