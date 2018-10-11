package server

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

type RequestInConversation struct {
	Method     string
	URL        *url.URL
	Headers    map[string][]string
	RemoteAddr string
	Body       string
	Form       map[string][]string
}

type CommandInConversation struct {
	Env    []string
	Path   string
	Args   []string
	Dir    string
	Stdout string
	Stderr string
}

type Conversation struct {
	Request RequestInConversation
	Command CommandInConversation
}

func buildUserScriptHandler(config *Config, cors, verbose, errorNoMatch bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if cors {
			if r.Method == "OPTIONS" {
				w.Header().Set("Access-Control-Allow-Origin", r.RemoteAddr)
				w.Header().Set("Access-Control-Allow-Headers", r.Header.Get("Access-Control-Request-Headers"))
				w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
				w.WriteHeader(http.StatusOK)
				return
			}
		}

		cRequest, cCommand, cResponse, found := config.SelectConfigItem(r.Method, r.URL.Path, r.Header)

		if !found {
			if errorNoMatch {
				w.WriteHeader(http.StatusInternalServerError)
			} else {
				w.WriteHeader(http.StatusOK)
			}
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

func StartHttpServer(addr string, config *Config, cors, tls bool, cert, key string, verbose, errorNoMatch bool) error {
	http.HandleFunc("/", buildUserScriptHandler(config, cors, verbose, errorNoMatch))
	if tls {
		if err := http.ListenAndServeTLS(addr, cert, key, nil); err != nil {
			return errors.Wrap(err, "failed to start https server")
		}
	} else {
		if err := http.ListenAndServe(addr, nil); err != nil {
			return errors.Wrap(err, "failed to start http server")
		}
	}
	return nil
}
