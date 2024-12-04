package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func purgeString() []string {
	file, err := os.Open("./2024/inputs/day_3.txt")
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	var tmp []string
	var purgedString []string
	var pattern = `mul\(\d+,\d+\)`
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			re := regexp.MustCompile(pattern)
			tmp = re.FindAllString(line, -1)
			purgedString = append(purgedString, tmp...)
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err) // Gérer une éventuelle erreur de lecture
	}
	return purgedString
}

func calculateTotal(list []string) int {
	var total = 0
	var cleanLine []string
	var pattern = `\d+,\d+`
	for _, elem := range list {
		re := regexp.MustCompile(pattern)
		cleanLine = re.FindAllString(elem, -1)
		for _, value := range cleanLine {
			tmp := strings.Split(value, ",")
			var nb1, _ = strconv.Atoi(tmp[0])
			var nb2, _ = strconv.Atoi(tmp[1])

			var mul = nb1 * nb2
			total += mul
		}
	}
	return total
}

func main() {
	fmt.Println("Part 1 - calculateTotal : ", calculateTotal(purgeString()))
}
