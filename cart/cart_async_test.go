package cart

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/bolshaaan/gelato"

	"github.com/bolshaaan/gelato/discount"
)

// Test that we can  use cart in many goroutines and have no data races
// run with -race
func TestCart_AsyncTotal(t *testing.T) {
	col := discount.NewCollection()
	col.SetItemDiscount("A", 10, 20)
	col.SetItemDiscount("B", 5, 20)
	col.SetItemDiscount("C", 20, 20)
	col.SetItemDiscount("D", 70, 20)
	col.SetItemDiscount("E", 90, 20)

	shoppers := 10
	wg := &sync.WaitGroup{}
	for i := 0; i < shoppers; i++ {
		wg.Add(1)

		go func(col *discount.Collection, wg *sync.WaitGroup) {
			defer wg.Done()

			cart := NewCart(col)
			cart.Scan(gelato.Item{
				SKU:   "A",
				Count: 100,
				Price: 3,
			})

			//fmt.Println(cart.Total())
			assert.Equal(t, 200, cart.Total())

		}(col, wg)
	}

	wg.Wait()

}
