package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var field []string
	field = append(field, strings.Repeat(".", 142)) // I'm lazy to control array limits
	for scanner.Scan() {
		ln := scanner.Text()
		field = append(field, "."+ln+".")
	}
	field = append(field, strings.Repeat(".", 142)) // I'm lazy to control array limits

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sum := 0
	gearRatio := 0

	for i, ln := range field {
		// re := regexp.MustCompile(`[^0123456789.]`) //p1
		re := regexp.MustCompile(`\*`) // p2
		fieldParts := re.FindAllStringIndex(ln, -1)
		for _, partIndex := range fieldParts {
			rei := regexp.MustCompile(`\d+`)
			aj := 0
			ngr := 1
			for offset := -1; offset < 2; offset++ {
				fnumi := rei.FindAllStringIndex(field[i+offset], -1)
				for _, nIndex := range fnumi {
					if nIndex[0] <= partIndex[1] && nIndex[1] >= partIndex[0] {
						li, _ := strconv.Atoi(string(field[i+offset][nIndex[0]:nIndex[1]]))
						sum += li
						ngr *= li
						aj++
					}
				}
			}
			if aj == 2 {
				gearRatio += ngr
			}
		}
	}
	fmt.Println("res:", sum)
	fmt.Println("res:", gearRatio)
}
