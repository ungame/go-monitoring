package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/ungame/go-monitoring/api/logger"
	"github.com/ungame/go-monitoring/api/random"
	"github.com/ungame/go-monitoring/client/pusher"
	"net/http"
	"sync"
	"time"
)

const (
	defaultConcurrent     = false
	defaultRequests       = 100
	defaultRepeat         = 5
	defaultDelay          = 30
	defaultPushGatewayURL = "http://localhost:9091/"
)

var (
	concurrent     bool
	requests       int
	repeat         int
	delay          int
	pushGatewayURL string
)

func init() {
	flag.BoolVar(&concurrent, "concurrent", defaultConcurrent, "concurrent clients requests")
	flag.IntVar(&requests, "requests", defaultRequests, "requests per minute")
	flag.IntVar(&repeat, "repeat", defaultRepeat, "repeat requests")
	flag.IntVar(&delay, "delay", defaultDelay, "delay in seconds after requests")
	flag.StringVar(&pushGatewayURL, "push_gateway_url", defaultPushGatewayURL, "prometheus push gateway url")
	flag.Parse()
}

func main() {

	if delay <= 0 {
		delay = defaultDelay
	}

	logger.Info("Making client requests...\n")
	logger.Info("Concurrent=%v, Requests=%v, Repeat=%v, Delay=%v\n", concurrent, requests, repeat, delay)
	logger.Info("Total Requests=%v\n", requests*repeat)
	logger.Info("Push Gateway: %s\n\n", pushGatewayURL)

	var wg *sync.WaitGroup
	if concurrent {
		wg = &sync.WaitGroup{}
		defer func() {
			logger.Info("Waiting clients...")
			wg.Wait()
		}()
	}

	pushService := pusher.New(pushGatewayURL)

	for i := 0; i < repeat; i++ {
		if concurrent {
			wg.Add(1)
			go flood(requests, pushService, wg)
		} else {
			flood(requests, pushService)
			time.Sleep(time.Second * time.Duration(delay))
		}
	}

	logger.Info("Finished clients.")
}

func flood(n int, pushService pusher.Pusher, wg ...*sync.WaitGroup) {
	var (
		uri     string
		randInt int
	)

	logger.Info("Starting flood...")

	for i := 0; i < n; i++ {
		uri = "/"
		randInt = random.Int()
		if randInt%100 == 0 {
			uri = "/down"
			pushService.Incr()
		} else if randInt%2 != 0 {
			uri = "/fail"
		}
		request(uri)
	}

	pushService.Push()

	for _, g := range wg {
		g.Done()
	}

	logger.Info("Flood finished.")
}

func request(uri string) {
	var (
		randInt = random.Intn(100) + 1
		wait    = time.Millisecond * time.Duration(randInt)
	)

	logger.Info("request: URI=%s, Wait=%s", uri, wait.String())

	var (
		ctx    = context.Background()
		cancel context.CancelFunc
	)
	//ctx, cancel = context.WithTimeout(ctx, time.Second*3)
	if cancel != nil {
		defer cancel()
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("http://localhost:8080%s", uri), nil)
	if err != nil {
		logger.Error("Create request error: %s", err.Error())
	}
	_, err = http.DefaultClient.Do(req)
	if err != nil {
		logger.Error("Do request error: %s", err.Error())
	}

	time.Sleep(wait)
}
