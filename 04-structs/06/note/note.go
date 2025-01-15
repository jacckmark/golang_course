package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	// here adding struct tag that would make it possible to name the
	// fields different when creating JSON object using json.Marshal method (important
	// to notice that by default it will do nothing, this can be used by some methods
	// external libraries like 'json' one, but without them it changes nothing in how
	// your struct works)
	Title     string `json:"titelllo"`
	Content   string
	CreatedAt time.Time
}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("invalid input")
	}

	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}

func (note Note) Display() {
	fmt.Printf("Your note titled '%v' has the following content:\n\n'%v'\n\n", note.Title, note.Content)
}

func (note Note) Save() error {
	// here we are replacing all not necessary data in the note title to make it
	// possible to save
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"
	// this method allows us to convert a struct into a JSON (this method produces
	// the byte code already so no need to convert it in 'WriteFile' method). Important
	// to no that the method Marshal will produce the data only from the fields that
	// are exported in struct (written in upperCase)
	json, err := json.Marshal(note)

	if err != nil {
		return err
	}

	// this method returns nil if all is good and we were able to write to file and
	// an error object if we failed
	return os.WriteFile(fileName, json, 0644)
}
