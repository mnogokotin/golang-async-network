package app1

import (
	"fmt"
	"github.com/mnogokotin/golang-packages/utils/async"
	"net/http"
)

func Run() {
	defer async.Timer("app 1 main")()

	urls := []string{
		"https://www.easyjet.com/",
		"https://www.skyscanner.de/",
		"https://facebook.com/",
		"https://wizzair.com/",
		"https://www.swiss.com/",
	}

	c := make(chan urlStatus)
	for _, url := range urls {
		go checkUrl(url, c)
	}

	result := make([]urlStatus, len(urls))
	for i := range result {
		result[i] = <-c
		fmt.Println(result[i].url, result[i].status)
	}
}

func checkUrl(url string, c chan urlStatus) {
	_, err := http.Get(url)
	if err != nil {
		c <- urlStatus{url, false}
	} else {
		c <- urlStatus{url, true}
	}
}

type urlStatus struct {
	url    string
	status bool
}
