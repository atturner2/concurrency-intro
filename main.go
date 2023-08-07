package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	lock  sync.Mutex
	count int
)

func main() {

	iterations := 10000

	for i := 0; i < iterations; i++ {
		go increment()
	}
	time.Sleep(1 * time.Second)
	fmt.Println("Resulted count is: ", count)
}

func increment() {
	lock.Lock()
	count++
	lock.Unlock()
}

// deadlock because we pass in the beeper by value, it needs to be a pointer
func attack(evilNinja string, beeper *sync.WaitGroup) {
	fmt.Println("Attacked evil ninja: ", evilNinja)
	beeper.Done()
}

func captainElect(ninja chan string, message string) {
	ninja <- message
}
