package main

import (
	"awesomeProject/errorhandling/filelistingserver/filelisting"
	"fmt"
	"log"
	"net/http"
	"os"
)

type appHandler func(writer http.ResponseWriter,
	request *http.Request) error

func errWrapper(handler appHandler) func(http.ResponseWriter,*http.Request){
		return func(writer http.ResponseWriter, request *http.Request) {

			// panic
			defer func() {
				r := recover()
				if r != nil {
					log.Printf("Panic: %v",r)
					http.Error(writer,http.StatusText(http.StatusInternalServerError),http.StatusInternalServerError)
				}
			}()
			err := handler(writer,request)
			if err != nil {
				fmt.Println(err.Error())

				// user error
				if  userErr,ok := err.(userError); ok {
					http.Error(writer, userErr.Message(),http.StatusBadRequest)
					return
				}

				// system error
				code := http.StatusOK
				switch {
				case os.IsNotExist(err):
					code = http.StatusNotFound
				case os.IsPermission(err):
					code = http.StatusForbidden
				default:
					code = http.StatusInternalServerError
				}
					http.Error(writer,
						http.StatusText(code), code)

			}
		}
}

type userError interface {
	error
	Message() string
}

func main() {
	http.HandleFunc("/list/", errWrapper(filelisting.HandlerFileList))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}

	//file, e := os.Open("fib.txt")
	//if e != nil {
	//	panic(e)
	//}
	//
	//scanner := bufio.NewScanner(file)
	//
	//if scanner.Scan() {
	//	fmt.Println(scanner.Text())
	//}
}
