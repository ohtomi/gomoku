package server

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func TestConversation_GetByKey(t *testing.T) {
	conversation := &Conversation{}

	values := map[string][]string{"key1": []string{"value1"}, "key2": []string{"value2a", "value2b"}}
	assertSlice(conversation.GetByKey(values, "key1"), []string{"value1"}, t)
	assertSlice(conversation.GetByKey(values, "key2"), []string{"value2a", "value2b"}, t)
	assertSlice(conversation.GetByKey(values, "key3"), nil, t)
}

func TestConversation_GetByIndex(t *testing.T) {
	conversation := &Conversation{}

	values := []string{"value1", "value2"}
	assertString(conversation.GetByIndex(values, 0), "value1", t)
	assertString(conversation.GetByIndex(values, 1), "value2", t)
	assertString(conversation.GetByIndex(values, 2), "", t)
}

func TestConversation_JoinWith(t *testing.T) {
	conversation := &Conversation{}

	values := []string{"value1", "value2"}
	assertString(conversation.JoinWith(values, ""), "value1value2", t)
	assertString(conversation.JoinWith(values, "_"), "value1_value2", t)
	assertString(conversation.JoinWith(values, "_!_"), "value1_!_value2", t)
}

func TestConversation_ReadFile(t *testing.T) {
	tempfile, err := createTempFile("template", []byte("content1\ncontent1\ncontent1"))
	if err != nil {
		t.Fatalf("unable to prepare test. cause: %q", err.Error())
	}
	defer os.Remove(tempfile.Name())

	conversation := &Conversation{
		Request: RequestInConversation{
			Form: map[string][]string{"uploaded-file": []string{tempfile.Name()}},
		},
	}

	assertString(conversation.ReadFile("not-uploaded-file"), "", t)
	assertString(conversation.ReadFile("uploaded-file"), "content1\ncontent1\ncontent1", t)
}

func TestConversation_ReadFiles(t *testing.T) {
	tempfile1, err := createTempFile("template", []byte("content1\ncontent1\ncontent1"))
	if err != nil {
		t.Fatalf("unable to prepare test. cause: %q", err.Error())
	}
	defer os.Remove(tempfile1.Name())

	tempfile2, err := createTempFile("template", []byte("content2\ncontent2\ncontent2"))
	if err != nil {
		t.Fatalf("unable to prepare test. cause: %q", err.Error())
	}
	defer os.Remove(tempfile2.Name())

	conversation := &Conversation{
		Request: RequestInConversation{
			Form: map[string][]string{"uploaded-file": []string{tempfile1.Name(), tempfile2.Name()}},
		},
	}

	assertString(conversation.ReadFile("not-uploaded-file"), "", t)
	assertSlice(conversation.ReadFiles("uploaded-file"),
		[]string{"content1\ncontent1\ncontent1", "content2\ncontent2\ncontent2"}, t)
}

func createTempFile(prefix string, content []byte) (*os.File, error) {
	tempfile, err := ioutil.TempFile("", prefix)
	if err != nil {
		return nil, err
	}

	if _, err := tempfile.Write(content); err != nil {
		return nil, err
	}

	if err := tempfile.Close(); err != nil {
		return nil, err
	}

	return tempfile, nil
}

func assertSlice(actual, expected []string, tb testing.TB) {
	tb.Helper()
	if !reflect.DeepEqual(actual, expected) {
		tb.Fatalf("got %+v, but expected %+v", actual, expected)
	}
}

func assertString(actual, expected string, tb testing.TB) {
	tb.Helper()
	if actual != expected {
		tb.Fatalf("got %+v, but expected %+v", actual, expected)
	}
}
