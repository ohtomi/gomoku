package server

import (
	"os"
	"testing"
)

func TestCreateScaffold(t *testing.T) {
	os.Chdir("..")
	os.RemoveAll("testdata/scaffold")

	CreateScaffold("testdata/scaffold")

	if _, err := os.Stat("testdata/scaffold"); err != nil {
		t.Fatal(err.Error())
	}
	if _, err := os.Stat("testdata/scaffold/gomoku.yml"); err != nil {
		t.Fatal(err.Error())
	}
	if _, err := os.Stat("testdata/scaffold/foo.py"); err != nil {
		t.Fatal(err.Error())
	}
	if _, err := os.Stat("testdata/scaffold/bar.py"); err != nil {
		t.Fatal(err.Error())
	}
	if _, err := os.Stat("testdata/scaffold/bar.tmpl"); err != nil {
		t.Fatal(err.Error())
	}
	if _, err := os.Stat("testdata/scaffold/baz.py"); err != nil {
		t.Fatal(err.Error())
	}
	if _, err := os.Stat("testdata/scaffold/static/html/index.html"); err != nil {
		t.Fatal(err.Error())
	}
	if _, err := os.Stat("testdata/scaffold/static/js/page.js"); err != nil {
		t.Fatal(err.Error())
	}
}
