package main

import (
	"fmt"
	"time"
)

type req struct {
	inp int
	res chan int
}

func newReq(i int) *req {
	r := req{}
	r.inp = i
	r.res = make(chan int)
	return &r
}

func main() {
	fmt.Println("stateful server")
	fmt.Println("ctrl-c to terminate")
	ch := svr()
	boss(5, ch)
	for {
		time.Sleep(time.Second)
	}
}

func now() string {
	return time.Now().Format("04.05.0000")
}

// server code
func svr() chan *req {
	state := 100
	ch := make(chan *req)
	go func() {
		for {
			r := <-ch
			serve(r, &state)
		}
	}()
	return ch
}
func serve(r *req, state *int) {
	fmt.Printf("%s svr accepted request: %d\n", now(), r.inp)
	time.Sleep(time.Second)
	*state = *state + r.inp
	r.res <- *state
	fmt.Printf("%s   svr responded to request: %d\n", now(), r.inp)
}

// client code
func boss(numReq int, svrCh chan *req) {
	for i := 0; i < numReq; i++ {
		request(i, svrCh)
	}
}
func request(inp int, ch chan *req) {
	rq := newReq(inp)
	ch <- rq
	res := <-rq.res
	fmt.Printf("%s client request: %d, response %d\n", now(), inp, res)
}
