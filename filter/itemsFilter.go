package filter

import (
	"fmt"

	"github.com/auenc/gTeller-core/items"
)

type ItemsFilter struct {
	ID       Condition
	Name     Condition
	Price    Condition
	Discount DiscountFilter
	Options  ItemOptionsFilter
	ImageURI Condition
}

//Filter is a method that filters a given slice of items given the ItemsFilters Conditions
func (filter *ItemsFilter) Filter(source []items.Item) []items.Item {
	filtered := make([]items.Item, 0)
	//Looping through original data
	for _, item := range source {
		if filter.Valid(item) {
			fmt.Println("Item Filter pass")
			filtered = append(filtered, item)
		} else {
			fmt.Println("Filter fail")
		}
	}
	fmt.Println("Returning with ", len(filtered))
	return filtered
}

//Valid returns whether a given item meets the filter
func (filter *ItemsFilter) Valid(item items.Item) bool {
	valid := true
	emptyCon := Condition{}
	emptyDis := DiscountFilter{}
	emptyOpt := ItemOptionsFilter{}
	fmt.Println("Pass 0")
	//Checking if there is a filter for Status ID
	if filter.ID != emptyCon {
		//Checking if Item matches ID filter.
		//IF it doesn't, return false.
		if !filter.ID.Valid(item.ID) {
			valid = false
			return valid
		}
	}
	fmt.Println("Pass 1")
	//Checking if there is a filter for Status ID
	if filter.Name != emptyCon {
		//Checking if Item matches Name filter.
		//IF it doesn't, return false.
		if !filter.Name.Valid(item.NameRaw) {
			valid = false
			return valid
		}
	}
	fmt.Println("Pass 2")
	//Checking if there is a filter for Status Price
	if filter.Price != emptyCon {
		//Checking if Item matches Price filter.
		//IF it doesn't, return false.
		if !filter.Price.Valid(item.Price) {
			valid = false
			return valid
		}
	}
	fmt.Println("Pass 3")
	//Checking if there is a filter for Status Discount
	if filter.Discount != emptyDis {
		//Checking if Item matches Discount filter.
		//IF it doesn't, return false.
		if !filter.Discount.Valid(item.DicountRaw) {
			valid = false
			return valid
		}
	}
	fmt.Println("Pass 4")
	//Checking if there is a filter for Status Options
	if filter.Options != emptyOpt {
		//Checking if Item matches Options filter.
		//IF it doesn't, return false.
		if !filter.Options.Valid(item.Options) {
			valid = false
			return valid
		}
	}
	fmt.Println("Pass 5")
	//Checking if there is a filter for Status ImageURI
	if filter.ImageURI != emptyCon {
		//Checking if Item matches ImageURI filter.
		//IF it doesn't, return false.
		if !filter.ImageURI.Valid(item.ImageURI) {
			valid = false
			return valid
		}
	}
	fmt.Println("Pass 6")
	return valid
}
