package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func eval(a,b int,op string) (int,error) {
	switch op {
	case "_+":
		return a + b,nil
	case "_-":
		return a - b,nil
	case "_*":
		return a * b,nil
	case "_%":
		return a % b,nil
	case "_/":
		 q,_  := div(a,b)
		return q,nil
	default:
		return 0,fmt.Errorf("unsupported operation: %s" + op)
	}
}

// 13 / 3 = 4 .. 1
func div(a,b int) (q,r int) {

	// q = a / b
	// r = a % b

	return a / b,a % b
}

const (
	a = iota
	b
	c
	d
	e
)

func apply(op func(int,int) int,a,b int) int  {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function: %s with args \n" + "(%d,%d) \n",opName,a,b)
	return op(a,b)
}

func Pow(a,b int) int {
	return int(math.Pow(float64(a),float64(b)))
}

func sum(numbers ...int) int{
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

func swap(a,b int) (int,int){
	return b,a
}

type user struct {
	username string
	email string
}

func (u *user) notify() string{
	return u.email + ": @" + u.username
}

type Duration int

func main() {
	if result,err := eval(3,2,"_+"); err != nil{
		fmt.Println("Error:" , err)
	}else {
		fmt.Println(result)
	}
	_,b := div(12, 3)
	fmt.Println(a,b)
	fmt.Println(apply(Pow,2,3))
	a,c := 3,4
	a,c = swap(a,c)
	fmt.Print(a,c)

	var uu user= user{"guanzheng","123456@qq.com"}

	fmt.Print(uu.notify())

	var duration Duration
	duration = Duration(2)

	fmt.Print(duration)
}
