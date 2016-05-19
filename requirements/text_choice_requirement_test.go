package requirements

import "testing"

func TestTextChoiceRequirement(t *testing.T) {
	reqUUID := "bob"

	var choices []UserInput
	appleInput := NewInputText("apple", reqUUID)
	orangeInput := NewInputText("orange", reqUUID)
	lemonInput := NewInputText("lemon", reqUUID)

	fakeInput := NewInputText("fake", reqUUID)

	choices = append(choices, &appleInput)
	choices = append(choices, &orangeInput)
	choices = append(choices, &lemonInput)

	req, _ := NewRequirementTextChoice(reqUUID, choices)
	err := req.Data(&orangeInput)
	if err != nil {
		t.Error("Error occured when setting correct user input", err)
	}
	if !req.Met() {
		t.Error("TestTextChouceRequirement::Requirement should be met, but isn't")
	}

	err = req.Data(&fakeInput)
	if err != nil {
		t.Error("Error occured when setting fake user input", err)
	}
	if req.Met() {
		t.Error("TestTextChouceRequirement::Requirement shouldn't be met but is")
	}
}
