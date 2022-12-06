package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func properAppend(a, b []string) []string { // GOD TIER THING
	la := len(a)
	c := make([]string, la, la+len(b))
	_ = copy(c, a)
	c = append(c, b...)
	return c
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var msp []string // Starting point
	var mrs []string // Rule set
	spread := false  // Starting point read

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sl := scanner.Text()
		sll := len(sl)

		if sll == 0 {
			mspl := len(msp)
			msp = append(msp[:mspl-1], msp[mspl:]...)
			spread = true
		} else if !spread {
			msp = append(msp, sl)
		} else if spread {
			mrs = append(mrs, sl)
		}
	}

	var field [][]string

	for i := range msp[0] {
		var fc []string
		if i%4 == 1 {
			for _, ln := range msp {
				rs := string(rune(ln[i]))
				if rs != " " { // Skip empty items in column
					fc = append(fc, string(rune(ln[i])))
				}
			}
			field = append(field, fc)
		}
	}

	fieldnew := make([][]string, len(field))
	for i := range field {
		fieldnew[i] = make([]string, len(field[i]))
		copy(fieldnew[i], field[i])
	}

	re := regexp.MustCompile(`\d+`)

	for _, rsl := range mrs {
		match := re.FindAllString(rsl, -1)
		count, _ := strconv.Atoi(match[0])
		from, _ := strconv.Atoi(match[1])
		to, _ := strconv.Atoi(match[2])
		for i := 0; i < count; i++ {
			field[to-1] = properAppend([]string{field[from-1][0]}, field[to-1])
			field[from-1] = properAppend(field[from-1][:0], field[from-1][1:])
		}
		fieldnew[to-1] = properAppend(fieldnew[from-1][:count], fieldnew[to-1])
		fieldnew[from-1] = properAppend(fieldnew[from-1][:0], fieldnew[from-1][count:])
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for _, col := range field {
		fmt.Print(col[0])
	}

	fmt.Print("\n")

	for _, col := range fieldnew {
		fmt.Print(col[0])
	}

	fmt.Print("\n")
}
