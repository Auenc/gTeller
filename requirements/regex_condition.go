package requirements

import (
	"encoding/json"
	"regexp"
)

//NewConditionRegex returns a new Regex Condition
func NewConditionRegex(uuid, data string) (RegexCondition, error) {
	var con RegexCondition
	reg, err := regexp.Compile(data)
	if err != nil {
		return con, err
	}
	return RegexCondition{kind: ConditionRegex, patternString: data, pattern: reg, id: uuid}, nil
}

//RegexCondition represents
type RegexCondition struct {
	id            string
	kind          int
	patternString string
	pattern       *regexp.Regexp
}
type regexConditionSave struct {
	ID      string
	Pattern string
}

//SetID take a guess
func (con *RegexCondition) SetID(n string) {
	con.id = n
}

//ID returns the Conditions ID
func (con *RegexCondition) ID() string {
	return con.id
}

//Condition returns the interface that has to be met
func (con *RegexCondition) Condition() interface{} {
	return con.patternString
}

//Load loads a condition based of a provided JSON byte slice
func (con *RegexCondition) Load(src []byte) error {
	var data regexConditionSave
	if err := json.Unmarshal(src, &data); err != nil {
		return err
	}

	//Change REGEX to work with being stored
	con.id = data.ID
	return nil
}

//Save returns a JSON string containing the data required to create this condition
func (con *RegexCondition) Save() (string, error) {
	d := regexConditionSave{ID: con.id, Pattern: con.patternString}
	b, err := json.Marshal(d)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

//Type returns the type of the RegexCondition
func (con *RegexCondition) Type() int {
	return con.kind
}

//Valid returns true if given input matches the previously stated regex pattern
func (con *RegexCondition) Valid(data interface{}) bool {
	if s, ok := data.(string); ok {
		return con.pattern.MatchString(s)
	}
	return false
}

//Name returns the string "Regex conditon"
func (con *RegexCondition) Name() string {
	return "Regex condition"
}
