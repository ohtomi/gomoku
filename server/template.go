package server

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"strings"
	"text/template"

	"gopkg.in/yaml.v2"
)

func (c *Conversation) GetByKey(values map[string][]string, key string) interface{} {
	return values[key]
}

func (c *Conversation) GetByIndex(values []string, index int) interface{} {
	return values[index]
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

func (r *RequestInConversation) BodyToJson() interface{} {
	var v interface{}
	dec := json.NewDecoder(strings.NewReader(r.Body))
	if err := dec.Decode(&v); err != nil {
		return nil
	}
	return v
}

func (r *RequestInConversation) BodyToYaml() interface{} {
	var v interface{}
	if err := yaml.Unmarshal([]byte(r.Body), &v); err != nil {
		return nil
	}
	return v
}

func (c *CommandInConversation) StdoutToJson() interface{} {
	var v interface{}
	dec := json.NewDecoder(strings.NewReader(c.Stdout))
	if err := dec.Decode(&v); err != nil {
		return nil
	}
	return v
}

func (c *CommandInConversation) StdoutToYaml() interface{} {
	var v interface{}
	if err := yaml.Unmarshal([]byte(c.Stdout), &v); err != nil {
		return nil
	}
	return v
}

func (c *CommandInConversation) StderrToJson() interface{} {
	var v interface{}
	dec := json.NewDecoder(strings.NewReader(c.Stderr))
	if err := dec.Decode(&v); err != nil {
		return nil
	}
	return v
}

func (c *CommandInConversation) StderrToYaml() interface{} {
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
