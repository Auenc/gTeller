package gTeller

import (
	"github.com/auenc/gTeller/discounts"
	"github.com/auenc/gTeller/filter"
)

//GetDiscountRequest represents the data that can be received by the GetItems api call
type ListDiscountRequest struct {
	Filter filter.DiscountFilter
}

//AddDiscountRequest represents the data that can be received by the AddShipping api call
//TODO Describe AddItemRequest more
type AddDiscountRequest struct {
	Type            int
	Percent         int
	Numerical       float64
	RequirementType string
	RequirementData string
	Code            string
}

//RemoveDiscountRequest represents the data that can be received by the RemoveItems api call
type RemoveDiscountRequest struct {
	IDs []string
}

//UpdateDiscountRequest represents the data that update item accepts.
type UpdateDiscountRequest struct {
	Discount []discounts.ParseableDiscount
}
