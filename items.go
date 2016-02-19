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
	Name       string
	Price      string
	DiscountID string
	Options    []string
	ImageURI   string
}
