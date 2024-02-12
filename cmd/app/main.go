package main

import (
	"github.com/mnogokotin/golang-async-network/internal/app1"
	"github.com/mnogokotin/golang-async-network/internal/app2"
	"github.com/mnogokotin/golang-async-network/internal/app3"
)

func main() {
	app1.Run()
	app2.Run()
	app3.Run()
}
