package test

import (
	"bytes"
	"daily_test_go/service"
	"testing"
	"time"
)

func TestCountDown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySleeper := &service.SpySleeper{Duration: 1 * time.Second}

	service.CountDown(buffer, spySleeper, 3, "Go")
	want := `3
2
1
Go`
	got := buffer.String()

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}

	if spySleeper.Calls != 4 {
		t.Errorf("not enough calls to sleeper, want 4 got %d", spySleeper.Calls)
	}
}
