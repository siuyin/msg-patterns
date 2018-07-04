package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("work-queue with buffered input channel")
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
	const bufSize = 100
	ch := make(chan int, bufSize)

	go func() {
		for {
			select {
			case x := <-ch:
				fmt.Printf("%s started processing: %v\n", now(), x)
				time.Sleep(time.Second)
				fmt.Printf("%s   finished processing: %v\n", now(), x)
			}
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
