package common

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Response[T any] struct {
	Data  *T    `json:"data,omitempty"`
	Count *uint `json:"count,omitempty"`
}

func HTTPError(w http.ResponseWriter, message string, httpStatus int) {
	response := Error{
		Code:    httpStatus,
		Message: message,
	}

	responseBytes, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Unable to marshal error response"))

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)

	if _, err := w.Write(responseBytes); err != nil {
		HTTPError(w, "Unable to write response", http.StatusInternalServerError)
	}
}

func HTTPOK[T any](w http.ResponseWriter, data T, count *uint) {
	response := Response[T]{
		Data:  &data,
		Count: count,
	}

	responseBytes, err := json.Marshal(response)
	if err != nil {
		HTTPError(w, "Unable to marshal response", http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseBytes)
}
