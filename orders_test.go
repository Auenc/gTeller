package gTeller

import (
	"net/http"
	"testing"
	"time"

	"github.com/auenc/gTeller/filter"
	"github.com/auenc/gTeller/orders"
	"github.com/auenc/gTeller/shipping"
)

func TestOrderList(t *testing.T) {
	gt := New("http://localhost:8081/api/", "nokey")

	req := gt.Orders()
	req.RequestData = ListOrderRequest{}
	aReq, err := req.Authorise("lewis", "password")
	if err != nil {
		t.Errorf("TestOrderList::Unexpected error authorising request %v", err)
	}
	resp, err := aReq.List()
	if err != nil {
		t.Errorf("TestOrderList::Unexpected error listing shipping types %v", err)
	}

	var orderL []orders.ParseableOrder

	err = resp.Get(&orderL)
	if err != nil {
		t.Errorf("TestOrderList::Error getting response data %v", err)
	}
	//Amount of Order within db
	expected := 0

	if len(orderL) == expected {
		t.Errorf("TestOrderList::Received an empty list")
	}
}

func TestOrderAdd(t *testing.T) {
	gt := New("http://localhost:8081/api/", "nokey")

	//New Order stuff
	name := "Bob"
	typeID := "2dfb9842-92e0-4d94-a0f5-f40259a4f080"
	town := "Bridgend"
	country := "apple"
	line1 := "Pencoed"
	city := "Amsterdamn"
	postCode := "SA10QT"
	phone := "07444444444"
	tracking := "bug"
	sd := shipping.ParseableShippingDetails{TypeID: typeID, Town: town,
		Country: country, Line1: line1, City: city, PostCode: postCode, Name: name,
		Phone: phone, Tracking: tracking}

	items := make([]orders.OrderItem, 1)
	items[0] = orders.OrderItem{ItemID: "355c7881-7c32-41e6-bbec-ea0f615bf9c5",
		ImageURL: "a.jpg", Quantity: 1}
	statID := "e3feec3f-9d27-425b-a535-eaea8c7d5d0c"
	notes := "a note"
	payed := false
	customerID := "1"
	timeCreated := time.Now().Unix()
	discountID := "8f5039f2-ec92-41a0-a8f9-3c12132873f2"
	archived := false
	reqData := AddOrderRequest{sd, items, statID, notes, payed, customerID,
		timeCreated, discountID, archived}

	//Creating inital request
	req := gt.Orders()
	//Creating authorised request
	aReq, err := req.Authorise("lewis", "password")
	if err != nil {
		t.Errorf("TestOrderAdd::Unexpected error creatign authorised request %v", err)
	}
	aReq.RequestData = reqData

	resp, err := aReq.Add()
	if err != nil {
		t.Errorf("TestDiffShipType::Unexpected error listing shipping types %v", err)
	}

	if resp.ResponseCode != http.StatusOK {
		t.Errorf("TestOrderAdd::Expected to receive status code %d but instead received %d",
			http.StatusOK, resp.ResponseCode)
	}
}

