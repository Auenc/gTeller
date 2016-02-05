package items

import "github.com/auenc/gTeller-core/items"

type Category struct {
	mId      string
	mName    string
	mItems   *[]items.Item
	mEnabled bool
}

func (cat *Category) addItem(item *items.Item) {
	//cat.mItems = append(cat.mItems, item)
}
