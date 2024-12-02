package helper

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadInputFromFileAndFormat(fileName string) (input1, input2 []int) {
	fmt.Println("Reading from file..")
	data, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatalf("Unable to read file: %v", err)
	}
	input1, input2 = formatInputToTwoLists(data)
	return
}

func formatInputToTwoLists(data []byte) (list1, list2 []int) {
	list1 = make([]int, 1000)
	list2 = make([]int, 1000)

	temp := strings.Fields(string(data))
	l := 0
	for i := 0; i < len(temp); i += 2 {
		list1[l], _ = strconv.Atoi(temp[i])
		list2[l], _ = strconv.Atoi(temp[i+1])
		l++
	}
	return
}
