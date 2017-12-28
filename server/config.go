package server

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

type Config []struct {
	Request  Request  `yaml:",omitempty"`
	Command  Command  `yaml:",omitempty"`
	Response Response `yaml:",omitempty"`
}

type Command struct {
	Env  []string `yaml:",omitempty"`
	Path string   `yaml:",omitempty"`
	Args []string `yaml:",omitempty"`
}

type Request struct {
	Method  string            `yaml:",omitempty"`
	Route   string            `yaml:",omitempty"`
	Headers map[string]string `yaml:",omitempty"`
}

type Response struct {
	Status   int               `yaml:",omitempty"`
	Headers  map[string]string `yaml:",omitempty"`
	Body     string            `yaml:",omitempty"`
	Template string            `yaml:",omitempty"`
	File     string            `yaml:",omitempty"`
}

func NewConfig(filename string) (*Config, error) {
	var (
		config Config
	)

	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	if err = yaml.Unmarshal(buf, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

func (c *Config) SelectConfigItem(method, path string, headers http.Header) (*Request, *Command, *Response) {
	for _, element := range *c {
		if len(element.Request.Method) != 0 {
			if !regexp.MustCompile(fmt.Sprintf("(?i)%s", element.Request.Method)).MatchString(method) {
				continue
			}
		}
		if len(element.Request.Route) != 0 {
			if !strings.HasSuffix(path, "/") {
				path = fmt.Sprintf("%s/", path)
			}
			if !regexp.MustCompile(fmt.Sprintf("^%s/", strings.TrimRight(element.Request.Route, "/"))).MatchString(path) {
				continue
			}
		}
		if len(element.Request.Headers) != 0 {
			if !matchHeaders(element.Request.Headers, headers) {
				continue
			}
		}
		return &element.Request, &element.Command, &element.Response
	}
	return nil, nil, nil
}

func matchHeaders(expected map[string]string, actual http.Header) bool {
	for k, v := range expected {
		if !regexp.MustCompile(fmt.Sprintf("^%s", v)).MatchString(actual.Get(k)) {
			return false
		}
	}
	return true
}

func (c *Config) SaveToFile(filename string) error {
	if _, err := os.Stat(filename); err == nil {
		return errors.New(fmt.Sprintf("unable to create file: %q already exists", filename))
	}

	fd, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer fd.Close()

	b, err := yaml.Marshal(c)
	if err != nil {
		return err
	}
	if _, err := fd.Write(b); err != nil {
		return err
	}

	return nil
}

func (c *Config) ToYaml() ([]byte, error) {
	return yaml.Marshal(c)
}
