package channels

import (
	"fmt"
	"sync"
	"time"
)

type counter struct {
	m sync.RWMutex
	v int
}

func (c *counter) inc() bool {
	c.m.Lock()
	defer c.m.Unlock()
	if c.v == max {
		return false
	}
	c.v++
	return true
}

func MutexSyncCounter() {

	done := make(chan struct{})
	var counter counter

	var wg sync.WaitGroup
	wg.Add(workers)

	for i := 0; i < workers; i++ {

		go func(id int) {
			var wc int
			defer wg.Done()
			for {
				select {
				case <-done:
					fmt.Printf("worker #%d finished, inc times %d\n", id, wc)
					return
				default:
					if counter.inc() {
						wc++
					}
				}
			}
		}(i)

	}
	for {
		time.Sleep(time.Millisecond * 100)
		counter.m.RLock()
		if counter.v == max {
			close(done)
			fmt.Println("max is reached")
			counter.m.RUnlock()
			break
		}
		fmt.Println(counter.v)
		counter.m.RUnlock()
	}

	wg.Wait()
	fmt.Println(counter.v)
}
