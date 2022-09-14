package service

import "time"

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls    int
	Duration time.Duration
}

func (s *SpySleeper) Sleep() {
	s.Calls++
	time.Sleep(s.Duration)
}
