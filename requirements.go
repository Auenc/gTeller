package gTeller

import (
	"github.com/auenc/gTeller/filter"
	"github.com/auenc/gTeller/requirements"
)

//ListRequirementRequest represents the data that can be received by the GetItems api call
type ListRequirementRequest struct {
	Filter filter.RequirementFilter
}

//AddRequirementRequest represents the data that can be received by the AddRequirement api call
type AddRequirementRequest struct {
	Type          int
	ConditionType int
	Condition     string
}

//RemoveEmailTemplateRequest represents the data that can be received by the RemoveItems api call
type RemoveRequirementRequest struct {
	IDs []string
}

//UpdateEmailTemplateRequest represents the data that update item accepts.
type UpdateRequirementRequest struct {
	Requirements []requirements.Requirement
}
