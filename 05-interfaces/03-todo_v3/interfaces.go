package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/interface/note"
	"example.com/interface/todo"
)

type saver interface {
	Save() error
}

type displayer interface {
	Display()
}

// we could just repeat the interface methods here, but it would make no sense
// instead we use the embedding and prepare the interface that combines these
// two interfaces into one below
// type outputtable interface {
// 	Save() error
// 	Display()
// }

type outputtable interface {
	saver
	displayer
}

func main() {
	// using the function that takes any argument in
	printSomething("this is the day")
	printSomething(1)
	printSomething(1.234)

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

	err = outputData(todo)

	if err != nil {
		return
	}

	err = outputData(userNote)

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

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("Saving succeeded!")

	return nil
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

// this is golang any interface, so this function can take any type of value here
// (you have to however cast the var into type before using)
func printSomething(value interface{}) {
	// we can use in go something called type switch which is pretty useful when we
	// work with any types. It allows us to check the type of out value and execute
	// code only for this type of value
	switch value.(type) {
	case int:
		fmt.Println("Integer: ", value)
	case float64:
		fmt.Println("Float64: ", value)
	case string:
		fmt.Println("String: ", value)
	}

	// we can also check for the type like this and get the typed value and boolean
	// with information if it is of given type
	typedVal, ok := value.(int)
	// typedVal will be of type int so we would have to cast it into expected type
	// before using later on
	fmt.Println(typedVal, ok)
}
