package api

import (
	"flag"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/ungame/go-monitoring/api/handlers"
	"github.com/ungame/go-monitoring/api/httpext"
	"github.com/ungame/go-monitoring/api/logger"
	"github.com/ungame/go-monitoring/api/middlewares"
	"net/http"
	"os"
	"os/signal"
)

var (
	host string
	port int
)

func init() {
	flag.StringVar(&host, "host", "127.0.0.1", "api host")
	flag.IntVar(&port, "port", 8080, "api port")
	flag.Parse()
}

func Run() {
	gracefulStop()
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/", middlewares.Logger(handlers.Up))
	http.HandleFunc("/down", middlewares.Logger(handlers.Down))
	http.HandleFunc("/logs", middlewares.Logger(handlers.Logs))
	httpext.Listen(port, func() {
		logger.Info("Listening http://%s:%d\n\n", host, port)
	})
}

func gracefulStop() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	go func() {
		for {
			select {
			case _ = <-stop:
				logger.Info("Stopped.")
				os.Exit(0)
			}
		}
	}()
}
