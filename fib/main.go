package main

import "fmt"

var quit = make(chan bool)

func main() {
	var command string
	ch := make(chan int)

	go fib(ch)

	for {
		num := <-ch
		fmt.Print(num)
		fmt.Scanf("%s", &command)
		if command == "quit" {
			quit <- true
			break
		}
	}
}

func fib(ch chan<- int) {
	x, y := 1, 1
	for {
		select {
		case ch <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("Done")
			break
		}
	}
}
