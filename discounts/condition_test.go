package discounts

import (
	"encoding/json"
	"fmt"
	"testing"
)

type counterData struct {
	Iteration int
	MaxIts    int
}

func loadCounterData(raw string) counterData {
	dataObj := counterData{}
	json.Unmarshal([]byte(raw), &dataObj)
	return dataObj
}

func testCounterConditionData(expectedIterations, expectedMaxIts int, data counterData, t *testing.T) {
	if data.Iteration != expectedIterations {
		t.Error("Unexpected value for CounterCondition.Iteration. Expected", expectedIterations, "recieved", data.Iteration)
	}
	if data.MaxIts != expectedMaxIts {
		fmt.Println("RAW DATA DUMP\n", data)
		t.Error("Unexpected value for CounterCondition.maxIts. Expected", expectedIterations, "recieved", data.MaxIts)
	}
}

func TestCounterCondition(t *testing.T) {
	//Creating counter condition as condition interface
	var con Condition
	con = &CounterCondition{CounterConditionType, 0, 10}

	//Testing if condition can be ran
	if !con.IsMet() {
		t.Error("IsMet returned false when it is expected to return true!")
	}

	//Testing Data()
	raw, err := con.Data()
	if err != nil {
		t.Error("Unexpected error while calling CounterCondition.Data()", err)
	}

	//Decoding data
	dataObj := loadCounterData(raw)
	//Testing if data is still the same
	testCounterConditionData(0, 10, dataObj, t)

	//Changing the values
	dataObj.Iteration = 3
	dataObj.MaxIts = 7

	//Marshaling json
	newraw, _ := json.Marshal(dataObj)

	//Parsing data
	con.ParseData(string(newraw))

	raw, err = con.Data()
	if err != nil {
		t.Error("Unexpected error while calling CounterCondition.Data()", err)
	}

	//Decoding data
	dataObj = loadCounterData(raw)
	testCounterConditionData(3, 7, dataObj, t)

	//Testing Use()

	con.Use()
	raw, err = con.Data()
	if err != nil {
		t.Error("Unexpected error while calling CounterCondition.Data()", err)
	}

	//Decoding data
	dataObj = loadCounterData(raw)
	testCounterConditionData(4, 7, dataObj, t)

	//Testing if the conditon will stop after multiple use calls
	con.Use() //5
	con.Use() //6
	con.Use() //7

	if con.IsMet() {
		t.Error("CounterCondition.IsMet() returned true when is should return false!")
	}
}
