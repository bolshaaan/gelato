/* package rules collects different rules of price creation */

package discount

import (
	"github.com/bolshaaan/gelato"
)

type ItemDiscounts map[gelato.SKU]ByItem

// Collection is collection of different price rules
type Collection struct {
	// TotalRules is rules for total price in cart
	ByTotalPrice ByTotalPrice
	// ByItem is rules by 1 item
	ByItem ItemDiscounts
}

// NewRules is rules constructor
func NewRules() *Collection {
	return &Collection{}
}

// SetTotalDiscount  adds total discount
func (c *Collection) SetTotalDiscount(percent, treshold int) {
	c.ByTotalPrice = *NewByTotalPrice(percent, treshold)
}

// SetItemDiscount adds item discount
func (c *Collection) SetItemDiscount(sku gelato.SKU, count, amount int) {
	if c.ByItem == nil {
		c.ByItem = make(ItemDiscounts)
	}

	c.ByItem[sku] = *NewByItem(count, amount)
}
