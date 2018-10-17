package server

import (
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

func buildUserScriptHandler(config *Config, cors, errorNoMatch bool, reporter Reporter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if reporter.IsEnabled() {
			reporter.Info("Incoming Request")
			for k, v := range r.Header {
				reporter.Infof("    header => %s: %s", k, v)
			}
			reporter.Infof("    remote => %s", r.RemoteAddr)
			reporter.Infof("    method => %s", r.Method)
			reporter.Infof("       url => %s", r.URL)
		}

		if cors {
			if r.Method == "OPTIONS" {
				w.Header().Set("Access-Control-Allow-Origin", r.RemoteAddr)
				w.Header().Set("Access-Control-Allow-Headers", r.Header.Get("Access-Control-Request-Headers"))
				w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
				w.WriteHeader(http.StatusOK)
				reporter.Infof("    <Preflight request>")
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
			reporter.Errorf("    <No match routes>")
			return
		}

		conversation := &Conversation{}

		if err := cRequest.Transform(conversation, r); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			reporter.Error(err.Error())
			return
		}

		if reporter.IsEnabled() {
			reporter.Infof("      body => %s", conversation.Request.Body)
			reporter.Infof("      form => %s", conversation.Request.Form)
		}

		if err := cCommand.Execute(conversation); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			reporter.Error(err.Error())
			return
		}

		if err := cResponse.Write(conversation, w); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			reporter.Error(err.Error())
			return
		}
	}
}

func StartHttpServer(addr string, config *Config, cors, tls bool, cert, key string, errorNoMatch bool, reporter Reporter) error {
	http.HandleFunc("/", buildUserScriptHandler(config, cors, errorNoMatch, reporter))
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
