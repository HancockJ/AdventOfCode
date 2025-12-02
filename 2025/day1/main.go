package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type rotation struct {
	direction string
	distance  int
}

func main() {
	// fmt.Println((100 + -5) % 100)
	rotations, err := loadData()
	if err != nil {
		fmt.Println("THE ELVES ARE GOING TO BE MAD AT ME!")
		panic(err)
	}
	// fmt.Println(rotations)

	dial := 50
	pw := 0
	for i := range rotations {
		// fmt.Println(dial)
		if rotations[i].direction == "L" {
			dial = ((dial - rotations[i].distance) + 100) % 100
		} else {
			dial = (dial + rotations[i].distance) % 100
		}
		if dial == 0 {
			pw += 1
		}
	}
	fmt.Printf("The password for silver is %d!\n", pw)

	fmt.Println("Performing 0x434C49434B password method now!")
	pw = 0
	dial = 50
	for i := range rotations {

		fmt.Println(dial)
		// each 100 is just another full loop, so it visits 0 and ends back up at same spot
		pw += rotations[i].distance / 100
		// now we can ignore full loops and just focus on the mod
		rotations[i].distance = rotations[i].distance % 100
		// fmt.Println(rotations[i].distance)
		if rotations[i].direction == "L" {
			tmp := dial - rotations[i].distance
			// if dial is already at 0, it never actually passes it again
			if tmp <= 0 && dial != 0 {
				pw += 1
			}
			dial = ((dial - rotations[i].distance) + 100) % 100
		} else {
			tmp := dial + rotations[i].distance
			if tmp > 99 {
				pw += 1
			}
			dial = (dial + rotations[i].distance) % 100
		}
	}
	fmt.Printf("The password for gold is %d!\n", pw)
}

func loadData() ([]rotation, error) {
	file, err := os.Open("data.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var rotations []rotation
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		currentDistance, err := strconv.Atoi(line[1:])
		if err != nil {
			return nil, err
		}

		rotations = append(rotations, rotation{
			direction: string(line[0]),
			distance:  currentDistance,
		})
	}

	return rotations, nil
}
