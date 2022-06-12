package handlers

import (
	"github.com/ungame/go-monitoring/api/httpext"
	"github.com/ungame/go-monitoring/api/logger"
	"net/http"
	"strings"
)

func Logs(w http.ResponseWriter, r *http.Request) {
	if status, err := httpext.Validate(r, http.MethodGet, "/logs"); err != nil {
		httpext.WriteError(w, status, err)
	} else {
		logs := logger.Logs()
		var body = strings.Join(logs, "")
		_, _ = w.Write([]byte(body))
	}
}
