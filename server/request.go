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

	var contentType string
	if header, ok := request.Header["content-type"]; ok {
		contentType = header[0]
	} else if header, ok := request.Header["Content-Type"]; ok {
		contentType = header[0]
	} else {
		contentType = "text/plain"
	}

	if err := r.readBody(conversation, request, contentType); err != nil {
		return err
	}
	if err := r.readForm(conversation, request, contentType); err != nil {
		return err
	}
	if err := r.readMultipartForm(conversation, request, contentType); err != nil {
		return err
	}

	return nil
}

func (r *Request) readBody(conversation *Conversation, request *http.Request, contentType string) error {
	if strings.HasPrefix(contentType, "application/x-www-form-urlencoded") || strings.HasPrefix(contentType, "multipart/form-data") {
		return nil
	}

	buf := &bytes.Buffer{}
	if _, err := buf.ReadFrom(request.Body); err != nil {
		return err
	}
	conversation.Request.Body = buf.String()

	return nil
}

func (r *Request) readForm(conversation *Conversation, request *http.Request, contentType string) error {
	if !strings.HasPrefix(contentType, "application/x-www-form-urlencoded") {
		return nil
	}

	if err := request.ParseForm(); err != nil {
		return err
	}
	conversation.Request.Form = request.Form

	return nil
}

func (r *Request) readMultipartForm(conversation *Conversation, request *http.Request, contentType string) error {
	if !strings.HasPrefix(contentType, "multipart/form-data") {
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
