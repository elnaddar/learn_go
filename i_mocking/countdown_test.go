package main

import (
	"bytes"
	"slices"
	"testing"
	"time"
)

type SpyCountdownOps struct {
	Calls []string
}

const (
	write = "write"
	sleep = "sleep"
)

func (s *SpyCountdownOps) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOps) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func TestCountdown(t *testing.T) {
	t.Run("prints 3 to go", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spySleeper := &SpyCountdownOps{}
		Countdown(buffer, spySleeper)

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("sleep before every point", func(t *testing.T) {
		spySleeper := &SpyCountdownOps{}
		Countdown(spySleeper, spySleeper)

		got := spySleeper.Calls
		want := []string{write, sleep, write, sleep, write, sleep, write}

		if !slices.Equal(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}
}
