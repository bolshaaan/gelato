/* package rules collects different rules of price creation */

package discount

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/bolshaaan/gelato"
	"github.com/pkg/errors"
)

type ItemDiscounts map[gelato.SKU]ByItem

// Collection is collection of different price rules
type Collection struct {
	// TotalRules is rules for total price in cart
	ByTotalPrice ByTotalPrice
	// ByItem is rules by 1 item
	ByItem ItemDiscounts
}

// NewCollection is rules constructor
func NewCollection() *Collection {
	return &Collection{}
}

// SetTotalDiscount  adds total discount
func (c *Collection) SetTotalDiscount(percent, treshold int) {
	c.ByTotalPrice = *NewByTotalPrice(percent, treshold)
}

// SetItemDiscount adds item discount
func (c *Collection) SetItemDiscount(sku gelato.SKU, count, amount int) {
	if c.ByItem == nil {
		c.ByItem = make(ItemDiscounts)
	}

	c.ByItem[sku] = *NewByItem(count, amount)
}

// LoadFromFile loads from file
// format of file is csv "SKU;count;amount"
func LoadFromFile(fileName string) (*Collection, error) {
	c := NewCollection()

	file, err := os.Open(fileName)
	if err != nil {
		return nil, errors.Wrap(err, "cannot open file")
	}

	defer file.Close()

	buf := bufio.NewReader(file)

	c.ByItem = make(ItemDiscounts)

	var lineNumber int
	var eof bool
	for !eof {
		l, err := buf.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				eof = true
			} else {
				return nil, errors.Wrap(err, "error reading file")
			}
		}

		r := strings.Split(l, ";")
		if len(r) != 3 {
			return nil, fmt.Errorf("file format is bad, line number: %d", lineNumber)
		}

		count, err := strconv.Atoi(r[1])
		if err != nil {
			return nil, errors.Wrap(err, "not integer")
		}

		amount, err := strconv.Atoi(r[2])
		if err != nil {
			return nil, errors.Wrap(err, "not integer")
		}

		if count < 1 {
			return nil, fmt.Errorf("count is less then 1, line number: %d", lineNumber)
		}

		if amount < 1 {
			return nil, fmt.Errorf("amount is less then 1, line number: %d", lineNumber)
		}

		if _, ok := c.ByItem[gelato.SKU(r[0])]; ok {
			return nil, fmt.Errorf("duplicate rule, line number: %d", lineNumber)
		}
		fmt.Println("count", count, "amount", amount)
		c.SetItemDiscount(gelato.SKU(r[0]), count, amount)

		lineNumber++
	}

	return c, nil
}
