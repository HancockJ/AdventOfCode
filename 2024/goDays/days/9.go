package days

import (
	"fmt"
	"os"
	"strconv"
)

func Nine() {
	input, err := os.ReadFile("days/9.txt")
	if err != nil {
		fmt.Println(err)
	}

	inputString := string(input)
	fmt.Println(inputString)
	// id := 0
	fileBuffer := make([]string, 0)
	files := make(map[int]int)
	for i, char := range inputString {
		if i%2 == 0 {
			// id := int(char - '0')
			files[i/2] = int(char - '0')
			for j := 0; j < int(char-'0'); j++ {
				fileBuffer = append(fileBuffer, fmt.Sprint(i/2))
			}
		} else {
			for j := 0; j < int(char-'0'); j++ {
				fileBuffer = append(fileBuffer, ".")
			}
		}
	}

	// for a, b := 0, len(fileBuffer)-1; a < b; {
	// 	fmt.Println(fileBuffer[a], fileBuffer[b])
	// 	if fileBuffer[a] == "." {
	// 		if fileBuffer[b] != "." {
	// 			tmp := fileBuffer[b]
	// 			fileBuffer[b] = "."
	// 			fileBuffer[a] = tmp
	// 			a++
	// 			b--
	// 		} else {
	// 			b--
	// 		}
	// 	} else {
	// 		a += 1
	// 	}
	// }
	// checkSum := 0
	// for i, file := range fileBuffer {
	// 	if file == "." {
	// 		break
	// 	}
	// 	weight, err := strconv.Atoi(file)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	checkSum += i * weight
	// }
	// fmt.Println(checkSum)

	// fmt.Println(files)
	for i := len(files) - 1; i >= 0; i-- {
		// try to move file of size files[i]
		window := files[i]
		count := 0
		for j := 0; j < len(fileBuffer)-1; j++ {
			if fileBuffer[j] == fmt.Sprint(i) {
				break
			}
			if fileBuffer[j] == "." {
				count += 1
				if count == window {
					// fmt.Println("replacing - ", i, j-count+1, j)
					// found a spot at fileBuffer[i] - window
					// clear the val
					for x := range fileBuffer {
						if fileBuffer[x] == fmt.Sprint(i) {
							//fmt.Println("replacing", x, i)
							fileBuffer[x] = "."
						}
					}
					//fmt.Println(fileBuffer)
					for swap := j - count + 1; swap <= j; swap++ {
						fileBuffer[swap] = fmt.Sprint(i)
					}
					//fmt.Println(fileBuffer)
					break
				}
			} else {
				count = 0
			}
		}
	}

	checkSum := 0
	for i, file := range fileBuffer {
		if file != "." {
			weight, err := strconv.Atoi(file)
			if err != nil {
				fmt.Println(err)
			}
			checkSum += i * weight
		}

	}
	fmt.Println(checkSum)
}
