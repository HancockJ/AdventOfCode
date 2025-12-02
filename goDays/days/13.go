package days

import (
	"fmt"
	"os"
)

func Thirteen() {
	// Read input
	input, err := os.ReadFile("days/00.txt")
	if err != nil {
		fmt.Println(err)
	}

	inputString := string(input)
	fmt.Println(inputString)
}
