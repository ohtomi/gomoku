package server

import (
	"net/http"
	"reflect"
	"testing"
)

func TestConfig_SelectConfigItem__no_item(t *testing.T) {
	tests := []struct {
		name   string
		config *Config
	}{
		{
			"not matched predicate",
			&Config{
				ConfigItem{
					Request: Request{
						Method:  "method.request",
						Route:   "/route/request",
						Headers: map[string]string{"key1": "value1"},
					},
					Command:  Command{},
					Response: Response{},
				},
			},
		},
	}

	method := "method.request"
	route := "/route/request"
	headers := http.Header{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			item, found := tt.config.SelectConfigItem(method, route, headers)

			if found {
				t.Fatalf("found wrong config item: %+v", item)
			}
		})
	}
}

func TestConfig_SelectConfigItem__last_item(t *testing.T) {
	tests := []struct {
		name   string
		config *Config
	}{
		{
			"not matched predicate",
			&Config{
				ConfigItem{
					Request:  Request{Method: "method.request"},
					Command:  Command{Path: "path.request"},
					Response: Response{File: "file.request"},
				},
				ConfigItem{
					Command:  Command{Path: "path.last"},
					Response: Response{File: "file.last"},
				},
			},
		},
	}

	method := "method.other"
	route := "/route/other"
	headers := http.Header{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			item, found := tt.config.SelectConfigItem(method, route, headers)

			if !found {
				t.Fatal("not found config item")
			}
			assertCommand(item.Command, Command{Path: "path.last"}, t)
			assertResponse(item.Response, Response{File: "file.last"}, t)
		})
	}
}

func TestConfig_SelectConfigItem__find_by_method(t *testing.T) {
	tests := []struct {
		name   string
		config *Config
	}{
		{
			"basic method predicate",
			&Config{
				ConfigItem{
					Request:  Request{Method: "method.request"},
					Command:  Command{Path: "path.request"},
					Response: Response{File: "file.request"},
				},
				ConfigItem{
					Command:  Command{Path: "path.last"},
					Response: Response{File: "file.last"},
				},
			},
		},
		{
			"method predicate ignores case sensitivity",
			&Config{
				ConfigItem{
					Request:  Request{Method: "METHOD.REQUEST"},
					Command:  Command{Path: "path.request"},
					Response: Response{File: "file.request"},
				},
				ConfigItem{
					Command:  Command{Path: "path.last"},
					Response: Response{File: "file.last"},
				},
			},
		},
		{
			"method predicate accepts regex pattern",
			&Config{
				ConfigItem{
					Request:  Request{Method: "method.request.extra|method.request"},
					Command:  Command{Path: "path.request"},
					Response: Response{File: "file.request"},
				},
				ConfigItem{
					Command:  Command{Path: "path.last"},
					Response: Response{File: "file.last"},
				},
			},
		},
	}

	method := "method.request"
	route := "/route/other"
	headers := http.Header{}
	headers.Set("key1", "value1")
	headers.Add("key1", "value1.extra")
	headers.Set("key2", "value2")
	headers.Add("key2", "value2.extra")

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			item, found := tt.config.SelectConfigItem(method, route, headers)

			if !found {
				t.Fatal("not found config item")
			}
			assertCommand(item.Command, Command{Path: "path.request"}, t)
			assertResponse(item.Response, Response{File: "file.request"}, t)
		})
	}
}

func TestConfig_SelectConfigItem__find_by_route(t *testing.T) {
	tests := []struct {
		name   string
		config *Config
	}{
		{
			"basic route predicate",
			&Config{
				ConfigItem{
					Request:  Request{Route: "/route/common"},
					Command:  Command{Path: "path.common"},
					Response: Response{File: "file.common"},
				},
				ConfigItem{
					Command:  Command{Path: "path.last"},
					Response: Response{File: "file.last"},
				},
			},
		},
		{
			"route predicate ignores trailing slash",
			&Config{
				ConfigItem{
					Request:  Request{Route: "/route/common/"},
					Command:  Command{Path: "path.common"},
					Response: Response{File: "file.common"},
				},
				ConfigItem{
					Command:  Command{Path: "path.last"},
					Response: Response{File: "file.last"},
				},
			},
		},
	}

	method := "method.other"
	route := "/route/common"
	headers := http.Header{}
	headers.Set("key1", "value1")
	headers.Add("key1", "value1.extra")
	headers.Set("key2", "value2")
	headers.Add("key2", "value2.extra")

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			item, found := tt.config.SelectConfigItem(method, route, headers)

			if !found {
				t.Fatal("not found config item")
			}
			assertCommand(item.Command, Command{Path: "path.common"}, t)
			assertResponse(item.Response, Response{File: "file.common"}, t)
		})
	}
}

func TestConfig_SelectConfigItem__find_by_headers(t *testing.T) {
	tests := []struct {
		name   string
		config *Config
	}{
		{
			"single entry in headers predicate",
			&Config{
				ConfigItem{
					Request:  Request{Headers: map[string]string{"key1": "value1"}},
					Command:  Command{Path: "path.request"},
					Response: Response{File: "file.request"},
				},
				ConfigItem{
					Command:  Command{Path: "path.last"},
					Response: Response{File: "file.last"},
				},
			},
		},
		{
			"some entries in headers predicate",
			&Config{
				ConfigItem{
					Request:  Request{Headers: map[string]string{"key1": "value1", "key2": "value2"}},
					Command:  Command{Path: "path.request"},
					Response: Response{File: "file.request"},
				},
				ConfigItem{
					Command:  Command{Path: "path.last"},
					Response: Response{File: "file.last"},
				},
			},
		},
	}

	method := "method.other"
	route := "/route/other"
	headers := http.Header{}
	headers.Set("key1", "value1")
	headers.Add("key1", "value1.extra")
	headers.Set("key2", "value2")
	headers.Add("key2", "value2.extra")

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			item, found := tt.config.SelectConfigItem(method, route, headers)

			if !found {
				t.Fatal("not found config item")
			}
			assertCommand(item.Command, Command{Path: "path.request"}, t)
			assertResponse(item.Response, Response{File: "file.request"}, t)
		})
	}
}

func assertCommand(actual, expected Command, tb testing.TB) {
	tb.Helper()
	if !reflect.DeepEqual(actual.Env, expected.Env) {
		tb.Fatalf("got %+v, but expected %+v", actual.Env, expected.Env)
	}
	if !reflect.DeepEqual(actual.Path, expected.Path) {
		tb.Fatalf("got %+v, but expected %+v", actual.Path, expected.Path)
	}
	if !reflect.DeepEqual(actual.Args, expected.Args) {
		tb.Fatalf("got %+v, but expected %+v", actual.Args, expected.Args)
	}
}

func assertResponse(actual, expected Response, tb testing.TB) {
	tb.Helper()
	if !reflect.DeepEqual(actual.Status, expected.Status) {
		tb.Fatalf("got %+v, but expected %+v", actual.Status, expected.Status)
	}
	if !reflect.DeepEqual(actual.Headers, expected.Headers) {
		tb.Fatalf("got %+v, but expected %+v", actual.Headers, expected.Headers)
	}
	if !reflect.DeepEqual(actual.Body, expected.Body) {
		tb.Fatalf("got %+v, but expected %+v", actual.Body, expected.Body)
	}
	if !reflect.DeepEqual(actual.Template, expected.Template) {
		tb.Fatalf("got %+v, but expected %+v", actual.Template, expected.Template)
	}
	if !reflect.DeepEqual(actual.File, expected.File) {
		tb.Fatalf("got %+v, but expected %+v", actual.File, expected.File)
	}
}
