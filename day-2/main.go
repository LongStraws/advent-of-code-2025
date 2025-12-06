package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func saveInputToFile(output int, outputPath string) {
	newFile, err := os.Create(outputPath)
	if err != nil {
		panic(err)
	}

	defer newFile.Close()

	bytes, err := newFile.WriteString(strconv.Itoa(output))

	if err != nil {
		panic(err)
	}

	fmt.Print("%d bytes have been written to the file", bytes)

}
func partOne(start int, end int) int {
	res := 0
	for i := start; i <= end; i++ {
		curString := strconv.Itoa(i)
		stringLen := len(curString)

		if stringLen%2 == 0 {
			middle := stringLen / 2

			if curString[:middle] == curString[middle:] {
				fmt.Println(curString)
				res += i
			}
		}

	}
	return res

}
func partTwo(start int, end int) int {
	res := 0
	for i := start; i <= end; i++ {
		curString := strconv.Itoa(i)
		stringLen := len(curString)

		for k := 1; k <= stringLen/2; k++ {
			addCur := true
			if stringLen%k == 0 {

				for j := k; j <= stringLen-k; j += k {
					if curString[j-k:j] != curString[j:j+k] {
						addCur = false
					}
				}

				if addCur {
					res += i
					break

				}

			}
		}

	}
	return res
}
func main() {
	filePath := "input.txt"

	contentBytes, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file")
		return
	}
	fileContent := string(contentBytes)
	oldRanges := strings.Split(fileContent, "\n")
	ranges := strings.Split(oldRanges[0], ",")

	var res int = 0
	for _, u := range ranges {

		tmp := strings.Split(u, "-")

		str1, str2 := tmp[0], tmp[1]
		start, err := strconv.Atoi(str1)
		if err != nil {
			fmt.Println("string parsing failed", err)
			return
		}
		end, err := strconv.Atoi(str2)

		res += partTwo(start, end)
		if err != nil {
			fmt.Println("string parsing failed", err)
			return
		}

	}
	saveInputToFile(res, "output.txt")
	fmt.Print("Success")
}
