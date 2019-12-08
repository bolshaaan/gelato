package gelato

// Item represents
type Item struct {
	// SKU is unique number of product
	SKU SKU
	// Count represents number of product
	Count int
	// Price is unit price of a product
	Price int
}

type ItemsBySKU map[SKU]Item
