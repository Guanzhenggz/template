package main

import (
	"fmt"
	"io/ioutil"
)

func grade(score int) string{
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong score input %d",score))
	case score > 60 :
		g = "F"
	case score > 70 :
		g = "C"
	case score > 80 :
		g = "B"
	case score > 90 :
		 g ="A"
	}

	return g
}

func main(){

	const filename = "./abc.txt"
	if contents,err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	}else {
		fmt.Printf("%s\n",contents)
	}

    fmt.Println(
    	grade(13),
		grade(56),
		grade(63),
		grade(72),
		grade(86),
		grade(93),
		//grade(102),
    	)

}
