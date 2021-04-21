package main


import(
	"fmt"
)

func  mergeSort(arr []int,res *[]int,start int,end int){
	if start >= end {
		return
	}

	len := end - start
	mid := (len >> 1) + start
	var start1 int = start
	var end1 int = mid
	var start2 int = mid + 1
	var end2 = end

	mergeSort(arr,res,start1,end1)
	mergeSort(arr,res,start2,end2)

	k := start

	for{
		if start1 > end1 || start2 > end2 {
			break
		}

		if arr[start1] < arr[start2] {
			(*res)[k] = arr[start1]
			start1 ++
		}else{
			(*res)[k] = arr[start2]
			start2 ++
		}
		k ++
	}

	for{
		if start1 > end1{
			break
		}
		(*res)[k] = arr[start1]
		k ++
		start1 ++
	}

	for{
		if start2 > end2 {
			break
		}
		(*res)[k] = arr[start2]
		k ++
		start2 ++
	}

	for k := start;k <= end;k ++ {
		arr[k] = (*res)[k]
	}
}

func main(){
	var arr []int = []int{5,4,7,2,8,9,1,4}
	var res []int = make([]int,8)
	mergeSort(arr,&res,0,len(arr) - 1)
	fmt.Print(res)
}
