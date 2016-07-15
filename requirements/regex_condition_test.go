package requirements

import "testing"

func TestRegextCondition(t *testing.T) {
	con, err := NewConditionRegex("bob", "apple")
	if err != nil {
		t.Error("Error creating condition", err)
	}

	if !con.Valid("apple") {
		t.Error("TestRegexCondition::Expected condition to be valid")
	}

	if con.Valid("fish") {
		t.Error("TestRegexCondition::Expected condition to be false!")
	}
}

func TestRegexWildcardQuery(t *testing.T) {
	con, err := NewConditionRegex("bob", ".")
	if err != nil {
		t.Error("Error creating condition", err)
	}

	if !con.Valid("apple") {
		t.Error("TestRegexCondition::Expected condition to be valid")
	}

	if !con.Valid("hello world") {
		t.Error("TestRegexCondition::Expected condition to be valid!")
	}
}
