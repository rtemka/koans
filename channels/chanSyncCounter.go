package channels

import (
	"fmt"
	"sync"
)

const workers = 10
const max = 10_000

func ChannelSyncCounter() {

	inc := make(chan int, workers)
	done := make(chan struct{})

	var wg sync.WaitGroup
	wg.Add(workers)

	for i := 0; i < workers; i++ {

		go func(id int) {
			var wc int

			for {
				select {
				case <-done:
					fmt.Printf("worker #%d finished, pass values %d\n", id, wc)
					wg.Done()
					return
				case inc <- 1:
					wc++
				}
			}
		}(i)

	}
	var c int
	for v := range inc {
		if c == max {
			close(done)
			fmt.Println("max is reached")
			break
		}
		c += v
		fmt.Println(c)
	}

	wg.Wait()
	fmt.Println(c)
}
