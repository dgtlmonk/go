package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		/*
			t.Helper() is needed to tell the test suite that this method is a helper.
			By doing this when it fails the line number reported
			will be in our function call rather than inside our test helper.
		*/
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying Sup to Chris", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Sup Chris! We're back!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying Sup Homie! We're back! when empty params", func(t *testing.T) {
		got := Hello("", "")
		want := "Sup Homie! We're back!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying Hola, Elodie in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola Elodie! We're back!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying Bonjour, Bonjour in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "French")
		want := "Bonjour Elodie! We're back!"

		assertCorrectMessage(t, got, want)
	})

}
