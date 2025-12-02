package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	start := 50
	timesAtZero := 0
	k := 100
	path := "input.txt"

	linesFile, err := os.Open(path)
	defer linesFile.Close()

	if err != nil {
		panic(err)
	}

	lineScanner := bufio.NewScanner(linesFile)

	if err != nil {
		panic(err)
	}

	for lineScanner.Scan() {
		lineContent := lineScanner.Text()
		direction := lineContent[0]
		number := lineContent[1:]

		if direction == 82 {
			convertedNumber, err := strconv.Atoi(number)
			if err != nil {
				panic(err)
			}

			start += convertedNumber
			timesAtZero += start / k
			start %= k
		} else {
			convertedNumber, err := strconv.Atoi(number)
			if err != nil {
				panic(err)
			}

			wasZero := start == 0
			start -= convertedNumber

			landedOnZero := start%k == 0
			if landedOnZero {
				timesAtZero++
			}

			for start < 0 {
				start += k
				timesAtZero++
			}

			if wasZero {
				timesAtZero--
			}
		}

	}

	fmt.Println(timesAtZero)
	saveInputToFile(timesAtZero, "output-part-2.txt")
	if scannerErr := lineScanner.Err(); scannerErr != nil {
		fmt.Print("Error during file reading")
	}

}
