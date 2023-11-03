package app1

import (
	"fmt"
	"net/http"
	"time"
)

func Run() {
	defer timer("main")()

	urls := []string{
		"https://www.easyjet.com/",
		"https://www.skyscanner.de/",
		"https://www.ryanair.com",
		"https://wizzair.com/",
		"https://www.swiss.com/",
		"https://google.com/",
		"https://youtube.com/",
		"https://facebook.com/",
		"https://instagram.com/",
		"https://reddit.com/",
		"https://baidu.com/",
		"https://wikipedia.org/",
		"https://yahoo.com/",
		"https://yandex.ru/",
		"https://whatsapp.com/",
		"https://xvideos.com/",
		"https://amazon.com/",
		"https://pornhub.com/",
		"https://tiktok.com/",
		"https://xnxx.com/",
	}

	c := make(chan urlStatus)
	for _, url := range urls {
		go checkUrl(url, c)
	}

	result := make([]urlStatus, len(urls))
	for i := range result {
		result[i] = <-c
		if result[i].status {
			fmt.Println(result[i].url, "ok")
		} else {
			fmt.Println(result[i].url, "down -----")
		}
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

func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}
