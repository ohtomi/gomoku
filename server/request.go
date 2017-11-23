package server

import (
	"net/http"
)

func (r *Request) Transform(conversation *Conversation, request *http.Request) error {
	conversation.Method = request.Method
	conversation.URL = request.URL.String()
	conversation.Headers = request.Header
	conversation.Form = request.Form
	conversation.RemoteAddr = request.RemoteAddr

	return nil
}
