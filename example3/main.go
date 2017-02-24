package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func request(url string) {
	fmt.Println("->", url)
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	// wg.Done()
}

func main() {

	// wg.Add(20)

	request("http://localhost:8080/sleep/2")
	request("http://localhost:8080/sleep/3")

	// wg.Wait()
	fmt.Println("Fim")
}
