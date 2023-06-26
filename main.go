package main

import (
	"fmt"
	"time"
)

func main() {

	start := time.Now()

	defer func() {
		fmt.Println(time.Since(start))
	}()

	evilNinjas := []string{"Tommy", "Johnny", "Bobby", "Andy"}

	for _, evilNinja := range evilNinjas {
		go attack(evilNinja) // adding go creates concurrency by creating seperate GoRoutines. We need channels to communicate between them
	}

	time.Sleep(time.Second * 2)

}

func attack(target string) {
	time.Sleep(time.Second)

	fmt.Println("Throwing ninja stars at ", target)

}
