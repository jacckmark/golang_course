package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/interface/note"
	"example.com/interface/todo"
)

// here creating an interface in golang (like in java whichever struct implements
// it is required to implement the methods that it specifies)
type saver interface {
	// this interface defines the method Save and its return type as error
	Save() error
	// you can also define the information which arguments the method takes
	// (you don't have to name the arguments only the type is important)
	// Save(int, string) error
}

// for interfaces that only define one method there is naming convention
// that it (interface) should be named like the method in it with added
// "er" postfix

func main() {
	title, content := getNoteData()
	todoText := getUserInput("Todo text: ")

	todo, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	todo.Display()
	err = saveData(todo)

	if err != nil {
		return
	}

	userNote.Display()
	err = saveData(userNote)

	if err != nil {
		return
	}
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	return text
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")

	return title, content
}

// here writing an function with interface which makes it possible for us to
// provide different structs for this method (here todo and note). These structs
// can be anything the only real requirement in golang is that they do implement
// specified method (both todo and note have method save that returns an error)
func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("Saving succeeded!")

	return nil
}
