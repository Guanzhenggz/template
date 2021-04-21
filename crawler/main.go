package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {

	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error:status code",resp.StatusCode)
		return
	}

	//转化GBK为UTF-8
	//utf8Reader := transform.NewReader(resp.Body, simplifiedchinese.GBK.NewDecoder())

	all, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n",all)
}
