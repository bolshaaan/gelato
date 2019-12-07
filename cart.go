package gelato

// Cart contains items with rules
type Cart struct {
	Items Items
	Rules Rules
}

func NewCart(rules Rules) *Cart {
	return &Cart{
		Items: make(Items),
		Rules: rules,
	}
}

func (c *Cart) Scan(item Item) {
	if v, ok := c.Items[item.SKU]; ok {
		v.Count += item.Count
	} else {
		c.Items[item.SKU] = item
	}
}

func (c *Cart) Total() int {

	// apply rules

	return 0
}
