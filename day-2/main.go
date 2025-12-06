package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func saveInputToFile(output int, outputPath string) {
	newFile, err := os.Create(outputPath)
	defer newFile.Close()

	if err != nil {
		panic(err)
	}

	bytes, err := newFile.WriteString(strconv.Itoa(output))

	if err != nil {
		panic(err)
	}

	fmt.Print("%d bytes have been written to the file", bytes)

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

		if err != nil {
			fmt.Println("string parsing failed", err)
			return
		}

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
		fmt.Println(res, start, end)
	}
	fmt.Print("Success")
}
