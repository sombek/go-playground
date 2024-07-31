package maps_and_dicts

import (
	"errors"
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, err := dictionary.Search("test")
		want := "this is just a test"
		if err != nil {
			t.Fatal("error")
		}
		assertStrings(t, got, want)
	})

	t.Run("Unknown Word", func(t *testing.T) {
		_, err := dictionary.Search("unkown")
		if err == nil {
			t.Fatal("Expected an error, but got nothing")
		}
		assertError(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}
	dictionary.Add("test", "this is just a test")

	want := "this is just a test"
	got, err := dictionary.Search("test")

	if err != nil {
		t.Fatal("should find added word:", err)
	}
	assertStrings(t, got, want)
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if !errors.Is(got, want) {
		t.Errorf("got '%q want %q", got, want)
	}
}
func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got '%q want %q", got, want)
	}
}
