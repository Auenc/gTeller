package status

import (
	"github.com/auenc/gTeller-core/email"
)

//Status is an object that represents an order status
type Status struct {
	ID            string
	Name          string
	EmailTemplate email.EmailTemplate
}
