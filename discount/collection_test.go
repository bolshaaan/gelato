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

		val, ok := c.ByItem.Load(gelato.SKU("0"))
		assert.True(t, ok)
		v, ok := val.(ByItem)
		assert.True(t, ok)

		assert.Equal(t, 100, v.Count)
		assert.Equal(t, 123, v.Amount)

	})
}
