package requirements

import (
	"errors"
	"fmt"
	"reflect"
)

//TextChoiceRequirement is a requirement that accepts a string input
type TextChoiceRequirement struct {
	id        string
	choices   []UserInput
	data      UserInput
	condition Condition
}

//NewRequirementTextChoice returns a new TextChoiceRequirement
func NewRequirementTextChoice(uuid string, choices []UserInput) (TextChoiceRequirement, error) {
	return TextChoiceRequirement{id: uuid, data: nil, choices: choices}, nil
}

//ID returns the ID of the Requirement
func (req *TextChoiceRequirement) ID() string {
	return req.id
}

//Type returns an interger representing the type of Requirement
func (req *TextChoiceRequirement) Type() int {
	return RequirementText
}

//GetCondition returns the Condition applied to the Requirement
func (req *TextChoiceRequirement) GetCondition() Condition {
	return req.condition
}

//HasOptions returns false
func (req *TextChoiceRequirement) HasOptions() bool {
	return false
}

//Options returns an empty list of UserInput as TextInput has no options
func (req *TextChoiceRequirement) Options() []UserInput {

	return req.choices
}

//Name returns the string "Text Requirement"
func (req *TextChoiceRequirement) Name() string {
	return "Text Choice Requirement"
}

//Data is a method that provides the TextChoiceRequirement with a given UserInput
func (req *TextChoiceRequirement) Data(dataList ...UserInput) error {
	if len(dataList) == 0 {
		return errors.New("No data provided")
	}

	data := dataList[0]

	if data.For() != req.id {
		req.data = nil
		return errors.New("Incompatible input given")
	}

	req.data = data

	return nil
}

//Supported returns true if TextChoiceRequirement supports the given Condition
func (req *TextChoiceRequirement) Supported(con Condition) bool {
	switch con.Type() {
	case ConditionEqual:
		return true
	default:
		return false
	}
}

//SetID sets the id of the Requirement to the specified storing
func (req *TextChoiceRequirement) SetID(n string) {
	req.id = n
}

//Condition states that input given for this requirement has to meet the given condition
func (req *TextChoiceRequirement) Condition(con Condition) error {
	if !req.Supported(con) {
		er := fmt.Sprintf("Condition %d not supported by %s", con.Type(), req.Name())
		return errors.New(er)
	}
	req.condition = con

	return nil
}

//Met returns true if the data != nil && data instanceof string && Condition.Valid(data)
func (req *TextChoiceRequirement) Met() bool {
	//emptyInput := new(UserInput)
	//If data is empty requirement not met
	if req.data == nil {
		return false
	}

	//Check if data is string
	if reflect.TypeOf(req.data.Data()).Name() != "string" {
		return false
	}
	//Checking if the input was an option specified.
	found := false
	for _, tmp := range req.choices {
		//fmt.Println("checking", tmp.Data(), "against", req.data.Data())
		if reflect.DeepEqual(tmp.Data(), req.data.Data()) {
			found = true
		}
	}
	//If the input was not an option specified, input has not met requirement
	if !found {
		//fmt.Println("Not found option", req.data.Data())
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

func (req *TextChoiceRequirement) GetData() UserInput {
	return req.data
}

func (req *TextChoiceRequirement) Parseable() (ParseableRequirement, error) {
	var parseable ParseableRequirement

	//Save uuid
	parseable.UUID = req.ID()
	//Save type
	parseable.Type = req.Type()

	//parseable.Reference = req.Reference()

	//If we have a condition - save condition uuid
	if req.GetCondition() != nil {
		parseable.ConditionType = req.GetCondition().Type()
		conSave, err := req.GetCondition().Save()
		if err != nil {
			return parseable, err
		}
		parseable.ConditionSave = conSave
	}

	//If we have input - save input
	if req.GetData() != nil {
		//fmt.Println("Requirements::", req.Name(), "::Has data!")
		inputSave, err := req.GetData().Save()
		if err != nil {
			return parseable, err
		}
		parseable.Data = inputSave
	} else {
		//fmt.Println("Requirements::", req.Name(), "::Has no data!")
	}

	options := make([]string, len(req.choices))
	for i, op := range req.choices {
		inputSave, err := op.Save()
		if err != nil {
			return parseable, err
		}
		options[i] = inputSave
	}

	parseable.Options = options

	//fmt.Println("Requirements::TextRequirement::Providing data", parseable.Data)

	return parseable, nil
}
