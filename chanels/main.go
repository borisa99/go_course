package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://golang.org",
		"https://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c {
		go func(ln string) {
			time.Sleep(5 * time.Second)
			checkLink(ln, c)
		}(l)
	}

}
func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "up and running!")
	c <- link
}
