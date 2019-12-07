package gelato

type Item struct {
	SKU   SKU
	Count int
	Price int
}

type Items map[SKU]Item
