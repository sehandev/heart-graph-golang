package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func checkOpenFile(err error) {
	if err != nil {
		log.Fatalf("open file error: %v", err)
		panic(err)
	}
}

func checkScanFile(err error) {
	if err != nil {
		log.Fatalf("scan file error: %v", err)
		panic(err)
	}
}

func readFile(fileName string) {
	// beginTime := time.Now()
	// fileName := "data/data.txt"

	fp, err := os.OpenFile(fileName, os.O_RDONLY, os.ModePerm)
	checkOpenFile(err)
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "사랑") {
			fmt.Println(line)
		}
	}

	checkScanFile(scanner.Err())

	// fmt.Println("scan file, time :", time.Since(beginTime).Seconds())
}
