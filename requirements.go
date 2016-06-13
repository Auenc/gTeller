package gTeller

import "github.com/auenc/gTeller/filter"

//ParseableRequirement is a data object returned by gTeller-core that can be converted
//into a requirement object
type ParseableRequirement struct {
	UUID          string
	Type          int
	ConditionType int
	Reference     string
	Condition     string
}

//ListRequirementRequest represents the data that can be received by the GetItems api call
type ListRequirementRequest struct {
	Filter filter.RequirementFilter
}

//AddRequirementRequest represents the data that can be received by the AddRequirement api call
type AddRequirementRequest struct {
	Type          int
	ConditionType int
	Condition     string
	Reference     string
}

//RemoveEmailTemplateRequest represents the data that can be received by the RemoveItems api call
type RemoveRequirementRequest struct {
	IDs []string
}

//UpdateEmailTemplateRequest represents the data that update item accepts.
type UpdateRequirementRequest struct {
	Requirements []ParseableRequirement
}
