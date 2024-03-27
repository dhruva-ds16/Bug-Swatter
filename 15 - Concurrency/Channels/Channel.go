package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func querySite(url string, requestBody chan int) {
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
	//fmt.Println(body)
	requestBody <- len(body)
}

func main() {
	//Declaring an unbuffered channel
	requestBody := make(chan int)

	wg.Add(1)
	go querySite("https://linkedin.com/in/dhruvasantosh", requestBody)

	//Read from channel
	fmt.Println(<-requestBody)

	wg.Wait()

	//Close channel
	close(requestBody)
}
