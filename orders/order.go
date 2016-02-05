package orders

import (
	"time"

	"github.com/auenc/gTeller-core/discounts"
	"github.com/auenc/gTeller-core/items"
	"github.com/auenc/gTeller-core/shipping"
	"github.com/auenc/gTeller-core/statuses"
)

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
	Items           []OrderItemInfo
	Status          status.Status
	Total           float64
	Notes           string
	Payed           bool
	CustomerID      string
	TimeCreated     time.Time
	Archived        bool
}

func (order *Order) Total(itemRepo items.ItemRepository) (float64, error) {
	var total float64
	for _, item := range order.Items {
		tmp, err := item.Price(itemRepo)
		if err != nil {
			return total, err
		}
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

func (order *Order) Info(itemRepo items.ItemRepository) (OrderInfo, error) {
	var info OrderInfo
	total, err := order.Total(itemRepo)
	if err != nil {
		return info, err
	}

	itemInfo := make([]OrderItemInfo, len(order.Items))
	for i, item := range order.Items {
		tmp, err := item.Info(itemRepo)
		if err != nil {
			return info, err
		}
		itemInfo[i] = tmp
	}

	info = OrderInfo{order.ID, order.FriendlyID, order.ShippingDetails, itemInfo, order.Status, total, order.Notes, order.Payed, order.CustomerID, order.TimeCreated, order.Archived}

	return info, nil
}
