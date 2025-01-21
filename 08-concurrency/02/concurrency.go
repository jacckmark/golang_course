package main

import (
	"fmt"
	"time"
)

func greet(phrase string, doneChan chan bool) {
	fmt.Println("Hello!", phrase)
	doneChan <- true
}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second) // simulate a slow, long-taking task
	fmt.Println("Hello!", phrase)
	doneChan <- true
}

func main() {
	// here using one channel for all our methods (which forces go to wait for
	// the functions to end)
	done := make(chan bool)
	go greet("Nice to meet you!", done)
	go slowGreet("How ... are ... you ...?", done)
	go greet("How are you?", done)
	go greet("I hope you're liking the course!", done)
	// here we need to specify for how many routines we wait, because otherwise we
	// have a situation where the routines race (first one to end the execution is
	// shown, but the rest gets ignored). We can either do this or add more channels
	// (one channel per routine)
	<-done
	<-done
	<-done
	<-done
	// this can be written using slice and we can store there channels (look in notes
	// 03)
}
