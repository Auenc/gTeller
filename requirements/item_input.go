package requirements

//NewInputItem accepts a string representing the data from the user
//and the string ID of the request
func NewInputItem(itemID, req string) ItemInput {
	return ItemInput{target: req, data: itemID}
}

//ItemInput is an object that holds string input from the user
type ItemInput struct {
	target string
	data   string
}

//For returns an int representing the Requirement the input is for
func (in *ItemInput) For() string {
	return in.target
}

//Data returns the data held by the TextInput object
func (in *ItemInput) Data() interface{} {
	return in.data
}
