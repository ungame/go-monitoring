package middlewares

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"net/http"
)

var (
	successCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "successfully_requests",
		Help: "Número total de requisições bem sucedidas",
	})

	failCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "failed_requests",
		Help: "Número total de requisições que falharam",
	})
)

func monitor(status int) {
	if status < http.StatusBadRequest {
		go func() { successCounter.Inc() }()
	} else {
		go func() { failCounter.Inc() }()
	}
}
