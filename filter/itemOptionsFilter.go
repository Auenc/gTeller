package filter

import "github.com/auenc/gTeller-core/items"

type ItemOptionFilter struct {
	ID    Condition
	Name  Condition
	Price Condition
}

type ItemOptionsFilter struct {
	ID      Condition
	Name    Condition
	Options ItemOptionFilter
}

//Valid returns true if the given categories match the filter
func (filter *ItemOptionsFilter) Valid(cats []items.OptionCategory) bool {
	valid := true
	emptyCon := Condition{}
	emptyOpt := ItemOptionFilter{}

	for _, cat := range cats {

		if filter.ID != emptyCon {
			//Checking if ItemOption matches ID filter.
			//IF it doesn't, return false.
			if !filter.ID.Valid(cat.ID) {
				valid = false
				return valid
			}
		}

		if filter.Name != emptyCon {
			//Checking if ItemOption matches Name filter.
			//IF it doesn't, return false.
			if !filter.Name.Valid(cat.Name) {
				valid = false
				return valid
			}
		}

		if filter.Options != emptyOpt {
			//Checking if ItemOption matches Options filter.
			//IF it doesn't, return false.
			if !filter.Options.Valid(cat.Options) {
				valid = false
				return valid
			}
		}
	}

	return valid
}

//Valid returns true if the given options match the filter
func (filter *ItemOptionFilter) Valid(options []items.Option) bool {
	valid := true
	emptyCon := Condition{}

	for _, option := range options {
		if filter.ID != emptyCon {
			//Checking if Option matches ID filter.
			//IF it doesn't, return false.
			if !filter.ID.Valid(option.ID) {
				valid = false
				return valid
			}
		}

		if filter.Name != emptyCon {
			//Checking if Option matches Name filter.
			//IF it doesn't, return false.
			if !filter.Name.Valid(option.Name) {
				valid = false
				return valid
			}
		}

		if filter.Price != emptyCon {
			//Checking if Option matches Price filter.
			//IF it doesn't, return false.
			if !filter.Price.Valid(option.Price) {
				valid = false
				return valid
			}
		}
	}

	return valid
}

//Filter filters out Options that do not match the filter and return the ones that do.
func (filter *ItemOptionFilter) Filter(source []items.Option) []items.Option {
	var filtered []items.Option

	//Looping through options
	for _, op := range source {
		if filter.Valid([]items.Option{op}) {
			filtered = append(filtered, op)
		}
	}

	return filtered
}

//Filter filters out OptionCategories that do not match the filter
func (filter *ItemOptionsFilter) Filter(source []items.OptionCategory) []items.OptionCategory {
	var filtered []items.OptionCategory

	//Looping through option categories
	for _, cat := range source {
		if filter.Valid([]items.OptionCategory{cat}) {
			filtered = append(filtered, cat)
		}
	}

	return filtered
}
