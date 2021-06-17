package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go process(ch1)
	go replicate(ch2)
	for i := 0; i < 2; i++ {
		select {
		case process := <-ch1:
			fmt.Println(process)
		case replicate := <-ch2:
			fmt.Println(replicate)
		}
	}
}

func process(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "Done processing"
}
func replicate(ch chan string) {
	time.Sleep(1 * time.Second)
	ch <- "Done replicating"
}

func send_and_read() {
	ch := make(chan string, 1)
	send(ch, "Hello World")
	read(ch)
}

func request() {
	start := time.Now()

	apis := []string{
		"https://management.azure.com",
		"https://dev.azure.com",
		"https://api.github.com",
		"https://outlook.office.com/",
		"https://api.somewhereintheinternet.com/",
		"https://graph.microsoft.com",
	}

	ch := make(chan string)

	for _, api := range apis {
		go checkAPI(api, ch)
	}

	for i := 0; i < len(apis); i++ {
		fmt.Printf(<-ch)
	}

	elapsed := time.Since(start)
	fmt.Printf("Done! It took %v seconds!\n", elapsed.Seconds())

}

func checkAPI(api string, ch chan string) {
	_, err := http.Get(api)
	if err != nil {
		ch <- fmt.Sprintf("ERROR: %s is down\n", api)
	} else {
		ch <- fmt.Sprintf("SUCCESS: %s is up and runngin\n", api)
	}
	return
}

func send(ch chan<- string, message string) {
	fmt.Printf("Sending: %#v\n", message)
	ch <- message

}

func read(ch <-chan string) {
	fmt.Printf("Read: %#v\n", <-ch)

}
