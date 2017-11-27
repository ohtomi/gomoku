package server

import (
	"fmt"
	"io/ioutil"
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
			Command:  Command{Path: "python3", Args: []string{"-m", "foo", "{{ .URL }}"}, Env: []string{"GOMOKU=gomoku", "METHOD={{ .Method }}"}},
			Response: Response{Status: 200, Headers: map[string]string{"content-type": "application/json; charset=utf-8"}, Body: "{\"greeting\": \"{{ .CommandResult.Stdout }}\"}"},
		},
		{
			Request:  Request{Route: "/bar"},
			Command:  Command{Path: "python3", Args: []string{"-m", "bar"}},
			Response: Response{Template: "bar.tmpl"},
		},
		{
			Request:  Request{Route: "/baz"},
			Command:  Command{Path: "python3", Args: []string{"-m", "baz"}},
			Response: Response{File: "baz.txt"},
		},
	}

	if err := config.SaveToFile(filepath.Join(dirname, "gomoku.yml")); err != nil {
		return err
	}

	if err := unpackFromBindataToFile(filepath.Join(dirname, "foo.py"), "foo.py"); err != nil {
		return err
	}

	if err := unpackFromBindataToFile(filepath.Join(dirname, "bar.py"), "bar.py"); err != nil {
		return err
	}

	if err := unpackFromBindataToFile(filepath.Join(dirname, "bar.tmpl"), "bar.tmpl"); err != nil {
		return err
	}

	if err := unpackFromBindataToFile(filepath.Join(dirname, "baz.py"), "baz.py"); err != nil {
		return err
	}

	return nil
}

func unpackFromBindataToFile(filename, resource string) error {
	if _, err := os.Stat(filename); err == nil {
		return errors.New(fmt.Sprintf("unable to create file: %q already exists", filename))
	}

	dest, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer dest.Close()

	src, err := Assets.Open(resource)
	if err != nil {
		return errors.Wrap(err, resource)
	}
	defer src.Close()

	buf, err := ioutil.ReadAll(src)
	if err != nil {
		return err
	}

	if _, err := dest.Write(buf); err != nil {
		return err
	}

	return nil
}
