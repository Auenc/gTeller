package requirements

import "testing"

func TestTextInputNoCondition(t *testing.T) {
	reqID := "bob"
	in := NewInputText("lewis", reqID)

	req, err := NewRequirementText(reqID)
	if err != nil {
		t.Error("TestTextInputEqualCondition::Error creating requirement", err)
	}

	err = req.Data(&in)
	if err != nil {
		t.Error("TestTextInputNoCondition::Error providing requirement with input", err)
	}
	if !req.Met() {
		t.Errorf("TestTextInputNoCondition::Expected %#v.Met() returned false", req)
	}
}

func TestTextInputEqualCondition(t *testing.T) {
	reqID := "bob"
	pass := NewInputText("lewis", reqID)
	fail := NewInputText("apple", reqID)

	req, err := NewRequirementText(reqID)
	if err != nil {
		t.Error("TestTextInputEqualCondition::Error creating requirement", err)
	}
	con, err := NewConditionEqual("bob", "lewis")
	if err != nil {
		t.Error("TestTextInputEqualCondition::Error creating condition", err)
	}
	err = req.Condition(&con)
	if err != nil {
		t.Error("TestTextInputEqualCondition::Error setting condition to requirement", err)
		return
	}

	err = req.Data(&pass)
	if err != nil {
		t.Error("TestTextInputEqualCondition::Error providing requirement with input", err)
	}
	if !req.Met() {
		t.Errorf("TestTextInputEqualCondition::Expected %#v.Met() to return true", req)
	}

	err = req.Data(&fail)
	if err != nil {
		t.Error("TestTextInputEqualCondition::Error providing requirement with input", err)
	}
	if req.Met() {
		t.Errorf("TestTextInputEqualCondition::Expected %#v.Met() to return false", req)
	}
}
