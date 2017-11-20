package server

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config []struct {
	Request  Request
	Command  Command
	Response Response
}

type Command struct {
	Path string
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
