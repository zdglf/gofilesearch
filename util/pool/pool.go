package pool

import (
	"sync"
)

type Pool struct {
	queue chan int
	wg    *sync.WaitGroup
}

// New 新建一个协程池
func NewPool(size int) *Pool {
	if size <= 0 {
		size = 1
	}
	return &Pool{
		queue: make(chan int, size),
		wg:    &sync.WaitGroup{},
	}
}

// Add 新增一个执行
func (p *Pool) Add(delta int) {
	if delta <= 0 {
		return
	}
	for i := 0; i < delta; i++ {
		p.queue <- 1
	}
	p.wg.Add(delta)
}

// Done 执行完成减一
func (p *Pool) Done() {
	<-p.queue
	p.wg.Done()
}

// Wait 等待Goroutine执行完毕
func (p *Pool) Wait() {
	p.wg.Wait()
}
