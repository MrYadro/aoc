package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func cycle(cycleCount, sum *int, xsum int, img *[240]string) {
	draw(cycleCount, xsum, img)
	move(cycleCount, sum, xsum)
}

func move(cycleCount, sum *int, xsum int) {
	*cycleCount++
	if (*cycleCount+20)%40 == 0 {
		*sum += *cycleCount * xsum
	}
}

func draw(cycleCount *int, xsum int, img *[240]string) {
	if *cycleCount%40 >= xsum-1 && *cycleCount%40 <= xsum+1 {
		img[*cycleCount] = "#"
	} else {
		img[*cycleCount] = "."
	}
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	cycleCount := 0
	xsum := 1
	sum := 0

	img := [240]string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sl := scanner.Text()
		sla := strings.Split(sl, " ")
		if sla[0] == "addx" {
			cycle(&cycleCount, &sum, xsum, &img)
			cycle(&cycleCount, &sum, xsum, &img)
			add, _ := strconv.Atoi(sla[1])
			xsum += add
		}
		if sla[0] == "noop" {
			cycle(&cycleCount, &sum, xsum, &img)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for i := range img {
		if i%40 == 0 {
			fmt.Print("\n")
		}
		fmt.Print(img[i])
	}

	fmt.Print("\n")

	fmt.Println(sum)
}
