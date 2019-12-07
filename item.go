package gelato

type Item struct {
	SKU   SKU
	Count int
}

type Items map[SKU]Item
