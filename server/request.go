package server

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

func (r *Request) Transform(conversation *Conversation, request *http.Request) error {
	conversation.Request.Method = request.Method
	conversation.Request.URL = request.URL
	conversation.Request.Headers = request.Header
	conversation.Request.RemoteAddr = request.RemoteAddr

	if err := readRequestBody(conversation, request); err != nil {
		return err
	}
	if err := readRequestForm(conversation, request); err != nil {
		return err
	}
	if err := readRequestMultipartForm(conversation, request); err != nil {
		return err
	}

	return nil
}

func readRequestBody(conversation *Conversation, request *http.Request) error {
	if strings.HasPrefix(request.Header.Get("content-type"), "application/x-www-form-urlencoded") || strings.HasPrefix(request.Header.Get("content-type"), "multipart/form-data") {
		return nil
	}

	buf := &bytes.Buffer{}
	if _, err := buf.ReadFrom(request.Body); err != nil {
		return err
	}
	conversation.Request.Body = buf.String()

	return nil
}

func readRequestForm(conversation *Conversation, request *http.Request) error {
	if !strings.HasPrefix(request.Header.Get("content-type"), "application/x-www-form-urlencoded") {
		return nil
	}

	if err := request.ParseForm(); err != nil {
		return err
	}
	conversation.Request.Form = request.Form

	return nil
}

func readRequestMultipartForm(conversation *Conversation, request *http.Request) error {
	if !strings.HasPrefix(request.Header.Get("content-type"), "multipart/form-data") {
		return nil
	}

	if err := request.ParseMultipartForm(1024); err != nil {
		return err
	}
	conversation.Request.Form = request.Form
	for name, files := range request.MultipartForm.File {
		tempfiles := make([]string, len(files))
		for i, file := range files {
			dst, err := ioutil.TempFile("", "gomoku")
			if err != nil {
				return err
			}
			defer dst.Close()
			src, err := file.Open()
			if err != nil {
				return err
			}
			defer src.Close()
			if _, err := io.Copy(dst, src); err != nil {
				return err
			}
			tempfiles[i] = dst.Name()
		}
		conversation.Request.Form[name] = tempfiles
	}

	return nil
}
