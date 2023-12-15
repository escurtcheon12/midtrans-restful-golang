package middleware

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

type Middleware struct {
	Handler http.Handler
}

func NewMiddleware(handler http.Handler) *Middleware {
	return &Middleware{
		Handler: handler,
	}
}

func (middleware *Middleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	middleware.Handler.ServeHTTP(writer, request)
	logrus.Info("Running middleware")
}
