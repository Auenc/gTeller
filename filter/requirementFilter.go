package filter

import "github.com/auenc/gTeller-core/requirements"

type RequirementFilter struct {
	Type      Condition
	ID        Condition
	Condition RequirementConditionFilter
	Name      Condition
	Options   UserInputFilter
	Reference Condition
}

func (filter *RequirementFilter) Valid(reqs ...requirements.Requirement) bool {
	emptyCon := Condition{}
	emptyConFilter := RequirementConditionFilter{}
	valid := true
	for _, req := range reqs {
		//Type condition
		if filter.Type != emptyCon {
			//Checking if requirementCondtion matches ConditionType filter.
			//IF it doesn't, return false.
			if !filter.Type.Valid(req.Type()) {
				valid = false
				return valid
			}
		}

		//ID condition
		if filter.ID != emptyCon {
			//Checking if requirementCondtion matches ConditionType filter.
			//IF it doesn't, return false.
			if !filter.ID.Valid(req.ID()) {
				valid = false
				return valid
			}
		}
		//Condition condition
		if filter.Condition != emptyConFilter {
			//Checking if requirementCondtion matches ConditionType filter.
			//IF it doesn't, return false.
			if !filter.Condition.Valid(req.GetCondition()) {
				valid = false
				return valid
			}
		}
		//Name condition
		if filter.Name != emptyCon {
			//Checking if requirementCondtion matches ConditionType filter.
			//IF it doesn't, return false.
			if !filter.Name.Valid(req.Name()) {
				valid = false
				return valid
			}
		}

		//Reference condition
		if filter.Reference != emptyCon {
			//Checking if requirementCondtion matches ConditionType filter.
			//IF it doesn't, return false.
			if !filter.Reference.Valid(req.Reference()) {
				valid = false
				return valid
			}
		}

		for _, op := range req.Options() {
			if !filter.Options.Valid(op) {
				valid = false
				return valid
			}
		}
	}

	return valid
}

func (filter *RequirementFilter) Filter(reqs []requirements.Requirement) []requirements.Requirement {
	var valid []requirements.Requirement

	for _, req := range reqs {
		if filter.Valid(req) {
			valid = append(valid, req)
		}
	}
	return valid
}
