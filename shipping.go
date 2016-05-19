package gTeller

import (
	"github.com/auenc/gTeller/filter"
	"github.com/auenc/gTeller/shipping"
)

//ShippingType objects
type ListShippingTypeRequest struct {
	Filter filter.ShippingTypeFilter
}

type ListShippingTypeResponse struct {
	Items []shipping.ShippingType
}

type AddShippingTypeRequest struct {
	Name  string
	Price float64
}

type UpdateShippingTypeRequest struct {
	ShippingType []shipping.ShippingType
}

type RemoveShippingTypeRequest struct {
	IDs []string
}

//ShippingDetails objects
type ListShippingDetailsRequest struct {
	Filter filter.ShippingDetailsFilter
}

type ListShippingDetailsResponse struct {
	Details []shipping.ShippingDetails
}

type AddShippingDetailsRequest struct {
	TypeID   string
	Town     string
	Country  string
	Line1    string
	City     string
	PostCode string
	Name     string
	Phone    string
	Tracking string
}

type UpdateShippingDetailsRequest struct {
	ShippingDetails []shipping.ShippingDetails
}

type RemoveShippingDetailsRequest struct {
	IDs []string
}
