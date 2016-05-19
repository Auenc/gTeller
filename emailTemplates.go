package gTeller

import (
	"github.com/auenc/gTeller/email"
	"github.com/auenc/gTeller/filter"
)

//ListEmailTemplateRequest represents the data that can be received by the GetItems api call
type ListEmailTemplateRequest struct {
	Filter filter.EmailTemplateFilter
}

//AddEmailTemplateRequest represents the data that can be received by the AddShipping api call
//TODO Describe AddItemRequest more
type AddEmailTemplateRequest struct {
	Name        string
	Subject     string
	Content     string
	SenderEmail string
	SenderName  string
}

//RemoveEmailTemplateRequest represents the data that can be received by the RemoveItems api call
//TODO Describe RemoveItemsRequest more
type RemoveEmailTemplateRequest struct {
	IDs []string
}

//UpdateEmailTemplateRequest represents the data that update item accepts.
type UpdateEmailTemplateRequest struct {
	EmailTemplate []email.EmailTemplate
}
