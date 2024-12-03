package inputformatter

import (
	"strconv"
	"strings"
)

func FormatInputTo2DArray(input []string) (list [][]int) {
	list = make([][]int, len(input))
	for i := range len(list) {
		temp := strings.Split(input[i], " ")
		list[i] = make([]int, len(temp))

		for j, val := range temp {
			list[i][j], _ = strconv.Atoi(val)
		}
	}
	return
}
