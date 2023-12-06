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

var reps = map[string]string{
	"one":   "o1e",
	"two":   "t2o",
	"three": "t3e",
	"four":  "f4r",
	"five":  "f5e",
	"six":   "s6x",
	"seven": "s7n",
	"eight": "e8t",
	"nine":  "n9e",
}

func replaceLast(s string) string {
	last := 0
	gRep := ""
	for rep := range reps {
		li := strings.LastIndex(s, rep)
		lp := li + len(rep)
		if last < lp && li > 0 {
			last = lp
			gRep = rep
		}
	}
	return s[:last-len(gRep)] + reps[gRep] + s[last:]
}

func replaceFirst(s string) string {
	first := len(s)
	gRep := ""
	for rep := range reps {
		fi := strings.Index(s, rep)
		if first > fi && fi >= 0 {
			first = fi
			gRep = rep
		}
	}
	return s[:first] + reps[gRep] + s[first+len(gRep):]
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		ln := scanner.Text()

		sl := replaceLast(ln)
		sl = replaceFirst(sl)

		re := regexp.MustCompile(`[^0-9]`)
		s := re.ReplaceAllString(sl, ``)

		first := s[0]
		last := s[len(s)-1]

		fi, _ := strconv.Atoi(string(first))
		li, _ := strconv.Atoi(string(last))
		sum += fi*10 + li
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("res:", sum)
}
