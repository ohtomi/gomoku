package server

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

func CreateScaffold(dirname string) error {
	if _, err := os.Stat(dirname); err == nil {
		return errors.New(fmt.Sprintf("unable to create directory: %q already exists", dirname))
	}

	if err := os.MkdirAll(dirname, os.ModePerm); err != nil {
		return err
	}

	config := &Config{
		{
			Request:  Request{Route: "/foo", Method: "get|post"},
			Command:  Command{Path: "python3", Args: []string{"-m", "foo.py", "{{ URL }}"}},
			Response: Response{Status: 200, Headers: map[string]string{"content-type": "application/json; charset=utf-8"}, Body: "{\"greeting\": \"{{ .CommandResult.Stdout }}\""},
		},
		{
			Request:  Request{Route: "/bar"},
			Command:  Command{Path: "python3", Args: []string{"-m", "bar.py"}},
			Response: Response{Template: "bar.tmpl"},
		},
		{
			Request:  Request{Route: "/baz"},
			Command:  Command{Path: "python3", Args: []string{"-m", "baz.py"}},
			Response: Response{File: "baz.txt"},
		},
	}

	if err := config.SaveToFile(filepath.Join(dirname, "gomoku.yml")); err != nil {
		return err
	}

	return nil
}
