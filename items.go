package gTeller

import (
	"github.com/auenc/gTeller/filter"
	"github.com/auenc/gTeller/items"
)

type ListItemsRequest struct {
	Filter filter.ItemsFilter
}

type ListItemsResponse struct {
	Items []items.Item
}

type AddItemRequest struct {
	Name         string
	Price        string
	DiscountID   string
	Requirements []string
	ImageURI     string
	Hidden       bool
}

//RemoveItemsRequest represents the data that can be received by the RemoveItems api call
//TODO Describe RemoveItemsRequest more
type RemoveItemsRequest struct {
	IDs []string
}

//UpdateItemsRequest represents the data that update item accepts.
type UpdateItemsRequest struct {
	Items []items.Item
}
