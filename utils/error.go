package utils

import (
	"net/http"
)

// ReturnJsonResponse function for returning data in JSON format
func ReturnJsonResponse(res http.ResponseWriter, httpCode int, resMessage []byte) {
	res.Header().Set("Content-type", "application/json")
	res.WriteHeader(httpCode)
	res.Write(resMessage)
}

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}
