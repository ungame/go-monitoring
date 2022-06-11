package httpext

import (
	"fmt"
	"net/http"
)

func Validate(r *http.Request, method, uri string) (int, error) {
	if r.URL.Path != uri {
		return http.StatusNotFound, fmt.Errorf("%s: %s", http.StatusText(http.StatusNotFound), r.RequestURI)
	}
	if r.Method != method {
		return http.StatusMethodNotAllowed, fmt.Errorf("%s: %s", http.StatusText(http.StatusMethodNotAllowed), r.Method)
	}
	return 0, nil
}
