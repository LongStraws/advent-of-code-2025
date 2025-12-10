package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type tupleDigitIndex struct {
	digit int
	index int
}

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

	fmt.Printf("%d bytes have been written to the file\n", bytes)
}

func main() {
	lines, err := readInputLines("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	part1Result := part1(lines)
	fmt.Printf("Part 1 result: %d\n", part1Result)

	part2Result := part2(lines)

	fmt.Printf("Part 2 result: %d\n", part2Result)

	saveInputToFile(part1Result, "output.txt")
	fmt.Println("Success")
}

func readInputLines(filePath string) ([]string, error) {
	contentBytes, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	content := strings.TrimSpace(string(contentBytes))
	if content == "" {
		return []string{}, nil
	}
	return strings.Split(content, "\n"), nil
}

func part1(batteries []string) int {
	total := 0
	for _, battery := range batteries {
		if battery == "" {
			continue
		}
		pairValue := strongestPairValue(battery)
		total += pairValue
	}
	return total
}

func part2(batteries []string) int {

	res := 0
	for _, battery := range batteries {

		tupleSlice := []int{}

		for index, volt := range battery {
			digit, _ := runeToDigit(volt)
			rem := len(battery) - index
			for len(tupleSlice) > 0 && tupleSlice[len(tupleSlice)-1] < digit && len(tupleSlice)-1+rem >= 12 {

				tupleSlice = tupleSlice[:len(tupleSlice)-1]
			}

			if len(tupleSlice) < 12 {
				tupleSlice = append(tupleSlice, digit)
			}

		}

		start := 0
		curRes := ""
		for start < 12 {
			curRes += strconv.Itoa(tupleSlice[start])
			start++
		}

		num, err := strconv.Atoi(curRes)
		if err != nil {
			fmt.Println("ERROR!")
			return 0
		}

		res += num
		fmt.Println(tupleSlice, battery)
	}

	fmt.Println("res is: ", res)
	return res
}

func strongestPairValue(battery string) int {
	first, second := -1, -1
	for _, r := range battery {
		digit, ok := runeToDigit(r)
		if !ok {
			continue
		}

		if digit > first {
			second = first
			first = digit
		} else if digit > second {
			second = digit
		}
	}

	if first < 0 || second < 0 {
		return 0
	}
	return combineDigits(first, second)
}

func runeToDigit(r rune) (int, bool) {
	if r < '0' || r > '9' {
		return 0, false
	}
	return int(r - '0'), true
}

func combineDigits(high, low int) int {
	return high*10 + low
}

func getMinIndexAndVal(tupleSlice []int) (int, int) {
	slices.Reverse(tupleSlice)
	for index, num := range tupleSlice {

		if num == slices.Min(tupleSlice) {
			return index, num
		}
	}

	return -1, -1
}
