package items

import (
	"github.com/auenc/gTeller/filter"
	"github.com/auenc/gTeller/users"
)

type ItemRepository interface {
	Add(item Item, auth users.User) error
	Item(id string) (Item, error)
	Items(filter filter.ItemsFilter) ([]Item, error)
	Update(items []Item, auth users.User) error
	Remove(items []Item, auth users.User) error
	//	RemoveOption(id string) error
}

/*
type MemItemRepository struct {
	items       []Item
	itemOptions []OptionCategory
	logger      logging.Logger
}

func (repo *MemItemRepository) Add(item Item) error {
	repo.items = append(repo.items, item)
	return nil
}

func (repo *MemItemRepository) Item(id string) (*Item, error) {
	for _, item := range repo.items {
		if item.Id() == id {
			return &item, nil
		}
	}
	return nil, errors.New("Item not found") //TODO add error type for missing item
}

func (repo *MemItemRepository) Items() ([]Item, error) {
	if len(repo.items) > 0 {
		return repo.items, nil
	} else {
		repo.logger.LogLine("Call to Items received but no items exist")
		return nil, errors.New("There are no items stored")
	}
}

func (repo *MemItemRepository) Remove(id string) error {
	var removed bool
	for i, item := range repo.items {
		if item.Id() == id {
			repo.items = append(repo.items[:i], repo.items[i+1:]...)
			removed = true
			break
		}
	}
	if !removed {
		return errors.New("Item not found") //TODO add error type for missing item
	} else {
		return nil
	}
}

func (repo *MemItemRepository) AddOption(option OptionCategory) error {
	repo.itemOptions = append(repo.itemOptions, option)
	return nil
}

func (repo *MemItemRepository) Option(id string) (OptionCategory, error) {
	optionCat := OptionCategory{}
	for _, option := range repo.itemOptions {
		if option.ID == id {
			return option, nil
		}
	}
	return optionCat, errors.New("Option not found")
}

func (repo *MemItemRepository) Options() ([]OptionCategory, error) {
	if len(repo.itemOptions) > 0 {
		return repo.itemOptions, nil
	}
	return nil, errors.New("No options")
}

func (repo *MemItemRepository) RemoveOption(id string) error {
	var removed bool
	for i, option := range repo.items {
		if option.ID == id {
			repo.itemOptions = append(repo.itemOptions[:i], repo.itemOptions[i+1:]...)
			removed = true
			break
		}
	}
	if removed {
		return nil
	}
	return errors.New("Option not found")
}
*/
