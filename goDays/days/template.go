package days

import (
	"fmt"
	"os"
)

func template() {
	// Read input
	input, err := os.ReadFile("days/00.txt")
	if err != nil {
		fmt.Println(err)
	}

	inputString := string(input)
	fmt.Println(inputString)
}
