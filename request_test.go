package gTeller

import (
	"testing"

	"github.com/auenc/gTeller/filter"
)

func TestItems(t *testing.T) {
	req := Request{Host: "http://localhost:8081/api/"}
	items, err := req.Items(filter.ItemsFilter{})
	if err != nil {
		t.Errorf("TestItems::Unexpected error getting items %s", err)
	}

	if len(items) != 3 {
		t.Errorf("TestItems::Expected to get 3 items but got %d", len(items))
	}
}
