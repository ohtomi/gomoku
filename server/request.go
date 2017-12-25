package server

import (
	"io"
	"io/ioutil"
	"net/http"
)

func (r *Request) Transform(conversation *Conversation, request *http.Request) error {
	conversation.Request.Method = request.Method
	conversation.Request.URL = request.URL
	conversation.Request.Headers = request.Header
	conversation.Request.RemoteAddr = request.RemoteAddr

	if r.Method != "GET" && r.MultiPart {
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
	} else {
		if err := request.ParseForm(); err != nil {
			return err
		}
		conversation.Request.Form = request.Form
	}

	return nil
}
