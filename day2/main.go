package main

import (
	"log"

	"github.com/varbrad/advent-of-code-2019/shared"
	"github.com/varbrad/advent-of-code-2019/utils"
)

func main() {
	input, err := utils.ReadInputToIntegerListWithSeperator("day2/input", ",")

	if err != nil {
		log.Fatal(err)
	}

	utils.Part1(Day2Part1(input))
}

func Day2Part1(input []int) int {
	program := shared.NewIntcode(input)
	program.SetValue(1, 12)
	program.SetValue(2, 2)
	return program.Run().GetValue(0)
}
