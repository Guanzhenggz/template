package main

import (
	"fmt"
	"runtime"
	"time"
)

func cal(a int,b int){
	c := a + b
	fmt.Printf("%d + %d = %d\n",a,b,c)
}

func addele(a []int,i int) {
	defer func() { // 匿名函数捕获错误
		error := recover()
		if error != nil{
			fmt.Println("add ele fail")
		}
	}()
	a[i] = i
	fmt.Println(a)
}

func main() {
	//for i := 0; i < 1000; i ++ {
	//	go func(i int) {
	//		for {
	//			fmt.Printf("Hello from goroutine %d\n",i)
	//		}
	//	}(i)
	//}
	//time.Sleep(time.Millisecond * 3)

	fmt.Println("Logic CPU Num :",runtime.NumCPU())

	for i := 0;i < 10;i ++ {
		go cal(i,i + 1)
	}

	time.Sleep(time.Millisecond * 2)

	Array := make([]int,4)
	for i := 0;i < 10;i ++ {
		go addele(Array,i)
	}

	time.Sleep(time.Millisecond * 3)

}
