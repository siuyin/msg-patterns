package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("work-queue with unbuffered input channel")
	fmt.Println("Ctrl-C to terminate program")
	in := worker()
	boss(5, in)
	// sleep until cancelled by user
	for {
		time.Sleep(time.Second)
	}
}

func now() string {
	return time.Now().Format("04:05.0000")
}
func worker() chan<- int {
	ch := make(chan int)
	q := []int{} // our queue
	m := &sync.Mutex{}

	// queuer goroutine
	go func() {
		for {
			select {
			case i := <-ch:
				// queue incoming requests
				m.Lock()
				q = append(q, i)
				fmt.Printf("%v queued: %v\n", now(), i)
				m.Unlock()
			}
		}
	}()

	// worker goroutine
	go func() {
		for {
			m.Lock()
			if len(q) == 0 {
				m.Unlock()
				time.Sleep(time.Second)
				continue
			}
			var x int
			x, q = q[0], q[1:]
			m.Unlock()
			fmt.Printf("%s started processing: %v\n", now(), x)
			time.Sleep(time.Second)
			fmt.Printf("%s   finished processing: %v\n", now(), x)
		}
	}()
	return ch
}

func boss(numJobs int, wCh chan<- int) {
	go func() {
		fmt.Printf("%s started bossing\n", now())
		for i := 0; i < numJobs; i++ {
			wCh <- i
		}
		fmt.Printf("%s   finished bossing\n", now())
	}()
}
