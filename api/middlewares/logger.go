package middlewares

import (
	"github.com/ungame/go-monitoring/api/logger"
	"net/http"
	"time"
)

type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (s *statusRecorder) WriteHeader(status int) {
	s.status = status
	s.ResponseWriter.WriteHeader(status)
}

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		var (
			sr     = &statusRecorder{ResponseWriter: writer}
			method = request.Method
			host   = request.Host
			uri    = request.RequestURI
			start  = time.Now()
		)

		next(sr, request)

		monitor(sr.status)

		logger.Info("%s %s%s %d %s\n", method, host, uri, sr.status, time.Since(start).String())
	}
}
