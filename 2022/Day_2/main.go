package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Part 1 - Result encrypted guide : ", encryptedGuide())
	fmt.Println("Part 2 - Result decrypted guide : ", decryptedGuide())
}

func encryptedGuide() int {
	var total = 0

	file, err := os.Open("./2022/inputs/day_2.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		turn := calculResultTurn(line)
		total = total + turn
	}
	return total
}

func decryptedGuide() int {
	var total = 0

	file, err := os.Open("./2022/inputs/day_2.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		turn := chooseShapeToPlay(line)
		total = total + turn
	}
	return total
}

func calculResultTurn(turn string) int {
	var result = 0

	switch turn {
	case "A X":
		result = 4
	case "A Y":
		result = 8
	case "A Z":
		result = 3

	case "B X":
		result = 1
	case "B Y":
		result = 5
	case "B Z":
		result = 9

	case "C X":
		result = 7
	case "C Y":
		result = 2
	case "C Z":
		result = 6
	default:
		result = 0
	}
	return result
}

func chooseShapeToPlay(turn string) int {
	var result = 0

	switch turn {
	case "A X":
		result = 3
	case "A Y":
		result = 4
	case "A Z":
		result = 8

	case "B X":
		result = 1
	case "B Y":
		result = 5
	case "B Z":
		result = 9

	case "C X":
		result = 2
	case "C Y":
		result = 6
	case "C Z":
		result = 7
	default:
		result = 0
	}
	return result
}
