package requirements

import "errors"

const (
	//RequirementText is an integer that represents a Text Requirement
	RequirementText = 1
	//RequirementItemChoice is an integer that represents a Item Choice Requirement
	RequirementItemChoice = 2
	//RequirementItems is a requirement that states the input must be the list of specifed inputs
	RequirementItems = 3
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
	Data(...UserInput) error
	Met() bool
	Condition(Condition) error
	Supported(Condition) bool
	GetCondition() Condition
	GetData() UserInput
	Type() int
	ID() string
	SetID(string)
	Parseable() (ParseableRequirement, error)
}

type ParseableRequirement struct {
	UUID          string
	Reference     string
	Type          int
	ConditionType int
	ConditionSave string
	Data          string
	Options       []string
}

func (parse *ParseableRequirement) Parse() (Requirement, error) {
	var req Requirement

	req, err := LoadRequirement(parse.Type)
	if err != nil {
		return req, err
	}

	req.SetReference(parse.Reference)

	req.SetID(parse.UUID)

	/*	Load condition type									*/
	con, err := LoadCondition(parse.ConditionType)
	if err != nil {
		return req, err
	}

	/*	Load data into condition						*/
	err = con.Load(parse.UUID, []byte(parse.ConditionSave))
	if err != nil {
		return req, err
	}

	/*	Load condition into requirement			*/
	err = req.Condition(con)
	if err != nil {
		return req, err
	}

	/*	Load UserInput	*/
	in := NewInputText("", parse.UUID)
	err = in.Load(parse.Data)
	if err != nil {
		return req, err
	}

	/*	Load options		*/
	options := make([]UserInput, len(parse.Options))
	for i, op := range parse.Options {
		in := NewInputText("", parse.UUID)
		err = in.Load(op)
		if err != nil {
			return req, err
		}
		options[i] = &in
	}

	req.SetOptions(options)

	/*	Apply data		*/
	err = req.Data(&in)
	if err != nil {
		return req, err
	}

	return req, nil
}

func LoadRequirement(t int) (Requirement, error) {
	switch t {
	case RequirementText:
		return &TextRequirement{}, nil
	case RequirementItemChoice:
		return &ItemChoiceRequirement{}, nil
	case RequirementItems:
		return &ItemsRequirement{}, nil
	default:
		return nil, errors.New("Unknown type")
	}
}
