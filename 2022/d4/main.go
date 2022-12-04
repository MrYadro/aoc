package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sumfull := 0
	sumover := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sl := scanner.Text()

		re := regexp.MustCompile(`\d+`)

		match := re.FindAllString(sl, -1)

		p11, _ := strconv.Atoi(match[0])
		p12, _ := strconv.Atoi(match[1])
		p21, _ := strconv.Atoi(match[2])
		p22, _ := strconv.Atoi(match[3])

		isin := (p11 <= p21 && p12 >= p22) || (p21 <= p11 && p22 >= p12)
		isover := (p21 <= p11 && p22 >= p11) || (p21 <= p12 && p22 >= p12)

		if isin {
			sumfull++
		}

		if isover || isin {
			sumover++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Sum full: ", sumfull)
	fmt.Println("Sum overlap: ", sumover)
}
