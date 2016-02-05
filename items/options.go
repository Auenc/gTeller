package items

import (
	"errors"
)

type ByPrice []Option

type OptionCategory struct {
	ID      string
	Name    string
	Options []Option
}
type Option struct {
	ID    string
	Name  string
	Price float64
}

func (cat *OptionCategory) Option(id string) (*Option, error) {
	for _, option := range cat.Options {
		if option.ID == id {
			return &option, nil
		}
	}
	return nil, errors.New("Could not find option with id" + id)
}

func (op ByPrice) Len() int {
	return len(op)
}
func (op ByPrice) Less(i, j int) bool {
	return op[i].Price < op[j].Price
}
func (op ByPrice) Swap(i, j int) {
	op[i], op[j] = op[j], op[i]
}
