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
