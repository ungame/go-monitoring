package httpext

import (
	"encoding/json"
	"log"
	"net/http"
)

func Write(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set(HeaderContentType, MimeJSON)
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Println("error on encode json:", err.Error())
	}
}

func WriteError(w http.ResponseWriter, statusCode int, err error) {
	e := new(JSONError)
	if err != nil {
		e.Error = err.Error()
	}
	Write(w, statusCode, e)
}
