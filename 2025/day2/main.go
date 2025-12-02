package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	strRanges, err := loadData()
	if err != nil {
		panic(err)
	}
	splits := strings.Split(strRanges, ",")
	ranges := make([][]int, len(splits))
	for i := range splits {
		tmp := strings.Split(splits[i], "-")
		a, _ := strconv.Atoi(tmp[0])
		b, _ := strconv.Atoi(tmp[1])
		ranges[i] = []int{a, b}
	}
	// fmt.Println(ranges)

	invalid := 0
	for i := range ranges {
		for j := ranges[i][0]; j <= ranges[i][1]; j++ {
			// fmt.Println(j)
			if isSequence(strconv.Itoa(j)) {
				fmt.Println(j, "- found invalid")
				invalid += j
			}
		}
	}
	fmt.Printf("invalid sum is %d\n", invalid)

}

func isSequence(str string) bool {
	// fmt.Println("SEQ CHECK", str)
	for i := 1; i <= len(str)/2; i++ {
		if len(str)%i != 0 {
			continue
		}
		sub := str[:i]
		broke := false
		// fmt.Print(sub, "=>")
		for j := i; j <= len(str)-i; j += i {
			tmp := str[j : j+i]
			// fmt.Print(tmp, ",")
			if sub != tmp {
				broke = true
				break
			}
		}
		// fmt.Println()
		if !broke {
			return true
		}
	}
	return false
}

func loadData() (string, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		return scanner.Text(), nil
	}
	return "", nil
}
