package orders

import (
	"time"

	"github.com/auenc/gTeller/discounts"
	"github.com/auenc/gTeller/shipping"
	"github.com/auenc/gTeller/statuses"
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

//Parse parses data within ParseableOrder and returns an Order object
func (parse *ParseableOrder) Parse() Order {
	t := time.Unix(parse.TimeCreated, 0)
	dis := parse.Discount.Parse()
	return Order{ID: parse.ID, FriendlyID: parse.FriendlyID,
		ShippingDetails: parse.ShippingDetails, Items: parse.Items,
		Status: parse.Status, Notes: parse.Notes, Payed: parse.Payed,
		CustomerID: parse.CustomerID, TimeCreated: t,
		Discount: dis, Archived: parse.Archived}
}

//Parsable returns an instance of ParseableOrder derived from the Order object
func (order *Order) Parsable() (ParseableOrder, error) {
	var p ParseableOrder

	tmp := &order.Discount
	pDis, err := tmp.Parsable()
	if err != nil {
		return p, err
	}

	return ParseableOrder{ID: order.ID, FriendlyID: order.FriendlyID,
		ShippingDetails: order.ShippingDetails, Items: order.Items,
		Status: order.Status, Notes: order.Notes, Payed: order.Payed,
		CustomerID: order.CustomerID, TimeCreated: order.TimeCreated.Unix(),
		Discount: pDis, Archived: order.Archived}, nil
}

/*func (order *Order) Total(itemRepo items.ItemRepository) (float64, error) {
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
*/
