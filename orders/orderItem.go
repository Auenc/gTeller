package orders

import (
	"github.com/auenc/gTeller-core/items"
)

type OrderItem struct {
	UUID     string
	ItemID   string
	Options  map[string]string //map[OPTION_UUID]CHOICE_UUID
	ImageURL string
	Quantity int64
}

type OrderItemInfo struct {
	UUID     string
	Name     string
	Options  map[string]string
	Price    float64
	ImageURL string
	Quantity int64
}

func (oi *OrderItem) Info(itemRepo items.ItemRepository) (OrderItemInfo, error) {
	var info OrderItemInfo
	//Getting price
	price, err := oi.Price(itemRepo)
	if err != nil {
		return info, err
	}
	//Getting the item name
	item, err := itemRepo.Item(oi.ItemID)
	if err != nil {
		return info, err
	}
	name := item.NameRaw

	//Getting option names
	options := make(map[string]string, len(oi.Options))
	for key, val := range oi.Options {
		//Getting option category
		optionCat, err := itemRepo.Option(key)
		if err != nil {
			return info, err
		}
		optionCatName := optionCat.Name
		//Getting option choice
		option, err := optionCat.Option(val)
		if err != nil {
			return info, err
		}
		optionName := option.Name
		//Adding option to map
		options[optionCatName] = optionName
	}
	info = OrderItemInfo{oi.UUID, name, options, price, oi.ImageURL, oi.Quantity}
	return info, nil
}

func (oi *OrderItem) Price(itemRepo items.ItemRepository) (float64, error) {
	var price float64
	item, err := itemRepo.Item(oi.ItemID)
	if err != nil {
		return price, err
	}

	//Getting item price
	itemPrice, err := item.Price()
	if err != nil {
		return price, err
	}
	//Adding items price to total
	price += itemPrice

	for key, val := range oi.Options {
		//Getting option category
		optionCat, err := itemRepo.Option(key)
		if err != nil {
			return -1, err
		}
		//Getting option choice
		option, err := optionCat.Option(val)
		if err != nil {
			return -1, err
		}
		//Adding options price to total
		price += option.Price
	}
	return price * float64(oi.Quantity), nil
}
