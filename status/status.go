package status

import (
	"github.com/auenc/gTeller/email"
)

//Status is an object that represents an order status
type Status struct {
	ID            string
	Name          string
	EmailTemplate email.EmailTemplate
}
