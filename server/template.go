package server

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"gopkg.in/yaml.v2"
)

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
