package day3

import (
	helper "aoc/Helper"
	"regexp"
	"strconv"
	"strings"
)

func Day3Challenge() (part1, part2 int) {
	inputStringArray := helper.FileReader("day3/input.txt")

	inputString := strings.Join(inputStringArray, " ")
	part1 = multiplyValidInstruction(inputString)
	part2 = multiplyDosNotDonts(inputString)
	return
}

func multiplyValidInstruction(input string) (product int) {
	matches := findValidMatches(input)
	product = 0
	for _, match := range matches {
		x, err := strconv.Atoi(string(match[1]))
		if err != nil {
			panic(err)
		}
		y, err := strconv.Atoi(string(match[2]))
		if err != nil {
			panic(err)
		}
		product += x * y
	}
	return
}

func findValidMatches(input string) (matches [][]string) {
	mulRegex := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	matches = mulRegex.FindAllStringSubmatch(input, len(input))
	return
}

func multiplyDosNotDonts(input string) (product int) {
	result := ""
	flag := true
	dontOffset := len("don't()")
	doOffset := len("do()")

	for len(input) > 0 {
		dontIndex := strings.Index(input, "don't()")
		doIndex := strings.Index(input, "do()")

		if dontIndex == -1 && doIndex == -1 { // both not present, line is valid
			if flag {
				result += input
			}
			break
		}
		// dont exists but either do doesn't or do is present after don't
		if dontIndex != -1 && (doIndex == -1 || dontIndex < doIndex) {
			if flag {
				result += input[:dontIndex]
			}
			flag = false
			input = input[dontIndex+dontOffset:]
		} else {
			if flag {
				result += input[:doIndex]
			}
			flag = true
			input = input[doIndex+doOffset:]
		}
	}

	return multiplyValidInstruction(result)
}
