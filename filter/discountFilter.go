package filter

import "github.com/auenc/gTeller/discounts"

type DiscountFilter struct {
	ID          Condition
	Type        Condition
	Percent     Condition
	Numerical   Condition
	Requirement DiscountConditionFilter
	Code        Condition
}

func (filter *DiscountFilter) Valid(discount *discounts.Discount) bool {
	valid := true
	emptyCon := Condition{}
	emptyDis := DiscountConditionFilter{}

	if filter.ID != emptyCon {
		//Checking if discount matches ID filter.
		//IF it doesn't, return false.
		if !filter.ID.Valid(discount.ID) {
			valid = false
			return valid
		}
	}

	if filter.Type != emptyCon {
		//Checking if discount matches Type filter.
		//IF it doesn't, return false.
		if !filter.Type.Valid(discount.Type) {
			valid = false
			return valid
		}
	}

	if filter.Percent != emptyCon {
		//Checking if discount matches Percent filter.
		//IF it doesn't, return false.
		if !filter.Percent.Valid(discount.Percent) {
			valid = false
			return valid
		}
	}

	if filter.Numerical != emptyCon {
		//Checking if discount matches Numerical filter.
		//IF it doesn't, return false.
		if !filter.Numerical.Valid(discount.Numerical) {
			valid = false
			return valid
		}
	}

	if filter.Requirement != emptyDis {
		//Checking if discount matches Requirement filter.
		//IF it doesn't, return false.
		if !filter.Requirement.Valid(discount.Requirement) {
			valid = false
			return valid
		}
	}
	if filter.Code != emptyCon {
		//Checking if discount matches Code filter.
		//IF it doesn't, return false.
		if !filter.Code.Valid(discount.Code) {
			valid = false
			return valid
		}
	}
	return valid
}

func (filter *DiscountFilter) Filter(source []discounts.Discount) []discounts.Discount {
	filtered := make([]discounts.Discount, 0)
	//Loop through each Discount
	for _, discount := range source {
		if filter.Valid(&discount) {
			//Adding discount to filtered list
			filtered = append(filtered, discount)
		}
	}

	return filtered
}
