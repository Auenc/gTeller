package filter

import (
	"encoding/json"
	"log"
	"os"
	"testing"

	"github.com/auenc/gTeller-core/config"
	"github.com/auenc/gTeller-core/databases"
	"github.com/auenc/gTeller-core/requirements"
)

var (
	db databases.Database
)

func loadSettings() (config.Config, error) {
	var settings config.Config
	//Reading the file
	configFile, err := os.Open("/etc/gteller/" + "config.json")
	if err != nil {
		log.Println(err)
		return settings, err
	}
	//Parsing the json
	jsonParser := json.NewDecoder(configFile)
	if err = jsonParser.Decode(&settings); err != nil {
		log.Println("Error parsing settings", err.Error())
		return settings, err
	}

	//log.Println("Settings = ", settings)

	return settings, nil
}

func init() {
	config, _ := loadSettings()
	config.Testing = true
	db, _ = databases.New(config)
}

func TestBasicFilter(t *testing.T) {
	//Note: This is an ID from an order within the testing DB.
	//If that record does not exist, this test will fail
	desiredID := "6c323cf3-7e1a-4ff9-aaff-ebfcb05963a3"

	//Creating Filter
	idFilter := Condition{ConditionEquals, desiredID}
	filter := OrderFilter{ID: idFilter}

	orders, err := db.OrderRepo().Orders()
	if err != nil {
		t.Errorf("TestBasicFilter::Unexpected error while getting orders from repo %v", err)
		return
	}

	filtered := filter.Filter(orders)
	if len(filtered) != 1 {
		t.Errorf("TestBasicFilter::Expected filtered list length to be 1 but instead it was %d", len(filtered))
		return
	}

	if filtered[0].ID != desiredID {
		t.Errorf("TestBasicFilter::Expected order.ID to be %s but instead it was %s", desiredID, filtered[0].ID)
	}

}

func TestBasicItemFilter(t *testing.T) {
	//Note: This is an ID from an item within the testing DB.
	//If that record does not exist, this test will fail
	desiredID := "355c7881-7c32-41e6-bbec-ea0f615bf9c5"

	//Creating Filter
	emptyCon := Condition{}
	idFilter := Condition{ConditionEquals, desiredID}
	filter := ItemsFilter{idFilter, emptyCon, emptyCon,
		DiscountFilter{}, RequirementFilter{}, emptyCon}

	items, err := db.ItemRepo().Items()
	if err != nil {
		t.Errorf("TestBasicItemFilter::Unexpected error while getting orders from repo %v", err)
		return
	}

	filtered := filter.Filter(items)
	if len(filtered) != 1 {
		t.Errorf("TestBasicItemFilter::Expected filtered list length to be 1 but instead it was %d", len(filtered))
		return
	}

	if filtered[0].ID != desiredID {
		t.Errorf("TestBasicItemFilter::Expected order.ID to be %s but instead it was %s", desiredID, filtered[0].ID)
	}

}

func TestBasicDiscountFilter(t *testing.T) {
	//Note: This is an ID from an item within the testing DB.
	//If that record does not exist, this test will fail
	desiredID := "8f5039f2-ec92-41a0-a8f9-3c12132873f2"

	//Creating Filter
	emptyCon := Condition{}
	idFilter := Condition{ConditionEquals, desiredID}
	filter := DiscountFilter{idFilter, emptyCon, emptyCon,
		emptyCon, DiscountConditionFilter{}, emptyCon}

	discounts, err := db.DiscountRepo().Discounts()
	if err != nil {
		t.Errorf("TestBasicDiscountFilter::Unexpected error while getting orders from repo %v", err)
		return
	}

	filtered := filter.Filter(discounts)
	if len(filtered) != 1 {
		t.Errorf("TestBasicDiscountFilter::Expected filtered list length to be 1 but instead it was %d", len(filtered))
		return
	}

	if filtered[0].ID != desiredID {
		t.Errorf("TestBasicDiscountFilter::Expected order.ID to be %s but instead it was %s", desiredID, filtered[0].ID)
	}

}

