package server

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
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
	} else if len(r.File) != 0 {
		fd, err := os.Open(r.File)
		if err != nil {
			return err
		}
		defer fd.Close()
		content, err := ioutil.ReadAll(fd)
		if err != nil {
			return err
		}
		buf.Write(content)
	} else {
		buf.WriteString("")
	}

	for key, value := range r.Headers {
		value, err := ApplyTemplateText("header", value, conversation)
		if err != nil {
			return err
		}
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
