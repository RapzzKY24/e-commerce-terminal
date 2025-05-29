package cart

type Cart struct {
	Items []Item
}

type Item struct {
	Name     string
	Quantity int
	Price    int
}

func (c *Cart) AddItem(item Item) {
	c.Items = append(c.Items, item)
}

func (c *Cart) RemoveItem(name string) {
	for i, item := range c.Items {
		if item.Name == name {
			c.Items = append(c.Items[:i], c.Items[i+1:]...)
			return
		}
	}
}

func (c *Cart) UpdateItem(name string, quantity int) {
	for i, item := range c.Items {
		if item.Name == name {
			c.Items[i].Quantity = quantity
			return
		}
	}
}

func (c *Cart) ClearCart() {
	c.Items = []Item{}
}
