package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

// A WaitGroup is used to wait for the program to finish goroutines
var wg sync.WaitGroup

// Function that will be called as a goroutine
func querySite(url string) {
	//Tell the WaitGroup that the goroutine is complete
	defer wg.Done()

	fmt.Println("HTTP Request: ", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	fmt.Println("Read Body: ", url)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Body Length: ", len(body))
}

func main() {
	//add a count of two, one for each goroutine
	wg.Add(2)
	go querySite("https://google.com")
	go querySite("https://stackoverflow.com")

	wg.Wait()
	fmt.Println("Terminating Program")
}
