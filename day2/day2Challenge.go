package day2

import (
	helper "aoc/Helper"
	inputFormatter "aoc/Helper/InputFormatter"
	"math"
)

func Day2Challenge() (part1 int) {
	inputString := helper.FileReader("day2/input.txt")
	arr := inputFormatter.FormatInputTo2DArray(inputString)

	part1 = SafeReport(arr)

	return
}

func SafeReport(inputArr [][]int) (count int) {

	for i := range len(inputArr) {
		if strictlyIncreasing(inputArr[i]) || strictlyDecreasing(inputArr[i]) {
			if maxDiffInReports(inputArr[i]) {
				count++
			}
		}
	}
	return
}

func strictlyIncreasing(report []int) bool {
	for i := range len(report) - 1 {
		if report[i] >= report[i+1] {
			return false
		}
	}
	return true
}

func strictlyDecreasing(report []int) bool {
	for i := range len(report) - 1 {
		if report[i] <= report[i+1] {
			return false
		}
	}
	return true
}

func maxDiffInReports(report []int) bool {
	for i := range len(report) - 1 {
		if math.Abs(float64(report[i]-report[i+1])) > 3 {
			return false
		}
	}
	return true
}
