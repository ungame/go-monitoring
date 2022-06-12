package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/ungame/go-monitoring/api/logger"
	"github.com/ungame/go-monitoring/api/random"
	"net/http"
	"sync"
	"time"
)

const (
	defaultConcurrent = false
	defaultRequests   = 100
	defaultRepeat     = 5
	defaultDelay      = 30
)

var (
	concurrent bool
	requests   int
	repeat     int
	delay      int
)

func init() {
	flag.BoolVar(&concurrent, "concurrent", defaultConcurrent, "concurrent clients requests")
	flag.IntVar(&requests, "requests", defaultRequests, "requests per minute")
	flag.IntVar(&repeat, "repeat", defaultRepeat, "repeat requests")
	flag.IntVar(&delay, "delay", defaultDelay, "delay in seconds after requests")
	flag.Parse()
}

func main() {

	if delay <= 0 {
		delay = defaultDelay
	}

	logger.Info("Making client requests...\n")
	logger.Info("Concurrent=%v, Requests=%v, Repeat=%v, Delay=%v\n", concurrent, requests, repeat, delay)
	logger.Info("Total Requests=%v\n\n", requests*repeat)

	var wg *sync.WaitGroup
	if concurrent {
		wg = &sync.WaitGroup{}
		defer func() {
			logger.Info("Waiting clients...")
			wg.Wait()
		}()
	}

	for i := 0; i < repeat; i++ {
		if concurrent {
			wg.Add(1)
			go flood(requests, wg)
		} else {
			flood(requests)
			time.Sleep(time.Second * time.Duration(delay))
		}
	}

	logger.Info("Finished clients.")
}

func flood(n int, wg ...*sync.WaitGroup) {
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
		} else if randInt%2 != 0 {
			uri = "/fail"
		}
		request(uri)
	}

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
