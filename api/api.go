package api

import (
	"flag"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/ungame/go-monitoring/api/handlers"
	"github.com/ungame/go-monitoring/api/httpext"
	"github.com/ungame/go-monitoring/api/middlewares"
	"log"
	"net/http"
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
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/", middlewares.Logger(handlers.Up))
	http.HandleFunc("/down", middlewares.Logger(handlers.Down))
	httpext.Listen(port, func() {
		log.Printf("Listening %s:%d\n\n", host, port)
	})
}
