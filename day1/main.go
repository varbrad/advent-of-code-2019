package main

import (
	"log"

	"github.com/varbrad/advent-of-code-2019/utils"
)

func main() {
	input, err := utils.ReadInputToIntegerList("day1/input")

	if err != nil {
		log.Fatal(err)
	}

	utils.Part1(Day1Part1(input))
}

func Day1Part1(input []int) int {
	var sum int

	for _, mass := range input {
		sum += calculateFuel(mass)
	}

	return sum
}

func calculateFuel(mass int) int {
	return (mass / 3) - 2
}
