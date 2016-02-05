package items

/*
func TestPrice(t *testing.T) {
	item := Item{"1", "Apple", 25.00, nil, "anapple.jpg"}

	p, err := item.Price()
	if err != nil {
		t.Error("Error occured, ", err)
	}
	if p != 25.00 {
		t.Error("Expected 25.00, got ", p)
	}

	discount := Discount{"per", DISCOUNT_TYPE_PERCENTAGE, 50, 0}

	item = Item{"1", "Orange", 50, &discount, "anorange.jpg"}
	p, err = item.Price()
	if err != nil {
		t.Error("Error occured, ", err)
	}
	if p != 25.00 {
		t.Error("Expected 25.00 got ", p)
	}
}

func TestName(t *testing.T) {
	item := Item{"1", "Apple", 25.00, nil, "anapple.jpg"}

	name := item.Name()
	if name != "Apple" {
		t.Error("Expecting Apple got ", name)
	}
}

func TestID(t *testing.T) {
	item := Item{"ab", "Apple", 25.00, nil, "anapple.jpg"}

	id := item.Id()
	if id != "ab" {
		t.Error("Expecting to receive the ID \"ab\" but received", id)
	}
}

func TestSetName(t *testing.T) {
	item := Item{"1", "Apple", 25.00, nil, "anapple.jpg"}

	item.SetName("Fish")
	name := item.Name()
	if name != "Fish" {
		t.Error("Expecting Fish got ", name)
	}
}

func TestImage(t *testing.T) {
	item := Item{"1", "Apple", 25.00, nil, "anapple.jpg"}

	image := item.Image()

	if image != "anapple.jpg" {
		t.Error("Expected anapple.jpg got ", image)
	}
}

func TestSetImage(t *testing.T) {
	item := Item{"1", "Apple", 25.00, nil, "anapple.jpg"}

	item.SetImage("apple2.jpg")
	name := item.Image()
	if name != "apple2.jpg" {
		t.Error("Expected apple2.jpg got ", name)
	}
}

func TestDiscount(t *testing.T) {
	item := Item{"1", "Apple", 25.00, nil, "anapple.jpg"}

	discount := item.Discount()

	if discount != nil {
		t.Error("Expected nil got", discount)
	}

	dis1 := Discount{"per", DISCOUNT_TYPE_PERCENTAGE, 50, 0}
	item.SetDiscount(&dis1)

	discount = item.Discount()
	if discount != &dis1 {
		t.Errorf("Expected %p got ", discount)
	}
}
*/
