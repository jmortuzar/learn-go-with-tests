package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Matias")

	got := buffer.String()
	want := "Hello, Matias"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
