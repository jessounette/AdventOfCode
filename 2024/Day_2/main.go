package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

func isMonotonic(nums []int) bool {
	// Vérifier si la suite est monotone : croissante ou décroissante
	isIncreasing := true
	isDecreasing := true

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
	// Vérifier si les différences entre les éléments respectent les limites [1, 3]
	for i := 0; i < len(nums)-1; i++ {
		actualValue := nums[i]
		nextValue := nums[i+1]
		subtractValue := abs(nextValue - actualValue)
		if subtractValue < 1 || subtractValue > 3 {
			return false
		}
	}
	return true
}

// Fonction pour calculer la valeur absolue
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func isValidAfterRemoval(nums []int) bool {
	// Itération sur la liste
	for i := 0; i < len(nums); i++ {
		// On stock la valeur que l'on va supprimer pour la ré-insérer ensuite dans la liste
		val := nums[i]

		// Suppression de l'élément en utilisant slice
		if i < len(nums)-1 {
			nums = append(nums[:i], nums[i+1:]...)
		} else {
			// Si i est le dernier élément, on ne fait que couper la slice
			nums = nums[:i]
		}
		// Si la liste de vient valide en enlevant cet élement, on arrête le traitement
		if isMonotonic(nums) && lineIsSafe(nums) {
			return true
		}
		// On remet la valeur supprimée à la position i
		nums = append(nums[:i], append([]int{val}, nums[i:]...)...)
	}
	// La suppression d'un seul élement ne rends pas la liste valide
	return false
}

func checkSafetyReports_part1(result [][]int) int {
	safe := 0
	for _, line := range result {
		if lineIsSafe(line) && isMonotonic(line) {
			safe++
		}
	}
	return safe
}

func checkSafetyReports_part2(result [][]int) int {
	safe := 0
	// On parcourt chaque ligne
	for _, line := range result {
		// On vérifie d'abord si la ligne entière est valide
		if isMonotonic(line) && lineIsSafe(line) {
			safe++
		} else {
			// Sinon on vérifie si la ligne est valide après avoir enlevé un seul élément
			if isValidAfterRemoval(line) {
				safe++
			}
		}
	}
	return safe
}

func main() {
	table := parseFile()
	fmt.Println("Part 1 - checkSafetyReports : ", checkSafetyReports_part1(table))
	fmt.Println("Part 2 - checkSafetyReports : ", checkSafetyReports_part2(table))
}
