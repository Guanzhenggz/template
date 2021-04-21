package main

import "fmt"

func main1() {
	m := map[string]string {
		"name" : "ccmouse",
		"course" : "golang",
		"site" : "imooc",
		"quality" : " notbad",
	}

	m2 := make(map[string]int) // m2 = empty map

	var m3 map[string]int // m3 = nil

	fmt.Println(m,m2,m3)

	fmt.Println("Traversing map ")
	for k,v := range m { // _,v
		fmt.Println(k,v)
	}

	fmt.Println("Getting values")
	if courseName,ok := m["course"]; ok {
		fmt.Println(courseName,ok)
	}else {
		fmt.Println("key does not exists")
	}

	fmt.Println("Delete value")
	if name,ok := m["name"]; ok{
		fmt.Println(name,ok)
		delete(m,"name")
	}

}
