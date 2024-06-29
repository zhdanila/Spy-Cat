package handler

import (
	"log"
	"net/http"
)

func NewErrorResponse(w http.ResponseWriter, statusCode int, errorString string) {
	log.Println(errorString)
	w.WriteHeader(statusCode)
	w.Write([]byte(errorString))
}