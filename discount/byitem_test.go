package discount

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestByItem_Apply(t *testing.T) {

	tests := []struct {
		byItem    ByItem
		itemCount int
		iPrice    int
		expected  int
		message   string
	}{
		{ByItem{10, 30}, 100, 20, 300, "simple case"},
		{ByItem{10, 30}, 101, 20, 320, "simple case + 1 more"},
	}

	for _, test := range tests {
		t.Run(test.message, func(t *testing.T) {
			assert.Equal(t,
				test.byItem.Apply(test.itemCount, test.iPrice),
				test.expected,
			)

		})
	}

}
