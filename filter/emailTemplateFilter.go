package filter

import "github.com/auenc/gTeller/email"

type EmailTemplateFilter struct {
	ID          Condition
	Name        Condition
	Subject     Condition
	Content     Condition
	SenderEmail Condition
	SenderName  Condition
}

//Valid returns true if the given EmailTemplate matches the filter
func (filter *EmailTemplateFilter) Valid(temp email.EmailTemplate) bool {
	valid := true
	emptyCon := Condition{}

	//Checking if there is a filter for Status ID
	if filter.ID != emptyCon {
		//Checking if Status matches ID filter.
		//IF it doesn't, return false.
		if !filter.ID.Valid(temp.ID) {
			valid = false
			return valid
		}
	}

	//Checking if there is a filter for Status Name
	if filter.Name != emptyCon {
		//Checking if Status matches Name filter.
		//IF it doesn't, return false.
		if !filter.Name.Valid(temp.Name) {
			valid = false
			return valid
		}
	}

	//Checking if there is a filter for Status Subject
	if filter.Subject != emptyCon {
		//Checking if Status matches Subject filter.
		//IF it doesn't, return false.
		if !filter.Subject.Valid(temp.Subject) {
			valid = false
			return valid
		}
	}

	//Checking if there is a filter for Status Content
	if filter.Content != emptyCon {
		//Checking if Status matches Content filter.
		//IF it doesn't, return false.
		if !filter.Content.Valid(temp.Content) {
			valid = false
			return valid
		}
	}

	//Checking if there is a filter for Status SenderEmail
	if filter.SenderEmail != emptyCon {
		//Checking if Status matches SenderEmail filter.
		//IF it doesn't, return false.
		if !filter.SenderEmail.Valid(temp.SenderEmail) {
			valid = false
			return valid
		}
	}

	//Checking if there is a filter for Status SenderName
	if filter.SenderName != emptyCon {
		//Checking if Status matches SenderName filter.
		//IF it doesn't, return false.
		if !filter.SenderName.Valid(temp.SenderName) {
			valid = false
			return valid
		}
	}

	return valid
}

//Filter filters out any EmailTemplates that do not meet the filter within a given slice
func (filter *EmailTemplateFilter) Filter(source []email.EmailTemplate) []email.EmailTemplate {
	var filtered []email.EmailTemplate

	//Lopping through source EmailTemplates
	for _, email := range source {
		if filter.Valid(email) {
			filtered = append(filtered, email)
		}
	}

	return filtered
}
