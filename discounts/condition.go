package discounts

import "encoding/json"

const (
	//CounterConditionType is a string representing the type of a counter condition
	CounterConditionType = "COUNTER"
)

//Condition is an interface that allows for multiple types of Discount conditions
//A Discount condition states whether the discount is usable.
type Condition interface {
	//IsMet returns whether the Discount object is useable
	IsMet() bool
	//Type returns a string representing the Conditions type
	Type() string
	//Data returns a string representing the data required to run the condition
	//The string should be in a format that allows for saving within a mysql text
	//field
	Data() (string, error)
	//ParseData parsed the inputted json string to load the data values of a conditon
	ParseData(string)

	//Use is a method that should be called when the condition has returned true and
	//you wish to use the conditon
	Use()
}

//ConditionParseable is an object that can hold the data for a condition over json
type ConditionParseable struct {
	ConditionType string
	Data          string
}

//CounterCondition is a condition that represents x < y where x is the number of
//times of times the conditon is used and y is some constant representing the max
//amount of times the contion can be ran
type CounterCondition struct {
	ConditionType string
	Iteration     int
	MaxIts        int
}

//LoadCondition takes a type and data as string and returns the correct condition object
func LoadCondition(conType, conData string) Condition {
	var con Condition
	if conType == CounterConditionType {
		con = &CounterCondition{conType, 0, 0}
		con.ParseData(conData)
	}
	return con
}

//IsMet returns true if CounterCondition.iteration < CounterCondition.maxIts
func (con *CounterCondition) IsMet() bool {
	return con.Iteration < con.MaxIts
}

//Type returns a string representing the CounterCondition type
func (con *CounterCondition) Type() string {
	return con.ConditionType
}

//Data returns a json string containing the current iteration and the maximum
//iteration. This json string is intended to be loaded by the ParseData() method
func (con *CounterCondition) Data() (string, error) {
	var data string
	dataObj := struct {
		Iteration int
		MaxIts    int
	}{
		con.Iteration,
		con.MaxIts,
	}
	marsh, err := json.Marshal(dataObj)
	if err != nil {
		return data, err
	}
	return string(marsh), nil
}

//ParseData parses the json string provided, extracting the current iteration
//and maximum iterations possible
func (con *CounterCondition) ParseData(raw string) {
	dataObj := struct {
		Iteration int
		MaxIts    int
	}{}

	json.Unmarshal([]byte(raw), &dataObj)
	con.Iteration = dataObj.Iteration
	con.MaxIts = dataObj.MaxIts
}

//Use should be called when the condition is used
func (con *CounterCondition) Use() {
	if con.IsMet() {
		con.Iteration = con.Iteration + 1
	}
}
