package server

import (
	"fmt"
	"io/ioutil"
	"regexp"

	"gopkg.in/yaml.v2"
)

type Config []struct {
	Request  Request
	Command  Command
	Response Response
}

type Command struct {
	Path string
	Args []string
}

type Request struct {
	Route  string
	Method string
}

type Response struct {
	Status   int
	Headers  map[string]string
	Body     string
	Template string
	File     string
}

func NewConfig(yamlFile string) (*Config, error) {
	var (
		config Config
	)

	buf, err := ioutil.ReadFile(yamlFile)
	if err != nil {
		return nil, err
	}

	if err = yaml.Unmarshal(buf, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

func (c *Config) SelectConfigItem(method, path string) (*Request, *Command, *Response) {
	for _, element := range *c {
		if len(element.Request.Method) != 0 {
			if !regexp.MustCompile(fmt.Sprintf("(?i)%s", element.Request.Method)).MatchString(method) {
				continue
			}
		}
		if len(element.Request.Route) != 0 {
			if !regexp.MustCompile(element.Request.Route).MatchString(path) {
				continue
			}
		}
		return &element.Request, &element.Command, &element.Response
	}
	return nil, nil, nil
}
