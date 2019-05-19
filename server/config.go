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
	Upgrade  Upgrade  `yaml:",omitempty"`
	Robots   Robots   `yaml:",omitempty"`
	Command  Command  `yaml:",omitempty"`
	Response Response `yaml:",omitempty"`
}

type Request struct {
	Method  string            `yaml:",omitempty"`
	Route   string            `yaml:",omitempty"`
	Headers map[string]string `yaml:",omitempty"`
}

type Upgrade struct {
	Protocol string `yaml:",omitempty"`
}

type Robots []RobotItem

type RobotItem struct {
	Source    SourceOrSink
	Transform Transform
	Sink      SourceOrSink
}

type SourceOrSink struct {
	Type     int
	Body     string `yaml:",omitempty"`
	Template string `yaml:",omitempty"`
	File     string `yaml:",omitempty"`
}

type Transform struct {
	Env  []string `yaml:",omitempty"`
	Path string   `yaml:",omitempty"`
	Args []string `yaml:",omitempty"`
}

type Command struct {
	Env  []string `yaml:",omitempty"`
	Path string   `yaml:",omitempty"`
	Args []string `yaml:",omitempty"`
}

type Response struct {
	Status   int               `yaml:",omitempty"`
	Headers  map[string]string `yaml:",omitempty"`
	Cookies  []Cookie          `yaml:",omitempty"`
	Body     string            `yaml:",omitempty"`
	Template string            `yaml:",omitempty"`
	File     string            `yaml:",omitempty"`
}

type Cookie struct {
	Name  string
	Value string

	Path     string `yaml:",omitempty"`
	Domain   string `yaml:",omitempty"`
	Expires  string `yaml:",omitempty"`
	MaxAge   int    `yaml:",omitempty"`
	Secure   bool   `yaml:",omitempty"`
	HttpOnly bool   `yaml:",omitempty"`
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

func (c *Config) SelectConfigItem(method, route string, headers http.Header) (*ConfigItem, bool) {
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
		return &element, true
	}
	return nil, false
}

func matchHeaders(want map[string]string, got http.Header) bool {
	for k1, v1 := range want {
		if v2array, ok := got[http.CanonicalHeaderKey(k1)]; ok {
			if !containsValue(v1, v2array) {
				return false
			}
		} else {
			return false
		}
	}
	return true
}

func containsValue(want string, got []string) bool {
	matcher := regexp.MustCompile(fmt.Sprintf("^%s", want))
	for _, v := range got {
		if matcher.MatchString(v) {
			return true
		}
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
