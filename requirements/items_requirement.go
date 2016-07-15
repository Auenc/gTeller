package requirements

import (
	"errors"
	"fmt"
	"reflect"
)

//ItemsRequirement is a requirement that accepts a string input
type ItemsRequirement struct {
	id        string
	reference string
	choices   []UserInput
	data      UserInput
	condition Condition
}

//NewRequirementItems returns a new TextChoiceRequirement
func NewRequirementItems(uuid string, choices []UserInput) (ItemsRequirement, error) {
	return ItemsRequirement{id: uuid, data: nil, choices: choices}, nil
}

//Reference returns the friendly reference that the requirement is known as
func (req *ItemsRequirement) Reference() string {
	return req.reference
}

//SetReference sets the requirements reference to the specified string
func (req *ItemsRequirement) SetReference(ref string) {
	req.reference = ref
}

//ID returns the ID of the Requirement
func (req *ItemsRequirement) ID() string {
	return req.id
}

//Type returns an interger representing the type of Requirement
func (req *ItemsRequirement) Type() int {
	return RequirementItems
}

//GetCondition returns the Condition applied to the Requirement
func (req *ItemsRequirement) GetCondition() Condition {
	return req.condition
}

//HasOptions returns false
func (req *ItemsRequirement) HasOptions() bool {
	return true
}

//Options returns a list of the items possible to choose
func (req *ItemsRequirement) Options() []UserInput {

	return req.choices
}

//SetOptions sets the options associated with ItemsRequirement to the specified inputs
func (req *ItemsRequirement) SetOptions(options []UserInput) error {
	req.choices = options
	if req.GetCondition() != nil {
		cond := req.GetCondition()

		err := cond.SetTargets(req.ID(), options)
		if err != nil {
			return err
		}
	}
	return nil
}

//Name returns the string "Item choice requirement"
func (req *ItemsRequirement) Name() string {
	return "Items Requirement"
}

//Data is a method that provides the ItemsRequirement with a given UserInput
func (req *ItemsRequirement) Data(dataList ...UserInput) error {
	if len(dataList) == 0 {
		return errors.New("No data given")
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
func (req *ItemsRequirement) Supported(con Condition) bool {
	switch con.Type() {
	case ConditonMultipleEquals:
		return true
	default:
		return false
	}
}

//SetID sets the id of the Requirement to the specified storing
func (req *ItemsRequirement) SetID(n string) {
	req.id = n
}

//Condition states that input given for this requirement has to meet the given condition
func (req *ItemsRequirement) Condition(con Condition) error {
	if !req.Supported(con) {
		er := fmt.Sprintf("Condition %d not supported by %s", con.Type(), req.Name())
		return errors.New(er)
	}
	req.condition = con

	return nil
}

func (req *ItemsRequirement) GetData() UserInput {
	return req.data
}

//Met returns true if the data != nil && data instanceof string && Condition.Valid(data)
func (req *ItemsRequirement) Met() bool {
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

func (req *ItemsRequirement) Parseable() (ParseableRequirement, error) {
	var parseable ParseableRequirement

	//Save uuid
	parseable.UUID = req.ID()
	//Save type
	parseable.Type = req.Type()

	parseable.Reference = req.Reference()

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
		fmt.Println("Requirements::", req.Name(), "::Has data!")
		inputSave, err := req.GetData().Save()
		if err != nil {
			return parseable, err
		}
		parseable.Data = inputSave
	} else {
		fmt.Println("Requirements::", req.Name(), "::Has no data!")
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

	fmt.Println("Requirements::TextRequirement::Providing data", parseable.Data)

	return parseable, nil
}
