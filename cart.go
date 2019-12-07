package gelato

import "fmt"

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
		c.Items[item.SKU] = v
	} else {
		c.Items[item.SKU] = item
	}
}

func (c *Cart) Total() int {
	var totalPrice int
	fmt.Println(c.Items)
	for sku, item := range c.Items {

		rule, ok := c.Rules.ItemRules[sku]
		if !ok {
			totalPrice += item.Price * item.Count
			continue
		}

		totalPrice += rule.Apply(item)
	}

	// execute rule on total price
	totalPrice = c.Rules.TotalPriceRule.Apply(totalPrice)

	return totalPrice
}
