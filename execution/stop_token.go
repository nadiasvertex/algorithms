package execution

import (
	"sync"
	"sync/atomic"
)

type StopToken struct {
	stopped    atomic.Bool
	stopLock   sync.Mutex
	stopSignal chan bool
}

func makeStopToken() *StopToken {
	return &StopToken{
		stopSignal: make(chan bool, 1),
	}
}

func (s *StopToken) Stop() {
	if !s.stopped.Load() {
		s.stopLock.Lock()
		defer s.stopLock.Unlock()
		if !s.stopped.Load() {
			s.stopped.Store(true)
			close(s.stopSignal)
		}
	}
}

func (s *StopToken) Stopped() bool {
	return s.stopped.Load()
}

func (s *StopToken) StopChannel() chan bool {
	return s.stopSignal
}
