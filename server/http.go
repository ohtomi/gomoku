package server

import (
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

type CommandResult struct {
	Env    []string
	Path   string
	Args   []string
	Dir    string
	Stdout string
	Stderr string
}

type Conversation struct {
	Method     string
	URL        string
	Headers    map[string][]string
	Form       map[string][]string
	RemoteAddr string

	CommandResult *CommandResult
}

func buildHandleFunc(config *Config, verbose bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cRequest, cCommand, cResponse := config.SelectConfigItem(r.Method, r.URL.Path)

		if cRequest == nil || cCommand == nil || cResponse == nil {
			w.WriteHeader(http.StatusOK)
			return
		}

		conversation := &Conversation{}

		if err := cRequest.Transform(conversation, r); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			if verbose {
				fmt.Fprintf(w, err.Error())
			}
			return
		}

		if err := cCommand.Execute(conversation); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			if verbose {
				fmt.Fprintf(w, err.Error())
			}
			return
		}

		if err := cResponse.Write(conversation, w); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			if verbose {
				fmt.Fprintf(w, err.Error())
			}
			return
		}
	}
}

func StartHttpServer(addr string, config *Config, verbose bool) error {
	http.HandleFunc("/", buildHandleFunc(config, verbose))
	if err := http.ListenAndServe(addr, nil); err != nil {
		return errors.Wrap(err, "failed to start http server")
	}
	return nil
}
