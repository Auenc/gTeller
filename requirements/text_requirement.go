package requirements

import (
	"errors"
	"fmt"
	"reflect"
)

//TextRequirement is a requirement that accepts a string input
type TextRequirement struct {
	id        string
	data      UserInput
	reference string
	condition Condition
}

//NewRequirementText returns a new TextRequirement
func NewRequirementText(uuid string) (TextRequirement, error) {
	return TextRequirement{id: uuid, data: nil}, nil
}

//Reference returns the friendly reference the requirement is known as
func (req *TextRequirement) Reference() string {
	return req.reference
}

//SetReference sets the requirements reference to the specified string
func (req *TextRequirement) SetReference(ref string) {
	req.reference = ref
}

func (req *TextRequirement) GetData() UserInput {
	return req.data
}

//ID returns the ID of the Requirement
func (req *TextRequirement) ID() string {
	return req.id
}

//Type returns an interger representing the type of Requirement
func (req *TextRequirement) Type() int {
	return RequirementText
}

//GetCondition returns the Condition applied to the Requirement
func (req *TextRequirement) GetCondition() Condition {
	return req.condition
}

//HasOptions returns false
func (req *TextRequirement) HasOptions() bool {
	return false
}

//Options returns an empty list of UserInput as TextInput has no options
func (req *TextRequirement) Options() []UserInput {
	var options []UserInput
	return options
}

//SetOptions does nothing for TextRequirement as TextRequirement has no options
func (req *TextRequirement) SetOptions(options []UserInput) error {
	return nil
}

//Name returns the string "Text Requirement"
func (req *TextRequirement) Name() string {
	return "Text Requirement"
}

//Data is a method that provides the TextRequirement with a given UserInput
func (req *TextRequirement) Data(dataList ...UserInput) error {
	if len(dataList) == 0 {
		return errors.New("No data provided")
	}

	data := dataList[0]

	if data.For() != req.id {
		return errors.New("Incompatible input given")
	}

	req.data = data

	return nil
}

//Supported returns true if TextRequirement supports the given Condition
func (req *TextRequirement) Supported(con Condition) bool {
	if con != nil {
		switch con.Type() {
		case ConditionEqual:
			return true
		default:
			return false
		}
	}
	return false
}

//SetID sets the id of the Requirement to the specified storing
func (req *TextRequirement) SetID(n string) {
	req.id = n
}

//Condition states that input given for this requirement has to meet the given condition
func (req *TextRequirement) Condition(con Condition) error {
	if con != nil {
		if !req.Supported(con) {
			er := fmt.Sprintf("Condition %d not supported by %s", con.Type(), req.Name())
			return errors.New(er)
		}
		req.condition = con
	}

	return nil
}

//Met returns true if the data != nil && data instanceof string && Condition.Valid(data)
func (req *TextRequirement) Met() bool {
	//emptyInput := new(UserInput)
	//If data is empty requirement not met
	if req.data == nil {
		return false
	}

	//Check if data is string
	if reflect.TypeOf(req.data.Data()).Name() != "string" {
		return false
	}

	//If requirement has a condition
	if req.condition != nil {
		//If condition is not met
		if !req.condition.Valid(req.data.Data()) {
			return false
		}
	}
	//Success
	return true
}
