package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("synchronous fan-out example")
	workCh := make(chan int)
	worker("A", workCh)
	// worker("B", workCh)
	// worker("C", workCh)
	startTime := time.Now()
	boss(5, workCh)
	bossFinishedTime := time.Now()
	fmt.Printf("boss duration: %v\n", bossFinishedTime.Sub(startTime).Seconds())
	time.Sleep(2 * time.Second)
}

func worker(id string, ch <-chan int) {
	go func() {
		for {
			jobID := <-ch
			fmt.Printf("%s worker %s accepted job: %d\n", time.Now().Format("04:05.000"), id, jobID)
			time.Sleep(time.Second)
			fmt.Printf("%s   worker %s completed job: %d\n", time.Now().Format("04:05.000"), id, jobID)
		}
	}()
}

func boss(numJobs int, workerCh chan<- int) {
	for i := 0; i < numJobs; i++ {
		workerCh <- i
	}
}
