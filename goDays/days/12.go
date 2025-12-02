package days

import (
	"aoc2024/helpers"
	"fmt"
)

var regionCount int
var region map[string]int
var regionArea map[int]int

// var regionPerimeter map[int]int

func Twelve() {
	//// import and parse data
	farm, err := helpers.Get2DArray("days/12.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, row := range farm {
		fmt.Println(row)
	}
	fmt.Println("plots to check:", len(farm)*len(farm[0]))

	region = make(map[string]int)
	regionArea = make(map[int]int)

	regionCount = 0
	// loop through entire farm
	for y := 0; y < len(farm); y++ {
		for x := 0; x < len(farm[y]); x++ {
			// checking farm[y][x]
			hashValue := getRHash(x, y)
			_, exists := region[hashValue]
			if !exists {
				// Create new region
				fmt.Println("New region at ", x, y)
				addToRegion(farm, farm[y][x], x, y)
				regionCount += 1
			}
		}
	}
	fmt.Println(region)
	fmt.Println(len(region))
	fmt.Println(regionArea)
}

func getRHash(x, y int) string {
	return fmt.Sprint(x, ",", y)
}

// addToRegion does a DFS to add all plots in current region
func addToRegion(farm [][]string, currentRegion string, x, y int) {
	// check if in bounds
	if x < 0 || x >= len(farm[0]) || y < 0 || y >= len(farm) {
		return
	}

	// skip if already in a region
	hashValue := getRHash(x, y)
	_, exists := region[hashValue]
	if exists {
		return
	}

	// check if it should be added to current Region
	if farm[y][x] == currentRegion {
		// add to region
		region[getRHash(x, y)] = regionCount
		val, exists := regionArea[regionCount]
		if exists {
			regionArea[regionCount] = val + 1
		} else {
			regionArea[regionCount] = 1
		}
	} else {
		// add to region
		region[getRHash(x, y)] = -1
		val, exists := regionArea[regionCount]
		if exists {
			regionArea[regionCount] = val + 1
		} else {
			regionArea[regionCount] = 1
		}
	}

	// try to add left or right
	addToRegion(farm, currentRegion, x+1, y)
	addToRegion(farm, currentRegion, x, y+1)
	addToRegion(farm, currentRegion, x-1, y)
	addToRegion(farm, currentRegion, x, y-1)
	// calculate perimeter
}
