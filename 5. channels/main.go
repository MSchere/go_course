package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	urls := []string{
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://go.dev",
		"https://amazon.com",
	}
	ch := make(chan string)
	for _, url := range urls {
		go checkStatus(url, ch)
	}

	for u := range ch {
		go func(url string) { // Function literal (like a js anonymous function)
			time.Sleep(5 * time.Second)
			checkStatus(url, ch)
		}(u)
	}
}

func checkStatus(url string, ch chan string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Println(url, "might be down!")
		ch <- url
		return
	}
	fmt.Println(url, "is up")
	ch <- url
}
