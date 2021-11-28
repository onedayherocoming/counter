package sync_counter

import (
	"sync"
	"time"
)

type ISyncCounter interface {
	Get(string) int
	Init()
	Flush2broker(int, func())
	Incr(string, int)
}

func NewSyncCounter() ISyncCounter {
	sc := &SyncCounter{}
	sc.Init()
	return sc
}

type SyncCounter struct {
	buckets map[string]int
	mutex   sync.Mutex
}

func (s *SyncCounter) Init() {
	s.buckets = make(map[string]int)
}

func (s *SyncCounter) Get(key string) int {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.buckets[key]
}

func (s *SyncCounter) Incr(key string, value int) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.buckets[key] += value
}

func (s *SyncCounter) Flush2broker(mills int, f func()) {
	go func() {
		t := time.Tick(time.Duration(mills) * time.Millisecond)
		for {
			<-t
			f()
			s.Init()
		}
	}()
}
