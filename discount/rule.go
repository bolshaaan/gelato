package discount

import "github.com/bolshaaan/gelato"

// Rules collects methods to work with discount
type Rules interface {
	GetItemDiscount(sku gelato.SKU) *ByItem
	GetTotalDiscount() *ByTotalPrice
	SetTotalDiscount(percent, treshold int)
	SetItemDiscount(sku gelato.SKU, count, amount int)
}
