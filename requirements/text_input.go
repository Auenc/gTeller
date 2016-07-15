package requirements

import (
	"encoding/json"
	"fmt"
)

//NewInputText accepts a string representing the data from the user
//and the string ID of the request
func NewInputText(text, req string) TextInput {
	return TextInput{target: req, data: text}
}

//TextInput is an object that holds string input from the user
type TextInput struct {
	target string
	data   string
}

//For returns an int representing the Requirement the input is for
func (in *TextInput) For() string {
	return in.target
}

//Data returns the data held by the TextInput object
func (in *TextInput) Data() interface{} {
	return in.data
}

//Save generates a JSON string to represent the input
func (in *TextInput) Save() (string, error) {
	var save string

	bdata, err := json.Marshal(in.Data())
	if err != nil {
		return save, err
	}

	save = string(bdata)

	fmt.Println("Requirements::UserInput::TextInput::Saving data", save, "from", in.Data())

	return save, nil
}

//Load loads a user input based of a string source
func (in *TextInput) Load(src string) error {
	var data string
	err := json.Unmarshal([]byte(src), &data)
	if err != nil {
		return err
	}

	in.data = data

	return nil
}
