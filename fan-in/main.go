package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("synchronous fan-in example")
	reportCh := make(chan string)
	worker("A", reportCh)
	worker("B", reportCh)
	worker("C", reportCh)
	boss(reportCh)
}

func worker(id string, ch chan<- string) {
	go func() {
		fmt.Printf("%s worker %s started job\n", time.Now().Format("04:05.000"), id)
		time.Sleep(time.Second)
		ch <- id
		fmt.Printf("%s   worker %s reported completed job\n", time.Now().Format("04:05.000"), id)
	}()
}

func boss(reportCh <-chan string) {
	const timeOut = 2 * time.Second
	goHomeTimer := time.NewTimer(timeOut)
mainLoop:
	for {
		if !goHomeTimer.Stop() {
			<-goHomeTimer.C
		}
		goHomeTimer.Reset(timeOut)
		select {
		case s := <-reportCh:
			fmt.Printf("%s boss got report from: %s\n", time.Now().Format("04:05.000"), s)
			time.Sleep(time.Second)
			fmt.Printf("%s   boss read report from: %s\n", time.Now().Format("04:05.000"), s)
		case <-goHomeTimer.C:
			fmt.Printf("%s boss is going home\n", time.Now().Format("04:05.000"))
			break mainLoop
		}
	}
}
