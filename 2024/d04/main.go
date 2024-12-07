package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	toFind := "XMAS"

	first := toFind[0:1]

	mid := toFind[2:3]

	var field []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		field = append(field, line)
	}

	// fmt.Println(field)

	lenY := len(field)
	lenX := len(field[0])

	// fmt.Println(lenX)
	// fmt.Println(lenY)

	var sum int
	var sumX int

	for x := 0; x < lenX; x++ {
		for y := 0; y < lenY; y++ {
			if field[y][x:x+1] == first {
				for dx := -1; dx <= 1; dx++ {
					for dy := -1; dy <= 1; dy++ {
						if dy != 0 || dx != 0 {
							if y+dy*3 >= 0 && y+dy*3 < lenY && x+dx*3 >= 0 && x+dx*3 < lenX {
								if field[y][x:x+1] == "X" && field[y+dy*1][x+dx*1:x+dx*1+1] == "M" && field[y+dy*2][x+dx*2:x+dx*2+1] == "A" && field[y+dy*3][x+dx*3:x+dx*3+1] == "S" {
									sum++
								}
							}
						}
					}
				}
			}
		}
	}

	for x := 1; x < lenX-1; x++ {
		for y := 1; y < lenY-1; y++ {
			if field[y][x:x+1] == mid {
				if ((field[y-1][x-1:x-1+1] == "M" && field[y+1][x+1:x+1+1] == "S") || (field[y-1][x-1:x-1+1] == "S" && field[y+1][x+1:x+1+1] == "M")) &&
					((field[y+1][x-1:x-1+1] == "M" && field[y-1][x+1:x+1+1] == "S") || (field[y+1][x-1:x-1+1] == "S" && field[y-1][x+1:x+1+1] == "M")) {
					sumX++
				}
			}
		}
	}

	fmt.Println(sum)
	fmt.Println(sumX)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
