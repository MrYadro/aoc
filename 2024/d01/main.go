package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var ls []int
	var rs []int
	var sum int
	var sim int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(line)

		result := strings.Split(line, "   ")
		lv, _ := strconv.Atoi(result[0])
		rv, _ := strconv.Atoi(result[1])
		ls = append(ls, lv)
		rs = append(rs, rv)
	}

	sort.Ints(ls)
	sort.Ints(rs)

	for i, lv := range ls {
		rv := rs[i]
		if lv > rv {
			sum += lv - rv
		} else {
			sum += rv - lv
		}
		for _, rv := range rs {
			if rv > lv {
				continue
			}
			if rv == lv {
				sim += rv
			}
		}
	}

	fmt.Println(sum)
	fmt.Println(sim)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
