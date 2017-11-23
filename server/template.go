package server

import (
	"bytes"
	"text/template"
)

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
