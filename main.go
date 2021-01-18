package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	response, err := http.Get("http://www.zhenai.com/zhenghun/")
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		fmt.Print("Error: status code", response.StatusCode)
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", body)
	}
}
