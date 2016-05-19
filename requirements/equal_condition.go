package requirements

import (
	"encoding/json"
	"reflect"
)

//NewConditionEqual returns a new Equal Condition
func NewConditionEqual(uuid string, data interface{}) (EqualCondition, error) {
	return EqualCondition{id: uuid, kind: ConditionEqual, target: data}, nil
}

//EqualCondition represents
type EqualCondition struct {
	id     string
	kind   int
	target interface{}
}

type equalConditionSave struct {
	ID     string
	Target interface{}
}

//SetID take a guess
func (con *EqualCondition) SetID(n string) {
	con.id = n
}

//ID returns the Conditions ID
func (con *EqualCondition) ID() string {
	return con.id
}

//Condition returns the interface that has to be met
func (con *EqualCondition) Condition() interface{} {
	return con.target
}

//Load loads a condition based of a provided JSON byte slice
func (con *EqualCondition) Load(src []byte) error {
	var data equalConditionSave
	if err := json.Unmarshal(src, &data); err != nil {
		return err
	}

	con.target = data.Target
	con.id = data.ID
	return nil
}

//Save returns a JSON string containing the data required to create this condition
func (con *EqualCondition) Save() (string, error) {
	d := equalConditionSave{ID: con.id, Target: con.target}
	b, err := json.Marshal(d)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

//Type returns the type of the EqualCondition
func (con *EqualCondition) Type() int {
	return con.kind
}

//Valid returns true if data is int && data == target
func (con *EqualCondition) Valid(data interface{}) bool {
	return reflect.DeepEqual(con.target, data)
}

//Name returns the string "Equal conditon"
func (con *EqualCondition) Name() string {
	return "Equal condition"
}
