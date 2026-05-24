package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links { // working as an initial spark for the infinite loop
		go checkLink(link, c)
	}

	for l := range c {
		// fmt.Println(<-c) // for eg, if the c data is empty, then the main go routine pauses, go scheduler schedules another go routine, the cpu never stops here
		// go checkLink(l, c) // can spawn millions of go routines for concurrent requests

		go func(link string) {
			time.Sleep(5 * time.Second) // just to introduce some throttling so that the main loop can breathe a bit
			checkLink(l, c)             // refer case of M:N multiplexing (m go routines references n threads)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	// time.Sleep(5 * time.Second) // this is for not spamming the urls with continous requests
	// not needed though
	if err != nil {
		fmt.Println(link, "might be down")
		c <- link
		return
	}

	fmt.Println(link, "is up")
	c <- link // always pushing the actual link for infinite http get requests
}
