package days

import (
	"aoc2024/helpers"
	"fmt"
)

func Eight() {
	//// import and parse data
	fMap, err := helpers.Get2DArray("days/8.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	antiNodes := make([][]bool, len(fMap))
	for i := range antiNodes {
		antiNodes[i] = make([]bool, len(fMap[0]))
	}

	//// use data
	for y := range fMap {
		for x := range fMap[y] {
			current := fMap[y][x]
			if current != "." {
				antiNodes[y][x] = true
				setAntiNodes(fMap, antiNodes, x, y, current)
			}
		}
	}

	//// print results
	printStrMap(fMap)
	fmt.Println("------------------------")
	printMap(antiNodes)
}

// setAntiNodes takes in current antenna and adds in anti nodes for each
func setAntiNodes(fMap [][]string, aNodes [][]bool, x, y int, char string) {
	for i := y; i < len(fMap); i++ {
		for j := 0; j < len(fMap[0]); j++ {
			if !(i == y && j <= x) {
				if char == fMap[i][j] {
					// drop anti nodes
					xStep := x - j
					yStep := y - i
					newX, newY := j-xStep, i-yStep
					for newY >= 0 && newY < len(fMap) && newX >= 0 && newX < len(fMap[0]) {
						aNodes[newY][newX] = true
						newX, newY = newX-xStep, newY-yStep
					}
					newX, newY = x+xStep, y+yStep
					for newY >= 0 && newY < len(fMap) && newX >= 0 && newX < len(fMap[0]) {
						aNodes[newY][newX] = true
						newX, newY = newX+xStep, newY+yStep
					}
				}
			}
		}
	}
}

// printMap prints the antiNodes map and total count
func printMap(aNodes [][]bool) {
	fmt.Println("printing anti-node map:")
	count := 0
	for _, row := range aNodes {
		for _, val := range row {
			if val {
				count += 1
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println("------------------------")
	fmt.Println("Total anti nodes:", count)
}

// printStrMap prints the actual intenna map
func printStrMap(chars [][]string) {
	fmt.Println("Printing antenna map")
	for _, row := range chars {
		for _, val := range row {
			fmt.Print(val)
		}
		fmt.Println()
	}
}
