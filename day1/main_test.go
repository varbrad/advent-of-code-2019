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

func TestDay1Part2(t *testing.T) {
	t.Run(`works on a small example`, func(t *testing.T) {
		numbers := []int{
			14,     // 2
			1969,   // 966
			100756, // 50346
		}

		expected := 51314
		actual := Day1Part2(numbers)

		assert.Equal(t, expected, actual)
	})
}
