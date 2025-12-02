package days

import (
	"fmt"
	"os"
	"strings"
)

var moves map[rune][]int

func Fifteen() {
	// Read input
	input, err := os.ReadFile("days/15.txt")
	if err != nil {
		fmt.Println(err)
	}
	inputString := string(input)

	// Split the input into grid and movements parts
	parts := strings.Split(inputString, "\n\n")
	gridText := parts[0]
	movements := strings.ReplaceAll(parts[1], "\n", "")

	var grid [][]rune
	line := make([]rune, 0)
	for _, char := range gridText {
		c := char
		if c == '\n' {
			grid = append(grid, line)
			line = make([]rune, 0)
		} else {
			line = append(line, c)
		}
	}
	grid = append(grid, line)

	// loop through map
	var location []int
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == '@' {
				location = []int{x, y}
			}
		}
	}

	moves = map[rune][]int{
		'>': {0, 1},  // Move right: x + 1
		'<': {0, -1}, // Move left: x - 1
		'^': {-1, 0}, // Move up: y - 1
		'v': {1, 0},  // Move down: y + 1
	}
	printGrid(grid)
	var newGrid [][]rune
	for y := 0; y < len(grid); y++ {
		newGrid = append(newGrid, []rune{})
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == '@' {
				newGrid[y] = append(newGrid[y], '@')
				newGrid[y] = append(newGrid[y], '.')
			}
			if grid[y][x] == 'O' {
				newGrid[y] = append(newGrid[y], '[')
				newGrid[y] = append(newGrid[y], ']')
			}
			if grid[y][x] == '.' {
				newGrid[y] = append(newGrid[y], '.')
				newGrid[y] = append(newGrid[y], '.')
			}
			if grid[y][x] == '#' {
				newGrid[y] = append(newGrid[y], '#')
				newGrid[y] = append(newGrid[y], '#')
			}
		}
	}

	// part 1

	for _, move := range movements {
		nextLocation := []int{location[0] + moves[move][0], location[1] + moves[move][1]}
		if nextLocation[0] >= 0 && nextLocation[0] < len(grid) && nextLocation[1] >= 0 && nextLocation[1] < len(grid[0]) && grid[nextLocation[0]][nextLocation[1]] == '.' {
			grid[location[0]][location[1]] = '.'
			grid[nextLocation[0]][nextLocation[1]] = '@'
			location = nextLocation
			continue
		}
		if grid[nextLocation[0]][nextLocation[1]] == 'O' {
			for x, y := nextLocation[0], nextLocation[1]; x < len(grid[0]) && x >= 0 && y < len(grid) && y >= 0; x, y = x+moves[move][0], y+moves[move][1] {
				if grid[x][y] == '@' {
					continue
				}
				if grid[x][y] == '#' {
					break
				}
				if grid[x][y] == '.' {
					grid[x][y] = 'O'
					grid[location[0]][location[1]] = '.'
					grid[nextLocation[0]][nextLocation[1]] = '@'
					location = nextLocation
					break
				}
			}
		}
	}

	totalGps := 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == 'O' {
				//fmt.Println(y, x)
				totalGps += (y * 100) + x
			}
		}
	}
	fmt.Println(totalGps)

	// part 2
	grid = newGrid
	// loop through map
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == '@' {
				location = []int{x, y}
				break
			}
		}
	}
	fmt.Println(location)
	for _, move := range movements {
		printGrid(grid)
		fmt.Println("Move: ", string(move), ":")
		nextLocation := []int{location[0] + moves[move][0], location[1] + moves[move][1]}
		fmt.Println("Next location is ", string(grid[nextLocation[0]][nextLocation[1]]), nextLocation[0], nextLocation[1])
		if grid[nextLocation[0]][nextLocation[1]] == '.' {
			fmt.Println("Moving")
			grid[location[0]][location[1]] = '.'
			grid[nextLocation[0]][nextLocation[1]] = '@'
			location = nextLocation
			continue
		}
		// left/right move
		if move == '>' || move == '<' {
			if grid[nextLocation[0]][nextLocation[1]] == '[' || grid[nextLocation[0]][nextLocation[1]] == ']' {
				// box in the way, try to move
				for x, y := nextLocation[0], nextLocation[1]; x < len(grid[0]) && x >= 0 && y < len(grid) && y >= 0; x, y = x+moves[move][0], y+moves[move][1] {
					if grid[x][y] == '@' {
						continue
					}
					if grid[x][y] == '#' {
						break
					}
					if grid[x][y] == '.' {
						// move is possible
						grid[nextLocation[0]][nextLocation[1]] = '@'
						grid[location[0]][location[1]] = '.'

						for i := location[0] + 1; i <= x; i++ {
							if i%2 == 0 {
								grid[nextLocation[0]][i] = '['
							} else {
								grid[nextLocation[0]][i] = ']'
							}
						}
						location = nextLocation
						break
					}
				}
			}
		}
		// up/down move
		if move == 'V' || move == '^' {
			if grid[nextLocation[0]][nextLocation[1]] == '[' || grid[nextLocation[0]][nextLocation[1]] == ']' {
				// box in the way, try to move
				for x, y := nextLocation[0], nextLocation[1]; x < len(grid[0]) && x >= 0 && y < len(grid) && y >= 0; x, y = x+moves[move][0], y+moves[move][1] {
					if grid[x][y] == '@' {
						continue
					}
					if grid[x][y] == '#' {
						break
					}
					if grid[x][y] == '.' {
						grid[location[0]][location[1]] = '.'
						grid[x][y] = 'O'

						grid[nextLocation[0]][nextLocation[1]] = '@'
						location = nextLocation
						break
					}
				}
			}
		}
	}
	printGrid(grid)
}

func printGrid(grid [][]rune) {
	for _, row := range grid {
		for _, col := range row {
			fmt.Print(string(col))
		}
		fmt.Println()
	}
}
