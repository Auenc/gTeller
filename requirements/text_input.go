package requirements

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
