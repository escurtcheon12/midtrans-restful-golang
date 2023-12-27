package app

import (
	"fmt"
	"midtrans-go/controllers"
	"midtrans-go/exception"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(authController controllers.AuthController, midtransController controllers.MidtransController, ordersController controllers.OrdersController, notificationsController controllers.NotificationsController, userController controllers.UserController) *httprouter.Router {
	router := httprouter.New()

	router.POST("/api/midtrans/create-transaction", (midtransController.ChargeTransaction))
	router.POST("/api/midtrans/cancel-transaction", (midtransController.CancelTransaction))
	router.POST("/api/midtrans/refund-transaction", (midtransController.RefundTransaction))
	router.POST("/api/midtrans/capture-transaction", (midtransController.VerifyPayment))
	router.GET("/api/midtrans/get-transaction-status", (midtransController.GetTransactionStatus))
	router.POST("/api/midtrans/callback-notif", (midtransController.Notification))

	router.GET("/api/jwt/generate", authController.CreateJWT)

	router.GET("/api/orders", ordersController.Get)
	router.GET("/api/orders/:orderId", ordersController.GetById)
	router.POST("/api/orders/create", ordersController.Create)

	router.GET("/api/notifications", notificationsController.Get)
	router.GET("/api/notifications/:notificationId", notificationsController.GetById)
	router.POST("/api/notifications/create", notificationsController.Create)

	router.GET("/api/user", userController.Get)
	router.GET("/api/user/:userId", userController.GetById)
	router.POST("/api/user/create", userController.Create)

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
