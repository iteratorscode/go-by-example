package main

import (
	"fmt"
	"time"
)

func main() {
	const jobNums = 5
	jobs := make(chan int, jobNums)
	results := make(chan int, jobNums)

	for i := 0; i < 3; i++ {
		go func(id int, job chan int, result chan int) {
			for j := range job {
				fmt.Println("worker", id, "started job", j)
				time.Sleep(time.Second)
				fmt.Println("worker", id, "finished job", j)
				result <- j * 2
			}
		}(i, jobs, results)
	}

	for i := 1; i <= jobNums; i++ {
		jobs <- i
	}

	close(jobs)
	for i := 0; i < jobNums; i++ {
		<-results
	}
}
