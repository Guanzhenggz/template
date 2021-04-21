package main

import "fmt"

func printArray(arr *[9]int){
	arr[0] = 100
	for i,v := range arr {
		fmt.Print(i,v)
	}
}

func main() {
	var arr1 [9]int
	arr2 := [3]int{1,3,5}
	arr3 := [...]int{1,2,3,4,5,6,7,8,9}
	var grid [4][5]int
	fmt.Println(arr1,arr2,arr3)
	fmt.Println(grid)

	//for i := 0 ; i < len(arr3) ; i ++ {
	//	fmt.Print(arr3[i])
	//}

	for j,v := range arr3 {
		fmt.Print(j,v ,",")
	}

	fmt.Println(&arr1);

	printArray(&arr1)
	printArray(&arr3)
	
}
