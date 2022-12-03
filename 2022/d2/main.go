package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var mp = map[string]int{ // Points map
	"l": 0,
	"d": 3,
	"w": 6,
}

var ms = map[string]int{ // Picked shape score map
	"r": 1,
	"p": 2,
	"s": 3,
}

func crs(om, pm string) int { // Calculate real score
	rs := map[string]map[string]int{
		"r": {
			"r": mp["d"] + ms[pm],
			"p": mp["w"] + ms[pm],
			"s": mp["l"] + ms[pm],
		},
		"p": {
			"r": mp["l"] + ms[pm],
			"p": mp["d"] + ms[pm],
			"s": mp["w"] + ms[pm],
		},
		"s": {
			"r": mp["w"] + ms[pm],
			"p": mp["l"] + ms[pm],
			"s": mp["d"] + ms[pm],
		},
	}
	return rs[om][pm]
}

func cps(om, rs string) int { // Calculate predicted score
	ps := map[string]map[string]int{
		"r": {
			"d": mp["d"] + ms["r"],
			"w": mp["w"] + ms["p"],
			"l": mp["l"] + ms["s"],
		},
		"p": {
			"d": mp["d"] + ms["p"],
			"w": mp["w"] + ms["s"],
			"l": mp["l"] + ms["r"],
		},
		"s": {
			"d": mp["d"] + ms["s"],
			"w": mp["w"] + ms["r"],
			"l": mp["l"] + ms["p"],
		},
	}
	return ps[om][rs]
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	mm := map[byte]string{
		65: "r",
		66: "p",
		67: "s",
		88: "r",
		89: "p",
		90: "s",
	}

	mr := map[byte]string{
		88: "l",
		89: "d",
		90: "w",
	}

	sumr := 0
	sump := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		st := scanner.Text()
		om := mm[st[0]] // Opponent move
		pm := mm[st[2]] // Player move
		pr := mr[st[2]] // Predicted result
		sumr += crs(om, pm)
		sump += cps(om, pr)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Real sum: ", sumr)
	fmt.Println("Predicted sum: ", sump)
}
