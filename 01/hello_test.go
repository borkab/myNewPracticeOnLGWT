package main

import "testing"

func TestHello(t *testing.T) {

	t.Run("with name", func(t *testing.T) {
		b := "Programmer Girl"
		got := Hello(b, "")
		want := "Hello, Programmer Girl"
		asseretCorrectMessage(t, got, want)
	})

	t.Run("with an empty string", func(t *testing.T) {
		c := ""
		got := Hello(c, "")
		want := "Hello, World"
		asseretCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		d := "Programmer Chica"
		sp := "Spanish"
		got := Hello(d, sp)
		want := "Hola, Programmer Chica"
		asseretCorrectMessage(t, got, want)

	})

	t.Run("in French", func(t *testing.T) {
		e := "Programmer Modmoiselle"
		fr := "French"
		got := Hello(e, fr)
		want := "Bonjour, Programmer Modmoiselle"
		asseretCorrectMessage(t, got, want)
	})

	t.Run("in Hungarian", func(t *testing.T) {
		e := "Programozo Lany"
		hun := "Hungarian"
		got := Hello(e, hun)
		want := "Szia, Programozo Lany"
		asseretCorrectMessage(t, got, want)
	})
}

func asseretCorrectMessage(t testing.TB, got, want string) {
	t.Helper() //is needed to tell the test suite that this method is a helper.
	//doing this when it fails the line number reported will be in our function call
	//rather than inside our test helper
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
