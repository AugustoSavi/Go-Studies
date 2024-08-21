package main

import (
	"fmt"
	"time"
)

func worker(workerId int, data chan int){
	for x := range data{
		fmt.Printf("worker %d got %d\n", workerId, x)
		time.Sleep(time.Second)
	}
}

func main()  { // goroutine 1
	ch := make(chan int)
	qtdWorkers := 5

	for v := range qtdWorkers {
		go worker(v, ch)
	}

	for i := range 10 {
		ch <- i
	}
}
