package logger

import (
	"fmt"
	"github.com/ungame/timeutils"
	"log"
	"os"
	"sync"
	"time"
)

type PrintFormatter func(template string, values ...any)

type MemoryLogger interface {
	Logf(template string, values ...any)
}

type memoryLogger struct {
	prefix string
	logger *log.Logger
}

func New(prefix string) MemoryLogger {
	return &memoryLogger{
		prefix: prefix,
		logger: log.New(os.Stdout, prefix, log.LstdFlags),
	}
}

func (l *memoryLogger) Logf(template string, values ...any) {
	l.onPrintf(l.logger.Printf, template, values...)
}

func (l *memoryLogger) onPrintf(pf PrintFormatter, template string, values ...any) {
	go store(template, values...)
	pf(template, values...)
}

func store(template string, values ...any) {
	var (
		now    = time.Now().Format(timeutils.DateTimeFormat)
		info   = now + " " + fmt.Sprintf(template, values...)
		cached []string
	)

	items, ok := logs.Load(logsKey)
	if !ok {
		cached = make([]string, 0)
	} else {
		cached = items.([]string)
	}
	cached = append(cached, info)
	logs.Store(logsKey, cached)
}

const logsKey = "logs"

var (
	infoLogger  MemoryLogger
	errorLogger MemoryLogger
	logs        *sync.Map
)

func init() {
	logs = &sync.Map{}
	logs.Store(logsKey, []string{})
	infoLogger = New("[INFO]  ")
	errorLogger = New("[ERROR] ")
}

func Info(template string, values ...any) {
	infoLogger.Logf(template, values...)
}

func Error(template string, values ...any) {
	errorLogger.Logf(template, values...)
}

func Logs() []string {
	items, ok := logs.Load(logsKey)
	if !ok {
		return nil
	}
	return items.([]string)
}
