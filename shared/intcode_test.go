package shared

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntcode(t *testing.T) {
	t.Run(`works on a small example with ADD #1`, func(t *testing.T) {
		input := []int{1, 0, 0, 0, 99}

		expected := []int{2, 0, 0, 0, 99}
		actual := NewIntcode(input).Run().GetMemory()

		assert.Equal(t, expected, actual)
	})

	t.Run(`works on a small example with ADD #2`, func(t *testing.T) {
		input := []int{1, 1, 1, 4, 99, 5, 6, 0, 99}

		expected := []int{30, 1, 1, 4, 2, 5, 6, 0, 99}
		actual := NewIntcode(input).Run().GetMemory()

		assert.Equal(t, expected, actual)
	})

	t.Run(`works on a small example with MULTIPLY #1`, func(t *testing.T) {
		input := []int{2, 3, 0, 3, 99}

		expected := []int{2, 3, 0, 6, 99}
		actual := NewIntcode(input).Run().GetMemory()

		assert.Equal(t, expected, actual)
	})

	t.Run(`works on a small example with MULTIPLY #2`, func(t *testing.T) {
		input := []int{2, 4, 4, 5, 99, 0}

		expected := []int{2, 4, 4, 5, 99, 9801}
		actual := NewIntcode(input).Run().GetMemory()

		assert.Equal(t, expected, actual)
	})
}
