package orders

import (
	"time"

	"github.com/auenc/gTeller/discounts"
	"github.com/auenc/gTeller/shipping"
	"github.com/auenc/gTeller/statuses"
)

//ParseableOrder is a data object that can be parsed into an Order object.
type ParseableOrder struct {
	ID              string
	FriendlyID      string
	ShippingDetails shipping.ShippingDetails
	Items           []OrderItem
	Status          status.Status
	Notes           string
	Payed           bool
	CustomerID      string
	TimeCreated     int64
	Discount        discounts.ParseableDiscount
	Archived        bool
}

type Order struct {
	ID              string
	FriendlyID      string
	ShippingDetails shipping.ShippingDetails
	Items           []OrderItem
	Status          status.Status
	Notes           string
	Payed           bool
	CustomerID      string
	TimeCreated     time.Time
	Discount        discounts.Discount
	Archived        bool
}

type OrderInfo struct {
	ID              string
	FriendlyID      string
	ShippingDetails shipping.ShippingDetails
	Items           []OrderItem
	Status          status.Status
	Total           float64
	Notes           string
	Payed           bool
	CustomerID      string
	TimeCreated     time.Time
	Archived        bool
}

func (order *Order) Total() (float64, error) {
	var total float64
	for _, item := range order.Items {
		tmp := item.Price

		total += tmp
	}

	total += order.ShippingDetails.Type.Price

	//Checking if a discount should be applied
	tmp := discounts.Discount{}
	if order.Discount != tmp {
		tmp, err := order.Discount.Calculate(total)
		if err != nil {
			return total, err
		}
		total = tmp
	}

	return total, nil
}

func (order *Order) Parseable() (ParseableOrder, error) {
	var p ParseableOrder
	timeCreated := order.TimeCreated.Unix()
	parseDiscount, err := order.Discount.Parsable()
	if err != nil {
		return p, err
	}
	p = ParseableOrder{ID: order.ID, FriendlyID: order.FriendlyID,
		ShippingDetails: order.ShippingDetails, Items: order.Items,
		Status: order.Status, Notes: order.Notes, Payed: order.Payed,
		CustomerID: order.CustomerID, TimeCreated: timeCreated,
		Discount: parseDiscount, Archived: order.Archived}
	return p, nil
}

func (order *ParseableOrder) Parse() Order {
	discount := order.Discount.Parse()
	created := time.Unix(order.TimeCreated, 0)

	orderA := Order{ID: order.ID, FriendlyID: order.FriendlyID,
		ShippingDetails: order.ShippingDetails, Items: order.Items, Status: order.Status,
		Notes: order.Notes, Payed: order.Payed, CustomerID: order.CustomerID,
		TimeCreated: created, Archived: order.Archived, Discount: discount}

	return orderA
}

func (order *Order) Info() (OrderInfo, error) {
	var info OrderInfo
	total, err := order.Total()
	if err != nil {
		return info, err
	}

	info = OrderInfo{order.ID, order.FriendlyID, order.ShippingDetails, order.Items, order.Status, total, order.Notes, order.Payed, order.CustomerID, order.TimeCreated, order.Archived}

	return info, nil
}
