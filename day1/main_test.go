package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay1Part1(t *testing.T) {
	t.Run(`works on a small example`, func(t *testing.T) {
		numbers := []int{
			12,     // 2
			14,     // 2
			1969,   // 654
			100756, // 33583
		}

		expected := 34241
		actual := Day1Part1(numbers)

		assert.Equal(t, expected, actual)
	})
}
