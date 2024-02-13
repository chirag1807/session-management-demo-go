package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sessionmanagement/error"
)

func ErrorGenerator(w http.ResponseWriter, err error) {
	var response interface{}

	fmt.Println(err.Error())

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if error, ok := err.(errorhandling.CustomError); ok {
		response = errorhandling.CustomError{
			StatusCode: error.StatusCode,
			Message:    error.Message,
		}
		w.WriteHeader(error.StatusCode)

	} else {
		response = errorhandling.CustomError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Internal Server Error",
		}
		w.WriteHeader(http.StatusInternalServerError)
	}
	
	json.NewEncoder(w).Encode(response)
}