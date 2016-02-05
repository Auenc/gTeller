package items

import (
	//"errors"
	"strconv"

	"github.com/auenc/gTeller-core/discounts"
)

type Item struct {
	ID         string              `form:"id" binding:"required"`
	NameRaw    string              `form:"name" binding:"required"`
	PriceRaw   string              `form:"price" binding:"required"`
	DicountRaw *discounts.Discount `form:"-"`
	Options    []OptionCategory
	ImageURI   string `form:"image" binding:"required"`
}

//ID is a function to return the items id in the string format.
func (item *Item) Id() string {
	return item.ID
}

//Valid returns a boolean whether or not the item is valid
func (item *Item) Valid() bool {
	_, err := item.Price()
	if err != nil {
		return false
	}
	return true
}

//Price is a method that will return the chargable price of an item.
//If the item has a discount applied to it, the discount will be applied.
func (item *Item) Price() (float64, error) {
	var price float64
	var err error
	/*	if the item has a discount*/
	if item.DicountRaw != nil {
		/*	Get the discounted price*/
		float, err := strconv.ParseFloat(item.PriceRaw, 64)
		if err != nil {
			return 0.0, err
		}
		price, err = item.DicountRaw.Calculate(float)
	} else {
		float, err := strconv.ParseFloat(item.PriceRaw, 64)
		if err != nil {
			return 0.0, err
		}
		price = float
	}
	return price, err
}

/**
* 	A method to change the name of the item
	@param string The new name of the item
*/
func (item *Item) SetName(name string) {
	item.NameRaw = name
}

/**
*	A method to return the item's name
	@return string The name of the item
*/
func (item *Item) Name() string {
	return item.NameRaw
}

/**
*	A method that returns the item's image URI
	@return string The URI pointing to the item's image
*/
func (item *Item) Image() string {
	return item.ImageURI
}

/**
*	A method that sets
 */
func (item *Item) SetImage(URI string) {
	item.ImageURI = URI
}

/**
* 	A method that returns the discount object associated
	with the item.
	@return Discount The discount applied to the item
*/
func (item *Item) Discount() *discounts.Discount {
	return item.DicountRaw
}

/**
* A method that is used to add a specified discount to an Item
	@param *Discount A pointer to the discount item
*/
func (item *Item) SetDiscount(discount *discounts.Discount) {
	item.DicountRaw = discount
}

func NewItem(id string, name string, price string, discount *discounts.Discount, imageURI string) Item {
	return Item{id, name, price, discount, nil, imageURI}
}
