package main

import "testing"

func TestHello(t *testing.T) {

	t.Run("Saying Hello to people", func(t *testing.T) {
		got := Hello("Chris", "English")
		want := "Hello, Chris"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("Say 'Hello, world!!!' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, world!!!"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("Saying Hello to people in spanish", func(t *testing.T) {
		got := Hello("Chris", "Spanish")
		want := "Hola, Chris"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

}
