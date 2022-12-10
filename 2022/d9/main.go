package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func norm(num float64) float64 {
	if num != 0 {
		if num > 0 {
			num = 1
		} else {
			num = -1
		}
	}
	return num
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	moveMap := map[string]complex128{
		"U": 1,
		"D": -1,
		"R": 1i,
		"L": -1i,
	}

	moveSet1 := make(map[complex128]bool)
	moveSet10 := make(map[complex128]bool)

	curHeadPos := 0i

	rope := [10]complex128{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sl := scanner.Text()
		sla := strings.Split(sl, " ")
		curMove := sla[0]
		moveDist, _ := strconv.Atoi(sla[1])
		for i := 0; i < moveDist; i++ {
			curHeadPos += moveMap[curMove]
			rope[0] = curHeadPos
			for i := range rope {
				if i > 0 {
					diff := rope[i-1] - rope[i]
					diffUD := real(diff)
					diffLR := imag(diff)
					if math.Abs(diffUD)+math.Abs(diffLR) >= 2 {
						rope[i] += complex(norm(diffUD), norm(diffLR))
					}
				}
			}
			_, ex := moveSet10[rope[9]]
			if !ex {
				moveSet10[rope[9]] = true
			}

			_, ex = moveSet1[rope[1]]
			if !ex {
				moveSet1[rope[1]] = true
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(len(moveSet1))
	fmt.Println(len(moveSet10))
}
