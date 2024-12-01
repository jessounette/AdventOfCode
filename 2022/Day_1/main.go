package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	fmt.Println("Part 1 - MaxCal : ", countMaxCal())
	fmt.Println("Part 2 - Top 3 maxCal : ", countTop3MaxCal())
}

func countMaxCal() int {
	file, err := os.Open("./2022/inputs/day_1.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	var maxCal = 0
	var cal = 0

	for scanner.Scan() {
		var line = scanner.Text()
		if line != "" {
			var value, _ = strconv.Atoi(line)
			cal = cal + value
		} else {
			if maxCal < cal {
				maxCal = cal
			}
			cal, _ = strconv.Atoi(line)
		}
		if maxCal < cal {
			maxCal = cal
		}
	}
	return maxCal
}

func countTop3MaxCal() int {
	file, err := os.Open("./2022/inputs/day_1.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	var maxCal []int
	var cal = 0

	for scanner.Scan() {
		var line = scanner.Text()
		if line != "" {
			var value, _ = strconv.Atoi(line)
			cal = cal + value
		} else {
			maxCal = append(maxCal, cal)
			cal, _ = strconv.Atoi(line)
		}
	}
	maxCal = append(maxCal, cal)
	sort.Ints(maxCal)
	var top3 = maxCal[len(maxCal)-1] + maxCal[len(maxCal)-2] + maxCal[len(maxCal)-3]
	return top3
}
