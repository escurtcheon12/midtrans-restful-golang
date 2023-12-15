package helper

import (
	"github.com/sirupsen/logrus"
)

func PanicIfError(message string, err error) {
	if err != nil {
		logrus.Error(message)
		panic(err)
	}
}
