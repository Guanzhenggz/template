package main

import (
	"fmt"
)

func tryRecover() {
	defer func() {
		r := recover()
		if err,ok := r.(error); ok {
			fmt.Println("Error occurred:",err)
		}else {
			panic(fmt.Sprintf("I don't konw what to do:%v \n",r))
		}
	}()

// 	panic(errors.New("this is an error"))
	a := 0
	b := 5 / a
	fmt.Println(b)
}

func main()  {

}

