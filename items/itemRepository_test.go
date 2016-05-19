package items

/*
func TestAdd(t *testing.T) {
	repo := MemItemRepository{}
	item := Item{"1", "Apple", 25.00, nil, "anapple.jpg"}
	repo.Add(item)

	reqItem, err := repo.Item("1")
	if err != nil {
		t.Error("Error received when searching for added item", err)
	}
	if reqItem.Id() != "1" {
		t.Error("Expecting to receive an item with the ID \"1\" but received", reqItem.Id())
	}
}

func TestItems(t *testing.T) {
	repo := MemItemRepository{}
	itemOne := Item{"1", "Apple", 25.00, nil, "anapple.jpg"}
	itemTwo := Item{"2", "Orange", 25.00, nil, "anorange.jpg"}

	_, err := repo.Items()
	if err == nil {
		t.Error("Expected to receive an error indicating an empty item repo but an error was not thrown")
	}

	repo.Add(itemOne)
	itemListOne, err := repo.Items()
	if err != nil {
		t.Error("Error received when calling ItemRepository.Items", err)
	}
	if len(*itemListOne) != 1 {
		t.Error("Expected to received a list of items with the length 1 but received a list with the length", len(*itemListOne))
	}

	repo.Add(itemTwo)
	itemListTwo, err := repo.Items()
	if err != nil {
		t.Error("Error received when calling ItemRepository.Items", err)
	}
	if len(*itemListTwo) != 2 {
		t.Error("Expected to received a list of items with the length 2 but received a list with the length", len(*itemListTwo))
	}
}

func TestRemove(t *testing.T) {
	repo := MemItemRepository{}
	itemOne := Item{"1", "Apple", 25.00, nil, "anapple.jpg"}
	itemTwo := Item{"2", "Orange", 25.00, nil, "anorange.jpg"}

	repo.Add(itemOne)
	repo.Add(itemTwo)
	itemList, err := repo.Items()
	if err != nil {
		t.Error("Received an error when trying to call list items", err)
	}
	if len(*itemList) != 2 {
		t.Error("Expected to received a list of items with the length 2 but received a list with the length", len(*itemList))
	}
	item, err := repo.Item("1")
	if err != nil {
		t.Error("Error received when trying to get the item with the id \"1\"", err)
	}
	if item.Name() != "Apple" {
		t.Error("Expecting to receive an item with the name \"Apple\" but instead received an item with the name", item.Name())
	}
	err = repo.Remove("1")
	if err != nil {
		t.Error("Received an error when trying to remove an item", err)
	}

	_, err = repo.Item("1")
	if err == nil {
		t.Error("Expected to receive an item not found error but no error was thrown")
	}
	itemList, err = repo.Items()
	if err != nil {
		t.Error("Received an error when trying to call list items", err)
	}
	if len(*itemList) != 1 {
		t.Error("An item was removed from the repo yet received a list of items bigger than expected: len = ", len(*itemList))
	}
}
*/
