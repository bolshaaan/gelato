package cart

import (
	"testing"

	"github.com/bolshaaan/gelato"

	"github.com/bolshaaan/gelato/discount"

	"github.com/magiconair/properties/assert"
)

func TestCart_Total(t *testing.T) {

	t.Run("Test item price rules", func(t *testing.T) {
		tests := []struct {
			inItems        []gelato.Item
			inRules        discount.Collection
			expectedAmount int
			desc           string
		}{
			{
				[]gelato.Item{},
				discount.Collection{ByItem: discount.ItemDiscounts{}},
				0,
				"Zero items - return 0",
			},
			{
				[]gelato.Item{{SKU: "A", Count: 3, Price: 13}},
				discount.Collection{ByItem: discount.ItemDiscounts{}},
				39,
				"Apply empty rules - no discount",
			},
			{
				[]gelato.Item{{SKU: "A", Count: 3, Price: 13}},
				discount.Collection{ByItem: discount.ItemDiscounts{
					gelato.SKU("A"): discount.ByItem{Count: 3, Amount: 30},
				}},
				30,
				"Buy 3 items for 30 dollars",
			},
			{
				[]gelato.Item{{SKU: "A", Count: 2, Price: 13}},
				discount.Collection{ByItem: discount.ItemDiscounts{
					gelato.SKU("A"): discount.ByItem{Count: 3, Amount: 30},
				}},
				26,
				"Have only 2 items for 26 dollars - no discount",
			},
			{
				[]gelato.Item{{SKU: "A", Count: 4, Price: 13}},
				discount.Collection{ByItem: discount.ItemDiscounts{
					gelato.SKU("A"): discount.ByItem{Count: 3, Amount: 30},
				}},
				43, // 30 + 13
				"By 4 items - 3 discount 1 - no",
			},
			{
				[]gelato.Item{
					{SKU: "A", Count: 1, Price: 13},
					{SKU: "B", Count: 1, Price: 7},
					{SKU: "A", Count: 1, Price: 13},
					{SKU: "C", Count: 1, Price: 3},
					{SKU: "A", Count: 1, Price: 13},
				},
				discount.Collection{ByItem: discount.ItemDiscounts{
					gelato.SKU("A"): discount.ByItem{Count: 3, Amount: 30},
					gelato.SKU("B"): discount.ByItem{Count: 2, Amount: 10},
				}},
				40, // A = 30  + B = 7 + C = 3
				"Different items - A = 30  + B = 7 + C = 3  = 40 ",
			},
		}

		for _, test := range tests {
			t.Run(test.desc, func(t *testing.T) {
				cart := NewCart(test.inRules)

				for _, item := range test.inItems {
					cart.Scan(item)
				}

				assert.Equal(t, cart.Total(), test.expectedAmount)
			})
		}

	})

	t.Run("Test total price rules", func(t *testing.T) {
		tests := []struct {
			inItems        []gelato.Item
			inRules        discount.Collection
			expectedAmount int
			desc           string
		}{
			{
				[]gelato.Item{
					{SKU: "A", Count: 1, Price: 100},
					{SKU: "B", Count: 1, Price: 100},
					{SKU: "C", Count: 1, Price: 100},
				},
				discount.Collection{ByTotalPrice: discount.ByTotalPrice{
					DiscountPercent: 50,
					AmountThreshold: 100,
				}},
				150,
				"Zero items - return 0",
			},
		}

		for _, test := range tests {
			t.Run(test.desc, func(t *testing.T) {
				cart := NewCart(test.inRules)

				for _, item := range test.inItems {
					cart.Scan(item)
				}

				assert.Equal(t, cart.Total(), test.expectedAmount)
			})
		}

	})

}
