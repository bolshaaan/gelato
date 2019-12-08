package cart

import (
	"github.com/bolshaaan/gelato"
	"github.com/bolshaaan/gelato/discount"
)

// make sure Cart implements checkout interface
var _ gelato.Checkout = &Cart{}

// Cart contains items with rules
type Cart struct {
	ItemsBySKU gelato.ItemsBySKU
	Discount   discount.Rules
}

// NewCart generates new Cart
func NewCart(rulesDiscount discount.Rules) *Cart {
	return &Cart{
		ItemsBySKU: make(gelato.ItemsBySKU),
		Discount:   rulesDiscount,
	}
}

// Scan adds new item to cart
func (c *Cart) Scan(item gelato.Item) {
	if v, ok := c.ItemsBySKU[item.SKU]; ok {
		v.Count += item.Count
		c.ItemsBySKU[item.SKU] = v
	} else {
		c.ItemsBySKU[item.SKU] = item
	}
}

// Total calculates total sum of products
// in cart applying all discount rules
func (c *Cart) Total() int {
	var totalPrice int

	for sku, item := range c.ItemsBySKU {

		disc := c.Discount.GetItemDiscount(sku)
		if disc == nil {
			totalPrice += item.Price * item.Count
			continue
		}

		totalPrice += disc.Apply(item.Count, item.Price)
	}

	// execute rule on total price
	totalPrice = c.Discount.GetTotalDiscount().Apply(totalPrice)

	return totalPrice
}
