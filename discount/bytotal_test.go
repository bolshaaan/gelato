package discount

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestByTotalPrice_Apply(t *testing.T) {

	tests := []struct {
		byTotalPrice ByTotalPrice
		inAmount     int
		expected     int
		message      string
	}{
		{ByTotalPrice{10, 30}, 100, 90, "simple case"},
	}
	for _, test := range tests {
		t.Run(test.message, func(t *testing.T) {
			assert.Equal(t,
				test.byTotalPrice.Apply(test.inAmount),
				test.expected,
			)
		})
	}

}
