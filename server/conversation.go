package server

import (
	"net/http"
	"net/url"
)

type Conversation struct {
	Request RequestInConversation
	Command CommandInConversation
}

type RequestInConversation struct {
	Method     string
	URL        *url.URL
	Headers    map[string][]string
	RemoteAddr string
	Body       string
	Form       map[string][]string
}

type CommandInConversation struct {
	Env    []string
	Path   string
	Args   []string
	Dir    string
	Stdout string
	Stderr string
}

func TryConversation(request *Request, command *Command, response *Response, reporter Reporter, w http.ResponseWriter, r *http.Request) {
	conversation := &Conversation{}

	if request != nil {
		if err := request.Transform(conversation, r); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			reporter.Error(err.Error())
			return
		}

		if reporter.IsEnabled() {
			reporter.Infof("      body => %s", conversation.Request.Body)
			reporter.Infof("      form => %s", conversation.Request.Form)
		}
	}

	if command != nil {
		if err := command.Execute(conversation); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			reporter.Error(err.Error())
			return
		}
	}
	if response != nil {
		if err := response.Write(conversation, w); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			reporter.Error(err.Error())
			return
		}
	}
}
