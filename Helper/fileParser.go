package helper

import (
	"log"
	"os"
	"strings"
)

func FileReader(fileName string) (inputString []string) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatalf("Unable to read file: %v", err)
	}
	inputString = strings.Split(string(data), "\n")
	return
}
