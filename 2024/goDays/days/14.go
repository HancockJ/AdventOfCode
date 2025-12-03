package days

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func Fourteen() {
	file, err := os.Open("days/14.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	var robots [][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var px, py, vx, vy int

		// Parse line like "p=9,5 v=-3,-3"
		_, err := fmt.Sscanf(line, "p=%d,%d v=%d,%d", &px, &py, &vx, &vy)
		if err != nil {
			fmt.Println("Error parsing:", err)
			continue
		}
		robots = append(robots, []int{px, py, vx, vy})
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	// Print results
	// for _, robot := range robots {
	// 	fmt.Println(robot)
	// }

	wide := 101
	tall := 103

	fmt.Println("Initial State")
	// printRobotMap(robots, tall, wide)
	for i := 0; i <= 7000; i += 1 {
		if fullRow(getRobotMap(robots, tall, wide)) > 32 {
			fmt.Println("ANSWER ---->", i)
			printRobotMap(robots, tall, wide)
			fmt.Println("After", i, " seconds")
			time.Sleep(100 * time.Millisecond)
		}
		robots = moveRobots(robots, tall, wide)
	}
	// printRobotMap(robots, tall, wide)
	roboMap := getRobotMap(robots, tall, wide)
	middleX := wide / 2
	middleY := tall / 2
	// fmt.Println(middleX, middleY)
	quadrantCount := []int{0, 0, 0, 0}
	for i := range roboMap {
		for j := range roboMap[i] {
			if i < middleY && j < middleX {
				// q0
				quadrantCount[0] += roboMap[i][j]
			}
			if i < middleY && j > middleX {
				// q1
				quadrantCount[1] += roboMap[i][j]
			}
			if i > middleY && j < middleX {
				// q2
				quadrantCount[2] += roboMap[i][j]
			}
			if i > middleY && j > middleX {
				// q3
				quadrantCount[3] += roboMap[i][j]
			}
		}
	}
	// fmt.Println(quadrantCount)
	total := 1
	for _, q := range quadrantCount {
		total *= q
	}
	// fmt.Println(total)
}

func moveRobots(robots [][]int, tall, wide int) [][]int {
	for i, robot := range robots {
		// fmt.Println(robots[i])
		if (robot[0] + robot[2]) < 0 {
			robots[i][0] = wide + (robot[0] + robot[2])
		} else {
			robots[i][0] = (robot[0] + robot[2]) % wide
		}
		if (robot[1] + robot[3]) < 0 {
			robots[i][1] = tall + (robot[1] + robot[3])
		} else {
			robots[i][1] = (robot[1] + robot[3]) % tall
		}
		// fmt.Println(robots[i])
	}
	return robots
}

func printRobotMap(robots [][]int, tall, wide int) {
	// Create a 2D map to track robot counts
	grid := make([][]int, tall)
	for i := range grid {
		grid[i] = make([]int, wide)
	}

	// Fill the grid with robot counts
	for _, robot := range robots {
		x, y := robot[0], robot[1]
		if y >= 0 && y < tall && x >= 0 && x < wide {
			grid[y][x]++
		}
	}

	// Print the grid
	for y := 0; y < tall; y++ {
		for x := 0; x < wide; x++ {
			if grid[y][x] > 0 {
				fmt.Print(grid[y][x])
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func getRobotMap(robots [][]int, tall, wide int) [][]int {
	// Create a 2D map to track robot counts
	grid := make([][]int, tall)
	for i := range grid {
		grid[i] = make([]int, wide)
	}

	// Fill the grid with robot counts
	for _, robot := range robots {
		x, y := robot[0], robot[1]
		if y >= 0 && y < tall && x >= 0 && x < wide {
			grid[y][x]++
		}
	}

	return grid
}

func fullRow(robotMap [][]int) int {
	largestRow := 0
	for _, row := range robotMap {
		rowTotal := 0
		for i := range row {
			rowTotal += row[i]
		}
		if rowTotal > largestRow {
			largestRow = rowTotal
		}
	}
	return largestRow
}