func TestBasicStatusFilter(t *testing.T) {
	//Note: This is an ID from an item within the testing DB.
	//If that record does not exist, this test will fail
	desiredID := "e3feec3f-9d27-425b-a535-eaea8c7d5d0c"

	//Creating Filter
	emptyCon := Condition{}
	idFilter := Condition{ConditionEquals, desiredID}
	filter := StatusFilter{idFilter, emptyCon, EmailTemplateFilter{}}

	stats, err := db.StatusRepo().Statuses()
	if err != nil {
		t.Errorf("TestBasicStatusFilter::Unexpected error while getting orders from repo %v", err)
		return
	}

	filtered := filter.Filter(stats)
	if len(filtered) != 1 {
		t.Errorf("TestBasicStatusFilter::Expected filtered list length to be 1 but instead it was %d", len(filtered))
		return
	}

	if filtered[0].ID != desiredID {
		t.Errorf("TestBasicStatusFilter::Expected order.ID to be %s but instead it was %s", desiredID, filtered[0].ID)
	}

}

func TestBasicShippingTypeFilter(t *testing.T) {
	//Note: This is an ID from an item within the testing DB.
	//If that record does not exist, this test will fail
	desiredID := "2dfb9842-92e0-4d94-a0f5-f40259a4f080"

	//Creating Filter
	emptyCon := Condition{}
	idFilter := Condition{ConditionEquals, desiredID}
	filter := ShippingTypeFilter{idFilter, emptyCon, emptyCon}

	types, err := db.ShippingRepo().Types()
	if err != nil {
		t.Errorf("TestBasicShippingtypeFilter::Unexpected error while getting orders from repo %v", err)
		return
	}

	filtered := filter.Filter(types)
	if len(filtered) != 1 {
		t.Errorf("TestBasicShippingtypeFilter::Expected filtered list length to be 1 but instead it was %d", len(filtered))
		return
	}

	if filtered[0].ID != desiredID {
		t.Errorf("TestBasicShippingtypeFilter::Expected order.ID to be %s but instead it was %s", desiredID, filtered[0].ID)
	}

}

func TestBasicShippingDetailsFilter(t *testing.T) {
	//Note: This is an ID from an item within the testing DB.
	//If that record does not exist, this test will fail
	desiredID := "2c8b0901-713d-4a1f-b95b-95bbf1ebbb75"

	//Creating Filter
	emptyCon := Condition{}
	idFilter := Condition{ConditionEquals, desiredID}
	filter := ShippingDetailsFilter{idFilter, ShippingTypeFilter{}, emptyCon,
		emptyCon, emptyCon, emptyCon, emptyCon, emptyCon, emptyCon}

	dets, err := db.ShippingRepo().Details()
	if err != nil {
		t.Errorf("TestBasicShippingDetailsFilter::Unexpected error while getting details from repo %v",
			err)
		return
	}

	filtered := filter.Filter(dets)
	if len(filtered) != 1 {
		t.Errorf("TestBasicShippingDetailsFilter::Expected filtered list length to be 1 but instead it was %d",
			len(filtered))
		return
	}

	if filtered[0].ID != desiredID {
		t.Errorf("TestBasicShippingDetailsFilter::Expected order.ID to be %s but instead it was %s",
			desiredID, filtered[0].ID)
	}

}

func TestBasicOrderItemFilter(t *testing.T) {
	//Note: This is an ID from an item within the testing DB.
	//If that record does not exist, this test will fail
	orderID := "b1b85295-9cbf-48a0-9f1e-cf8f99f86fe3"
	desiredID := "1eb5c8e2-83fb-4d77-9d63-dca12a2b644d"

	//Creating Filter
	emptyCon := Condition{}
	idFilter := Condition{ConditionEquals, desiredID}
	filter := OrderItemFilter{idFilter, emptyCon, emptyCon, emptyCon}

	dets, err := db.OrderRepo().OrderItems(orderID)
	if err != nil {
		t.Errorf("TestBasicOrderItemFilter::Unexpected error while getting orders from repo %v",
			err)
		return
	}

	filtered := filter.Filter(dets)
	if len(filtered) != 1 {
		t.Errorf("TestBasicOrderItemFilter::Expected filtered list length to be 1 but instead it was %d",
			len(filtered))
		return
	}

	if filtered[0].UUID != desiredID {
		t.Errorf("TestBasicOrderItemFilter::Expected order.ID to be %s but instead it was %s",
			desiredID, filtered[0].UUID)
	}

}

