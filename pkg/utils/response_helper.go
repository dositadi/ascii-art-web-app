package utils

import "net/http"

func ErrorResponse(w http.ResponseWriter, errorMessage []byte, header int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(header)
	w.Write(errorMessage)
}
