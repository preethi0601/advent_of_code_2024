package day2

import (
	helper "aoc/Helper"
	inputFormatter "aoc/Helper/InputFormatter"
	"math"
)

func Day2Challenge() (part1, part2 int) {
	inputString := helper.FileReader("day2/input.txt")
	arr := inputFormatter.FormatInputTo2DArray(inputString)

	part1 = SafeReport(arr)
	part2 = safeReportDampener(arr)
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

func safeReportDampener(inputArr [][]int) (count int) {
	for i := range len(inputArr) {
		for j := range len(inputArr[i]) {
			indexRemovedArr := removeIndex(inputArr[i], j)
			if strictlyIncreasing(indexRemovedArr) || strictlyDecreasing(indexRemovedArr) {
				if maxDiffInReports(indexRemovedArr) {
					count++
					break
				}
			}
		}
	}
	return
}

func removeIndex(arr []int, index int) (modified []int) {
	arrCopy := make([]int, len(arr))
	copy(arrCopy, arr)
	return append(arrCopy[:index], arrCopy[index+1:]...)

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
