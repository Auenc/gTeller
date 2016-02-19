package gTeller

import (
	"github.com/auenc/gTeller/filter"
	"github.com/auenc/gTeller/orders"
	"github.com/auenc/gTeller/shipping"
)

type ListOrderRequest struct {
	Filter filter.OrderFilter
}

type ListOrderResponse struct {
	Items []orders.Order
}

type AddOrderRequest struct {
	ShippingDetails shipping.ParseableShippingDetails
	Items           []orders.OrderItem
	StatusID        string
	Notes           string
	Payed           bool
	CustomerID      string
	TimeCreated     int64
	DiscountID      string
	Archived        bool
}

type UpdateOrderRequest struct {
	Orders []orders.Order
}

type RemoveOrderRequest struct {
	IDs []string
}
