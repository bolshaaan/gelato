package gelato

// RulePricer calculates discount
type RulePricer interface {
	Price(SKU string) int
}

type RuleTotaler interface {
}

// TotalPriceRule  makes rule for total prices
type TotalPriceRule struct {
	// DiscountPercent from 0 to 100
	DiscountPercent int
	// AmountThreshold
	AmountThreshold int
}

// ItemRule shows price Amount for Count items
type ItemRule struct {
	Count  int
	Amount int
}
type ItemRules map[SKU]ItemRule

func (r *TotalPriceRule) Apply(totalPrice int) int {
	// Make discount if AmountTreshold achieved
	if totalPrice < r.AmountThreshold {
		return totalPrice
	}
	// apply discount
	totalPrice = totalPrice - (totalPrice*r.DiscountPercent)/100.0
	return totalPrice
}

func (r *ItemRule) Apply(item Item) (totalPrice int) {
	mod := item.Count % r.Count
	totalPrice += mod * item.Price
	//fmt.Println("mdo", mod, "price", item.Price, "totalPrice", totalPrice)
	totalPrice += (item.Count - mod) / r.Count * r.Amount
	//fmt.Println("mdo", mod, "price", item.Price, "totalPrice", totalPrice, "diff", item.Count-mod, item.Count)
	return
}

// Rules is collection of different price rules
type Rules struct {
	// TotalRules is rules for total price in cart
	TotalPriceRule TotalPriceRule
	// ItemRules is rules by 1 item
	ItemRules ItemRules
}

// NewRules is rules constructor
func NewRules() *Rules {
	return &Rules{}
}
