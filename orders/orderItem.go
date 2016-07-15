package orders

import (
	"fmt"

	"github.com/auenc/gTeller/requirements"
)

type OrderItem struct {
	UUID         string
	Name         string
	ItemID       string
	Requirements []requirements.Requirement
	Price        float64
	ImageURL     string
	Quantity     int64
}

type OrderItemInfo struct {
	UUID         string
	Name         string
	Requirements []requirements.Requirement
	Price        float64
	ImageURL     string
	Quantity     int64
}

type ParseableOrderItem struct {
	UUID         string
	ItemID       string
	Requirements []requirements.ParseableRequirement
	Price        float64
	ImageURL     string
	Quantity     int64
}

func (item *OrderItem) GetPrice() float64 {
	return item.Price
}
func (item *OrderItemInfo) SetPrice(price float64) {
	item.Price = price
}

func (item *OrderItem) Parseable() (ParseableOrderItem, error) {
	var parseable ParseableOrderItem

	parseable.UUID = item.UUID
	parseable.ItemID = item.ItemID
	parseable.Price = item.Price
	parseable.ImageURL = item.ImageURL
	parseable.Quantity = item.Quantity

	parseReqs := make([]requirements.ParseableRequirement, len(item.Requirements))
	fmt.Println("OrderItem::Parseable::Looping through", len(item.Requirements),
		"requirements")
	for i, req := range item.Requirements {
		tmp, err := req.Parseable()
		if err != nil {
			return parseable, err
		}

		parseReqs[i] = tmp
	}
	fmt.Println("OrderItem::Parseable::No errors occured. Created", len(parseReqs), "parseable requirements")

	parseable.Requirements = parseReqs

	return parseable, nil
}

func (item *ParseableOrderItem) Parse() (OrderItem, error) {
	var oItem OrderItem

	oItem.UUID = item.UUID
	oItem.ItemID = item.ItemID
	oItem.Price = item.Price
	oItem.ImageURL = item.ImageURL
	oItem.Quantity = item.Quantity

	reqs := make([]requirements.Requirement, len(item.Requirements))
	for i, el := range item.Requirements {
		tmp, err := el.Parse()
		if err != nil {
			return oItem, err
		}
		reqs[i] = tmp
	}

	oItem.Requirements = reqs

	return oItem, nil
}
