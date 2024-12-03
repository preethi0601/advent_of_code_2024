package day1

import (
	helper "aoc/Helper"
	inputFormatter "aoc/Helper/InputFormatter"
	"math"
	"sort"
)

func Day1Challenge() (part1 int, part2 int) {
	inputString := helper.FileReader("day1/input.txt")
	arr1, arr2 := inputFormatter.FormatInputToTwoLists(inputString)

	part1 = int(PairsAndDistancesPart1(arr1, arr2))
	part2 = PairsAndDistancesPart2(arr1, arr2)

	return
}
func PairsAndDistancesPart1(arr1, arr2 []int) (result float64) {
	sort.Ints(arr1)
	sort.Ints(arr2)
	result = 0

	for i := range len(arr1) {
		result += math.Abs(float64(arr1[i] - arr2[i]))
	}
	return
}

func PairsAndDistancesPart2(arr1, arr2 []int) (result int) {
	m := make(map[int]int)
	result = 0
	for i := range len(arr2) {
		m[arr2[i]] = m[arr2[i]] + 1
	}
	for i := range len(arr1) {
		result += arr1[i] * m[arr1[i]]
	}
	return
}