func TestOrderUpdate(t *testing.T) {
	//Creating gTeller object
	gt := New("http://localhost:8081/api/", "nokey")

	//Making sure there is at least one test status
	TestOrderAdd(t)

	//Creating OrderFilter to get all types with test name
	nameCondition := filter.Condition{Type: filter.ConditionEquals, Value: "a note"}
	f1 := filter.OrderFilter{Notes: nameCondition}

	//Getting shipping types
	req := gt.Orders()
	req.RequestData = ListOrderRequest{Filter: f1}
	aReq, err := req.Authorise("lewis", "password")
	if err != nil {
		t.Errorf("TestOrderUpdate::Unexpected error authorising request %v", err)
	}
	resp, err := aReq.List()
	if err != nil {
		t.Errorf("TestOrderUpdate::Unexpected error listing shipping types %v", err)
	}

	var types []orders.ParseableOrder

	err = resp.Get(&types)
	if err != nil {
		t.Errorf("TestOrderUpdate::Error getting response data %v", err)
	}

	newName := "TEST ORDER UPDATED"

	uO := make([]orders.Order, len(types))
	//Updating types to new name
	for i, t := range types {
		tmp := t.Parse()

		tmp.Notes = newName
		uO[i] = tmp
	}

	//Updating gTeller
	req = gt.Orders()
	aReq, err = req.Authorise("lewis", "password")
	if err != nil {
		t.Errorf("TestOrderUpdate::Unexpected error while creatign authorised request %v", err)
	}

	//Setting update request data
	aReq.RequestData = UpdateOrderRequest{uO}

	//Updating
	resp, err = aReq.Update()
	if err != nil {
		t.Errorf("TestOrderUpdate::Unexpected error while sending Update request %v", err)
	}

	if resp.ResponseCode != http.StatusOK {
		t.Errorf("TestOrderUpdate::Received Non-OK response status %d", resp.ResponseCode)
	}

	//Checking if update occured
	nameCondition = filter.Condition{Type: filter.ConditionEquals, Value: newName}
	f2 := filter.OrderFilter{Notes: nameCondition}

	//Getting shipping types
	req = gt.Orders()
	req.RequestData = ListOrderRequest{Filter: f2}
	aReq, err = req.Authorise("lewis", "password")
	if err != nil {
		t.Errorf("TestOrderUpdate::Unexpected error authorising request %v", err)
	}
	resp, err = aReq.List()
	if err != nil {
		t.Errorf("TestOrderUpdate::Unexpected error listing shipping types %v", err)
	}

	var ut []orders.ParseableOrder

	err = resp.Get(&ut)
	if err != nil {
		t.Errorf("TestOrderUpdate::Error getting response data %v", err)
	}

	if len(ut) < len(types) {
		t.Error("TestOrderUpdate::Not all items were updated")
	}
}

func TestOrderRemove(t *testing.T) {
	//Creating gTeller object
	gt := New("http://localhost:8081/api/", "nokey")

	//Making sure there is something to delete
	TestOrderUpdate(t)

	//Getting types to delete

	//Creating OrderFilter to get all types with test name
	nameCondition := filter.Condition{Type: filter.ConditionEquals, Value: "TEST ORDER UPDATED"}
	f1 := filter.OrderFilter{Notes: nameCondition}

	//Getting shipping types
	req := gt.Orders()
	req.RequestData = ListOrderRequest{Filter: f1}
	aReq, err := req.Authorise("lewis", "password")
	if err != nil {
		t.Errorf("TestOrderRemove::Unexpected error authorising request %v", err)
	}
	resp, err := aReq.List()
	if err != nil {
		t.Errorf("TestOrderRemove::Unexpected error listing shipping types %v", err)
	}

	var types []orders.ParseableOrder

	err = resp.Get(&types)
	if err != nil {
		t.Errorf("TestOrderRemove::Error getting response data %v", err)
	}

	rIDs := make([]string, len(types))
	for i, tmp := range types {
		rIDs[i] = tmp.ID
	}

	//Creating remove request
	req = gt.Orders()
	aReq, err = req.Authorise("lewis", "password")
	if err != nil {
		t.Errorf("TestOrderRemove::Unexpected error while creatign authorised request %v", err)
	}
	//Giving request the ids to remove
	aReq.RequestData = RemoveOrderRequest{IDs: rIDs}

	//Removing
	resp, err = aReq.Remove()
	if err != nil {
		t.Errorf("TestOrderRemove::Unexpected error while sending Remove request %v", err)
	}

	if resp.ResponseCode != http.StatusOK {
		t.Errorf("TestOrderRemove::Received Non-OK response status %d", resp.ResponseCode)
	}

	//Confirming types were removed
	req = gt.Orders()
	req.RequestData = ListOrderRequest{Filter: f1}
	aReq, err = req.Authorise("lewis", "password")
	if err != nil {
		t.Errorf("TestOrderRemove::Unexpected error authorising request %v", err)
	}
	resp, err = aReq.List()
	if err != nil {
		t.Errorf("TestOrderRemove::Unexpected error listing shipping types %v", err)
	}

	var cTypes []orders.ParseableOrder
	err = resp.Get(&cTypes)
	if err != nil {
		t.Errorf("TestOrderRemove::Error getting response data %v", err)
	}

	if len(cTypes) != 0 {
		t.Errorf("TestOrderRemove::Expected there to be no types left")
	}

}
