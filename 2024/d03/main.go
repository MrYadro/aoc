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

	var sum = 0
	var sum2 = 0

	var text string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		text += line
	}

	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	fi := re.FindAllStringSubmatch(text, -1)
	for _, v := range fi {
		m1, _ := strconv.Atoi(v[1])
		m2, _ := strconv.Atoi(v[2])
		sum += m1 * m2
	}

	reF := regexp.MustCompile(`don't\(\).*?do\(\)`)
	res := reF.ReplaceAllString(text, "")

	fmt.Println(res)

	rem := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	fi = rem.FindAllStringSubmatch(res, -1)
	for _, v := range fi {
		m21, _ := strconv.Atoi(v[1])
		m22, _ := strconv.Atoi(v[2])
		sum2 += m21 * m22
	}

	fmt.Println(sum)
	fmt.Println(sum2)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
