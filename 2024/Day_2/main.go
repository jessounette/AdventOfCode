package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	table := parseFile()
	fmt.Println("Part 1 - checkSafetyReports : ", checkSafetyReports(table))
}

func parseFile() (table [][]int) {
	file, err := os.Open("./2024/inputs/day_2.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			// Découper la ligne en champs
			fields := strings.Fields(line)

			// Convertir chaque champ en entier
			var row []int
			for _, field := range fields {
				num, err := strconv.Atoi(field)
				if err != nil {
					panic(err) // Gérer les erreurs de conversion si nécessaire
				}
				row = append(row, num)
			}
			// Ajouter la ligne convertie au tableau 2D
			table = append(table, row)
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err) // Gérer une éventuelle erreur de lecture
	}

	return table
}

func isMonotonic(nums []int) (isMonotonic bool) {
	var isIncreasing = true
	var isDecreasing = true

	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			isIncreasing = false
		}
		if nums[i] > nums[i-1] {
			isDecreasing = false
		}
	}
	// Une suite est monotone si elle est soit croissante, soit décroissante, soit constante
	return isIncreasing || isDecreasing
}

func lineIsSafe(nums []int) bool {
	subtractValue := 0
	for i := 0; i < len(nums)-1; i++ {
		actualValue := nums[i]
		nextValue := nums[i+1]
		if actualValue >= nextValue {
			subtractValue = actualValue - nextValue
		} else {
			subtractValue = nextValue - actualValue
		}
		if subtractValue < 1 || subtractValue > 3 {
			return false
		}
	}
	return true
}

func checkSafetyReports(result [][]int) int {
	safe := 0
	for _, line := range result {
		if lineIsSafe(line) && isMonotonic(line) {
			safe++
		}
	}
	return safe
}
