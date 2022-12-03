package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	max := []int{0, 0, 0}
	sum := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ist, _ := strconv.Atoi(scanner.Text())
		if ist != 0 {
			sum += ist
		} else {
			if sum > max[0] {
				max[0] = sum
			}
			sort.Ints(max)
			sum = 0
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Max: ", max[2])
	fmt.Println("Sum of top three: ", max[0]+max[1]+max[2])
}
