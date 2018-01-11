package server

import (
	"net/http"
	"reflect"
	"testing"
)

func TestConfig_SelectConfigItem__no_item(t *testing.T) {
	config := &Config{
		ConfigItem{Request{Method: "method1", Route: "/route1", Headers: map[string]string{"key1": "value1"}}, Command{}, Response{}},
	}

	method := "method1"
	route := "/route1"
	headers := http.Header{}

	req, cmd, res, found := config.SelectConfigItem(method, route, headers)

	if found {
		t.Fatalf("found wrong config item. req: %+v, cmd: %+v, res: %+v", req, cmd, res)
	}
}

func TestConfig_SelectConfigItem__last_item(t *testing.T) {
	config := &Config{
		ConfigItem{Request{Method: "method2"}, Command{Path: "path2"}, Response{Status: 2}},
		ConfigItem{Request{}, Command{Path: "path1"}, Response{Status: 1}},
	}

	method := "method1"
	route := "/route1"
	headers := http.Header{}

	_, cmd, res, found := config.SelectConfigItem(method, route, headers)

	if !found {
		t.Fatal("not found config item")
	}
	assertCommand(cmd, &Command{Path: "path1"}, t)
	assertResponse(res, &Response{Status: 1}, t)
}

func TestConfig_SelectConfigItem__find_by_method__plain(t *testing.T) {
	config := &Config{
		ConfigItem{Request{Method: "method2"}, Command{Path: "path2"}, Response{Status: 2}},
		ConfigItem{Request{}, Command{Path: "path1"}, Response{Status: 1}},
	}

	method := "method2"
	route := "/route2"
	headers := http.Header{}
	headers.Set("key1", "value0")
	headers.Add("key1", "value1")
	headers.Set("key2", "value2")
	headers.Add("key2", "value3")

	_, cmd, res, found := config.SelectConfigItem(method, route, headers)

	if !found {
		t.Fatal("not found config item")
	}
	assertCommand(cmd, &Command{Path: "path2"}, t)
	assertResponse(res, &Response{Status: 2}, t)
}

func TestConfig_SelectConfigItem__find_by_method__ignore_case(t *testing.T) {
	config := &Config{
		ConfigItem{Request{Method: "METHOD2"}, Command{Path: "path2"}, Response{Status: 2}},
		ConfigItem{Request{}, Command{Path: "path1"}, Response{Status: 1}},
	}

	method := "method2"
	route := "/route2"
	headers := http.Header{}
	headers.Set("key1", "value0")
	headers.Add("key1", "value1")
	headers.Set("key2", "value2")
	headers.Add("key2", "value3")

	_, cmd, res, found := config.SelectConfigItem(method, route, headers)

	if !found {
		t.Fatal("not found config item")
	}
	assertCommand(cmd, &Command{Path: "path2"}, t)
	assertResponse(res, &Response{Status: 2}, t)
}

func TestConfig_SelectConfigItem__find_by_method__regex(t *testing.T) {
	config := &Config{
		ConfigItem{Request{Method: "method2x|method2"}, Command{Path: "path2"}, Response{Status: 2}},
		ConfigItem{Request{}, Command{Path: "path1"}, Response{Status: 1}},
	}

	method := "method2"
	route := "/route2"
	headers := http.Header{}
	headers.Set("key1", "value0")
	headers.Add("key1", "value1")
	headers.Set("key2", "value2")
	headers.Add("key2", "value3")

	_, cmd, res, found := config.SelectConfigItem(method, route, headers)

	if !found {
		t.Fatal("not found config item")
	}
	assertCommand(cmd, &Command{Path: "path2"}, t)
	assertResponse(res, &Response{Status: 2}, t)
}

func TestConfig_SelectConfigItem__find_by_route__file(t *testing.T) {
	config := &Config{
		ConfigItem{Request{Route: "/route2"}, Command{Path: "path2"}, Response{Status: 2}},
		ConfigItem{Request{}, Command{Path: "path1"}, Response{Status: 1}},
	}

	method := "method2"
	route := "/route2"
	headers := http.Header{}
	headers.Set("key1", "value0")
	headers.Add("key1", "value1")
	headers.Set("key2", "value2")
	headers.Add("key2", "value3")

	_, cmd, res, found := config.SelectConfigItem(method, route, headers)

	if !found {
		t.Fatal("not found config item")
	}
	assertCommand(cmd, &Command{Path: "path2"}, t)
	assertResponse(res, &Response{Status: 2}, t)
}

