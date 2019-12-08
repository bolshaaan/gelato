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
	Discount *discount.Collection
}

// NewCart generates new Cart
func NewCart(c *discount.Collection) *Cart {
	return &Cart{
		Items:    make(gelato.Items),
		Discount: c,
	}
}

// Scan adds new item to cart
func (c *Cart) Scan(item gelato.Item) {
	if v, ok := c.Items[item.SKU]; ok {
		v.Count += item.Count
		c.Items[item.SKU] = v
	} else {
		c.Items[item.SKU] = item
	}
}

// Total calculates total sum of products
// in cart applying all discount rules
func (c *Cart) Total() int {
	var totalPrice int

	for sku, item := range c.Items {

		loadedByItem, ok := c.Discount.ByItem.Load(sku)
		if !ok {
			totalPrice += item.Price * item.Count
			continue
		}
		// do not check error, because values are always discount.ByItem type
		disc := loadedByItem.(discount.ByItem)
		totalPrice += disc.Apply(item.Count, item.Price)
	}

	// execute rule on total price
	totalPrice = c.Discount.ByTotalPrice.Apply(totalPrice)

	return totalPrice
}
