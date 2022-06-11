package handlers

import (
	"errors"
	"github.com/ungame/go-monitoring/api/httpext"
	"log"
	"net/http"
)

func Down(w http.ResponseWriter, r *http.Request) {
	if status, err := httpext.Validate(r, http.MethodGet, "/down"); err != nil {
		httpext.WriteError(w, status, err)
	} else {
		log.Fatal(errors.New(http.StatusText(http.StatusInternalServerError)))
	}
}
