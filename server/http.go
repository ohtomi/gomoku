package server

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"regexp"

	"github.com/pkg/errors"
)

func buildHandleFunc(config *Config) http.HandlerFunc {
	selector := func(path string) (*Request, *Command, *Response) {
		for _, element := range *config {
			if len(element.Request.Route) == 0 {
				continue
			}
			if !regexp.MustCompile(element.Request.Route).MatchString(path) {
				continue
			}
			return &element.Request, &element.Command, &element.Response
		}
		return nil, nil, nil
	}

	return func(w http.ResponseWriter, r *http.Request) {
		cRequest, cCommand, cResponse := selector(r.URL.Path)

		if cRequest == nil || cCommand == nil || cResponse == nil {
			w.WriteHeader(http.StatusOK)
			return
		}

		result := map[string]interface{}{
			"Command": cCommand.Path,
			"Return":  "foo! bar! baz!, hoge? fuga?",
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

		if cResponse.Status == 0 {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(cResponse.Status)
		}

		for key, value := range cResponse.Headers {
			w.Header().Add(key, value)
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
