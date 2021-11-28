package counter

import "time"

type ICounter interface {
	Get(string) chan int
	Init()
	Flush2broker(int, func())
	Incr(string, int)
}

type readQ struct {
	bucket string
	res    chan int
}

type incrQ struct {
	bucket string
	count  int
}

type Counter struct {
	buckets map[string]int
	incrQ   chan incrQ
	readQ   chan readQ
	sumQ    chan chan int
}

func NewCounter() ICounter {
	c := Counter{}
	c.Init()
	go c.run()
	return &c
}

func (c *Counter) Init() {
	c.buckets = make(map[string]int, 10000)
	c.incrQ = make(chan incrQ, 10000000)
	c.readQ = make(chan readQ, 10000000)
}

func (c Counter) run() {
	for {
		select {
		case a := <-c.readQ:
			a.res <- c.buckets[a.bucket]
		case a := <-c.incrQ:
			c.buckets[a.bucket] += a.count
		default:

		}
	}
}

func (c Counter) Get(bucket string) chan int {
	res := make(chan int, 1)
	c.readQ <- readQ{bucket: bucket, res: res}
	return res
}

func (c Counter) Incr(bucket string, count int) {
	c.incrQ <- incrQ{bucket: bucket, count: count}
}

func (c *Counter) Flush2broker(mills int, f func()) {
	go func() {
		t := time.Tick(time.Duration(mills) * time.Millisecond)
		for {
			<-t
			f()
			c.Init()
		}
	}()
}
