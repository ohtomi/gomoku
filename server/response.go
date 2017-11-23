package server

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func (r *Response) Write(conversation *Conversation, writer http.ResponseWriter) error {
	var content string
	if len(r.Body) != 0 {
		applied, err := ApplyTemplateText("body", r.Body, conversation)
		if err != nil {
			return err
		}
		content = applied
	} else if len(r.Template) != 0 {
		applied, err := ApplyTemplateFile("template", r.Template, conversation)
		if err != nil {
			return err
		}
		content = applied
	} else if len(r.File) != 0 {
		read, err := readFile(r.File)
		if err != nil {
			return err
		}
		content = string(read[:len(read)])
	} else {
		content = ""
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

	fmt.Fprintf(writer, content)

	return nil
}

func readFile(filename string) ([]byte, error) {
	fd, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer fd.Close()
	return ioutil.ReadAll(fd)
}
