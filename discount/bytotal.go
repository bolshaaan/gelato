package discount

// ByTotalPrice  makes rule for total prices
type ByTotalPrice struct {
	// DiscountPercent from 0 to 100
	DiscountPercent int
	// AmountThreshold
	AmountThreshold int
}

// Apply discount total price
func (r *ByTotalPrice) Apply(totalPrice int) int {
	// Make discount if AmountTreshold achieved
	if totalPrice < r.AmountThreshold {
		return totalPrice
	}
	// apply discount
	totalPrice = totalPrice - (totalPrice*r.DiscountPercent)/100.0
	return totalPrice
}

// NewByTotalPrice creates ByTotalPrice rule
func NewByTotalPrice(discountPercent, amountThreshold int) *ByTotalPrice {
	return &ByTotalPrice{discountPercent, amountThreshold}
}
