package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{} //the Buffer type from the bytes package implements the Writer interface, because it has the method Write(p []byte)(n int, err error)
	//So we'll use it in our test to send in as our Writer and then we can check what was written to it after we invoke Greet()
	Greet(&buffer, "Baby Shark")

	got := buffer.String()
	want := "Hello, Baby Shark"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
