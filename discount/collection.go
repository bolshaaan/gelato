/* package rules collects different rules of price creation */

package discount

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/bolshaaan/gelato"
	"github.com/pkg/errors"
)

// Collection is collection of different price rules
// ByItem stores in sync.Map
// this primitive is useful because collection will be only once updated (at the start of application)
// and ByItem can be read by many goroutines so we reduce lock contention with sync.Map (instead simple
// map with Mutexes)
type Collection struct {
	// ByItem is rules by 1 item
	ByItem *sync.Map
	// TotalRules is rules for total price in cart
	ByTotalPrice ByTotalPrice
}

// NewCollection is rules constructor
func NewCollection() *Collection {
	return &Collection{
		ByItem: &sync.Map{},
	}
}

// SetTotalDiscount  adds total discount
func (c *Collection) SetTotalDiscount(percent, treshold int) {
	c.ByTotalPrice = *NewByTotalPrice(percent, treshold)
}

// SetItemDiscount adds item discount
func (c *Collection) SetItemDiscount(sku gelato.SKU, count, amount int) {
	if c.ByItem == nil {
		c.ByItem = &sync.Map{}
	}

	c.ByItem.Store(sku, *NewByItem(count, amount))
}

// GetItemDiscount returns discount by sku
// returns nil if no discount for item
func (c *Collection) GetItemDiscount(sku gelato.SKU) *ByItem {

	loadedByItem, ok := c.ByItem.Load(sku)
	if !ok {
		return nil
	}

	// do not check error, because values are always discount.ByItem type
	disc, ok := loadedByItem.(ByItem)
	if !ok {
		// strange, maybe log about it
		return nil
	}

	return &disc
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

	if c.ByItem == nil {
		c.ByItem = &sync.Map{}
	}

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

		if _, ok := c.ByItem.Load(gelato.SKU(r[0])); ok {
			return nil, fmt.Errorf("duplicate rule, line number: %d", lineNumber)
		}

		c.SetItemDiscount(gelato.SKU(r[0]), count, amount)

		lineNumber++
	}

	return c, nil
}
