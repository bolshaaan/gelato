package cart

import (
	"github.com/bolshaaan/gelato"
	"github.com/bolshaaan/gelato/discount"
)

// make sure Cart implements checkout interface
var _ gelato.Checkout = &Cart{}

// Cart contains items with rules
type Cart struct {
	Items    gelato.Items
	Discount discount.Collection
}

func NewCart(c discount.Collection) *Cart {
	return &Cart{
		Items:    make(gelato.Items),
		Discount: c,
	}
}

func (c *Cart) Scan(item gelato.Item) {
	if v, ok := c.Items[item.SKU]; ok {
		v.Count += item.Count
		c.Items[item.SKU] = v
	} else {
		c.Items[item.SKU] = item
	}
}

func (c *Cart) Total() int {
	var totalPrice int

	for sku, item := range c.Items {

		disc, ok := c.Discount.ByItem[sku]
		if !ok {
			totalPrice += item.Price * item.Count
			continue
		}

		totalPrice += disc.Apply(item.Count, item.Price)
	}

	// execute rule on total price
	totalPrice = c.Discount.ByTotalPrice.Apply(totalPrice)

	return totalPrice
}
