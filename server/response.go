package server

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func (r *Response) Write(conversation *Conversation, writer http.ResponseWriter) error {
	var content []byte

	if len(r.Body) != 0 {
		applied, err := ApplyTemplateText("body", r.Body, conversation)
		if err != nil {
			return err
		}
		content = []byte(applied)
	} else if len(r.Template) != 0 {
		template, err := ApplyTemplateText("template", r.Template, conversation)
		if err != nil {
			return err
		}
		applied, err := ApplyTemplateFile("template", template, conversation)
		if err != nil {
			return err
		}
		content = []byte(applied)
	} else if len(r.File) != 0 {
		file, err := ApplyTemplateText("file", r.File, conversation)
		if err != nil {
			return err
		}
		read, err := readFile(file)
		if err != nil {
			return err
		}
		content = read
	} else {
		content = []byte{}
	}

	for key, value := range r.Headers {
		value, err := ApplyTemplateText("header", value, conversation)
		if err != nil {
			return err
		}
		writer.Header().Set(key, value)
	}
	writer.Header().Set("content-length", fmt.Sprintf("%d", len(content)))

	if r.Status == 0 {
		writer.WriteHeader(http.StatusOK)
	} else {
		writer.WriteHeader(r.Status)
	}

	if len(content) != 0 {
		writer.Write(content)
	}

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
