package gelato

type Checkout interface {
	Scan(item Item)
	Total() int
}
