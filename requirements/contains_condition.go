package requirements

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

//NewConditionEqual returns a new Equal Condition
func NewConditionContains(uuid string, data []UserInput) (ContainsCondition, error) {
	return ContainsCondition{id: uuid, kind: ConditionContains, targets: data}, nil
}

//EqualCondition represents
type ContainsCondition struct {
	id      string
	kind    int
	targets []UserInput
}

type containsConditionSave struct {
	ID      string
	Targets []string
}

//SetID take a guess
func (con *ContainsCondition) SetID(n string) {
	con.id = n
}

//ID returns the Conditions ID
func (con *ContainsCondition) ID() string {
	return con.id
}

//Condition returns the interface that has to be met
func (con *ContainsCondition) Condition() interface{} {
	return con.targets[0]
}

//Targets returns the list of inputs that the condition can match
func (con *ContainsCondition) Targets() []UserInput {
	return con.targets
}

//SetTargets sets the conditions targets to the specified list of strings
func (con *ContainsCondition) SetTargets(req string, targets []UserInput) error {

	con.targets = targets
	return nil
}

//Load loads a condition based of a provided JSON byte slice
func (con *ContainsCondition) Load(req string, src []byte) error {
	var data containsConditionSave
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
func (con *ContainsCondition) Save() (string, error) {
	list := make([]string, len(con.targets))
	for i, el := range con.targets {
		if tmp, ok := el.Data().(string); ok {
			list[i] = tmp
			fmt.Println("Saving string", tmp)
		} else {
			typeString := reflect.TypeOf(el.Data())
			errString := fmt.Sprintf("Targets are not strings! instead they are %s\n", typeString)
			return "", errors.New(errString)
		}
	}

	d := containsConditionSave{ID: con.id, Targets: list}
	b, err := json.Marshal(d)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

//Type returns the type of the EqualCondition
func (con *ContainsCondition) Type() int {
	return con.kind
}

//Valid returns true if data is int && data == target
func (con *ContainsCondition) Valid(dataList ...interface{}) bool {
	valid := false
	if len(dataList) == 0 {
		return false
	}

	data := dataList[0]

	for _, el := range con.targets {
		if reflect.DeepEqual(el, data) {
			valid = true
		}
	}

	return valid
}

//Name returns the string "Equal conditon"
func (con *ContainsCondition) Name() string {
	return "Contains condition"
}
