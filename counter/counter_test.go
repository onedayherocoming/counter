package counter

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestCounter_IncrAndGet(t *testing.T) {
	counter := NewCounter()
	times := 100
	for i := 0; i < times; i++ {
		counter.Incr(strconv.Itoa(i), 1)
		fmt.Println(i)
	}
	// 由于是异步的，所以需要稍微等一会
	time.Sleep(1 * time.Millisecond)
	// check
	for i := 0; i < times; i++ {
		res := counter.Get(strconv.Itoa(i))
		if res != 1 {
			t.Fatalf("%d not equal %d\n", res, i)
		}
	}
}

func TestCount_Flush2broker(t *testing.T) {
	f := func() {
		fmt.Println("I'm calling")
	}
	counter := NewCounter()
	counter.Flush2broker(1000, f)
	time.Sleep(10 * time.Second)
}
