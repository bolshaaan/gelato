package gelato

type Checkout interface {
	Scan(item Item)
	Checkout() int
}
