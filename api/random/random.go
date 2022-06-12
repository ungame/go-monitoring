package random

import (
	"math/rand"
	"time"
)

func Random() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

func Int() int {
	return Random().Int()
}

func Intn(max int) int {
	return Random().Intn(max)
}
