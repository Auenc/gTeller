package requirements

import "errors"

const (
	//RequirementText is an integer that represents a Text Requirement
	RequirementText = 1
	//RequirementItemChoice is an integer that represents a Item Choice Requirement
	RequirementItemChoice = 2
)

//Requirement is an interface that allows for complex requirements that can be
//attached to items in order to specify certain inputs from the customer.
type Requirement interface {
	Name() string
	Reference() string
	SetReference(string)
	HasOptions() bool
	Options() []UserInput
	SetOptions([]UserInput) error
	Data(UserInput) error
	Met() bool
	Condition(Condition) error
	Supported(Condition) bool
	GetCondition() Condition
	GetData() UserInput
	Type() int
	ID() string
	SetID(string)
}

func LoadRequirement(t int) (Requirement, error) {
	switch t {
	case RequirementText:
		return &TextRequirement{}, nil
	case RequirementItemChoice:
		return &ItemChoiceRequirement{}, nil
	default:
		return nil, errors.New("Unknown type")
	}
}
