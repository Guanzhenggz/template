package main

import (
	"fmt"
	"time"
	"unicode/utf8"
)

var l = make([]time.Time,1)

func recursionCount (arr []time.Time,t int) []time.Time {

	if t == len(arr){
		return make([]time.Time,0)
	}

	for _,date := range l {
		if date.Equal(arr[t]) {
			break
		}else {
			l = append(l,arr[t])
			break
		}
	}

	return recursionCount(arr,t + 1)
}

func execptRepeat (arr []time.Time) ([]time.Time,error) {
	m := make(map[time.Time] int)
	slice := make([]time.Time,0)
	for _,value := range arr {
		m[value] ++
	}

	for key,_ := range m {
		slice = append(slice,key)
	}
	return slice,nil
}

func main() {
	s := "GO语言 Yes !"
	fmt.Println(len(s))
	for _,b := range []byte(s) {
		fmt.Printf("%x ", b)
	}
	fmt.Println()

	for i,ch := range s { // ch is a rune
		fmt.Printf("(%d %x) " , i , ch)
	}
	fmt.Println()

	fmt.Println("Rune count:",utf8.RuneCountInString(s))

	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}
	fmt.Println()

	for i,ch := range []rune(s) {
		fmt.Printf("(%d  %c)",i,ch)
	}
	fmt.Println()


	a1 := []string{"a","b","c"}
	a2 := [] string{"c","d","e"}
	fmt.Println(append(a1,a2...))

	d := make([]time.Time,0)
	parseTime := time.Date(2019,time.Month(12),1,0,0,0,0,time.UTC)
	parseTime2, _ := time.Parse("2006-01-02", "2019-06-09")
	parseTime3, _ := time.Parse("2006-01-02", "2019-06-09")
	parseTime4, _ := time.Parse("2006-01-02", "2019-06-09")
	parseTime5, _ := time.Parse("2006-01-02", "2019-06-09")
	parseTime6, _ := time.Parse("2006-01-02", "2015-06-09")
	fmt.Println(parseTime)
	d = append(d,parseTime2)
	d = append(d,parseTime3)
	d = append(d,parseTime4)
	d = append(d,parseTime5)
	d = append(d,parseTime6)
	fmt.Println(parseTime)
	fmt.Println(d)
	fmt.Println(execptRepeat(d))
	fmt.Println(l)


}
