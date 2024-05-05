package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan bool)
	go greet("Nice to meet you!", channel)
	go greet("How are you?", channel)
	go slowGreet("How ... are ... you ...?", channel)
	go greet("I hope you're liking the course!", channel)
	for range channel {

	}
}

func greet(phrase string, ch chan bool) {
	fmt.Println("Hello!", phrase)
	ch <- true
}

func slowGreet(phrase string, ch chan bool) {
	time.Sleep(3 * time.Second) // simulate a slow, long-taking task
	fmt.Println("Hello!", phrase)
	ch <- true
	close(ch)
}
