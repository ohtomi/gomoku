package server

import (
	"net/http"

	"github.com/pkg/errors"
)

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
			origin := r.Header.Get("Origin")
			if len(origin) != 0 {
				w.Header().Set("Access-Control-Allow-Origin", origin)
			} else {
				w.Header().Set("Access-Control-Allow-Origin", "*")
			}
			if r.Method == "OPTIONS" {
				w.Header().Set("Access-Control-Allow-Headers", r.Header.Get("Access-Control-Request-Headers"))
				w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
				w.WriteHeader(http.StatusOK)
				reporter.Infof("    <Preflight request>")
				return
			}
		}

		cUpgrade, cRequest, cCommand, cResponse, found := config.SelectConfigItem(r.Method, r.URL.Path, r.Header)

		if !found {
			if errorNoMatch {
				w.WriteHeader(http.StatusInternalServerError)
			} else {
				w.WriteHeader(http.StatusOK)
			}
			reporter.Errorf("    <No match routes>")
		} else if cUpgrade != nil {
			RunRobots(cUpgrade, reporter, w, r)
		} else {
			DoConversation(cRequest, cCommand, cResponse, reporter, w, r)
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
