package server

import (
	"bytes"
	"fmt"
	"net/http"
	"text/template"
)

func (r *Response) Write(conversation *Conversation, writer http.ResponseWriter) error {
	buf := &bytes.Buffer{}
	if len(r.Body) != 0 {
		t, err := template.New("body").Parse(r.Body)
		if err != nil {
			return err
		}
		if err := t.Execute(buf, conversation); err != nil {
			return err
		}
	} else if len(r.Template) != 0 {
		t, err := template.ParseFiles(r.Template)
		if err != nil {
			return err
		}
		if err := t.Execute(buf, conversation); err != nil {
			return err
		}
	} else {
		buf.WriteString("")
	}

	for key, value := range r.Headers {
		writer.Header().Set(key, value)
	}

	if r.Status == 0 {
		writer.WriteHeader(http.StatusOK)
	} else {
		writer.WriteHeader(r.Status)
	}

	fmt.Fprintf(writer, buf.String())

	return nil
}
