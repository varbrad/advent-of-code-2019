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
	utils.Part2(Day2Part2(input))
}

func Day2Part1(input []int) int {
	program := shared.NewIntcode(input)
	program.SetAddress(1, 12)
	program.SetAddress(2, 2)
	return program.Run().GetAddress(0)
}

func Day2Part2(input []int) int {
	program := shared.NewIntcode(input)
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			program.Reset()
			program.SetAddress(1, noun)
			program.SetAddress(2, verb)
			program.Run()
			if program.GetAddress(0) == 19690720 {
				return 100*noun + verb
			}
		}
	}

	return -1
}
