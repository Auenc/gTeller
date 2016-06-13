package requirements

import (
	"encoding/json"
	"errors"
	"reflect"
)

//NewConditionMultipleEquals returns a new Equal Condition
func NewConditionMultipleEquals(uuid string, data []UserInput) (MultipleEqualsCondition, error) {
	return MultipleEqualsCondition{id: uuid, kind: ConditonMultipleEquals, targets: data}, nil
}

//MultipleEqualsCondition represents
type MultipleEqualsCondition struct {
	id      string
	kind    int
	targets []UserInput
}

//ConditionMultipleEqualsSave is the struct used for saving MultipleEqualsCondition
type ConditionMultipleEqualsSave struct {
	ID      string
	Targets []string
}

//SetID take a guess
func (con *MultipleEqualsCondition) SetID(n string) {
	con.id = n
}

//ID returns the Conditions ID
func (con *MultipleEqualsCondition) ID() string {
	return con.id
}

//Condition returns the interface that has to be met
func (con *MultipleEqualsCondition) Condition() interface{} {
	return con.targets[0]
}

//SetTargets sets the conditions targets to the specified list of strings
func (con *MultipleEqualsCondition) SetTargets(req string, targets []UserInput) error {

	con.targets = targets
	return nil
}

//Load loads a condition based of a provided JSON byte slice
func (con *MultipleEqualsCondition) Load(req string, src []byte) error {
	var data ConditionMultipleEqualsSave
	if err := json.Unmarshal(src, &data); err != nil {
		return err
	}

	inputs := make([]UserInput, len(data.Targets))
	for i, el := range data.Targets {
		in := NewInputText(el, req)
		inputs[i] = &in
	}

	con.targets = inputs
	con.id = data.ID
	return nil
}

//Save returns a JSON string containing the data required to create this condition
func (con *MultipleEqualsCondition) Save() (string, error) {
	list := make([]string, len(con.targets))
	for i, el := range con.targets {
		if tmp, ok := el.Data().(string); ok {
			list[i] = tmp
		} else {
			return "", errors.New("Targets are not strings!")
		}
	}

	d := ConditionMultipleEqualsSave{ID: con.id, Targets: list}
	b, err := json.Marshal(d)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

//Type returns the type of the EqualCondition
func (con *MultipleEqualsCondition) Type() int {
	return con.kind
}

//Valid returns true if data is int && data == target
func (con *MultipleEqualsCondition) Valid(dataList ...interface{}) bool {
	valid := true
	if len(dataList) != len(con.targets) {
		return false
	}

	found := make([]bool, len(con.targets))

	for i, el := range con.targets {
		if found[i] {
			continue
		}
		for _, data := range dataList {
			if reflect.DeepEqual(el, data) {
				found[i] = true
			}
		}
	}

	for _, el := range found {
		if !el {
			valid = false
		}
	}

	return valid
}

//Name returns the string "Equal conditon"
func (con *MultipleEqualsCondition) Name() string {
	return "Multiple equals condition"
}
