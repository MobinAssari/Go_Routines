package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan bool)
	greet("Nice to meet you!")
	greet("How are you?")
	go slowGreet("How ... are ... you ...?", channel)
	greet("I hope you're liking the course!")
	<-channel
}

func greet(phrase string) {
	fmt.Println("Hello!", phrase)
}

func slowGreet(phrase string, ch chan bool) {
	time.Sleep(3 * time.Second) // simulate a slow, long-taking task
	fmt.Println("Hello!", phrase)
	ch <- true
}
