package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	folderPath := ""
	fileSizeMap := make(map[string]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sl := scanner.Text()
		if strings.HasPrefix(sl, "$ cd") {
			folderInfo := strings.Split(sl, " ")
			folderPath = path.Join(folderPath, folderInfo[2])
		} else if !strings.HasPrefix(sl, "dir") {
			fileInfo := strings.Split(sl, " ")
			fileSize, _ := strconv.Atoi(fileInfo[0])
			for currentDir := folderPath; currentDir != "/"; currentDir = path.Dir(currentDir) {
				fileSizeMap[currentDir] += fileSize
			}
			fileSizeMap["/"] += fileSize
		}
	}

	sum := 0
	enoughSpace := 30000000
	maxSpace := 70000000
	toDelete := fileSizeMap["/"]

	for _, folderSize := range fileSizeMap {
		if folderSize < 100000 {
			sum += folderSize
		}
		if maxSpace-(fileSizeMap["/"]-folderSize) >= enoughSpace && folderSize < toDelete {
			toDelete = folderSize
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(sum)
	fmt.Println(toDelete)
}
