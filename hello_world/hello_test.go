package main

import "testing"

func TestHello(t *testing.T) {

	t.Run("Say hello to Ben", func(t *testing.T) {
		assetMessage(t, Hello("Ben", ""), "Hello Ben!")
	})

	t.Run("Say hello to Ben with some white spaces", func(t *testing.T) {
		assetMessage(t, Hello(" Ben ", ""), "Hello Ben!")
	})

	t.Run("Say hello to Anonymous", func(t *testing.T) {
		assetMessage(t, Hello("", ""), "Hello World!")
	})

	t.Run("Say hello to empty string", func(t *testing.T) {
		assetMessage(t, Hello("   ", ""), "Hello World!")
	})

	t.Run("Say hello in unknown language", func(t *testing.T) {
		assetMessage(t, Hello("Karl", "BLA"), "Hello Karl!")
	})

	t.Run("Say hola to Johann string", func(t *testing.T) {
		assetMessage(t, Hello("Johann", "ES"), "Hola Johann!")
	})

	t.Run("Say hallo to Johannes string", func(t *testing.T) {
		assetMessage(t, Hello("Johannes", " de "), "Hallo Johannes!")
	})

	t.Run("Say hola to Johann string", func(t *testing.T) {
		assetMessage(t, Hello("Johann", " ES "), "Hola Johann!")
	})

	t.Run("Say bonjour to Paul string with lower case", func(t *testing.T) {
		assetMessage(t, Hello("Paul", "fr"), "Bonjour Paul!")
	})

	t.Run("Say bonjour to empty string", func(t *testing.T) {
		assetMessage(t, Hello("Paul", "FR"), "Bonjour Paul!")
	})

}

func assetMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Got %q want %q", got, want)
	}

}
