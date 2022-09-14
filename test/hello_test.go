package test

import (
	"daily_test_go/service"
	"testing"
)

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got string, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%q', want '%q'", got, want)
		}
	}

	t.Run("Say hello to Fan", func(t *testing.T) {
		got := service.Hello("Fan", "French")
		want := "Bonjour, Fan"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Say hello to empty", func(t *testing.T) {
		got := service.Hello("Fan", "")
		want := "Hello, Fan"
		assertCorrectMessage(t, got, want)
	})
}
