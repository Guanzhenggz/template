package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i ++
		}
	}()
	return out
}

func Worker(id int,c chan int) {
	//for {
	//	n,ok := <-c
	//	if !ok {
	//		break
	//	}
	for n := range c {
		fmt.Printf("Worker %d received %d\n", id, n)
	}
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go Worker(id,c)
	return c
}

func main() {
	var c1, c2 = generator(), generator()
	var worker = createWorker(0)

	var values []int
	tm := time.After(10 * time.Second)
	tk := time.Tick(time.Second)
	for {
		var activeWorker chan<- int
		var activeValue int
		if len(values) > 0 {
			activeWorker = worker
			activeValue = values[0]
		}

		select {
		case n := <-c1:
			values = append(values,n)
		case n := <-c2:
			values = append(values,n)
		case activeWorker <- activeValue:
			values = values[1:]
		case <- tk:
			fmt.Println("queue len = ",len(values))
		case <- time.After(800 * time.Millisecond):
			fmt.Println("timeout")
		case <- tm:
			fmt.Println("bye")
			return
			

			//default:
			//	fmt.Println("no Value print out")
			//}

		}

	}
}