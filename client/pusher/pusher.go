package pusher

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
	"github.com/ungame/go-monitoring/api/logger"
)

var outageCounter = prometheus.NewCounter(prometheus.CounterOpts{
	Name: "outage_requests",
	Help: "Número de requisições que causaram indisponibilidade da aplicação",
})

func init() {
	prometheus.MustRegister(outageCounter)
}

type Pusher interface {
	Incr()
	Push()
}

type pusherImpl struct {
	client *push.Pusher
}

func New(gatewayURL string) Pusher {
	return &pusherImpl{
		client: push.New(gatewayURL, "clients"),
	}
}

func (p *pusherImpl) Incr() {
	outageCounter.Inc()
}

func (p *pusherImpl) Push() {
	err := p.client.Collector(outageCounter).Push()
	if err != nil {
		logger.Error("error on send outage counter: %s\n", err.Error())
	} else {
		logger.Info("outage counter sent: %v\n", outageCounter.Desc())
	}
}
