package dotenv

import (
	"github.com/j3llc/dotenv"
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	err := dotenv.Load()
	if err != nil {
		t.Error(err)
	}
	v := os.Getenv("foo")
	if v == "" {
		t.Errorf("Variable foo not set")
	}
	if v != "bar" {
		t.Errorf("Expected bar got %s", v)
	}
}

func TestLoadPath(t *testing.T) {
	err := dotenv.LoadPath("./.env")
	if err != nil {
		t.Error(err)
	}
	v := os.Getenv("foo")
	if v == "" {
		t.Errorf("Variable foo not set")
	}
}

func ExampleLoad() {
	if err := dotenv.Load(); err != nil {
		// handle error
	}
}

func ExampleLoadPath() {
	if err := dotenv.LoadPath("path_to_env"); err != nil {
		// hande error
	}
}
