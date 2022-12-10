package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func countDistinct(s string, n int) int {
	start := 0
	for i := 0; i < n; {
		sp := s[start : start+n]
		sym := sp[i]
		if strings.IndexByte(sp, sym) != strings.LastIndexByte(sp, sym) {
			start += i + 1
			i = 0
		} else {
			i++
		}
	}
	return start + n
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	sl := scanner.Text()

	count4 := countDistinct(sl, 4)
	count14 := countDistinct(sl, 14)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(count4)
	fmt.Println(count14)

}
