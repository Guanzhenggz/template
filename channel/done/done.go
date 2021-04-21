package main

import (
	"fmt"
	"sync"
)

func doWorker(id int,w worker) {
	//for {
	//	n,ok := <-c
	//	if !ok {
	//		break
	//	}
	for n := range w.in {
		fmt.Printf("doWorker %d received %c\n", id, n)
		//go func() {
		//	done <- true
		//}()
		w.done()
	}
}

type worker struct {
	in chan int
	done func()
}

func createWorker(id int,wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		done: func() {
			wg.Done()
		},
	}
	go doWorker(id,w)
	return w
}


//func chanDemo() {
//	var workers [10]worker
//	for i := 0; i < 10; i ++ {
//		workers[i] = createWorker(i)
//	}
//
//	for i,worker := range workers {
//		worker.in <- 'a' + i
//	}
//
//	for i,worker := range workers {
//		worker.in <- 'A' + i
//	}
//
//	// wait for all of them
//	for _,worker := range workers {
//		<- worker.done
//		<- worker.done
//	}
//
//}

func chanDemo2() {
	var workers [10]worker
	var wg sync.WaitGroup

	for i := 0; i < 10; i ++ {
		workers[i] = createWorker(i,&wg)
	}

	wg.Add(20)

	for i,worker := range workers {
		worker.in <- 'a' + i
	}

	for i,worker := range workers {
		worker.in <- 'A' + i
	}

	wg.Wait()

}

//func bufferedChannel() {
//	c := make(chan int,3)
//	go doWorker(0,c)
//	c <- 'a'
//	c <- 'b'
//	c <- 'c'
//	c <- 'd'
//	time.Sleep(time.Millisecond)
//}
//
//func channelClose() {
//	c := make(chan int)
//	go doWorker(0,c)
//	c <- 'a'
//	c <- 'b'
//	c <- 'c'
//	c <- 'd'
//	close(c)
//	time.Sleep(time.Millisecond)
//}

func main() {

	fmt.Println("Channel as first-class citizen")
	chanDemo2()
	//fmt.Println("Buffered channel")
	//bufferedChannel()
	//fmt.Println("Channel close and range")
	//channelClose()
}
