package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func cpp(ch rune) int { // Calculate priority points
	val := rune(ch)
	if ch > 96 {
		val = ch - 96
	} else {
		val = ch - 38
	}
	return int(val)
}

func ol(sl string) int { // One liner handler
	sllen := len(sl)
	sp1 := sl[:sllen/2]
	sp2 := sl[sllen/2:]
	for _, ch := range sp1 {
		cp := strings.IndexRune(sp2, ch)
		if cp != -1 {
			return cpp(ch)
		}
	}
	return 0
}

func tl(sl []string) int {
	cru := make([]rune, 0, 10)
	for _, ch := range sl[0] {
		cp := strings.IndexRune(sl[1], ch)
		if cp != -1 {
			cru = append(cru, ch)
		}
	}
	for _, ch := range cru {
		cp := strings.IndexRune(sl[2], ch)
		if cp != -1 {
			return cpp(ch)
		}
	}
	return 0
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sumol := 0
	sumtl := 0
	tlg := make([]string, 0, 3) // Three liner group

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sl := scanner.Text()
		sumol += ol(sl)
		tlg = append(tlg, sl)
		if len(tlg)%3 == 0 {
			sumtl += tl(tlg)
			tlg = tlg[:0]
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Sum oneliner:", sumol)
	fmt.Println("Sum threeliner:", sumtl)
}
