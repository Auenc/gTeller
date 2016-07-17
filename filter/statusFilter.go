package filter

import (
	statuses "github.com/auenc/gTeller/status"
)

type StatusFilter struct {
	ID            Condition
	Name          Condition
	EmailTemplate EmailTemplateFilter
}

//Valid returns true if the given status matches the filter
func (filter *StatusFilter) Valid(stat statuses.Status) bool {
	valid := true
	emptyCon := Condition{}
	emptyEmail := EmailTemplateFilter{}

	if filter.ID != emptyCon {
		//Checking if Status matches ID filter.
		//IF it doesn't, return false.
		if !filter.ID.Valid(stat.ID) {
			valid = false
			return valid
		}
	}

	if filter.Name != emptyCon {
		//Checking if Status matches Name filter.
		//IF it doesn't, return false.
		if !filter.Name.Valid(stat.Name) {
			valid = false
			return valid
		}
	}

	if filter.EmailTemplate != emptyEmail {
		//Checking if Status matches EmailTemplate filter.
		//IF it doesn't, return false.
		if !filter.EmailTemplate.Valid(stat.EmailTemplate) {
			valid = false
			return valid
		}
	}

	return valid
}

//Filter loops through the given Status list and returns a list of statuses that match the filter
func (filter *StatusFilter) Filter(source []statuses.Status) []statuses.Status {
	filtered := make([]statuses.Status, 0)

	//Looping through statuses
	for _, stat := range source {
		if filter.Valid(stat) {
			filtered = append(filtered, stat)
		}
	}

	return filtered
}
