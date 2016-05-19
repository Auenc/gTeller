package requirements

import "errors"

const (
	//RequirementText is an integer that represents a Text Requirement
	RequirementText = 1
)

//Requirement is an interface that allows for complex requirements that can be
//attached to items in order to specify certain inputs from the customer.
type Requirement interface {
	Name() string
	HasOptions() bool
	Options() []UserInput
	Data(UserInput) error
	Met() bool
	Condition(Condition) error
	Supported(Condition) bool
	GetCondition() Condition
	Type() int
	ID() string
	SetID(string)
}

func LoadRequirement(t int) (Requirement, error) {
	switch t {
	case RequirementText:
		return &TextRequirement{}, nil
	default:
		return nil, errors.New("Unknown type")
	}
}
