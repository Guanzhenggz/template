package main

import (
	"fmt"
)

func bubbleSort(arr *[7]int) {
	for i:= 0;i < len(arr) - 1;i ++ {
		for  j := 0;j < len(arr) - i - 1;j ++ {
			if arr[j] > arr[j + 1] {
				var tmp int = arr[j]
				arr[j] = arr[j + 1]
				arr[j + 1] = tmp
			}
		}
	}
}

func main(){
	arr := [...]int{2,4,1,6,8,4,6}
	bubbleSort(&arr)
	fmt.Print(arr)
}


