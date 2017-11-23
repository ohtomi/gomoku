package server

import (
	"bytes"
	"fmt"
	"net/http"
	"text/template"
)

func (r *Response) Write(result map[string]interface{}, w http.ResponseWriter) error {
	var body string
	if len(r.Body) != 0 {
		buf := &bytes.Buffer{}
		t := template.Must(template.New("body").Parse(r.Body))
		if err := t.Execute(buf, result); err != nil {
			return err
		}
		body = buf.String()
	} else if len(r.Template) != 0 {
		buf := &bytes.Buffer{}
		t := template.Must(template.ParseFiles(r.Template))
		if err := t.Execute(buf, result); err != nil {
			return err
		}
		body = buf.String()
	} else {
		body = ""
	}

	for key, value := range r.Headers {
		w.Header().Set(key, value)
	}

	if r.Status == 0 {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(r.Status)
	}

	fmt.Fprintf(w, body)

	return nil
}
