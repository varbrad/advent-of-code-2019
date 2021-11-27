package shared

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntcode(t *testing.T) {
	t.Run(`works on a small example`, func(t *testing.T) {
		input := []int{1, 0, 0, 0, 99}

		expected := []int{2, 0, 0, 0, 99}
		actual := NewIntcode(input).Run().GetMemory()

		assert.Equal(t, expected, actual)
	})
}
