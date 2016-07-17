package filter

import "github.com/auenc/gTeller/shipping"

type ShippingTypeFilter struct {
	ID    Condition
	Name  Condition
	Price Condition
}

//Valid returns true if the given ShippingType matches the filter
func (filter *ShippingTypeFilter) Valid(shipType shipping.ShippingType) bool {
	valid := true
	emptyCon := Condition{}

	if filter.ID != emptyCon {
		//Checking if ShippingType matches ID filter.
		//IF it doesn't, return false.
		if !filter.ID.Valid(shipType.ID) {
			valid = false
			return valid
		}
	}

	if filter.Name != emptyCon {
		//Checking if ShippingType matches Name filter.
		//IF it doesn't, return false.
		if !filter.Name.Valid(shipType.Name) {
			valid = false
			return valid
		}
	}

	if filter.Price != emptyCon {
		//Checking if ShippingType matches Price filter.
		//IF it doesn't, return false.
		if !filter.Price.Valid(shipType.Price) {
			valid = false
			return valid
		}
	}

	return valid
}

//Filter filters out ShippingTypes that do not match the filter from a given slice
func (filter ShippingTypeFilter) Filter(source []shipping.ShippingType) []shipping.ShippingType {
	filtered := make([]shipping.ShippingType, 0)

	//Lopping through shipping types
	for _, ship := range source {
		if filter.Valid(ship) {
			filtered = append(filtered, ship)
		}
	}

	return filtered
}
