package filter

import (
	"github.com/auenc/gTeller/discounts"
)

type DiscountConditionFilter struct {
	ConditionType Condition
}

func (filter *DiscountConditionFilter) Valid(con discounts.Condition) bool {
	valid := true
	emptyCon := Condition{}

	if filter.ConditionType != emptyCon {
		//Checking if discountCondtion matches ConditionType filter.
		//IF it doesn't, return false.
		if !filter.ConditionType.Valid(con.Type()) {
			valid = false
			return valid
		}
	}

	return valid
}
