package main

import (
	"log"

	"github.com/varbrad/advent-of-code-2019/utils"
)

func main() {
	input, err := utils.ReadInputToIntegerList("day2/input")

	if err != nil {
		log.Fatal(err)
	}

	utils.Part1(Day2Part1(input))
}

func Day2Part1(input []int) int {
	return 0
}
