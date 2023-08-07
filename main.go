package main

import (
	"fmt"
	"sync"
)

func main() {

	var beeper sync.WaitGroup
	evilNinjas := []string{"Tommy", "Johnny", "Bobby"}
	beeper.Add(len(evilNinjas))
	for _, evilNinja := range evilNinjas {
		go attack(evilNinja, &beeper)
	}

	beeper.Wait()
	fmt.Println("Mission completed")
	//maine executes before the other one can
}

// deadlock because we pass in the beeper by value, it needs to be a pointer
func attack(evilNinja string, beeper *sync.WaitGroup) {
	fmt.Println("Attacked evil ninja: ", evilNinja)
	beeper.Done()
}

func captainElect(ninja chan string, message string) {
	ninja <- message
}
