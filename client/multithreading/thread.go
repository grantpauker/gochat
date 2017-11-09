package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(s)
	}

}

func main() {
	go say("second")
	for {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("1")
	}

}
