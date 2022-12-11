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

type monke struct {
	items        []int
	op           string
	opVal        int
	testDiv      int
	toMonkeTrue  int
	toMonkeFalse int
	inTimes      int
}

func monkeBusiness(monkes []monke, rounds, lcm int) int {
	monkes = append([]monke{}, monkes...)
	for i := 0; i < rounds; i++ {
		for j, monk := range monkes {
			for _, item := range monk.items {
				monkes[j].inTimes++

				op := monk.op
				wLevel := item

				if op == "+" {
					wLevel += monk.opVal
				} else if op == "*" {
					wLevel *= monk.opVal
				} else if op == "^" {
					wLevel *= wLevel
				}

				if rounds == 20 {
					wLevel = wLevel / 3
				} else if rounds == 10000 {
					wLevel = wLevel % lcm
				}

				if wLevel%monk.testDiv == 0 {
					monkes[monk.toMonkeTrue].items = append(monkes[monk.toMonkeTrue].items, wLevel)
				} else {
					monkes[monk.toMonkeFalse].items = append(monkes[monk.toMonkeFalse].items, wLevel)
				}
			}
			monkes[j].items = nil
		}
	}

	insp := []int{}

	for i := range monkes {
		insp = append(insp, monkes[i].inTimes)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(insp)))

	return insp[0] * insp[1]
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	monkes := []monke{}

	scanner := bufio.NewScanner(file)

	curMonke := monke{}

	lcm := 1

	var monkeName int

	for scanner.Scan() {
		sl := scanner.Text()
		ts := strings.TrimSpace(sl)
		ts = strings.NewReplacer(", ", ",", "* old", "^ 2").Replace(ts)
		_, err := fmt.Sscanf(ts, "Monkey %d:", &monkeName)
		if err == nil {
			curMonke = monke{}
		}
		var tItems string
		_, err = fmt.Sscanf(ts, "Starting items: %s", &tItems)
		if err == nil {
			ti := strings.Split(tItems, ",")
			for i := 0; i < len(ti); i++ {
				tItem, _ := strconv.Atoi(ti[i])
				curMonke.items = append(curMonke.items, tItem)
			}
		}
		var op string
		var opVal int
		_, err = fmt.Sscanf(ts, "Operation: new = old %s %d", &op, &opVal)
		if err == nil {
			curMonke.op = op
			curMonke.opVal = opVal
		}
		var testDiv int
		_, err = fmt.Sscanf(ts, "Test: divisible by %d", &testDiv)
		if err == nil {
			lcm *= testDiv
			curMonke.testDiv = testDiv
		}
		var toMonkeTrue int
		_, err = fmt.Sscanf(ts, "If true: throw to monkey %d", &toMonkeTrue)
		if err == nil {
			curMonke.toMonkeTrue = toMonkeTrue
		}
		var toMonkeFalse int
		_, err = fmt.Sscanf(ts, "If false: throw to monkey %d", &toMonkeFalse)
		if err == nil {
			curMonke.toMonkeFalse = toMonkeFalse
			monkes = append(monkes, curMonke)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sum := monkeBusiness(monkes, 20, lcm)

	fmt.Println(sum)

	sum = monkeBusiness(monkes, 10000, lcm)

	fmt.Println(sum)
}
