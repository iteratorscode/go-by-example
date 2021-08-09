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

	// non blocking channel operation: select -- default
	messagess := make(chan string)
	signals := make(chan bool)
	select {
	case msg := <-messagess:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}
	msgs := "hi"
	select {
	// Here msg cannot be sent to the messages channel, because the channel has no buffer and there is no receiver.
	case messagess <- msgs:
		fmt.Println("sent message", msgs)
	default:
		fmt.Println("no message sent")
	}
	select {
	case msg := <-messagess:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}

	jobs := make(chan int, 5)
	doneAll := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				doneAll <- true
				return
			}
		}
	}()

	for i := 0; i <= 3; i++ {
		jobs <- i
		fmt.Println("sent job", i)
	}

	close(jobs)
	fmt.Println("sent all jobs")
	<-doneAll
}
