package inputformatter

import (
	"strconv"
	"strings"
)

func FormatInputToTwoLists(data []string) (list1, list2 []int) {
	list1 = make([]int, 1000)
	list2 = make([]int, 1000)

	for i := 0; i < len(data); i++ {
		pairs := strings.Split(data[i], "   ")
		list1[i], _ = strconv.Atoi(pairs[0])
		list2[i], _ = strconv.Atoi(pairs[1])
	}
	return
}
