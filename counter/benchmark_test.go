package counter

import (
	"strconv"
	"sync"
	"testing"
)

func BenchmarkSyncCounter_Incr(b *testing.B) {
	counter := NewCounter()
	for i := 0; i < b.N; i++ {
		counter.Incr(strconv.Itoa(i), 1)
	}
}

func BenchmarkSyncCounter_Get(b *testing.B) {
	counter := NewCounter()
	cnt := 100
	for i := 0; i < cnt; i++ {
		counter.Incr(strconv.Itoa(i), 1)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		counter.Get(strconv.Itoa(i % cnt))
	}
}

func BenchmarkSyncCounter_IncrPara(b *testing.B) {
	cnt := 2
	counter := NewCounter()
	w := sync.WaitGroup{}
	w.Add(cnt)
	for i := 0; i < cnt; i++ {
		go func() {
			for i := 0; i < b.N; i++ {
				counter.Incr(strconv.Itoa(i), 1)
			}
			w.Done()
		}()
	}
	w.Wait()
}

func BenchmarkSyncCounter_GetPara(b *testing.B) {
	cnt := 2
	counter := NewCounter()
	w := sync.WaitGroup{}
	w.Add(cnt)

	N := 100
	for i := 0; i < N; i++ {
		counter.Incr(strconv.Itoa(i), 1)
	}
	b.ResetTimer()
	for i := 0; i < cnt; i++ {
		go func() {
			for i := 0; i < b.N; i++ {
				counter.Get(strconv.Itoa(i % N))
			}
			w.Done()
		}()
	}
	w.Wait()
}
