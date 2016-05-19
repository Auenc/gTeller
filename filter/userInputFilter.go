package filter

import "github.com/auenc/gTeller-core/requirements"

type UserInputFilter struct {
	For  Condition
	Data Condition
}

func (filter *UserInputFilter) Valid(input requirements.UserInput) bool {
	emptyCon := Condition{}
	valid := true

	//Type condition
	if filter.For != emptyCon {
		//Checking if requirementCondtion matches ConditionType filter.
		//IF it doesn't, return false.
		if !filter.For.Valid(input.For()) {
			valid = false
			return valid
		}
	}

	//ID condition
	if filter.Data != emptyCon {
		//Checking if requirementCondtion matches ConditionType filter.
		//IF it doesn't, return false.
		if !filter.Data.Valid(input.Data()) {
			valid = false
			return valid
		}
	}

	return valid
}

func (filter *UserInputFilter) Filter(inputs []requirements.UserInput) []requirements.UserInput {
	var valid []requirements.UserInput

	for _, input := range inputs {
		if filter.Valid(input) {
			valid = append(valid, input)
		}
	}
	return valid
}
