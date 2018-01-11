package server

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"regexp"
	"strings"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

type Config []ConfigItem

type ConfigItem struct {
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

func (c *Config) SelectConfigItem(method, route string, headers http.Header) (*Request, *Command, *Response, bool) {
	for _, element := range *c {
		if len(element.Request.Method) != 0 {
			if !regexp.MustCompile(fmt.Sprintf("(?i)%s", element.Request.Method)).MatchString(method) {
				continue
			}
		}
		if len(element.Request.Route) != 0 {
			if !strings.HasSuffix(route, "/") {
				route = fmt.Sprintf("%s/", route)
			}
			if !regexp.MustCompile(fmt.Sprintf("^%s/", strings.TrimRight(element.Request.Route, "/"))).MatchString(route) {
				continue
			}
		}
		if len(element.Request.Headers) != 0 {
			if !matchHeaders(element.Request.Headers, headers) {
				continue
			}
		}
		return &element.Request, &element.Command, &element.Response, true
	}
	return nil, nil, nil, false
}

func matchHeaders(expected map[string]string, actual http.Header) bool {
	for k1, v1 := range expected {
		if v2array, ok := actual[http.CanonicalHeaderKey(k1)]; ok {
			for _, v2 := range v2array {
				if regexp.MustCompile(fmt.Sprintf("^%s", v1)).MatchString(v2) {
					return true
				}
			}
		}
		return false
	}
	return false
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

func (c *Config) EqualTo(other *Config) bool {
	if len(*c) != len(*other) {
		return false
	}

	return reflect.DeepEqual(*c, *other)
}
