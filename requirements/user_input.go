package requirements

//UserInput is an interface that contains methods that enforce an object to be
//able to describe the data it holds, and what requirement the data is concerning
//as well as return the data
type UserInput interface {
	For() string
	Data() interface{}
}
