package day1

import (
	helper "aoc/Helper"
	"math"
	"sort"
)

func PairsAndDistances() (result float64) {
	arr1, arr2 := helper.ReadInputFromFileAndFormat("day1/input.txt")

	sort.Ints(arr1)
	sort.Ints(arr2)
	result = 0

	for i := range len(arr1) {
		result += math.Abs(float64(arr1[i] - arr2[i]))
	}
	return
}

func PairsAndDistancesPart2() (result int) {
	arr1, arr2 := helper.ReadInputFromFileAndFormat("day1/input.txt")
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
