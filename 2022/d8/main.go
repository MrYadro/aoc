package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fieldArr := [][]int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sl := scanner.Text()
		lineArrS := strings.Split(sl, "")
		lineArr := []int{}
		for _, el := range lineArrS {
			iel, _ := strconv.Atoi(el)
			lineArr = append(lineArr, iel)
		}
		fieldArr = append(fieldArr, lineArr)
	}

	sum := 0
	sumscore := 0

	flen := len(fieldArr)

	for i := range fieldArr {
		for j := range fieldArr[i] {
			if i == 0 || j == 0 || i == flen-1 || j == flen-1 {
				sum++
			} else {
				ok1 := false
				ok2 := false
				ok3 := false
				ok4 := false
				cansee1 := 0
				cansee2 := 0
				cansee3 := 0
				cansee4 := 0
				for k := j - 1; k >= 0; k-- {
					cansee1++
					if fieldArr[i][k] < fieldArr[i][j] {
						ok1 = true
					} else {
						ok1 = false
						break
					}
				}

				for k := j + 1; k <= flen-1; k++ {
					cansee2++
					if fieldArr[i][k] < fieldArr[i][j] {
						ok2 = true
					} else {
						ok2 = false
						break
					}
				}

				for k := i - 1; k >= 0; k-- {
					cansee3++
					if fieldArr[k][j] < fieldArr[i][j] {
						ok3 = true
					} else {
						ok3 = false
						break
					}
				}

				for k := i + 1; k <= flen-1; k++ {
					cansee4++
					if fieldArr[k][j] < fieldArr[i][j] {
						ok4 = true
					} else {
						ok4 = false
						break
					}
				}
				if ok1 || ok2 || ok3 || ok4 {
					sum++
				}
				sumscore = max(sumscore, cansee1*cansee2*cansee3*cansee4)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(sum)
	fmt.Println(sumscore)
}
