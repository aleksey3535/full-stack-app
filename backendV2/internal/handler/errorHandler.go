package handler

import (
	"encoding/json"
	"net/http"
)

type ErrorMessage struct {
	Error string `json:"error"`
}

func ErrorWrapper(err error) []byte {
	data := ErrorMessage {Error: err.Error()}
	message, _ := json.Marshal(data)
	return message
}

func InternalErrorHandler(w http.ResponseWriter) {
	data := ErrorMessage{Error: "something is wrong with the server"}
	message, _:= json.Marshal(data)
	w.WriteHeader(http.StatusInternalServerError)
	w.Write(message)
}

func InvalidDataHandler(w http.ResponseWriter) {
	data := ErrorMessage{Error: "invalid data"}
	message, _ := json.Marshal(data)
	w.WriteHeader(http.StatusBadRequest)
	w.Write(message)
}
