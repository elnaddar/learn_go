package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	t.Run("test greet", func(t *testing.T) {
		buffer := bytes.Buffer{}
		Greet(&buffer, "Ahmed")

		got := buffer.String()
		want := "Hello, Ahmed"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
