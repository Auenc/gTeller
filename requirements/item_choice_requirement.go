package requirements

import (
	"errors"
	"fmt"
	"reflect"
)

//ItemChoiceRequirement is a requirement that accepts a string input
type ItemChoiceRequirement struct {
	id        string
	reference string
	choices   []UserInput
	data      UserInput
	condition Condition
}

//NewRequirementItemChoice returns a new TextChoiceRequirement
func NewRequirementItemChoice(uuid string, choices []UserInput) (ItemChoiceRequirement, error) {
	return ItemChoiceRequirement{id: uuid, data: nil, choices: choices}, nil
}

//Reference returns the friendly reference that the requirement is known as
func (req *ItemChoiceRequirement) Reference() string {
	return req.reference
}

//SetReference sets the requirements reference to the specified string
func (req *ItemChoiceRequirement) SetReference(ref string) {
	req.reference = ref
}

//ID returns the ID of the Requirement
func (req *ItemChoiceRequirement) ID() string {
	return req.id
}

//Type returns an interger representing the type of Requirement
func (req *ItemChoiceRequirement) Type() int {
	return RequirementItemChoice
}

//GetCondition returns the Condition applied to the Requirement
func (req *ItemChoiceRequirement) GetCondition() Condition {
	return req.condition
}

//HasOptions returns false
func (req *ItemChoiceRequirement) HasOptions() bool {
	return true
}

//Options returns a list of the items possible to choose
func (req *ItemChoiceRequirement) Options() []UserInput {

	return req.choices
}

//SetOptions sets the options associated with ItemChoiceRequirement to the specified inputs
func (req *ItemChoiceRequirement) SetOptions(options []UserInput) error {
	req.choices = options
	if req.GetCondition() != nil {
		cond := req.GetCondition()

		fmt.Printf("Giving condition targets %v\n", options)
		err := cond.SetTargets(req.ID(), options)
		if err != nil {
			return err
		}
	}
	return nil
}

//Name returns the string "Item choice requirement"
func (req *ItemChoiceRequirement) Name() string {
	return "Item Choice Requirement"
}

//Data is a method that provides the ItemChoiceRequirement with a given UserInput
func (req *ItemChoiceRequirement) Data(dataList ...UserInput) error {
	if len(dataList) == 0 {
		return errors.New("No data provided")
	}
	data := dataList[0]

	if data.For() != req.id {
		req.data = nil
		return errors.New("Incompatible input given")
	}
	if req.HasOptions() {
		found := false
		for _, el := range req.choices {
			if reflect.DeepEqual(data.Data(), el.Data()) {
				found = true
			}
		}
		if !found {
			return errors.New("Option not found in choices")
		}
	}

	req.data = data

	return nil
}

//Supported returns true if TextChoiceRequirement supports the given Condition
func (req *ItemChoiceRequirement) Supported(con Condition) bool {
	switch con.Type() {
	case ConditionContains:
		return true
	default:
		return false
	}
}

//SetID sets the id of the Requirement to the specified storing
func (req *ItemChoiceRequirement) SetID(n string) {
	req.id = n
}

//Condition states that input given for this requirement has to meet the given condition
func (req *ItemChoiceRequirement) Condition(con Condition) error {
	if !req.Supported(con) {
		er := fmt.Sprintf("Condition %d not supported by %s", con.Type(), req.Name())
		return errors.New(er)
	}
	req.condition = con
	req.choices = con.Targets()

	return nil
}

func (req *ItemChoiceRequirement) GetData() UserInput {
	return req.data
}

//Met returns true if the data != nil && data instanceof string && Condition.Valid(data)
func (req *ItemChoiceRequirement) Met() bool {
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
		fmt.Println("checking", tmp.Data(), "against", req.data.Data())
		if reflect.DeepEqual(tmp.Data(), req.data.Data()) {
			found = true
		}
	}
	//If the input was not an option specified, input has not met requirement
	if !found {
		fmt.Println("Not found option", req.data.Data())
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
