package handlers

import (
	"github.com/ungame/go-monitoring/api/httpext"
	"github.com/ungame/go-monitoring/api/types"
	"net/http"
)

func Up(w http.ResponseWriter, r *http.Request) {
	if status, err := httpext.Validate(r, http.MethodGet, "/"); err != nil {
		httpext.WriteError(w, status, err)
	} else {
		httpext.Write(w, http.StatusOK, types.Map{"hello": "world!"})
	}
}
