package discount

// ByItem shows price Amount for Count items
type ByItem struct {
	Count  int
	Amount int
}

// Apply calculates new price depending on discount
func (r *ByItem) Apply(iCount, iPrice int) (totalPrice int) {
	mod := iCount % r.Count
	totalPrice += mod * iPrice
	//fmt.Println("mdo", mod, "price", item.Price, "totalPrice", totalPrice)
	totalPrice += (iCount - mod) / r.Count * r.Amount
	//fmt.Println("mdo", mod, "price", item.Price, "totalPrice", totalPrice, "diff", item.Count-mod, item.Count)
	return
}

// NewByItem creator of item rule
func NewByItem(count, amount int) *ByItem {
	return &ByItem{count, amount}
}
