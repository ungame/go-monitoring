package exit

import (
	"github.com/ungame/go-monitoring/api/logger"
	"os"
	"os/signal"
)

const (
	SUCCESS = 0
	FAILURE = 1
)

func Graceful() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	go func() {
		for {
			select {
			case _ = <-stop:
				logger.Info("Stopped.")
				os.Exit(SUCCESS)
			}
		}
	}()
}
