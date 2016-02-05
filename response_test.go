package gTeller

type TestStructure struct {
	Val int
}

/*func TestResponseGet(t *testing.T) {
	correct := TestStructure{Val: 1}
	response := Response{ResponseData: correct}

	test := &TestStructure{}
	err := response.Get(test)
	if err != nil {
		t.Error("TestResponseGet::Error calling respoinse.Get", err)
	}

	if *test != correct {
		t.Errorf("Expected to have a value of %#v but instead got %#v", correct, test)
	}
}
*/
