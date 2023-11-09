package async

import (
	"fmt"
	"time"
)

func Timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}

func Send(channel chan interface{}, data string, delay time.Duration) {
	time.Sleep(delay)
	channel <- data
}

func Receive(channel <-chan interface{}) {
	for data := range channel {
		fmt.Printf("<- %s\n", data)
	}
}
