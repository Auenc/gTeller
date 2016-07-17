package filter

import "github.com/auenc/gTeller/orders"

type OrderItemFilter struct {
	UUID     Condition
	ItemID   Condition
	ImageURL Condition
	Quantity Condition
}

//Valid returns true if the given OrderItem matches the filter
func (filter *OrderItemFilter) Valid(items []orders.OrderItem) bool {
	valid := true
	emptyCon := Condition{}

	for _, item := range items {
		if filter.UUID != emptyCon {
			//Checking if OrderItem matches UUID filter.
			//IF it doesn't, return false.
			if !filter.UUID.Valid(item.UUID) {
				valid = false
				return valid
			}
		}

		if filter.ItemID != emptyCon {
			//Checking if OrderItem matches ItemID filter.
			//IF it doesn't, return false.
			if !filter.ItemID.Valid(item.ItemID) {
				valid = false
				return valid
			}
		}

		if filter.ImageURL != emptyCon {
			//Checking if OrderItem matches ImageURL filter.
			//IF it doesn't, return false.
			if !filter.ImageURL.Valid(item.ImageURL) {
				valid = false
				return valid
			}
		}

		if filter.Quantity != emptyCon {
			//Checking if OrderItem matches Quantity filter.
			//IF it doesn't, return false.
			if !filter.Quantity.Valid(item.Quantity) {
				valid = false
				return valid
			}
		}
	}

	return valid
}

//Filter filters out any OrderItems that do not match the filter
func (filter *OrderItemFilter) Filter(source []orders.OrderItem) []orders.OrderItem {
	var filtered []orders.OrderItem

	//Loop through each orderItem
	for _, item := range source {
		if filter.Valid([]orders.OrderItem{item}) {
			filtered = append(filtered, item)
		}
	}
	return filtered
}
