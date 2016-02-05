package filter

import (
	"fmt"

	"github.com/auenc/gTeller-core/orders"
)

type OrderFilter struct {
	ID              Condition
	FriendlyID      Condition
	ShippingDetails ShippingDetailsFilter
	Items           OrderItemFilter
	Status          StatusFilter
	Notes           Condition
	Payed           Condition
	CustomerID      Condition
	TimeCreated     Condition
	Discount        DiscountFilter
	Archived        Condition
}

func (filter *OrderFilter) Filter(source []orders.Order) []orders.Order {
	var filtered []orders.Order
	emptyCon := Condition{}
	emptyShip := ShippingDetailsFilter{}
	emptyItems := OrderItemFilter{}
	emptyStat := StatusFilter{}
	emptyDis := DiscountFilter{}

	//Looping through original data
	for _, order := range source {
		if filter.ID != emptyCon {
			//Checking if order matches ID filter.
			//IF it doesn't, skip to next item.
			if !filter.ID.Valid(order.ID) {
				continue
			}
			fmt.Println("Filter pass")
		}

		if filter.FriendlyID != emptyCon {
			//Checking if order matches FriendlyID filter.
			//IF it doesn't, skip to next item.
			if !filter.FriendlyID.Valid(order.FriendlyID) {
				continue
			}
		}

		if filter.ShippingDetails != emptyShip {
			//Checking if order matches ShippingDetails filter.
			//IF it doesn't, skip to next item.
			if !filter.ShippingDetails.Valid(order.ShippingDetails) {
				continue
			}
		}

		if filter.Items != emptyItems {
			//Checking if order matches Items filter.
			//IF it doesn't, skip to next item.
			if !filter.Items.Valid(order.Items) {
				continue
			}
		}

		if filter.Status != emptyStat {
			//Checking if order matches Status filter.
			//IF it doesn't, skip to next item.
			if !filter.Status.Valid(order.Status) {
				continue
			}
		}

		if filter.Notes != emptyCon {
			//Checking if order matches Notes filter.
			//IF it doesn't, skip to next item.
			if !filter.Notes.Valid(order.Notes) {
				continue
			}
		}

		if filter.Payed != emptyCon {
			//Checking if order matches Payed filter.
			//IF it doesn't, skip to next item.
			if !filter.Payed.Valid(order.Payed) {
				continue
			}
		}

		if filter.CustomerID != emptyCon {
			//Checking if order matches CustomerID filter.
			//IF it doesn't, skip to next item.
			if !filter.CustomerID.Valid(order.CustomerID) {
				continue
			}
		}

		if filter.TimeCreated != emptyCon {
			//Checking if order matches TimeCreated filter.
			//IF it doesn't, skip to next item.
			if !filter.TimeCreated.Valid(order.TimeCreated) {
				continue
			}
		}

		if filter.Discount != emptyDis {
			//Checking if order matches Discount filter.
			//IF it doesn't, skip to next item.
			if !filter.Discount.Valid(&order.Discount) {
				continue
			}
		}

		if filter.Archived != emptyCon {
			//Checking if order matches Archived filter.
			//IF it doesn't, skip to next item.
			if !filter.Archived.Valid(order.Archived) {
				continue
			}
		}

		filtered = append(filtered, order)
	}

	return filtered
}
