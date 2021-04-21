package main

import (
	fib "awesomeProject/functional"
	"bufio"
	"errors"
	"fmt"
	"os"
)

func tryDefer() {
	for i := 0;i < 100;i ++ {
		defer fmt.Println(i)
	}
}

func writeFile(filename string) {
	file,err := os.OpenFile(
		filename,os.O_EXCL|os.O_CREATE,0666)

	err = errors.New("this is a custom error")
	if err != nil {
		if pathError,ok := err.(*os.PathError); !ok{
			panic(err)
		}else {
			fmt.Printf("%s,%s,%s \n",
				pathError.Op,
				pathError.Path,
				pathError.Err)
		}
		fmt.Println("Error:",err.Error())
		return
	}
	defer file.Close()

	write := bufio.NewWriter(file)
	defer write.Flush()

	f := fib.Fibonacci()

	for i := 0;i < 20; i ++ {
		fmt.Fprintln(write,f())
	}
}

func main(){
	writeFile("fib.txt")
}
