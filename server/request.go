package server

import (
	"net/http"
)

func (r *Request) Transform(conversation *Conversation, request *http.Request) error {
	conversation.Request.Method = request.Method
	conversation.Request.URL = request.URL
	conversation.Request.Headers = request.Header
	conversation.Request.Form = request.Form
	conversation.Request.RemoteAddr = request.RemoteAddr

	return nil
}
