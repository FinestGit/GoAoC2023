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
		fmt.Println(puzzleInput[i])
		possibleNums := walkString(puzzleInput[i])
		fmt.Println(possibleNums)

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

func walkString(stringToWalk string) []string {
	var possibleNum string
	var possibleNums []string
	i := 0
	j := len(stringToWalk)
	for i < j {
		temp := int(stringToWalk[i]) - 48
		if temp < 10 {
			possibleNums = append(possibleNums, string(stringToWalk[i]))
			possibleNum = ""
		} else {
			possibleNum += string(stringToWalk[i])
			checkedNum, isNum := checkStringForNum(possibleNum)
			if isNum {
				possibleNums = append(possibleNums, string(checkedNum))
				possibleNum = ""
			}
			if len(possibleNum) > 5 {
				possibleNum = ""
			}
		}
		i++
	}
	return possibleNums
}

func checkStringForNum(stringToCheck string) (string, bool) {
	if strings.Contains(stringToCheck, "one") {
		return "1", true
	}
	if strings.Contains(stringToCheck, "two") {
		return "2", true
	}
	if strings.Contains(stringToCheck, "three") {
		return "3", true
	}
	if strings.Contains(stringToCheck, "four") {
		return "4", true
	}
	if strings.Contains(stringToCheck, "five") {
		return "5", true
	}
	if strings.Contains(stringToCheck, "six") {
		return "6", true
	}
	if strings.Contains(stringToCheck, "seven") {
		return "7", true
	}
	if strings.Contains(stringToCheck, "eight") {
		return "8", true
	}
	if strings.Contains(stringToCheck, "nine") {
		return "9", true
	}
	return "0", false
}

func sumIntArray(numArray []int) int {
	var runningTotal int
	for i := 0; i < len(numArray); i++ {
		runningTotal += numArray[i]
	}
	return runningTotal
}
