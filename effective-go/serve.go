package main

import "fmt"

const MaxOutstanding = 10

type Request struct {
	args       []int
	f          func([]int) int
	resultChan chan int
}

func sum(a []int) (s int) {
	for _, v := range a {
		s += v
	}
	return
}

func handle(queue chan *Request) {

	fmt.Println("goroutine is running")
	for req := range queue {
		req.resultChan <- req.f(req.args)
	}
	fmt.Printf("queue is closed")
}

func Serve(clientRequests chan *Request, quit chan bool) {
	for i := 0; i < MaxOutstanding; i++ {
		go handle(clientRequests)
	}
	fmt.Println("Serve done")
	<-quit
}

func main() {
	clientRequests := make(chan *Request)
	quit := make(chan bool)

	go Serve(clientRequests, quit)

	request := &Request{[]int{1, 2, 3}, sum, make(chan int)}
	clientRequests <- request
	fmt.Printf("answer: %d\n", <-request.resultChan)
	quit <- true
}
