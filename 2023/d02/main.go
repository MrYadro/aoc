package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func getMax(s, col string) int {
	re := regexp.MustCompile(`(\d{1,2}) ` + col)
	sms := re.FindAllStringSubmatch(s, -1)
	max := 0

	for _, sm := range sms {
		rv, _ := strconv.Atoi(string(sm[1]))
		if rv > max {
			max = rv
		}
	}
	return max
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	gameID := 1
	sum := 0
	for scanner.Scan() {
		ln := scanner.Text()

		// re := regexp.MustCompile(`(1[3-9]|20) red|(1[4-9]|20) green|(1[5-9]|20) blue`) p1
		// im := re.MatchString(ln)

		// if !im {
		// 	sum += gameID
		// }

		sum += getMax(ln, "red") * getMax(ln, "green") * getMax(ln, "blue")

		gameID++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("res:", sum)
}
