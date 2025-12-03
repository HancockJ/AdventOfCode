package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	banks, err := loadData()
	if err != nil {
		panic(err)
	}
	total := 0
	// fmt.Println(banks)
	for _, bank := range banks {
		most := ""
		lastMostI := -1
		for i := 11; i >= 0; i-- {
			nextCandidate := 0
			candidateI := lastMostI
			for j := lastMostI + 1; j < len(bank)-i; j++ {
				if bank[j] > nextCandidate {
					nextCandidate = bank[j]
					candidateI = j
				}
			}
			most += strconv.Itoa(nextCandidate)
			lastMostI = candidateI
		}
		fmt.Println(most)
		tmp, _ := strconv.Atoi(most)
		total += tmp
	}
	fmt.Println("TOTAL ", total)
}

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// WHY DID MY BRUTE FORCE CHRISTMAS TREEE NOT WORK??? Oh.. I guess I don't have a few trillion years to wait. :(
// for a := range bank {
// 	for b := a + 1; b < len(bank); b++ {
// 		for c := b + 1; c < len(bank); c++ {
// 			for d := c + 1; d < len(bank); d++ {
// 				for e := d + 1; e < len(bank); e++ {
// 					for f := e + 1; f < len(bank); f++ {
// 						for g := f + 1; g < len(bank); g++ {
// 							for h := g + 1; h < len(bank); h++ {
// 								for i := h + 1; i < len(bank); i++ {
// 									for j := i + 1; j < len(bank); j++ {
// 										for k := j + 1; k < len(bank); k++ {
// 											for l := k + 1; l < len(bank); l++ {
// 												tmp, _ := strconv.Atoi(fmt.Sprintf("%d%d%d%d%d%d%d%d%d%d%d%d", bank[a], bank[b], bank[c], bank[d], bank[e], bank[f], bank[g], bank[h], bank[i], bank[j], bank[k], bank[l]))
// 												most = max(most, tmp)
// 											}
// 										}
// 									}
// 								}
// 							}
// 						}
// 					}
// 				}
// 			}
// 		}
// 	}
// }

func loadData() ([][]int, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	banks := [][]int{}

	for scanner.Scan() {
		tmp := []int{}
		line := scanner.Text()
		for i := range line {
			tmp = append(tmp, int(line[i]-'0'))
		}
		banks = append(banks, tmp)
	}

	return banks, nil
}
