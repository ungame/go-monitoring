package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/ungame/go-monitoring/api/logger"
	"math/rand"
	"net/http"
	"time"
)

var (
	concurrent bool
	requests   int
)

func init() {
	flag.BoolVar(&concurrent, "concurrent", false, "concurrent clients requests")
	flag.IntVar(&requests, "requests", 1000, "requests per minute")
	flag.Parse()
}

func main() {
	for {
		for i := 0; i < 2; i++ {
			if concurrent {
				go flood(requests)
			} else {
				flood(requests)
			}
		}
		time.Sleep(time.Minute)
	}
}

func flood(n int) {
	var (
		uri    string
		random int64
	)
	logger.Info("Starting flood...")
	for i := 0; i < n; i++ {
		uri = "/"
		random = rand.New(rand.NewSource(time.Now().UnixNano())).Int63()
		if random%1000 == 0 {
			uri = "/down"
		} else if random%2 != 0 {
			uri = "/fail"
		}
		request(uri)
		random = rand.New(rand.NewSource(time.Now().UnixNano())).Int63n(100) + 1
		time.Sleep(time.Millisecond * time.Duration(random))
	}
	logger.Info("Flood finished.")
}

func request(uri string) {
	logger.Info("request: %s", uri)
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
}
