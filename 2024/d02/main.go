package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func isSafe(s []int) bool {
	asc := s[1]-s[0] > 0

	for i := 1; i < len(s); i++ {
		diff := s[i] - s[i-1]

		if diff > 0 != asc {
			return false
		}

		if diff < 0 {
			diff = -diff
		}

		if diff < 1 || diff > 3 {
			return false
		}
	}
	return true
}

func isSafeSafe(s []int) bool {
	for i := 0; i < len(s); i++ {
		news := append([]int{}, s[:i]...)
		news = append(news, s[i+1:]...)

		if isSafe(news) {
			return true
		}
	}
	return false
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var safetyIndex int
	var safetySafetyIndex int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		numstr := strings.Split(line, " ")
		var nums []int
		for _, num := range numstr {
			n, _ := strconv.Atoi(num)
			nums = append(nums, n)
		}

		if isSafe(nums) {
			safetyIndex++
			safetySafetyIndex++
		} else {
			if isSafeSafe(nums) {
				safetySafetyIndex++
			}
		}
	}

	fmt.Println(safetyIndex)
	fmt.Println(safetySafetyIndex)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
