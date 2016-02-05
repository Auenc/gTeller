package filter

import (
	"testing"
)

func TestCondition(t *testing.T) {
	con := Condition{ConditionEquals, 10}
	if !con.Valid(10) {
		t.Errorf("TestCondition::Expected EqualsCondition.Valid(10) returned false ")
	}
}
