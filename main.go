package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	defer func() {
		fmt.Println(time.Since(now))
	}()

	smokeSignal := make(chan bool)
	evilNinja := "Tommy"

	go attack(evilNinja, smokeSignal)
	fmt.Println("here is the smoke signal: ", <-smokeSignal)
	//smokeSignal <- false
	//arrow to the left on the left hand side of the channel variable is to recieve the channel variable
	//this smoke signal can only be used by one ninja at a time, so if line 19 ins't commented out you have deadlock.
	//bufferred channels are the solution to this, on the next branch
}

func attack(target string, attacked chan bool) {
	time.Sleep(time.Second)

	fmt.Println("Throwing ninja stars at ", target)
	attacked <- true

}
