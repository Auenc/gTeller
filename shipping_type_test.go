package gTeller

import (
	"net/http"
	"testing"

	"github.com/auenc/gTeller/filter"
	"github.com/auenc/gTeller/shipping"
)

func TestShippingTypeList(t *testing.T) {
	gt := New("http://localhost:8081/api/", apiKey)

	req := gt.ShippingTypes()
	req.RequestData = ListShippingTypeRequest{}
	resp, err := req.List()
	if err != nil {
		t.Errorf("TestShippingTypeList::Unexpected error listing shipping types %v", err)
	}

	var types []shipping.ShippingType

	err = resp.Get(&types)
	if err != nil {
		t.Errorf("TestShippingTypeList::Error getting response data %v", err)
	}
	//Amount of ShippingType within db
	expected := 0

	if len(types) == expected {
		t.Errorf("TestShippingTypeList::Received an empty list")
	}
}

func TestShippingTypeAdd(t *testing.T) {
	gt := New("http://localhost:8081/api/", apiKey)

	//New ShippingType stuff
	name := "TEST SHIPPING TYPE"
	price := 10.00

	//Creating inital request
	req := gt.ShippingTypes()
	//Creating authorised request
	aReq, err := req.Authorise("lewis", "password", apiKey)
	if err != nil {
		t.Errorf("TestShippingTypeAdd::Unexpected error creatign authorised request %v", err)
	}
	aReq.RequestData = AddShippingTypeRequest{Name: name, Price: price}

	resp, err := aReq.Add()
	if err != nil {
		t.Errorf("TestDiffShipType::Unexpected error listing shipping types %v", err)
	}

	if resp.ResponseCode != http.StatusOK {
		t.Errorf("TestShippingTypeAdd::Expected to receive status code %d but instead received %d",
			http.StatusOK, resp.ResponseCode)
	}
}

func TestShippingTypeUpdate(t *testing.T) {
	//Creating gTeller object
	gt := New("http://localhost:8081/api/", apiKey)

	//Making sure there is at least one test status
	TestShippingTypeAdd(t)

	//Creating ShippingTypeFilter to get all types with test name
	nameCondition := filter.Condition{Type: filter.ConditionEquals, Value: "TEST SHIPPING TYPE"}
	f1 := filter.ShippingTypeFilter{Name: nameCondition}

	//Getting shipping types
	req := gt.ShippingTypes()
	req.RequestData = ListShippingTypeRequest{Filter: f1}
	resp, err := req.List()
	if err != nil {
		t.Errorf("TestShippingTypeUpdate::Unexpected error listing shipping types %v", err)
	}

	var types []shipping.ShippingType

	err = resp.Get(&types)
	if err != nil {
		t.Errorf("TestShippingTypeUpdate::Error getting response data %v", err)
	}

	newName := "TEST TYPE UPDATED"

	//Updating types to new name
	for i, t := range types {
		t.Name = newName
		types[i] = t
	}

	//Updating gTeller
	req = gt.ShippingTypes()
	aReq, err := req.Authorise("lewis", "password", apiKey)
	if err != nil {
		t.Errorf("TestShippingTypeUpdate::Unexpected error while creatign authorised request %v", err)
	}

	//Setting update request data
	aReq.RequestData = UpdateShippingTypeRequest{types}

	//Updating
	resp, err = aReq.Update()
	if err != nil {
		t.Errorf("TestShippingTypeUpdate::Unexpected error while sending Update request %v", err)
	}

	if resp.ResponseCode != http.StatusOK {
		t.Errorf("TestShippingTypeUpdate::Received Non-OK response status %d", resp.ResponseCode)
	}

	//Checking if update occured
	nameCondition = filter.Condition{Type: filter.ConditionEquals, Value: newName}
	f2 := filter.ShippingTypeFilter{Name: nameCondition}

	//Getting shipping types
	req = gt.ShippingTypes()
	req.RequestData = ListShippingTypeRequest{Filter: f2}
	resp, err = req.List()
	if err != nil {
		t.Errorf("TestShippingTypeUpdate::Unexpected error listing shipping types %v", err)
	}

	var ut []shipping.ShippingType

	err = resp.Get(&ut)
	if err != nil {
		t.Errorf("TestShippingTypeUpdate::Error getting response data %v", err)
	}

	if len(ut) < len(types) {
		t.Error("TestShippingTypeUpdate::Not all items were updated")
	}
}

func TestShippingTypeRemove(t *testing.T) {
	//Creating gTeller object
	gt := New("http://localhost:8081/api/", apiKey)

	//Making sure there is something to delete
	TestShippingTypeUpdate(t)

	//Getting types to delete

	//Creating ShippingTypeFilter to get all types with test name
	nameCondition := filter.Condition{Type: filter.ConditionEquals, Value: "TEST TYPE UPDATED"}
	f1 := filter.ShippingTypeFilter{Name: nameCondition}

	//Getting shipping types
	req := gt.ShippingTypes()
	req.RequestData = ListShippingTypeRequest{Filter: f1}
	resp, err := req.List()
	if err != nil {
		t.Errorf("TestShippingTypeRemove::Unexpected error listing shipping types %v", err)
	}

	var types []shipping.ShippingType

	err = resp.Get(&types)
	if err != nil {
		t.Errorf("TestShippingTypeRemove::Error getting response data %v", err)
	}

	rIDs := make([]string, len(types))
	for i, tmp := range types {
		rIDs[i] = tmp.ID
	}

	//Creating remove request
	req = gt.ShippingTypes()
	aReq, err := req.Authorise("lewis", "password", apiKey)
	if err != nil {
		t.Errorf("TestShippingTypeRemove::Unexpected error while creatign authorised request %v", err)
	}
	//Giving request the ids to remove
	aReq.RequestData = RemoveShippingTypeRequest{IDs: rIDs}

	//Removing
	resp, err = aReq.Remove()
	if err != nil {
		t.Errorf("TestShippingTypeRemove::Unexpected error while sending Remove request %v", err)
	}

	if resp.ResponseCode != http.StatusOK {
		t.Errorf("TestShippingTypeRemove::Received Non-OK response status %d", resp.ResponseCode)
	}

	//Confirming types were removed
	req = gt.ShippingTypes()
	req.RequestData = ListShippingTypeRequest{Filter: f1}
	resp, err = req.List()
	if err != nil {
		t.Errorf("TestShippingTypeRemove::Unexpected error listing shipping types %v", err)
	}

	var cTypes []shipping.ShippingType
	err = resp.Get(&cTypes)
	if err != nil {
		t.Errorf("TestShippingTypeRemove::Error getting response data %v", err)
	}

	if len(cTypes) != 0 {
		t.Errorf("TestShippingTypeRemove::Expected there to be no types left")
	}

}
