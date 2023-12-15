package helper

import (
	"encoding/json"
	"net/http"
)

func ReadFromRequestBody(request *http.Request, result interface{}) {
	decoder := json.NewDecoder(request.Body)
	defer request.Body.Close()

	err := decoder.Decode(result)
	PanicIfError("Cannot read request body", err)
}

func WriteToResponse(writer http.ResponseWriter, response interface{}) {
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)

	err := encoder.Encode(response)
	PanicIfError(
		"Cannot response json",
		err,
	)
}
