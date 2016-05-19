package items

import (
//
)

type ItemHandler struct {
	ItemRepo ItemRepository
}

type FormItem struct {
	Id    string `form:"id" binding:"required"`
	Name  string `form:"name" binding:"required"`
	Price string `form:"price" binding:"required"`
	Image string `form:"image" binding:"required"`
}

/*
func (handler *ItemHandler) Handle(m *martini.ClassicMartini) {
	//m.Get("/api/items/", handler.listItems)
	m.Get("/items/", handler.renderItems)
	m.Post("/admin/items/add/", binding.Bind(Item{}), handler.addItem)
}

func (handler *ItemHandler) renderItems(render render.Render) {
	items, err := handler.ItemRepo.Items()
	if err != nil {
		log.Fatal("Something went wrong")
	}
	render.HTML(200, "test", items)
}

func (handler *ItemHandler) addItem(item Item, render render.Render) {
	log.Println("Adding item", item)
	handler.ItemRepo.Add(item)
}

/*func (handler *ItemHandler) listItems(enc Encoder) (int, string) {
	items, err := handler.ItemRepo.Items()
	log.Println(items[0].Name())
	if err != nil {
		r.Text(200, "No items found")
	} else {
		r.JSON(200, &NewItem("1", "fish", 25.00, nil, "afish.png"))
	}

}
*/
