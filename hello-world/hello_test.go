package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("Saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "English")
		want := "Hello, Chris!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Say 'Hello, World!' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, World!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Say 'Hola, Elodie!' when an empty string is supplied", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Say 'Bonjour, Monde!' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "French")
		want := "Bonjour, Monde!"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