func TestConfig_SelectConfigItem__find_by_route__directory(t *testing.T) {
	config := &Config{
		ConfigItem{Request{Route: "/route2/"}, Command{Path: "path2"}, Response{Status: 2}},
		ConfigItem{Request{}, Command{Path: "path1"}, Response{Status: 1}},
	}

	method := "method2"
	route := "/route2"
	headers := http.Header{}
	headers.Set("key1", "value0")
	headers.Add("key1", "value1")
	headers.Set("key2", "value2")
	headers.Add("key2", "value3")

	_, cmd, res, found := config.SelectConfigItem(method, route, headers)

	if !found {
		t.Fatal("not found config item")
	}
	assertCommand(cmd, &Command{Path: "path2"}, t)
	assertResponse(res, &Response{Status: 2}, t)
}

func TestConfig_SelectConfigItem__find_by_headers__a_condition(t *testing.T) {
	config := &Config{
		ConfigItem{Request{Headers: map[string]string{"key1": "value1"}}, Command{Path: "path2"}, Response{Status: 2}},
		ConfigItem{Request{}, Command{Path: "path1"}, Response{Status: 1}},
	}

	method := "method2"
	route := "/route2"
	headers := http.Header{}
	headers.Set("key1", "value0")
	headers.Add("key1", "value1")
	headers.Set("key2", "value2")
	headers.Add("key2", "value3")

	_, cmd, res, found := config.SelectConfigItem(method, route, headers)

	if !found {
		t.Fatal("not found config item")
	}
	assertCommand(cmd, &Command{Path: "path2"}, t)
	assertResponse(res, &Response{Status: 2}, t)
}

func TestConfig_SelectConfigItem__find_by_headers__some_conditions(t *testing.T) {
	config := &Config{
		ConfigItem{Request{Headers: map[string]string{"key1": "value1", "key2": "value2"}}, Command{Path: "path2"}, Response{Status: 2}},
		ConfigItem{Request{}, Command{Path: "path1"}, Response{Status: 1}},
	}

	method := "method2"
	route := "/route2"
	headers := http.Header{}
	headers.Set("key1", "value0")
	headers.Add("key1", "value1")
	headers.Set("key2", "value2")
	headers.Add("key2", "value3")

	_, cmd, res, found := config.SelectConfigItem(method, route, headers)

	if !found {
		t.Fatal("not found config item")
	}
	assertCommand(cmd, &Command{Path: "path2"}, t)
	assertResponse(res, &Response{Status: 2}, t)
}

func assertCommand(actual, expected *Command, t *testing.T) {
	if !reflect.DeepEqual(actual.Env, expected.Env) {
		t.Fatalf("got %+v, but expected %+v", actual.Env, expected.Env)
	}
	if !reflect.DeepEqual(actual.Path, expected.Path) {
		t.Fatalf("got %+v, but expected %+v", actual.Path, expected.Path)
	}
	if !reflect.DeepEqual(actual.Args, expected.Args) {
		t.Fatalf("got %+v, but expected %+v", actual.Args, expected.Args)
	}
}

func assertResponse(actual, expected *Response, t *testing.T) {
	if !reflect.DeepEqual(actual.Status, expected.Status) {
		t.Fatalf("got %+v, but expected %+v", actual.Status, expected.Status)
	}
	if !reflect.DeepEqual(actual.Headers, expected.Headers) {
		t.Fatalf("got %+v, but expected %+v", actual.Headers, expected.Headers)
	}
	if !reflect.DeepEqual(actual.Body, expected.Body) {
		t.Fatalf("got %+v, but expected %+v", actual.Body, expected.Body)
	}
	if !reflect.DeepEqual(actual.Template, expected.Template) {
		t.Fatalf("got %+v, but expected %+v", actual.Template, expected.Template)
	}
	if !reflect.DeepEqual(actual.File, expected.File) {
		t.Fatalf("got %+v, but expected %+v", actual.File, expected.File)
	}
}
