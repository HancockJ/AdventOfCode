package days

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

func Eleven() {
	// set blinks
	blinks := 20

	// Read input
	input, err := os.ReadFile("days/11.txt")
	if err != nil {
		fmt.Println(err)
	}

	inputString := string(input)

	fmt.Println("Output:")

	stones := toIntArray(strings.Split(inputString, " "))

	fmt.Println("stone list")
	fmt.Println(stones)
	fmt.Println("-----------")

	// no go routines
	start := time.Now()
	var fullStones []int
	for _, stone := range stones {
		fullStones = append(fullStones, blinkThrough([]int{stone}, blinks)...)
	}
	originalTime := time.Since(start)
	fmt.Println("original")
	originalAnswer := len(fullStones)
	fmt.Println("stones:", originalAnswer, "time:", originalTime)
	fmt.Println("--------------")

	// go routines
	conStart := time.Now()
	conStones := blinkConcurrent(stones, blinks)
	conTime := time.Since(conStart)

	fmt.Println("concurrency")
	conAnswer := len(conStones)
	fmt.Println("stones:", conAnswer, "time:", conTime)
	fmt.Println("--------------")

	// recursion
	recStart := time.Now()
	recAnswer := blinkRecursion(stones, blinks)
	recStop := time.Since(recStart)

	fmt.Println("recursion")
	fmt.Println("stones:", recAnswer, "time:", recStop)
	fmt.Println("--------------")

}

func blink(stones []int) []int {
	var newStones []int
	for i := range stones {
		if stones[i] == 0 {
			newStones = append(newStones, 1)
		} else if len(fmt.Sprint(stones[i]))%2 == 0 {
			strStone := fmt.Sprint(stones[i])
			m := len(strStone) / 2
			left, right := strStone[:m], strStone[m:]
			num, err := strconv.Atoi(left)
			if err != nil {
				fmt.Println("Error:", err)
			}
			newStones = append(newStones, num)
			num, err = strconv.Atoi(right)
			if err != nil {
				fmt.Println("Error:", err)
			}
			newStones = append(newStones, num)
		} else {
			newStones = append(newStones, stones[i]*2024)
		}
	}
	return newStones
}

func toIntArray(strArray []string) []int {
	var intArray []int
	for _, str := range strArray {
		num, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			intArray = append(intArray, num)
		}
	}
	return intArray
}

func blinkThrough(stones []int, blinks int) []int {
	for i := 0; i < blinks; i++ {
		stones = blink(stones)
	}
	return stones
}

// blinkConcurrent processes the stones concurrently using goroutines and channels with optimized batching
func blinkConcurrent(stones []int, blinks int) []int {
	var wg sync.WaitGroup
	numWorkers := runtime.NumCPU()     // Number of goroutines equal to the number of CPUs
	ch := make(chan []int, numWorkers) // Channel to collect results

	// If there are fewer stones than workers, each worker should get at least one stone
	chunkSize := len(stones) / numWorkers
	if chunkSize == 0 {
		numWorkers = len(stones) // If there are fewer stones than workers, reduce workers to stones count
		chunkSize = 1            // Assign one stone per worker
	}

	// Split the work into numWorkers slices and process them concurrently
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)

		// Define the range of stones for each worker
		startIdx := i * chunkSize
		endIdx := startIdx + chunkSize
		if endIdx > len(stones) {
			endIdx = len(stones) // Ensure the last worker processes any remaining stones
		}

		// Launch a goroutine to process the chunk
		go func(start, end int) {
			defer wg.Done()
			partialStones := stones[start:end]

			// Perform all blinks on the slice in one go
			for i := 0; i < blinks; i++ {
				partialStones = blink(partialStones)
			}

			// Send the processed result to the channel
			ch <- partialStones
		}(startIdx, endIdx)
	}

	// Wait for all goroutines to finish processing
	wg.Wait()

	// Close the channel once all goroutines are done
	close(ch)

	// Combine all the results into one final slice
	var fullStones []int
	for result := range ch {
		fullStones = append(fullStones, result...)
	}

	return fullStones
}

var cache = map[string]int{}

// blinkRecursion will recursively execute blinks and use memoization to store previous states
func blinkRecursion(stones []int, blinks int) int {
	totalLength := 0
	for _, stone := range stones {
		totalLength += blinkR(stone, blinks)
	}
	return totalLength
}

// returns the len of the final array
func blinkR(stone int, depth int) int {
	// Base case, return 1
	if depth == 0 {
		return 1
	}

	// Check if we have seen this case and use that if so
	hashValue := getHash(stone, depth)
	res, exists := cache[hashValue]
	if exists {
		return res
	}

	// We haven't used this case. Run it and cache it
	if stone == 0 {
		answer := blinkR(1, depth-1)
		cache[hashValue] = answer
		return answer
	} else if len(fmt.Sprint(stone))%2 == 0 {
		strStone := fmt.Sprint(stone)
		m := len(strStone) / 2
		left, right := strStone[:m], strStone[m:]
		numLeft, err := strconv.Atoi(left)
		if err != nil {
			fmt.Println("Error:", err)
		}
		numRight, err := strconv.Atoi(right)
		if err != nil {
			fmt.Println("Error:", err)
		}
		answer := blinkR(numLeft, depth-1) + blinkR(numRight, depth-1)
		cache[hashValue] = answer
		return answer
	} else {
		answer := blinkR(stone*2024, depth-1)
		cache[hashValue] = answer
		return answer
	}
}

func getHash(val int, depth int) string {
	return fmt.Sprint(val, ",", depth)
}
