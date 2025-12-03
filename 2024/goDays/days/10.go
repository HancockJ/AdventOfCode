package days

import (
	"fmt"
	"os"
	"strconv"
)

var directions = [][]int{
	{-1, 0}, // up
	{1, 0},  // down
	{0, -1}, // left
	{0, 1},  // right
}

func Ten() {
	// Read input
	input, err := os.ReadFile("days/10.txt")
	if err != nil {
		fmt.Println(err)
	}

	inputString := string(input)
	fmt.Println("trail map")
	fmt.Println(inputString)
	fmt.Println("-----------")

	// Use input
	trailMap := make([][]int, 0)
	var line []int
	for _, char := range inputString {
		char := string(char)
		if char == "\n" {
			trailMap = append(trailMap, line)
			line = []int{}
		} else {
			temp, err := strconv.Atoi(fmt.Sprint(char))
			if err != nil {
				fmt.Println(err)
			}
			line = append(line, temp)
		}
	}
	trailMap = append(trailMap, line)
	// fmt.Println("trail map")
	// fmt.Println(trailMap)
	// fmt.Println("-----------")

	trailHeads := make([][]int, 0)
	for i := 0; i < len(trailMap); i++ {
		for j := 0; j < len(trailMap[i]); j++ {
			if trailMap[i][j] == 0 {
				trailHeads = append(trailHeads, []int{i, j})
			}
			// fmt.Printf("trailMap[%d][%d] = %d\n", i, j, trailMap[i][j])
		}
	}
	// fmt.Println("Trail heads:")
	// fmt.Println(trailHeads)
	// fmt.Println(len(trailHeads))
	totalScore := 0
	for _, trailHead := range trailHeads {
		totalScore += getTrailScore(trailMap, trailHead[1], trailHead[0], make(map[string]bool))
	}

	fmt.Println(totalScore)

}

func getTrailScore(trailMap [][]int, x, y int, visited map[string]bool) int {
	score := 0

	visitedKey := fmt.Sprintf("%d,%d", x, y)
	if trailMap[y][x] == 9 {
		visited[visitedKey] = true
		return 1
	}

	currentVal := trailMap[y][x]

	for _, dir := range directions {
		newY, newX := y+dir[0], x+dir[1]
		if newY >= 0 && newY < len(trailMap) && newX >= 0 && newX < len(trailMap[0]) {
			if trailMap[newY][newX] == currentVal+1 {
				score += getTrailScore(trailMap, newX, newY, visited)
			}
		}
	}
	return score
}
