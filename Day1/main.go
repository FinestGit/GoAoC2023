package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	puzzleInput := readPuzzleInput()
	numArray := getValuesFromInput(puzzleInput)
	total := sumIntArray(numArray)
	fmt.Println(total)
}

func readPuzzleInput() []string {
	var fileContents []string

	file, err := os.Open("puzzle_input.txt")
	check(err)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		trimmedString := trimString(scanner.Text())
		fileContents = append(fileContents, trimmedString)
	}

	return fileContents
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func trimString(stringToTrim string) string {
	return strings.Trim(stringToTrim, "\n")
}

func getValuesFromInput(puzzleInput []string) []int {
	var numStrings []int
	for i := 0; i < len(puzzleInput); i++ {
		var possibleNums []string
		byteString := []byte(puzzleInput[i])

		for j := 0; j < len(byteString); j++ {
			possibleNum := byteString[j] - 48
			if possibleNum < 10 {
				possibleNums = append(possibleNums, string(byteString[j]))
			}
		}

		if len(possibleNums) == 1 {
			num, err := strconv.Atoi(possibleNums[0] + possibleNums[0])
			check(err)
			numStrings = append(numStrings, num)
		} else {
			num, err := strconv.Atoi(possibleNums[0] + possibleNums[len(possibleNums)-1])
			check(err)
			numStrings = append(numStrings, num)
		}
	}

	return numStrings
}

func sumIntArray(numArray []int) int {
	var runningTotal int
	for i := 0; i < len(numArray); i++ {
		runningTotal += numArray[i]
	}
	return runningTotal
}
