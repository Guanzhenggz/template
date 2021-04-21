package fib

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// 1，1，2，3，5，8，13，...
//    a,b
//      b,a+b
func Fibonacci() IntGen {
	a,b := 0,1
	return func() int {
		a,b = b,a + b
		return a
	}
}

type IntGen func() int

func (g IntGen) Read(p []byte) (n int,err error){
	next := g()
	if next > 500 {
		return 0,io.EOF
	}
	s := fmt.Sprintf("%d\n",next)
	return strings.NewReader(s).Read(p)
}

func printFileContents(read io.Reader) {
	if read == nil {
		panic("read must define not nil!")
	}
	scanner := bufio.NewScanner(read)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

//func main() {
//	f := Fibonacci()
//	printFileContents(f)
//
//}
