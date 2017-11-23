package server

import (
	"bytes"
	"fmt"
	"html/template"
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

		var body string
		if len(cResponse.Body) != 0 {
			buf := &bytes.Buffer{}
			t := template.Must(template.New("body").Parse(cResponse.Body))
			if err := t.Execute(buf, result); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			body = buf.String()
		} else if len(cResponse.Template) != 0 {
			body = "TODO"
		} else {
			if v, ok := result["Return"]; ok {
				body = v.(string)
			} else {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}

		for key, value := range cResponse.Headers {
			w.Header().Set(key, value)
		}

		if cResponse.Status == 0 {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(cResponse.Status)
		}

		fmt.Fprintf(w, body)
	}
}

func StartHttpServer(addr string, config *Config) error {
	http.HandleFunc("/", buildHandleFunc(config))
	if err := http.ListenAndServe(addr, nil); err != nil {
		return errors.Wrap(err, "failed to start http server")
	}
	return nil
}
