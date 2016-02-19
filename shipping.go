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
