package requirements

import (
	"errors"
	"strconv"
)

const (
	//ConditionEqual represents the EqualCondition object
	ConditionEqual = 1
	//ConditionRegex represents the RegexCondition object
	ConditionRegex = 2
)

//Condition is an interface that exists to allow for complex conditions that can
//be attached to requirements.
//Eg, a TextRequirement may have a length condition that states Input n has must
//have X or more characters.
type Condition interface {
	ID() string
	SetID(string)
	Type() int
	Valid(interface{}) bool
	Condition() interface{}
	Name() string
	Save() (string, error)
	Load([]byte) error
}

//Conditions returns a slice of all available Condition objects
func Conditions() []Condition {
	var cons []Condition

	cons = append(cons, &EqualCondition{})
	cons = append(cons, &RegexCondition{})

	return cons
}

//LoadCondition attempts to load an EmptyCondition with a given type
func LoadCondition(t int) (Condition, error) {
	switch t {
	case ConditionEqual:
		return &EqualCondition{kind: t}, nil
	case ConditionRegex:
		return &RegexCondition{kind: t}, nil
	default:
		s := strconv.Itoa(t)
		return nil, errors.New("condition type not found " + s)
	}
}
