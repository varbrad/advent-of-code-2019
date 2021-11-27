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
	utils.Part2(Day1Part2(input))
}

func Day1Part1(input []int) int {
	var sum int

	for _, mass := range input {
		sum += calculateFuel(mass, false)
	}

	return sum
}

func Day1Part2(input []int) int {
	var sum int

	for _, mass := range input {
		sum += calculateFuel(mass, true)
	}

	return sum
}

func calculateFuel(mass int, recursive bool) int {
	fuel := mass/3 - 2
	if recursive {
		if fuel > 0 {
			return fuel + calculateFuel(fuel, true)
		} else {
			return 0
		}
	}
	return fuel
}
