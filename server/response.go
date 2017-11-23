package server

import (
	"bytes"
	"fmt"
	"net/http"
	"text/template"
)

func (r *Response) Write(conversation *Conversation, writer http.ResponseWriter) error {
	var body string
	if len(r.Body) != 0 {
		buf := &bytes.Buffer{}
		t := template.Must(template.New("body").Parse(r.Body))
		if err := t.Execute(buf, conversation); err != nil {
			return err
		}
		body = buf.String()
	} else if len(r.Template) != 0 {
		buf := &bytes.Buffer{}
		t := template.Must(template.ParseFiles(r.Template))
		if err := t.Execute(buf, conversation); err != nil {
			return err
		}
		body = buf.String()
	} else {
		body = ""
	}

	for key, value := range r.Headers {
		writer.Header().Set(key, value)
	}

	if r.Status == 0 {
		writer.WriteHeader(http.StatusOK)
	} else {
		writer.WriteHeader(r.Status)
	}

	fmt.Fprintf(writer, body)

	return nil
}
