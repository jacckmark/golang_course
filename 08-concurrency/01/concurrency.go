package main

import (
	"fmt"
	"time"
)

func greet(phrase string) {
	fmt.Println("Hello!", phrase)
}

// we needed to add the channel to this function to make it send the info about
// it's execution
func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second) // simulate a slow, long-taking task
	fmt.Println("Hello!", phrase)
	// here we need to use special character in golang to send data to our channel
	doneChan <- true
}

func main() {
	// to run code in parallel you have to add "go" before the code that you want to
	// run in the routine

	// we need also to use channel in go to make it (go routines) work, because
	// of the fact that go routines will execute so fast we will simply step out of
	// the function without seeing result of executing our functions
	done := make(chan bool)
	// here we pass the channel into our function
	go slowGreet("How ... are ... you ...?", done)
	// and we read from the channel (we can also get data from it by putting print
	// in front of our arrow)
	<-done
	// thanks to this block we have the code that forces go to wait for it's completion

	// this function calls will not run because they are non-blocking now and are
	// not using a channel (thus we skip over them and reach end of main function
	// which in turn end our program)
	go greet("How are you?")
	go greet("I hope you're liking the course!")
}
