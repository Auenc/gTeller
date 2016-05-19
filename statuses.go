package gTeller

import (
	"github.com/auenc/gTeller/email"
	"github.com/auenc/gTeller/filter"
	"github.com/auenc/gTeller/shipping"
	"github.com/auenc/gTeller/statuses"
)

//ShippingType objects
type ListStatusRequest struct {
	Filter filter.StatusFilter
}

type ListStatusResponse struct {
	Items []shipping.ShippingType
}

type AddStatusRequest struct {
	Name          string
	EmailTemplate email.EmailTemplate
}

type UpdateStatusRequest struct {
	Statuses []status.Status
}

type RemoveStatusRequest struct {
	IDs []string
}
