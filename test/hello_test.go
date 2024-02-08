package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("user name 1", func(t *testing.T) {
		name := "Aziz"
		want := "Hello " + name
		assertHello(t, name, want)
	})

	t.Run("user name 2", func(t *testing.T) {
		name := "Nurmat"
		want := "Hello " + name
		assertHello(t, name, want)
	})

	t.Run("empty case", func(t *testing.T) {
		name := ""
		want := "Hello World!"
		assertHello(t, name, want)
	})

	t.Run("number case", func(t *testing.T) {
		name := "234567"
		want := "Hello Adam"
		assertHello(t, name, want)
	})
}

func assertHello(t *testing.T, name, want string) {
	t.Helper()
	msg := Hello(name)
	if want != msg {
		t.Errorf("expectd %q, but got: %q", want, msg)
	}
}
