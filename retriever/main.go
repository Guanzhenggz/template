package main

import (
	"awesomeProject/retriever/mock"
	"time"

	//"awesomeProject/retriever/mock"
	"awesomeProject/retriever/real"
	"fmt"
)

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string,form map[string]string) string
}

func download(r Retriever) string {
	return r.Get(URL)
}

func post(poster Poster) {
	poster.Post(URL, map[string]string{
		"name" : "ccmouse" ,
		"course" : "golang",
	})
}

type RetrieverPoster interface {
	Retriever
	Poster
}

func session(s RetrieverPoster) string {
	//s.Get()
	s.Post(URL, map[string]string{
		"contents" : "another faked imooc.com",
	})
	return s.Get(URL)
}

func inspect(r Retriever) {

	fmt.Println("Inspecting",r)
	fmt.Printf("> %T %v \n",r,r)
	fmt.Print("> Type switch:")
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents:",v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgant:",v.UserAgent)

	}
}

func main() {
	var r Retriever
	retriever := mock.Retriever{"this is a fake imooc.com"}
	r = &retriever
	inspect(r)
	r = &real.Retriever{
		UserAgent:"Mozilla/5.0",
		TimeOut: time.Minute,
	}

	inspect(r)

	// Type assertion
	if mockRetriever, ok := r.(*mock.Retriever); ok {
		fmt.Println(mockRetriever.Contents)
	}else {
		fmt.Println("not a mock retriever")
	}

	fmt.Println("Try a session")
	fmt.Println(session(&retriever))

	// fmt.Println(download(r))

}

const(
	URL = "http://www.baidu.com"
)
