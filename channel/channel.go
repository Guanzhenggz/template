package main

import (
	"fmt"
	"time"
)

func Worker(id int,c chan int) {
	//for {
	//	n,ok := <-c
	//	if !ok {
	//		break
	//	}
	for n := range c {
		fmt.Printf("Worker %d received %c\n", id, n)
	}
}

type run struct{
	timeout chan<- time.Time
}

func createWorker(id int) chan<- int {

	c := make(chan int)
	go Worker(id,c)
	return c
}


func chanDemo() {
	var channels [10]chan<- int
	for i := 0; i < 10; i ++ {
		channels[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}
	time.Sleep(time.Millisecond)
}

func bufferedChannel() {
	c := make(chan int,3)
	go Worker(0,c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	time.Sleep(time.Millisecond)
}

func channelClose() {
	c := make(chan int,5)
	go Worker(0,c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	close(c)
	time.Sleep(time.Millisecond)

}

func main() {

	//fmt.Println("Channel as first-class citizen")
	//chanDemo()
	fmt.Println("Buffered channel")
	channelClose()
	//fmt.Println("Channel close and range")
	//channelClose()
}
