package helpers

import (
	"os"
)

func Get2DArray(fileName string) ([][]string, error) {
	input, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	inputString := string(input)

	var mapArr [][]string
	line := make([]string, 0)
	for _, char := range inputString {
		c := string(char)
		if c == "\n" {
			mapArr = append(mapArr, line)
			line = make([]string, 0)
		} else {
			line = append(line, c)
		}
	}
	mapArr = append(mapArr, line)
	return mapArr, nil
}
