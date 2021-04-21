package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"time"
)

var (
	aa = 3
	ss = "KKK"
)

func variable(){
	var a int
	var s string
	fmt.Printf("%d %q\n",a,s)
}

func variableInitialValue(){
	var a,b int = 3,4
	var s string = "abc"
	fmt.Println(a,b,s)
}

func variableTypeDeduction(){
	var a,b,c = 3,4,true
	var s = "abc"
	fmt.Println(a,b,s,c)
}

func variableShorter(){
	a,b,c,w := "Hello","This is a golang",true,"yes"
	fmt.Println(a,b,c,w)
}

func eular(){
	fmt.Println(cmplx.Pow(math.E,1i * math.Pi) + 1)
}

func triangle(){
	var a,b int = 3,4
	fmt.Println(calcTriangle(a,b))
}

func calcTriangle(a,b int) int {

	c := int(math.Sqrt(float64(a*a + b*b)))
	return c
}

func consts(){
	const filename = "Pride and Basied.txt"
	const a,b = 3,4
	var c int
	c = int(math.Sqrt((a*a + b*b)))
	fmt.Println(filename,c)
}

func enums(){
	const(
		cpp = iota
		_
		java
		python
		golang
	)
	fmt.Println(cpp,java,python,golang)

	/*
		b,kb,mb,gb,tb
	 */
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
	)

	fmt.Println(b,kb,mb,gb,tb)
}


func main(){
	fmt.Println("Hello World")
	//variable()
	//variableInitialValue()
	//variableTypeDeduction()
	//variableShorter()
	//fmt.Println(aa,ss)
	//eular()
	//consts()
	enums()

	fmt.Println(time.Now().Sub(time.Now().Add(time.Minute * -12)).Round(time.Duration(30)))
}