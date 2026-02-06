package main

import (
	"fmt"
	"time"
)

func greet(phrase string, doneChannel chan bool) {
	fmt.Println("Hello!", phrase)
	doneChannel <- true
}

func slowGreet(phrase string, doneChannel chan bool) {
	time.Sleep(3 * time.Second) // simulate a slow, long-taking task
	fmt.Println("Hello!", phrase)
	doneChannel <- true
}

func main() {
	doneChannels := make([]chan bool, 4)

	doneChannels[0] = make(chan bool)
	go greet("Nice to meet you!", doneChannels[0])

	doneChannels[1] = make(chan bool)
	go greet("How are you?", doneChannels[1])

	doneChannels[2] = make(chan bool)
	go slowGreet("How ... are ... you ...?", doneChannels[2])

	doneChannels[3] = make(chan bool)
	go greet("I hope you're liking the course!", doneChannels[3])

	for _, done := range doneChannels {
		<-done
	}
}
