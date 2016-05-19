package filter

import "github.com/auenc/gTeller-core/requirements"

type RequirementConditionFilter struct {
	Type      Condition
	ID        Condition
	Condition Condition
	Name      Condition
	Save      Condition
}

type emptyReqCon interface{}

func (filter *RequirementConditionFilter) Valid(con requirements.Condition) bool {
	emptyCon := Condition{}
	var emptyReqCon requirements.Condition
	valid := true
	if con != emptyReqCon {
		//Type condition
		if filter.Type != emptyCon {
			//Checking if requirementCondtion matches ConditionType filter.
			//IF it doesn't, return false.
			if !filter.Type.Valid(con.Type()) {
				valid = false
				return valid
			}
		}

		//ID condition
		if filter.ID != emptyCon {
			//Checking if requirementCondtion matches ConditionType filter.
			//IF it doesn't, return false.
			if !filter.ID.Valid(con.ID()) {
				valid = false
				return valid
			}
		}
		//Condition condition
		if filter.Condition != emptyCon && con.Condition() != emptyReqCon {
			//Checking if requirementCondtion matches ConditionType filter.
			//IF it doesn't, return false.
			if !filter.Condition.Valid(con.Condition()) {
				valid = false
				return valid
			}
		}
		//Name condition
		if filter.Name != emptyCon {
			//Checking if requirementCondtion matches ConditionType filter.
			//IF it doesn't, return false.
			if !filter.Name.Valid(con.Name()) {
				valid = false
				return valid
			}
		}
		//Save condition
		if filter.Save != emptyCon {
			//Checking if requirementCondtion matches ConditionType filter.
			//IF it doesn't, return false.
			tmp, err := con.Save()
			if err != nil {
				valid = false
				return valid
			}
			if !filter.Save.Valid(tmp) {
				valid = false
				return valid
			}
		}
	}

	return valid
}

func (filter *RequirementConditionFilter) Filter(cons []requirements.Condition) []requirements.Condition {
	var valid []requirements.Condition

	for _, con := range cons {
		if filter.Valid(con) {
			valid = append(valid, con)
		}
	}
	return valid
}
