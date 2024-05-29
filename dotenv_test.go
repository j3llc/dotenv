package dotenv

import (
	"testing"
)

func TestStartsWith(t *testing.T) {
	got := startsWith("#hello world", "#")
	want := true
	if got != want {
		t.Errorf("Got %v wanted %v", got, want)
	}
}

func TestRemoveWrapper(t *testing.T) {
	got := removeWrapper("\"hello world\"", '"')
	want := "hello world"
	if got != want {
		t.Errorf("Expected %q but got %q", want, got)
	}
}
