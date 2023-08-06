package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	width, height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (r Rectangle) Perimeter() float64 {
	return (2 * r.width) + (2 * r.height)
}

func main() {
	channel := make(chan string)
	//numRounds := 3
	go throwingNinjaStar(channel)
	for {
		message, open := <-channel
		if !open {
			break
		}
		fmt.Println(message)
	}
}

func throwingNinjaStar(channel chan string) {
	rand.Seed(time.Now().UnixNano())
	numRounds := 3
	for i := 0; i < numRounds; i++ {
		score := rand.Intn(10)
		channel <- fmt.Sprint("You scored: ", score)
	}
	close(channel)
}
