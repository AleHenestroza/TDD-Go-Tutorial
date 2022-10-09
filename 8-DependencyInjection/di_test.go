package dependencyinjection

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Alejandro")

	got := buffer.String()
	want := "Hello, Alejandro"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
