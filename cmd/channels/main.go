package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {

	messages := make(chan string)

	go func() {
		messages <- "ping"
	}()

	msg := <-messages
	fmt.Println(msg)

	mb := make(chan string, 2)

	go func() {
		mb <- "buffered"
		mb <- "channel"
	}()

	fmt.Println(<-mb)
	fmt.Println(<-mb)

	// 实现goroutine同步
	done := make(chan bool, 1)
	go worker(done)
	<-done

	// ping -> pong
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
