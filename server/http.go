package server

import (
	"fmt"
	"github.com/pkg/errors"
	"html"
	"net/http"
)

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func StartHttpServer(addr string) error {
	http.HandleFunc("/", handle)
	if err := http.ListenAndServe(addr, nil); err != nil {
		return errors.Wrap(err, "failed to start http server")
	}
	return nil
}
