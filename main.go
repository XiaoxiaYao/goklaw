package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
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
		printCityList(body)
	}
}

func printCityList(contents []byte) {
	result := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)
	matches := result.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		fmt.Printf("City: %s, URL: %s\n", m[2], m[1])
	}
	fmt.Printf("matches %d", len(matches))
}
