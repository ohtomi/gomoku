package server

import (
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

func buildHandleFunc(config *Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cRequest, cCommand, cResponse := config.SelectConfigItem(r.Method, r.URL.Path)

		if cRequest == nil || cCommand == nil || cResponse == nil {
			w.WriteHeader(http.StatusOK)
			return
		}

		result, err := cCommand.Execute()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			// TODO write error as response body?
			fmt.Fprintf(w, err.Error())
			return
		}

		if err := cResponse.Write(result, w); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			// TODO write error as response body?
			fmt.Fprintf(w, err.Error())
			return
		}
	}
}

func StartHttpServer(addr string, config *Config) error {
	http.HandleFunc("/", buildHandleFunc(config))
	if err := http.ListenAndServe(addr, nil); err != nil {
		return errors.Wrap(err, "failed to start http server")
	}
	return nil
}