func TestBasicItemOptionsFilter(t *testing.T) {
	dets, err := db.ItemRepo().Options()
	if err != nil {
		t.Errorf("TestBasicItemOptionsFilter::Unexpected error while getting orders from repo %v",
			err)
		return
	}

	desiredID := dets[0].Options[0].ID

	//Creating Filter
	emptyCon := Condition{}
	idFilter := Condition{ConditionEquals, desiredID}
	filter := ItemOptionFilter{idFilter, emptyCon, emptyCon}

	filtered := filter.Filter(dets[0].Options)
	if len(filtered) != 1 {
		t.Errorf("TestBasicItemOptionsFilter::Expected filtered list length to be 1 but instead it was %d",
			len(filtered))
		return
	}

	if filtered[0].ID != desiredID {
		t.Errorf("TestBasicItemOptionsFilter::Expected order.ID to be %s but instead it was %s",
			desiredID, filtered[0].ID)
	}

}

func TestBasicOptionCategoryFilter(t *testing.T) {
	dets, err := db.ItemRepo().Options()
	if err != nil {
		t.Errorf("TestBasicOptionCategoryFilter::Unexpected error while getting orders from repo %v",
			err)
		return
	}

	desiredID := dets[0].ID

	//Creating Filter
	emptyCon := Condition{}
	idFilter := Condition{ConditionEquals, desiredID}
	filter := ItemOptionsFilter{idFilter, emptyCon, ItemOptionFilter{}}

	filtered := filter.Filter(dets)
	if len(filtered) != 1 {
		t.Errorf("TestBasicOptionCategoryFilter::Expected filtered list length to be 1 but instead it was %d",
			len(filtered))
		return
	}

	if filtered[0].ID != desiredID {
		t.Errorf("TestBasicOptionCategoryFilter::Expected order.ID to be %s but instead it was %s",
			desiredID, filtered[0].ID)
	}
}

func TestBasicEmailTemplateFilter(t *testing.T) {
	dets, err := db.EmailRepo().Templates()
	if err != nil {
		t.Errorf("TestBasicEmailTemplateFilter::Unexpected error while getting orders from repo %v",
			err)
		return
	}

	desiredID := dets[0].ID

	//Creating Filter
	emptyCon := Condition{}
	idFilter := Condition{ConditionEquals, desiredID}
	filter := EmailTemplateFilter{idFilter, emptyCon, emptyCon, emptyCon, emptyCon, emptyCon}

	filtered := filter.Filter(dets)
	if len(filtered) != 1 {
		t.Errorf("TestBasicEmailTemplateFilter::Expected filtered list length to be 1 but instead it was %d",
			len(filtered))
		return
	}

	if filtered[0].ID != desiredID {
		t.Errorf("TestBasicEmailTemplateFilter::Expected order.ID to be %s but instead it was %s",
			desiredID, filtered[0].ID)
	}
}

func TestBasicRequirementConditionFilter(t *testing.T) {
	var cons []requirements.Condition
	equalCon, _ := requirements.NewConditionEqual("bob", "apple")
	fake1, _ := requirements.NewConditionRegex("bob", "apple")
	fake2, _ := requirements.NewConditionRegex("bob", "apple")
	fake3, _ := requirements.NewConditionRegex("bob", "apple")

	cons = append(cons, &fake1)
	cons = append(cons, &fake2)
	cons = append(cons, &fake3)
	cons = append(cons, &equalCon)

	typeFilter := Condition{ConditionEquals, requirements.ConditionEqual}
	filter := RequirementConditionFilter{Type: typeFilter}

	filtered := filter.Filter(cons)

	if len(filtered) != 1 {
		t.Errorf("Expected to only receive 1 condition but instead received %d", len(filtered))
	}
}

func TestBasicUserInputFilter(t *testing.T) {
	var inputs []requirements.UserInput
	textCon := requirements.NewInputText("apple", "bob")
	fake1 := requirements.NewInputText("apple", "jeff")
	fake2 := requirements.NewInputText("apple", "Kat")
	fake3 := requirements.NewInputText("apple", "Sioned")

	inputs = append(inputs, &textCon)
	inputs = append(inputs, &fake1)
	inputs = append(inputs, &fake2)
	inputs = append(inputs, &fake3)

	forCondition := Condition{ConditionEquals, "bob"}
	filter := UserInputFilter{For: forCondition}

	filtered := filter.Filter(inputs)

	if len(filtered) != 1 {
		t.Errorf("Expected to only receive 1 condition but instead received %d", len(filtered))
	}
}
