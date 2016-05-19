package orders

import "github.com/auenc/gTeller/requirements"

type OrderItem struct {
	UUID         string
	ItemID       string
	Requirements []requirements.Requirement
	Price        float64
	ImageURL     string
	Quantity     int64
}

type OrderItemInfo struct {
	UUID         string
	Name         string
	Requirements []requirements.Requirement
	Price        float64
	ImageURL     string
	Quantity     int64
}

func (item *OrderItem) GetPrice() float64 {
	return item.Price
}
func (item *OrderItemInfo) SetPrice(price float64) {
	item.Price = price
}
