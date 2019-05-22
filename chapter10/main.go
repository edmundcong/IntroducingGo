package main

import (
	"fmt"
	"time"
	)

// pinger can only send values to c channel
func pinger(c chan<- string) {
	for i := 0; ; i++ {
		c <- "ping" // send value ping to c
	}
}


// printer can only receive values from c channel
func printer(c <-chan string) {
	for {
		msg := <- c // msg takes value c ("ping")
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	var c chan string = make(chan string)

	go pinger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)
}
