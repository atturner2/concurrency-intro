package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string, 2)
	channel <- "First Message"
	channel <- "second Message"
	fmt.Println(<-channel)
	fmt.Println(<-channel)

}

func attack(target string, attacked chan bool) {
	time.Sleep(time.Second)

	fmt.Println("Throwing ninja stars at ", target)
	attacked <- true

}
