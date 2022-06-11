package httpext

import (
	"fmt"
	"log"
	"net/http"
)

type Port int

func (p Port) Addr() string {
	return fmt.Sprintf(":%d", p)
}

func Listen(port int, beforeListen ...func()) {
	for _, fn := range beforeListen {
		fn()
	}
	log.Fatalln(http.ListenAndServe(Port(port).Addr(), nil))
}
