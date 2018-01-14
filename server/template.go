package server

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"strings"
	"text/template"

	"gopkg.in/yaml.v2"
)

func (c *Conversation) GetByKey(values map[string][]string, key string) []string {
	if value, ok := values[key]; ok {
		return value
	} else {
		return nil
	}
}

func (c *Conversation) GetByIndex(values []string, index int) string {
	if index < len(values) {
		return values[index]
	} else {
		return ""
	}
}

func (c *Conversation) JoinWith(values []string, sep string) string {
	return strings.Join(values, sep)
}

func (c *Conversation) ReadFile(filename string) string {
	contents := c.ReadFiles(filename)
	if len(contents) == 0 {
		return ""
	}
	return contents[0]
}

func (c *Conversation) ReadFiles(filename string) []string {
	contents := make([]string, len(c.Request.Form[filename]))
	for i, tempfile := range c.Request.Form[filename] {
		if buf, err := ioutil.ReadFile(tempfile); err != nil {
			contents[i] = ""
		} else {
			contents[i] = string(buf)
		}
	}
	return contents
}

func (r *RequestInConversation) ParseBodyAsJson() interface{} {
	var v interface{}
	dec := json.NewDecoder(strings.NewReader(r.Body))
	if err := dec.Decode(&v); err != nil {
		return nil
	}
	return v
}

func (r *RequestInConversation) ParseBodyAsYaml() interface{} {
	var v interface{}
	if err := yaml.Unmarshal([]byte(r.Body), &v); err != nil {
		return nil
	}
	return v
}

func (c *CommandInConversation) ParseStdoutAsJson() interface{} {
	var v interface{}
	dec := json.NewDecoder(strings.NewReader(c.Stdout))
	if err := dec.Decode(&v); err != nil {
		return nil
	}
	return v
}

func (c *CommandInConversation) ParseStdoutAsYaml() interface{} {
	var v interface{}
	if err := yaml.Unmarshal([]byte(c.Stdout), &v); err != nil {
		return nil
	}
	return v
}

func (c *CommandInConversation) ParseStderrAsJson() interface{} {
	var v interface{}
	dec := json.NewDecoder(strings.NewReader(c.Stderr))
	if err := dec.Decode(&v); err != nil {
		return nil
	}
	return v
}

func (c *CommandInConversation) ParseStderrAsYaml() interface{} {
	var v interface{}
	if err := yaml.Unmarshal([]byte(c.Stderr), &v); err != nil {
		return nil
	}
	return v
}

func ApplyTemplateText(name, text string, data interface{}) (string, error) {
	buf := &bytes.Buffer{}
	t, err := template.New(name).Parse(text)
	if err != nil {
		return "", err
	}
	if err := t.Execute(buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func ApplyTemplateFile(name, filename string, data interface{}) (string, error) {
	buf := &bytes.Buffer{}
	t, err := template.ParseFiles(filename)
	if err != nil {
		return "", err
	}
	if err := t.Execute(buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}
