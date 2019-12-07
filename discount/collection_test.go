package discount

import (
	"testing"

	"github.com/bolshaaan/gelato"

	"github.com/stretchr/testify/assert"
)

func TestLoadFromFile(t *testing.T) {
	t.Run("Read file success case", func(t *testing.T) {
		c, err := LoadFromFile("testdata/discounts.csv")
		assert.NoError(t, err)
		assert.NotNil(t, c)

		assert.Equal(t, 100, c.ByItem[gelato.SKU("0")].Count)
		assert.Equal(t, 123, c.ByItem[gelato.SKU("0")].Amount)
	})
}
