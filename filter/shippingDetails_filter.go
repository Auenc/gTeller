package filter

import "github.com/auenc/gTeller/shipping"

type ShippingDetailsFilter struct {
	ID       Condition
	Type     ShippingTypeFilter
	Town     Condition
	Country  Condition
	Line1    Condition
	City     Condition
	PostCode Condition
	Name     Condition
	Phone    Condition
}

//Valid returns true if the given ShippingDetails matches the filter
func (filter *ShippingDetailsFilter) Valid(shippingDets shipping.ShippingDetails) bool {
	valid := true
	emptyCon := Condition{}
	emptyShip := ShippingTypeFilter{}

	if filter.ID != emptyCon {
		//Checking if ShippingDetails matches ID filter.
		//IF it doesn't, return false.
		if !filter.ID.Valid(shippingDets.ID) {
			valid = false
			return valid
		}
	}

	if filter.Type != emptyShip {
		//Checking if ShippingDetails matches Type filter.
		//IF it doesn't, return false.
		if !filter.Type.Valid(shippingDets.Type) {
			valid = false
			return valid
		}
	}

	if filter.Town != emptyCon {
		//Checking if ShippingDetails matches Town filter.
		//IF it doesn't, return false.
		if !filter.Town.Valid(shippingDets.Town) {
			valid = false
			return valid
		}
	}

	if filter.Country != emptyCon {
		//Checking if ShippingDetails matches Country filter.
		//IF it doesn't, return false.
		if !filter.Country.Valid(shippingDets.Country) {
			valid = false
			return valid
		}
	}

	if filter.Line1 != emptyCon {
		//Checking if ShippingDetails matches Line1 filter.
		//IF it doesn't, return false.
		if !filter.Line1.Valid(shippingDets.Line1) {
			valid = false
			return valid
		}
	}

	if filter.City != emptyCon {
		//Checking if ShippingDetails matches City filter.
		//IF it doesn't, return false.
		if !filter.City.Valid(shippingDets.City) {
			valid = false
			return valid
		}
	}

	if filter.PostCode != emptyCon {
		//Checking if ShippingDetails matches PostCode filter.
		//IF it doesn't, return false.
		if !filter.PostCode.Valid(shippingDets.PostCode) {
			valid = false
			return valid
		}
	}

	if filter.Name != emptyCon {
		//Checking if ShippingDetails matches Name filter.
		//IF it doesn't, return false.
		if !filter.Name.Valid(shippingDets.Name) {
			valid = false
			return valid
		}
	}

	if filter.Phone != emptyCon {
		//Checking if ShippingDetails matches Phone filter.
		//IF it doesn't, return false.
		if !filter.Phone.Valid(shippingDets.Phone) {
			valid = false
			return valid
		}
	}

	return valid
}

//Filter filters out ShippingDetails that do not meet the filter
func (filter *ShippingDetailsFilter) Filter(source []shipping.ShippingDetails) []shipping.ShippingDetails {
	var filtered []shipping.ShippingDetails

	//Lopping through ShippingDetails
	for _, ship := range source {
		if filter.Valid(ship) {
			filtered = append(filtered, ship)
		}
	}

	return filtered
}
